package usecase

import (
	"os"
	"time"

	"github.com/DaigoSugiyama0317/Echo-REST-API/model"
	"github.com/DaigoSugiyama0317/Echo-REST-API/repository"
	"github.com/DaigoSugiyama0317/Echo-REST-API/validator"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// ユーザーに関するユースケースのインターフェース定義
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error) // 新規ユーザー登録
	LogIn(user model.User) (string, error)              // ログインしJWTトークンを返す
}

// ユースケースの構造体（リポジトリとバリデータへの依存を持つ）
type userUsecase struct {
	ur repository.IUserRepository // データアクセス
	uv validator.IUserValidator   // 入力バリデーション
}

// ユースケースのコンストラクタ関数
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

// ユーザー登録処理
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	// 入力バリデーション
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	// パスワードをハッシュ化（bcrypt使用）
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	// ハッシュ化済みのユーザー情報を作成
	newUser := model.User{Email: user.Email, Password: string(hash)}
	// DBへ新規ユーザー登録
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	// レスポンス用に必要な情報のみ返す
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

// ログイン処理（トークン発行）
func (uu *userUsecase) LogIn(user model.User) (string, error) {
	// 入力バリデーション
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	// 入力されたメールアドレスでDBからユーザー取得
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// パスワードの照合（bcryptのハッシュ比較）
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// JWTトークン生成（有効期限12時間）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	// 環境変数にあるシークレットキーで署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	// トークン文字列を返す
	return tokenString, nil
}
