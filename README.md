# golang-echo-admin

## About

This Go repository was inspired from adding production level observability features to the Echo server framework in the development of microservices.
The implementation of the features included in this repository was inspired from the actuator capability that exists in Spring Boot.

#### What you might find helpful in this repo.
- [x] [echo](https://github.com/labstack/echo)(Framework) for handling requests
- [x] [golang-stats-api-handler](github.com/fukata/golang-stats-api-handler)
- [x] [Production-ready Features -- Spring](https://docs.spring.io/spring-boot/docs/current/reference/html/actuator.html)

## How to run

Install.

```shell
git clone git@github.com:imforster/golang-echo-admin

cd golang-echo-admin
```

Download dependencies.
```shell
go mod download
```

Run.
```shell
cd examples
make build
./adminTest
```
