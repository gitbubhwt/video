package db

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/mgo.v2"
	"reflect"
	"strconv"
	"strings"
	"time"
	. "video/common"
	log "video/logger"
)

var (
	reidsDB *redis.Client
	mysqlDB *xorm.Engine
	monDB   *mgo.Database
	conf    *SystemConf
)

type SystemConf struct {
	MysqlUser string "mysqlUser"
	MysqlPwd  string "mysqlPwd"
	MysqlIp   string "mysqlIp"
	MysqlPort string "mysqlPort"
	MysqlDb   string "mysqlDb"

	MonIp   string "monIp"
	MonPort string "monPort"
	MonDb   string "monDb"

	RedisAddr string "redisAddr"
	RedisPwd  string "redisPwd"
	RedisDb   string "redisDb"
}

func init() {
	if conf == nil {
		path := GetCurrentDirectory()
		index := strings.Index(path, SERVER_NAME)
		if index != -1 {
			bytes := []byte(path)
			path = string(bytes[:index+len([]byte(SERVER_NAME))])
		}
		dbConf := fmt.Sprintf(CONF_PATH, path, CONF_NAME)
		log.Info(dbConf)
		newconfig, err := config.NewConfig("ini", dbConf)
		if err != nil {
			log.Error("init config fail,err:", err)
			return
		}
		conf = new(SystemConf)
		v := reflect.ValueOf(conf).Elem()
		for i := 0; i < v.NumField(); i++ {
			fieldInfo := v.Type().Field(i)
			tag := fmt.Sprintf("%v", fieldInfo.Tag)
			if value := newconfig.String(tag); value != STRING_NULL {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			}
		}
	}
}

//初始化redis
func InitRedis() error {
	if reidsDB == nil {
		redisDb, _ := strconv.Atoi(conf.RedisDb)
		reidsDB = redis.NewClient(&redis.Options{
			Addr:     conf.RedisAddr,
			Password: conf.RedisPwd,
			DB:       redisDb,
		})
	}
	_, err := reidsDB.Ping().Result()
	if err != nil {
		log.Error("fail to connect redis, error:", err)
		return err
	}
	log.Info("succeed connect to redis")
	return nil
}

//获取redis 客户端
func GetRedisClient() *redis.Client {
	if reidsDB == nil {
		InitRedis()
	}
	return reidsDB
}

//初始化mysql
func InitMysql() error {
	if mysqlDB == nil {
		dbInfo := conf.MysqlUser + ":" + conf.MysqlPwd + "@(" + conf.MysqlIp + ":" + conf.MysqlPort + ")" + "/" + conf.MysqlDb
		engine, err := xorm.NewEngine("mysql", dbInfo+"?charset=utf8")
		if err != nil {
			log.Error("fail to connect mysql,err:", err)
			return err
		}
		engine.ShowSQL(true)
		engine.Sync2(new(Video))
		engine.Sync2(new(VideoPath))
		mysqlDB = engine
	}
	log.Info("succeed connect to mysql")
	return nil
}

//获取mysql客户端
func GetMysql() *xorm.Engine {
	if mysqlDB == nil {
		InitMysql()
	}
	return mysqlDB
}

//初始化mongo
func InitMongo() error {
	if monDB == nil {
		connstr := "mongodb://" + conf.MonIp + ":" + conf.MonPort
		session, err := mgo.DialWithTimeout(connstr, time.Second*5)
		if err != nil {
			log.Error("fail to connect mongo,err:", err)
			return err
		}
		session.SetMode(mgo.Monotonic, true)
		monDB = session.DB(conf.MonDb)
	}
	log.Info("succeed connect to mongo")
	return nil
}

//获取mongo客户端
func GetMongo() *mgo.Database {
	if monDB == nil {
		InitMongo()
	}
	return monDB
}
