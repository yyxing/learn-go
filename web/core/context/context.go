package context

import (
	"learn-go/web/core"
	"log"
	"reflect"
	"sort"
)

const (
	packageSeparator = "/"
)

// 上下文资源管理器
type ApplicationContext struct {
	context map[string]interface{}
	starterRegister
}

// 声明starter注册器
type starterRegister struct {
	starters []core.Starter
}

func (register starterRegister) Len() int { return len(register.starters) }
func (register starterRegister) Swap(i, j int) {
	starters := register.starters
	starters[i], starters[j] = starters[j], starters[i]
}
func (register starterRegister) Less(i, j int) bool {
	return register.starters[i].GetOrder() < register.starters[j].GetOrder()
}

// 注册starter到上下文中
func (context *ApplicationContext) Register(starter core.Starter) {
	if context.context == nil {
		context.context = make(map[string]interface{})
	}
	starterType := reflect.TypeOf(starter)
	starterPackageName := starterType.PkgPath()
	starterName := starterPackageName + packageSeparator + starterType.Name()
	context.context[starterName] = starter
	context.register(starter)
}

// 注册starter到上下文中
func (context *ApplicationContext) SortStarter() {
	sort.Sort(context.starterRegister)
}

// 注册器注册
func (register *starterRegister) register(starter core.Starter) {
	register.starters = append(register.starters, starter)
}

func (context ApplicationContext) GetAllStarters() []core.Starter {
	return context.starters
}

func (context *ApplicationContext) Set(key string, value interface{}) {
	context.context[key] = value
}

func (context *ApplicationContext) Get(key string) interface{} {
	value, ok := context.context[key]
	if !ok {
		log.Printf("%v load failed\n", value)
	}
	return value
}

func (context *ApplicationContext) Remove(key string) {
	delete(context.context, key)
}
