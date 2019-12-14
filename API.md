# cURL

### `CreateUser`
```
$ curl -H "Content-Type: application/json" \
-X POST "localhost:8080/twirp/pepeunlimited.users.UserService/CreateUser" \
-d '{"username": "ssimoo", "email": "simo@gmail.com", "password": "p4sw0rd"}'
```
### `GetUser`
```
$ curl -H "Content-Type: application/json" \
-H "Authorization: Bearer REPLACE_WITH_TOKEN" \
-X POST "localhost:8080/twirp/pepeunlimited.users.UserService/GetUser" \
-d '{}'
```
##### `NOTE: test without nginx-ingress controller`
```
$ curl -H "Content-Type: application/json" \
-H "X-JWT-UserId: 1" \
-X POST "localhost:8080/twirp/pepeunlimited.users.UserService/GetUser" \
-d '{}'
```