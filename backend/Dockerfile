# ============================================
# ビルドステージ
# ============================================
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 依存関係のダウンロード（キャッシュ効率化）
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# バイナリをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api

# ============================================
# 実行ステージ（軽量イメージ）
# ============================================
FROM alpine:latest

WORKDIR /app

# タイムゾーン設定とCA証明書
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Tokyo

# ビルドステージからバイナリをコピー
COPY --from=builder /app/main .

# migrationsディレクトリもコピー（必要に応じて）
COPY --from=builder /app/migrations ./migrations

# ポート公開
EXPOSE 8080

# 実行
CMD ["./main"]
