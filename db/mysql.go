package db

import (
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"video/common"
)

var engine *xorm.Engine

/**单列模式**/
func GetOneDb() (*xorm.Engine, error) {
	var err error
	if engine == nil {
		newconfig, err := config.NewConfig("ini", "conf/db.conf")
		userName := newconfig.String("mysqluser")
		password := newconfig.String("mysqlpass")
		dbName := newconfig.String("mysqldb")
		ip := newconfig.String("mysqlip")
		port := newconfig.String("mysqlport")

		dbInfo := userName + ":" + password + "@(" + ip + ":" + port + ")" + "/" + dbName
		engine, err = xorm.NewEngine("mysql", dbInfo+"?charset=utf8")
		if err != nil {
			return engine, err
		}
		engine.ShowSQL(true)
		engine.Sync2(new(common.Video))
	}
	return engine, err
}
