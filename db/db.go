package db

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// GO_ENVフラグをコマンドライン引数から取得
	var flag_env = flag.String("GO_ENV", "", "開発環境フラグ")
	flag.Parse()

	// 開発環境（"dev"）の場合、.envファイルを読み込んで環境変数を設定
	if *flag_env == "dev" {
		err := godotenv.Load() // .envファイルをロード
		if err != nil {
			log.Fatalln(err) // .envファイルの読み込みに失敗した場合はエラーログを出力して終了
		}
	}

	// PostgreSQLの接続情報を環境変数から取得し、接続URLを作成
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	// GORMを使用してPostgreSQLデータベースに接続します。
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err) // データベース接続に失敗した場合はエラーログを出力して終了
	}
	fmt.Println("Connected") // 接続成功のメッセージを表示
	return db                // データベース接続インスタンスを返します。
}

// 開いているデータベース接続を閉じる
func CloseDB(db *gorm.DB) {
	// GORMのDBインスタンスからネイティブのSQL DBインスタンスを取得
	sqlDB, _ := db.DB()
	// SQL DBインスタンスを閉じる
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err) // クローズに失敗した場合はエラーログを出力して終了
	}
}
