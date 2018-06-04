[![Build Status](https://travis-ci.org/magaldima/bizday.svg?branch=master)](https://travis-ci.org/magaldima/bizday)
# Bizday
plugin based business calendar microservice exposes a `gRPC` API to perform complex business date and time calculation logic.

### Prerequisites
- Go v1.10
- dep

## Getting Started
1. Build example holiday plugin and set holiday plugin env var
```
$ go build -o dist/holiday-go-grpc ./holiday/plugin-go-grpc
$ export HOLIDAY_PLUGIN="./dist/holiday-go-grpc"
```

2. Build example day count basis plugin and set dcb plugin env var
```
$ go build -o dist/dcb-go-grpc ./dcb/plugin-go-grpc
$ export DCB_PLUGIN="./dist/dcb-go-grpc"
```

3. Run the server
```
$ go run main.go
```

4. Run the client
```
$ go run example/main.go
```
