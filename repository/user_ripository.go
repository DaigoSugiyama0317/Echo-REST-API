package repository

import (
	"github.com/DaigoSugiyama0317/Echo-REST-API/model"
	"gorm.io/gorm"
)

// ユーザー関連のDB操作のインターフェース
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error // メールアドレスからユーザーを取得
	CreateUser(user *model.User) error                   // ユーザーの新規作成
}

// リポジトリの構造体（GORMのDB接続を保持）
type userRepository struct {
	db *gorm.DB
}

// コンストラクタ関数
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// メールアドレスからユーザーを取得
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	//emailから該当ユーザーを取得
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザーをDBに新規作成
func (ur *userRepository) CreateUser(user *model.User) error {
	//user作成
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
