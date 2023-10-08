# Hotel reservation backend

## Projects outline
- users -> book room from a hotel
- admins -> going to check reservation/bookings
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- ROOM -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Resources
Инициализация зависимостей
```
go mod init github.com/EvgeniyBudaev/golang-hotel-reservation
```

### Mongodb driver
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber
Documentation
```
https://gofiber.io
```

Installing gofiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```

## BCRYPT
Documentation
```
https://golang.org/x/crypto/bcrypt
```

Installing bcrypt
```
go get golang.org/x/crypto/bcrypt
```

Fix errors
```
go mod tidy
```

## JWT
Documentation
```
https://github.com/golang-jwt/jwt
https://golang-jwt.github.io/jwt/usage/create
https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac
```

Installing
```
go get -u github.com/golang-jwt/jwt/v5
```