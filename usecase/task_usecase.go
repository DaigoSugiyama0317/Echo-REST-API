package usecase

import (
	"github.com/DaigoSugiyama0317/Echo-REST-API/model"
	"github.com/DaigoSugiyama0317/Echo-REST-API/repository"
	"github.com/DaigoSugiyama0317/Echo-REST-API/validator"
)

// タスクに関連するユースケース（ビジネスロジック）を定義
type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)                            //ユーザーIDに基づいて全タスクを取得
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)                 //特定のタスクIDに基づいてタスクを取得
	CreateTask(task model.Task) (model.TaskResponse, error)                           //新しいタスクを作成
	UpdateTask(task model.Task, UserId uint, taskId uint) (model.TaskResponse, error) //既存のタスクを更新
	DeleteTask(userId uint, taskId uint) error                                        //タスクを削除
}

// taskUsecase 構造体は ITaskUsecase インターフェースを実装
type taskUsecase struct {
	tr repository.ITaskRepository //タスクに関するリポジトリ
	tv validator.ITaskValidator //タスクに関するバリデーション
}

//コンストラクタ関数
func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

//ユーザーIDに基づいてすべてのタスクを取得
func (tu taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	// リポジトリからタスクを取得
	if err := tu.tr.GerAllTasks(&tasks, userId); err != nil {
		return nil, err
	}

	// タスクをレスポンス形式に変換
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

//特定のタスクをIDで取得
func (tu taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	// リポジトリからタスクを取得
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}

	// タスクをレスポンス形式に変換
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

//新しいタスクを作成
func (tu taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	// タスクのバリデーション
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}

	// リポジトリでタスクを作成
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}

	// 作成されたタスクをレスポンス形式に変換
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

//既存のタスクを更新
func (tu taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	// タスクのバリデーション
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}

	// リポジトリでタスクを更新
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}

	// 更新されたタスクをレスポンス形式に変換
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

//タスクを削除
func (tu taskUsecase) DeleteTask(userId uint, taskId uint) error {
	// リポジトリでタスクを削除
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
