package validator

import (
	"github.com/DaigoSugiyama0317/Echo-REST-API/model"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type TaskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &TaskValidator{}
}

func (tv *TaskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task, // タイトルの検証
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),            // タイトルは必須項目
			validation.RuneLength(1, 10).Error("limited max 10 char"), // タイトルの長さは1文字以上10文字以下である必要がある
		),
	)
}
