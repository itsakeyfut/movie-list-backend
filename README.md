# Install Dependencies

## Gin

```sh
go get github.com/gin-gonic/gin
```

## dotenv

```sh
go get github.com/joho/godotenv
```

## GORM

```sh
go get gorm.io/gorm
```

```sh
go get gorm.io/driver/postgres
```

## CORS

```sh
go get github.com/gin-contrib/cors
```

## logrus

```sh
go get github.com/sirupsen/logrus
```

### Log Level

| メソッド       | 用途                       |
| -------------- | -------------------------- |
| `Log.Debugf()` | 開発時の詳細ログ           |
| `Log.Infof()`  | 通常の操作ログ             |
| `Log.Warnf()`  | 問題になりそうな状況       |
| `Log.Errorf()` | エラーが発生した場合       |
| `Log.Fatalf()` | 致命的エラー（アプリ終了） |

## JWT

```sh
go get github.com/golang-jwt/jwt/v5
```

## swag cli

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

## Gin-Swagger

```sh
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
go get github.com/swaggo/swag
```

## golangci-lint

```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

# Run

```sh
go run ./cmd/api/main.go
```

# Test healthcheck

```sh
curl http://localhost:8080/healthcheck
```

# Swagger UI

http://localhost:8090/swagger/index.html
