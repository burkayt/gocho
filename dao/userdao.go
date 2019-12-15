package dao

import (
	"gocho/models"
)

func User(id int) (*models.User, error) {
	user := models.User{}
	var err error
	err = db.Get(&user, "select * from User where ID = ?", id)
	return &user, err
}

func Users() (users []*models.User, err error) {
	err = db.Select(&users, "Select * from User")
	return
}

func CreateUser(user *models.User) (*models.User, error) {
	var err error
	tx := db.MustBegin()
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

func DeleteUser(id int) (err error) {
	tx := db.MustBegin()
	defer tx.Commit()

	_, err = tx.Exec("DELETE from User where id = ?", id)
	return
}

func FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	var err error
	err = db.Get(&user, "select * from User where EMAIL = ?", email)
	return &user, err
}
