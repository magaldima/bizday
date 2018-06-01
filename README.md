# Bizday
plugin based business calendar microservice exposes a `gRPC` API to perform complex business date and time calculation logic.

### Prerequisites
- Go v1.10
- dep

## Getting Started
1. Build example holiday plugin and set holiday plugin env var
```
$ go build -o dist/holiday-go-grpc ./holidays/plugin-go-grpc
$ export HOLIDAY_PLUGIN="./dist/holiday-go-grpc"
```

2. Build example calendar plugin and set calendar plugin env var
```
$ go build -o dist/calendar-go-grpc ./calendar/plugin-go-grpc
$ export CALENDAR_PLUGIN="./dist/calendar-go-grpc"
```

3. Run the server
```
$ go run main.go
```

4. Run the client
```
$ go run example/main.go
```
