FROM golang:1.24-alpine

RUN apk add --no-cache git curl

# airをインストール
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# air設定ファイルとGoソースをマウント
COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]
