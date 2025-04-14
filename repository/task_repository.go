package repository

import (
	"fmt"

	"github.com/DaigoSugiyama0317/Echo-REST-API/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// タスクに関するデータベース操作を定義
type ITaskRepository interface {
	GerAllTasks(tasks *[]model.Task, userId uint) error           //ユーザーIDに基づいてすべてのタスクを取得
	GetTaskById(task *model.Task, userId uint, taskId uint) error //特定のタスクIDに基づいてタスクを取得
	CreateTask(task *model.Task) error                            // 新しいタスクをデータベースに作成
	UpdateTask(task *model.Task, userId uint, taskId uint) error  //既存のタスクを更新
	DeleteTask(userId uint, taskId uint) error                    //特定のタスクを削除
}

// データベース操作を実行するためのリポジトリ
type taskRepository struct {
	db *gorm.DB
}

// コンストラクタ関数
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

// ユーザーIDに基づいてすべてのタスクを取得
func (tr *taskRepository) GerAllTasks(tasks *[]model.Task, userId uint) error {
	// ユーザーIDでフィルタリングし、タスクを並べ替えて取得
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// 特定のタスクIDに基づいてタスクを取得
func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	// ユーザーIDとタスクIDでタスクを取得
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

// 新しいタスクをデータベースに作成
func (tr *taskRepository) CreateTask(task *model.Task) error {
	// タスクをデータベースに作成
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// 既存のタスクを更新
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	// タスクIDとユーザーIDで指定されたタスクを更新
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	// 更新結果のエラーチェック
	if result.Error != nil {
		return result.Error
	}
	// 更新された行数が0の場合、タスクが存在しないとみなす
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

// 特定のタスクを削除
func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	// タスクIDとユーザーIDで指定されたタスクを削除
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	// 削除結果のエラーチェック
	if result.Error != nil {
		return result.Error
	}
	// 削除された行数が0の場合、タスクが存在しないとみなす
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
