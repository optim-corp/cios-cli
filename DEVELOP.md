# Development

Compliant with Golang build.
## How to Build

```shell
go build -o cios main.go
```

## Cross Compile

### Linux

```shell
GOOS=linux GOARCH=amd64 go build -o
```

### MacOS

```shell
GOOS=darwin GOARCH=amd64 go build -o
```

### Windows

```shell
GOOS=windows GOARCH=amd64 go build -o cios.exe
```

## 