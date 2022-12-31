# GoRestBoilerPlate
A simple REST API boilerplate for Go

## Why use this boilerplate
This project has the following features:
* Fast routing with Gin
* JWT middleware Included
* User authentication routes added
* GORM, Go ORM added with support for multiple databases
* Secure route authorization using user roles
* Simple project structure

## Getting Started
### SQL
To create a new SQL migration use this command
```
migrate create -ext sql -dir ./infrastructure/modules/gorm/migrations -seq NAME_FOR_MIGRATION
```

Manually run migrations
```
migrate -path infrastructure/modules/gorm/migrations -database "postgresql://dbuser:dbpass@localhost:5432/boilerplate?sslmode=disable" -verbose up
```

Rollback version
```
migrate -path infrastructure/modules/gorm/migrations -database "postgresql://dbuser:dbpass@localhost:5432/boilerplate?sslmode=disable" force 15
```

### NoSQL
To create a new NoSQL migration use this command
```
go run infrastructure/modules/mongo/utils/migrations/manager.go new NAME_FOR_MIGRATION
```

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.
```
go get ...
go run main.go
```

### Prerequisites

Go environment must be setup.

For protos you should install protobuf on your system and execute the following command to install the plugin for golang protos generator.
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@1.5.2
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@1.2.0
```

To build the protos use the command below
```
protoc infrastructure/modules/grpc/protos/raw/*.proto --go_out=infrastructure/modules/grpc/protos/build --go_opt=paths=source_relative --go-grpc_out=infrastructure/modules/grpc/protos/build --go-grpc_opt=paths=source_relative --proto_path=infrastructure/modules/grpc/protos/raw
```
   
## Built With

* [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
* [GORM](http://jinzhu.me/gorm/) - ORM library for Golang
 