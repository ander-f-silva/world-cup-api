# WORD CUP API

Projeto para aprender a fazer uma api com golang

Official Site: https://go.dev/

Documentation Site https://go.dev/doc/

## Libraries:

- GIN: Wrapper for protocol HTTP => https://github.com/gin-gonic/gin;
- GORM: Model ORM for Golang => https://gorm.io/index.html
- TESTIFY: Asserts and Mock => https://github.com/stretchr/testify 

## Utils: 

- GOLINT: Analyses Static => https://golangci-lint.run/

Command with docker:

````
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.43.0 golangci-lint run -v
````

- MOD for manager dependency (https://go.dev/blog/using-go-modules)

To do start dependecy 

````
 go mod init world-cup-api
 ````

## Install the dependencies 

GIN

````
 go get -u github.com/gin-gonic/gin
 ````

GORM and SQL LITE 3

````
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
````

## Rund the applications with Docker

````
docker run --rm --name world-cup-api -p 3000:3000 -v "$PWD":/usr/src/app -w /usr/src/app golang:1.17 go run src/main.go
````
