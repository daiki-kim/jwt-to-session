# jwt-to-session

12-6: ProfileまでのJWTのコードをsessionに切り替える形で作成。

- `auth/session.go`を追加。

- `middlewares/session_auth.go`を追加。

- `controllers/logout.go`を追加。

- `controllers/login.go`のtoken generate部分をset sessionに変更。

- `main.go`のrouterをsessionに切り替え。

- sessionの確認は`cookies.txt`にcookieを作成するcurlコマンドを使用。

## API Endpoint

サーバーを起動:

```bash
go run main.go
```

### Signup

```bash
curl --location --request POST '0.0.0.0:8080/api/v1/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "jun",
    "email": "test@alamoa.io",
    "password": "password"
}'
```

### Login

```bash
curl --location --request POST -c cookies.txt '0.0.0.0:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "jun",
    "email": "test@alamoa.io",
    "password": "password"
}'
```

### Profile

```bash
 curl --location --request GET -b cookies.txt '0.0.0.0:8080/api/v1/profile'
 ```

### Logout

```bash
curl --location --request POST -b cookies.txt -c cookies.txt '0.0.0.0:8080/api/v1/logout'
```
