# quarky
Automated deployment and verification of [hashbang-api](https://github.com/arctair/hashbang-api) to Kubernetes
## Run the tests
```
$ go test
```
or
```
$ nodemon
```
### Run the tests against a deployment
```
$ BASE_URL=https://quarky.arctair.com go test
```
## Run the server
```
$ go run .
$ curl localhost:5000
```
## Build a docker image
```
$ go build -o bin/quarky
$ docker build -t arctair/quarky .
```
