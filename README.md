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

# Run

```sh
go run ./cmd/api/main.go
```

# Test healthcheck

```sh
curl http://localhost:8080/healthcheck
```
