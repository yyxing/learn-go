package starter

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"learn-go/web/core"
	"learn-go/web/core/context"
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
	for key := range dbMap {
		return GetDB(key)
	}
	return nil
}

func GetDB(datasourceName string) *gorm.DB {
	if len(dbMap) < 1 {
		panic("least one datasource must be configured ")
	}
	return dbMap[datasourceName]
}
func Transaction(tx *gorm.DB, worker func() error) error {
	err := worker()
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// mysql连接配置
func (starter *DatasourceStarter) Init(context context.ApplicationContext) {
	config := GetConfig()
	starter.datasourceAssembly(config)
}

// 正式建立连接
func (starter *DatasourceStarter) Start(context context.ApplicationContext) {
	dbMap = make(map[string]*gorm.DB)
	config := GetConfig()
	loggerConfig := logger.Config{}
	for i, datasource := range starter.datasourceList {
		//driverName := datasource.Driver
		driverUrl := fmt.Sprintf(core.MysqlDriverFormatter, datasource.Username, datasource.Password, datasource.Url)
		if config.GetBool("logger.showSql") {
			loggerConfig.LogLevel = logger.Info
		}
		db, err := gorm.Open(mysql.Open(driverUrl), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.New(logrus.StandardLogger(), loggerConfig),
		})
		if err != nil {
			panic(err)
		}
		logrus.Infof("datasource %s：%s connect success", datasource.DatasourceName, datasource.Url)
		if len(datasource.DatasourceName) == 0 {
			if i == 0 {
				dbMap[defaultDatasourceName] = db
			} else {
				dbMap[defaultDatasourceName+"-"+string(rune(i))] = db
			}
			continue
		}
		dbMap[datasource.DatasourceName] = db
	}
	context.Set("dbMap", dbMap)
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
func (starter *DatasourceStarter) Finalize(context context.ApplicationContext) {
	for _, db := range dbMap {
		s, _ := db.DB()
		s.Close()
	}
}

func (starter *DatasourceStarter) GetOrder() int {
	return core.Int32Min + 2
}
