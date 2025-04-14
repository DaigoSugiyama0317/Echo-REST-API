package router

import (
	"net/http"
	"os"

	"github.com/DaigoSugiyama0317/Echo-REST-API/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Echoのインスタンスを作成、ルートやミドルウェアを設定
// 引数として受け取るucとtcは、ユーザーとタスクのコントローラインターフェース
func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()

	// CORSミドルウェアを設定し、特定のオリジンからのリクエストを許可
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")}, // フロントエンドのURLを許可
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"}, // 許可するHTTPメソッド
		AllowCredentials: true,                                     // クッキーの送信を許可
	})) // 許可するヘッダ

	// CSRF保護のためのミドルウェアを設定、クッキーの設定を行い、セキュリティを強化
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",                     // クッキーのパス
		CookieDomain:   os.Getenv("API_DOMAIN"), // APIドメイン
		CookieHTTPOnly: true,                    // クッキーのJavaScriptアクセスを無効化
		CookieSameSite: http.SameSiteNoneMode,   // クロスサイトリクエスト時にクッキーを送信する
		CookieMaxAge:   60,                      // クッキーの有効期限（秒）
	}))

	// ユーザー関連のエンドポイントを設定
	e.POST("/signup", uc.SignUp) // サインアップ
	e.POST("/login", uc.LogIn)   // ログイン
	e.POST("/logout", uc.LogOut) // ログアウト
	e.GET("/csrf", uc.CsrfToken) // CSRFトークン取得

	// タスク関連のエンドポイント
	// このグループ内のエンドポイントはJWT認証を使用
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")), // JWT署名に使用する秘密鍵
		TokenLookup: "cookie:token",              // トークンはクッキーから取得
	}))
	// タスク関連のエンドポイントを設定
	t.GET("", tc.GetAllTasks)           // すべてのタスクを取得
	t.GET("/:taskId", tc.GetTaskById)   // ID指定でタスクを取得
	t.POST("", tc.CreateTask)           // 新しいタスクを作成
	t.PUT("/:taskId", tc.UpdateTask)    // タスクを更新
	t.DELETE("/:taskId", tc.DeleteTask) // タスクを削除
	return e
}
