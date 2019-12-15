FROM golang:1.13 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM scratch

WORKDIR /app
COPY --from=builder /app/application.yaml .
COPY --from=builder /app/app .
EXPOSE 8080

CMD ["./app"]