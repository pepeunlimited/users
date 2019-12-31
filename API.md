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
-d '{"username": "ssimoo", "email": "ssimoo@gmail.com", "language": "fi"}'
```
### `VerifyForgotPassword`
```
$ curl -H "Content-Type: application/json" \
-X POST "http://api.dev.pepeunlimited.com/twirp/pepeunlimited.users.UserService/VerifyForgotPassword" \
-d '{"token": "nk_gZUjN9gTXIAYLIEm1ITyj7DCtV-861JTeu87HzeA="}'
```
### `ResetPassword`
```
$ curl -H "Content-Type: application/json" \
-X POST "http://api.dev.pepeunlimited.com/twirp/pepeunlimited.users.UserService/ResetPassword" \
-d '{"token": "nk_gZUjN9gTXIAYLIEm1ITyj7DCtV-861JTeu87HzeA=", "password":"newpw"}'
```
### `UpdatePassword`
```
$ curl -H "Content-Type: application/json" \
-H "Authorization: Bearer REPLACE_WITH_TOKEN" \
-X POST "http://api.dev.pepeunlimited.com/twirp/pepeunlimited.users.UserService/UpdatePassword" \
-d '{"current_password": "currpw", "new_password":"newpw"}'
```