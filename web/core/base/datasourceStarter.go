package base

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"learn-go/web/core"
)

var (
	dbMap                 map[string]*gorm.DB
	defaultDatasourceName = "master"
)

type DatasourceStarter struct {
	AbstractStarter
	datasourceList []datasource
}

type datasource struct {
	Url            string
	Driver         string
	Username       string
	Password       string
	DatasourceName string
}

// 获取默认的DB 若配置了多个 则选择第一个
func DefaultDB() *gorm.DB {
	return GetDB(defaultDatasourceName)
}

func GetDB(datasourceName string) *gorm.DB {
	if len(dbMap) < 1 {
		panic("least one datasource must be configured ")
	}
	return dbMap[datasourceName]
}

// mysql连接配置
func (starter *DatasourceStarter) Init(context core.ApplicationContext) {
	config, ok := context.Get(GlobalConfigKey).(viper.Viper)
	if !ok {
		panic("database config load failed")
	}
	starter.datasourceAssembly(config)
}

// 正式建立连接
func (starter *DatasourceStarter) Start(context core.ApplicationContext) {
	dbMap = make(map[string]*gorm.DB)
	for i, datasource := range starter.datasourceList {
		driverName := datasource.Driver
		driverUrl := fmt.Sprintf(core.MysqlDriverFormatter, datasource.Username, datasource.Password, datasource.Url)
		db, err := gorm.Open(driverName, driverUrl)
		if err != nil {
			panic(err)
		}
		if len(datasource.DatasourceName) == 0 {
			if i == 0 {
				dbMap[defaultDatasourceName] = db
			} else {
				dbMap[defaultDatasourceName+"-"+string(i)] = db
			}
			continue
		}
		dbMap[datasource.DatasourceName] = db
	}
}

// 将配置文件
func (starter *DatasourceStarter) datasourceAssembly(config viper.Viper) {
	var datasource []datasource
	err := config.UnmarshalKey("datasource", &datasource)
	if err != nil {
		panic("datasource config wrong format")
	}
	starter.datasourceList = datasource
}

// 关闭db
func (starter *DatasourceStarter) Finalize(context core.ApplicationContext) {
	for _, db := range dbMap {
		db.Close()
	}
}

func (starter *DatasourceStarter) GetOrder() int {
	return -1
}
