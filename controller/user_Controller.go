package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/DaigoSugiyama0317/Echo-REST-API/model"
	"github.com/DaigoSugiyama0317/Echo-REST-API/usecase"
	"github.com/labstack/echo/v4"
)

// ユーザーに関するコントローラーのインターフェース定義
// 各関数は HTTP リクエストを受け取って処理を行う
type IUserController interface {
	SignUP(c echo.Context) error    // ユーザー登録
	LogIn(c echo.Context) error     // ログイン
	LogOut(c echo.Context) error    // ログアウト
	CsrfToken(c echo.Context) error // CSRFトークンの取得
}

// コントローラー構造体：ユースケース層を保持して依存注入
type userController struct {
	uu usecase.IUserUsecase
}

// コントローラーのコンストラクタ関数
// ユースケースを受け取って、userController構造体を返す
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// ユーザー登録用のハンドラー
func (uc *userController) SignUP(c echo.Context) error {
	user := model.User{}
	// リクエストボディを user 構造体にバインド
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// ユースケース層に登録処理を呼び出し
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// 成功したら201 Createdでユーザー情報を返す
	return c.JSON(http.StatusCreated, userRes)
}

// ログイン処理のハンドラー
func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	// リクエストボディを user 構造体にバインド
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// ユースケース層でトークンの発行
	tokenString, err := uc.uu.LogIn(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// トークンをCookieとしてセット（セキュリティ設定付き）
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN") // ドメインは環境変数から取得
	cookie.Secure = true                    // HTTPS限定
	cookie.HttpOnly = true                  // JSからアクセス不可
	cookie.SameSite = http.SameSiteNoneMode // クロスサイト送信許可
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// ログアウト処理のハンドラー
func (uc *userController) LogOut(c echo.Context) error {
	// Cookieの内容を無効にして上書き（即時期限切れ）
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// CSRFトークンをJSONで返すハンドラー
func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
