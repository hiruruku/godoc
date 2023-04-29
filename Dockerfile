# ビルドステージ
FROM golang:1.20 AS builder
# 作業ディレクトリを設定
WORKDIR /app
# ソースコードと依存関係ファイルをコピー
COPY . .

# 依存関係のインストール
RUN go mod tidy
# RUN go get github.com/labstack/echo/v4

# アプリケーションのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

# 実行ステージ
FROM alpine:latest

# RUN apk --no-cache add ca-certificates
WORKDIR /root/


# 必要なCA証明書をコピー
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# ビルドステージから実行ファイルをコピー
COPY --from=builder /app/main ./
COPY --from=builder /app/md ./md

# 実行コマンド
CMD [ "./main"]

# ポートを公開
EXPOSE 8080
