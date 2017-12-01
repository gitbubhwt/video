package db

import (
	"video/common"

	"fmt"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

/**单列模式**/
func GetMysqlDb() error {
	if engine == nil {
		systemPath := GetValue(common.SYSTEM_CONFIG_KEY, common.SYSTEM_CONFIG_ROOT_PATH)
		dbConfPath := fmt.Sprintf(common.CONF_PATH, systemPath, common.DB_CONF_NAME)
		newconfig, err := config.NewConfig("ini", dbConfPath)
		userName := newconfig.String("mysqluser")
		password := newconfig.String("mysqlpass")
		dbName := newconfig.String("mysqldb")
		ip := newconfig.String("mysqlip")
		port := newconfig.String("mysqlport")

		dbInfo := userName + ":" + password + "@(" + ip + ":" + port + ")" + "/" + dbName
		engine, err = xorm.NewEngine("mysql", dbInfo+"?charset=utf8")
		if err != nil {
			return err
		}
		engine.ShowSQL(true)
		engine.Sync2(new(common.Video))
		engine.Sync2(new(common.VideoPath))
	}
	return nil
}

func GetMysql() *xorm.Engine {
	if engine == nil {
		GetMysqlDb()
	}
	return engine
}
