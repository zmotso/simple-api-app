# simple-api-app
simple-api-app is API for testing purpose. It returns json response.
You can define port by env variable `PORT` (default is 8080).

### Running
From source:

```sh
go run .
```

From docker image:

```sh
docker run -pd 8080:8080 zmotso/simple-api-app:0.0.1
```
