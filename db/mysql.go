package db

import (
	"errors"
	"video/common"

	"fmt"
	log "video/logger"

	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

/**单列模式**/
func GetMysqlDb() error {
	if engine == nil {
		keys := []string{"mysql_user", "mysql_pass", "mysql_db", "mysql_ip", "mysql_port"}
		mp := ReadConf(keys)
		if mp == nil {
			log.Error("Read mongo conf fail,err:value is nil")
			return errors.New("Read mongo conf fail,err:value is nil")
		}
		userName := mp[keys[0]]
		password := mp[keys[1]]
		dbName := mp[keys[2]]
		ip := mp[keys[3]]
		port := mp[keys[4]]

		dbInfo := userName + ":" + password + "@(" + ip + ":" + port + ")" + "/" + dbName
		engine, err := xorm.NewEngine("mysql", dbInfo+"?charset=utf8")
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

//读取配置文件信息
func ReadConf(keys []string) map[string]string {
	systemPath := GetValue(common.SYSTEM_CONFIG_KEY, common.SYSTEM_CONFIG_ROOT_PATH)
	dbConfPath := fmt.Sprintf(common.CONF_PATH, systemPath, common.DB_CONF_NAME)
	newconfig, err := config.NewConfig("ini", dbConfPath)
	if err != nil {
		log.Error("Read conf fail,err:", err)
		return nil
	}
	mp := make(map[string]string)
	for i := 0; i < len(keys); i++ {
		mp[keys[i]] = newconfig.String(keys[i])
	}
	return mp
}
