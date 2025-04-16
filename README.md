# 特徴
ユーザー登録機能を備えた簡易的なタスク管理アプリです。
Go言語を用いて開発しており、フレームワークにはEcho、ORMにgorm、データベースにPostgresDataBase、開発環境の構築にdockerを使用しています。
以下のリンクの講座を見てアプリを作成し、すべてdocker composeで動作するように変更しています。マルチステージビルドを採用しているので
https://www.udemy.com/course/echo-go-react-restapi/?couponCode=ST14MT150425G2

## 機能
### ユーザー管理機能
- 新規登録
- ログイン
- ログアウト

### タスク
- 一括取得
- idでの取得
- 新規作成
- 更新
- 削除

### バリデーション
- ユースケース層でのユーザー、タスク構造体のバリデーション
- emailの形式チェック
- その他文字列の文字数チェック

### セキュリティ対策
- CORS対策
- CSRF対策
- SQLインジェクション対策(gormの機能)

## 構成
### クリーンアーキテクチャ
以下に示すような、クリーンアーキテクチャーを採用した3層構造となっています。
コントローラー　<->　ユースケース <->　リポジトリ
それぞれの役割は、
- コントローラ : クライアントからのリクエストを処理してユースケース層に渡し、レスポンスをクライアントに返す。
- ユースケース : コントローラから受け取ったデータを処理、またはリポジトリ層を呼び出して処理を行う。
- リポジトリ 　: ユースケース層からの呼び出しでデータベースへの操作を行う。
となっている。

これらの層はインターフェイスを介して疎結合になっているので、変更やテストが容易になっています。
依存の流れは、コントローラー　->　ユースケース <-　リポジトリ　のように内側の層が外部の層に依存しないようになっています。

### フォルダの役割

## 使い方
###アプリの起動、停止方法
.envファイルを作成し、以下の内容を書き込みます。
PORT=8080
POSTGRES_USER=db_user
POSTGRES_PW=db_password
POSTGRES_DB=db_name
POSTGRES_PORT=5434
POSTGRES_HOST=db
SECRET=uu5pveql
GO_ENV=dev
API_DOMAIN=localhost
FE_URL=http://localhost:3000

dockerがインストールされているパソコンで、このdocker-compose.ymlが入っているディレクトリまで移動し、以下のコマンドを打ち込みます。
docker-compose up
これで、GoのWEBサーバーが起動しているコンテナとデータベースのコンテナが立ち上がります。
もしうまく立ち上がらなかった場合は、データベースのコンテナを先に立ち上げると上手くいくはずです。

試し終わったら docker-compose down ですべてのコンテナをストップできます。

### APIの使い方
- POST /signup
- POST /login
- POST /logout
- GET /csrf

- GET    /tasks  すべてのタスクを取得
- POST   /tasks  タスクの作成
- GET    /tasks/:taskid  task id からタスクの取得
- PUT    /tasks/:taskid  task id からタスクの更新
- DELETE /tasks/:taskid  task id からタスクの削除

### ユーザー登録からログインまでの流れ

## 改善点
