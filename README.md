# Bizday
plugin based business calendar microservice exposes a `gRPC` API to perform complex business date and time calculation logic.

## Getting Started

### Prerequisites
- Go v1.10
- dep

1. Build example holiday plugin
```
$ go build -o dist/holiday-go-grpc ./holidays/plugin-go-grpc
```

2. Run the server
```
go run main.go
```

3. Run the client
```
go run example/main.go
```
