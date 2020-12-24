package starter

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translator "github.com/go-playground/validator/v10/translations/zh"
	"github.com/sirupsen/logrus"
	"learn-go/web/core/context"
	"strings"
)

var defaultValidate *validator.Validate
var defaultTranslator ut.Translator

type ValidatorStarter struct {
	AbstractStarter
}

// 参数校验异常
type ConstraintViolationError struct {
	s string
}

func (e ConstraintViolationError) Error() string {
	return e.s
}
func GetDefaultValidator() *validator.Validate {
	return defaultValidate
}

func GetTranslator() ut.Translator {
	return defaultTranslator
}

// 参数校验
func ParamValidate(param interface{}) error {
	paramValidator := GetDefaultValidator()
	trans := GetTranslator()
	err := paramValidator.Struct(param)
	errMsg := strings.Builder{}
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, fieldError := range errs {
				msg := fieldError.Translate(trans)
				errMsg.WriteString(msg)
			}
		}
		return ConstraintViolationError{s: errMsg.String()}
	}
	return nil
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
	defaultValidate = validate
	defaultTranslator = trans
	context.Set("validate", validate)
}

func (starter *ValidatorStarter) GetOrder() int {
	return 0
}
