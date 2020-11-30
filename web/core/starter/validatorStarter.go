package starter

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translator "github.com/go-playground/validator/v10/translations/zh"
	"github.com/sirupsen/logrus"
	"learn-go/web/core"
	"learn-go/web/core/context"
)

type ValidatorStarter struct {
	AbstractStarter
}

func (starter *ValidatorStarter) Init(context context.ApplicationContext) {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := translator.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		logrus.Error("validator" + err.Error())
	}
	core.Validate = validate
	context.Set("validate", validate)
}

func (starter *ValidatorStarter) GetOrder() int {
	return 0
}
