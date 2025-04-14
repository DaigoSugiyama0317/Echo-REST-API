# ビルドステージ
FROM golang:1.24-alpine AS builder

# ビルドに必要なパッケージをインストール
RUN apk add --no-cache git gcc musl-dev

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係のファイルをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -trimpath -ldflags "-w -s" -o webapp

# 実行ステージ
FROM alpine:3.19

# 必要なランタイム依存関係をインストール
RUN apk --no-cache add ca-certificates tzdata

# タイムゾーンを設定
ENV TZ=Asia/Tokyo

# 非root ユーザーを作成
RUN adduser -D -g '' appuser

# 作業ディレクトリを設定
WORKDIR /app

# ビルドステージからバイナリをコピー
COPY --from=builder /app/webapp /app/
# 設定ファイルやその他の静的ファイルをコピー（必要に応じて）
COPY --from=builder /app/config /app/config
COPY --from=builder /app/static /app/static
COPY --from=builder /app/templates /app/templates

# アプリケーションディレクトリの所有者を変更
RUN chown -R appuser:appuser /app

# 非root ユーザーに切り替え
USER appuser

# アプリケーションのポートを公開
EXPOSE 8080

# コマンドを実行
CMD ["/app/webapp"]