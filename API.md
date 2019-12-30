# cURL

### `CreateUser`
```
$ curl -H "Content-Type: application/json" \
-X POST "http://api.dev.pepeunlimited.com/twirp/pepeunlimited.users.UserService/CreateUser" \
-d '{"username": "ssimoo", "email": "simo@gmail.com", "password": "p4sw0rd"}'
```
### `GetUser`
```
$ curl -H "Content-Type: application/json" \
-H "Authorization: Bearer REPLACE_WITH_TOKEN" \
-X POST "http://api.dev.pepeunlimited.com/twirp/pepeunlimited.users.UserService/GetUser" \
-d '{}'
```
### `VerifySignIn`
```
$ curl -H "Content-Type: application/json" \
-X POST "http://api.dev.pepeunlimited.com/twirp/pepeunlimited.users.UserService/VeriySignIn" \
-d '{"username": "ssimoo", "password": "p4sw0rd"}'
```
### `ForgotPassword`
```
$ curl -H "Content-Type: application/json" \
-X POST "http://api.dev.pepeunlimited.com/twirp/pepeunlimited.users.UserService/ForgotPassword" \
-d '{"username": "ssimoo", "language": "fi"}'
```