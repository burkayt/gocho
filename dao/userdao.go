package dao

import (
	"gocho/models"
)

type UserDao struct {
	Db Database
}

func NewUserDao(db Database) models.UserDao {
	return &UserDao{db}
}

func (userDao *UserDao) User(id int) (*models.User, error) {
	user := models.User{}
	var err error
	err = userDao.Db.GetConnection().Get(&user, "select * from User where ID = ?", id)
	return &user, err
}

func (userDao *UserDao) Users() (users []*models.User, err error) {
	err = userDao.Db.GetConnection().Select(&users, "Select * from User")
	return
}

func (userDao *UserDao) CreateUser(user *models.User) (*models.User, error) {
	var err error
	tx := userDao.Db.GetConnection().MustBegin()
	defer tx.Commit()

	exec, err := tx.NamedExec(`insert into User(NAME, SURNAME, PASSWORD, EMAIL) values (:name, :surname, :password, :email)`, map[string]interface{}{
		"name":     user.Name,
		"surname":  user.Surname,
		"password": user.Password,
		"email":    user.Email,
	})

	id, err := exec.LastInsertId()

	user.Id = int(id)

	return user, err

}

func (userDao *UserDao) DeleteUser(id int) (err error) {
	tx := userDao.Db.GetConnection().MustBegin()
	defer tx.Commit()

	_, err = tx.Exec("DELETE from User where id = ?", id)
	return
}

func (userDao *UserDao) FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	var err error
	err = userDao.Db.GetConnection().Get(&user, "select * from User where EMAIL = ?", email)
	return &user, err
}
