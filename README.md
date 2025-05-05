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

# Run

```sh
go run ./cmd/api/main.go
```

# Test healthcheck

```sh
curl http://localhost:8080/healthcheck
```
