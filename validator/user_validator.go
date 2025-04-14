package validator

import (
	"github.com/DaigoSugiyama0317/Echo-REST-API/model"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// ユーザー入力の検証に必要なメソッドを定義するインターフェース
type IUserValidator interface {
	UserValidate(user model.User) error // ユーザーのバリデーションを実行するメソッド
}

// IUserValidator インターフェースを実装する構造体
type userValidator struct{}

// コンストラクタ関数
func NewUserValidator() IUserValidator {
	return &userValidator{}
}

// ユーザーのフィールドをバリデーションするメソッド
// Ozzoバリデーションライブラリを使用して、Email と Password フィールドを検証
func (tv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user, // 構造体のバリデーションを実行
		validation.Field(
			&user.Email, // Email フィールドを検証
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"), // 1～30文字の範囲で制限
			is.Email.Error("is not valid email format"),               // 有効なメール形式かどうかチェック
		),
		validation.Field(
			&user.Password, // Password フィールドを検証
			validation.Required.Error("password is required"),               // 必須チェック
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"), // 6～30文字の範囲で制限
		),
	)
}
