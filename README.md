# Install Dependencies

## Gin

```sh
go get -u github.com/gin-gonic/gin
```

## dotenv

```sh
go get -u github.com/joho/godotenv
```

# Run

```sh
go run ./cmd/api/main.go
```

# Test healthcheck

```sh
curl http://localhost:8080/healthcheck
```
