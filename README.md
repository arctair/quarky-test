# quarky-test
Test commits for quarky's automated tests
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
$ BASE_URL=https://quarky-test.arctair.com go test
```
## Run the server
```
$ go run .
$ curl localhost:5000
```
## Build a docker image
```
$ sh build
$ docker build -t arctair/quarky-test:<scenario> .
```
