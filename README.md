## Simple api

Implementation api (Golang)

## Tools used

-  Gin
-  Mongo Driver
-  Swagger

## First start

Before the start you need to install docker and run it.

1.Clone repository

2.Install dependencies

```bash
go mod download
```

3.Start dev database

```bash
docker-compose --profile dev up -d 
```

4.Start app

```bash
make run-core
```

## Start prod

Before the start you need to install docker and run it.


```bash
docker-compose --profile prod up -d 
```

## Documentation

After the start applications (auth in env):

-  [Swagger](http://localhost:3000/swagger/index.html)

Generate new documentation:

```bash
swag init -g ./cmd/core/main.go -o docs
```

## Postman

If you want to test endpoints. You can import this file in Postman

- go-simple-api.postman_collection.json

## Structure

```
.
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── core
│       ├── auth
│       │   ├── controller.go
│       │   ├── repository.go
│       │   └── routes.go
│       ├── main.go
│       └── post
│           ├── controller.go
│           ├── repository.go
│           └── routes.go
├── config
│   └── init.go
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
└── utils
    ├── constants
    │   ├── collections.go
    │   └── error-messages.go
    ├── exception
    │   └── exception.go
    ├── helpers
    │   ├── get-context-data.go
    │   └── get-context-userid.go
    ├── middleware
    │   ├── base-auth-swagger.go
    │   ├── guard.go
    │   └── validator.go
    ├── models
    │   ├── auth.go
    │   ├── error.go
    │   ├── post.go
    │   └── user.go
    ├── schemas
    │   ├── post.go
    │   └── user.go
    ├── server
    │   └── app.go
    └── services
        └── auth.go

```

## Go version

- 1.23.1


