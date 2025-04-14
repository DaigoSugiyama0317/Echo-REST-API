package controller

import (
	"net/http"
	"strconv"

	"github.com/DaigoSugiyama0317/Echo-REST-API/model"
	"github.com/DaigoSugiyama0317/Echo-REST-API/usecase"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// ITaskController は、タスクに関連する操作を定義したインターフェース
type ITaskController interface {
	GetAllTasks(c echo.Context) error // すべてのタスクを取得
	GetTaskById(c echo.Context) error // IDによるタスクの取得
	CreateTask(c echo.Context) error  // タスクの作成
	UpdateTask(c echo.Context) error  // タスクの更新
	DeleteTask(c echo.Context) error  // タスクの削除
}

// タスクに関連する操作を実装する構造体
type taskController struct {
	tu usecase.ITaskUsecase
}

// コンストラクタ関数
func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu} // ユースケースのインターフェース
}

// ログインしているユーザーのタスクをすべて取得
func (tc taskController) GetAllTasks(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// ユーザーIDを基にタスクを取得
	taskRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		// エラーがあれば内部サーバーエラーを返す
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes) // 成功した場合、タスクのリストを返す
}

// 指定されたIDのタスクを取得
func (tc taskController) GetTaskById(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// パスパラメータからタスクIDを取得し、整数に変換
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// ユーザーIDとタスクIDを基にタスクを取得
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // エラーがあれば内部サーバーエラーを返す
	}
	return c.JSON(http.StatusOK, taskRes) // 成功した場合、タスク情報を返す
}

// 新しいタスクを作成
func (tc taskController) CreateTask(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// リクエストボディからタスク情報をバインド
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error()) // バインディングエラーがあれば、400 Bad Requestを返す
	}

	// ユーザーIDをタスクに設定
	task.UserId = uint(userId.(float64))
	// タスク作成処理を呼び出し
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // エラーがあれば内部サーバーエラーを返す
	}
	return c.JSON(http.StatusCreated, taskRes) // 成功した場合、作成したタスクを返す
}

// 指定されたIDのタスクを更新
func (tc taskController) UpdateTask(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// パスパラメータからタスクIDを取得し、整数に変換
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// リクエストボディからタスク情報をバインド
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error()) // バインディングエラーがあれば、400 Bad Requestを返す
	}

	// タスク更新処理を呼び出し
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // エラーがあれば内部サーバーエラーを返す
	}
	return c.JSON(http.StatusOK, taskRes) // 成功した場合、更新したタスク情報を返す
}

// 指定されたIDのタスクを削除
func (tc taskController) DeleteTask(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// パスパラメータからタスクIDを取得し、整数に変換
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// タスク削除処理を呼び出し
	err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // エラーがあれば内部サーバーエラーを返す
	}
	return c.NoContent(http.StatusNoContent) // 成功した場合、No Contentレスポンスを返す
}
