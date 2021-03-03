```shell
# Git Submodule Update

$ git submodule foreach git pull origin master
$ git submodule sync
$ git submodule update

# Golang Dependency

$ go install

# OpenAPI Generate

$ docker-compose up
$ go run main.go

# Cross Compile

# Linux
# GOOS=linux GOARCH=amd64
# Mac
# GOOS=darwin GOARCH=amd64

$ go build -o cios

# Windows
# GOOS=windows GOARCH=amd64

$ go build -o cios.exe


```