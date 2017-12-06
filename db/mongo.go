package db

import (
	"fmt"
	"os"
	"strings"
	"time"
	log "video/logger"

	"gopkg.in/mgo.v2"
)

var ms *mgo.Session

func GetMs() *mgo.Session {
	if ms == nil {
		InitMgo()
	}
	return ms
}

func InitMgo() {
	//mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]
	//mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
	keys := []string{"mongo_hostsports", "mongo_dbname", "mongo_userpass", "mongo_replicaset"}
	mp := ReadConf(keys)
	if mp == nil {
		log.Error("Read mysql conf fail,err:value is nil")
		return
	}
	userpass := mp[keys[0]]
	hostsports := mp[keys[1]]
	role := mp[keys[2]]
	replicaset := mp[keys[3]]
	connstr := "mongodb://"
	if !strings.EqualFold("", userpass) {
		connstr += userpass + "@"
	}
	connstr += fmt.Sprintf("%s/%s", hostsports, role)
	if !strings.EqualFold("", replicaset) {
		connstr += "?replicaSet=" + replicaset
	}
	session, err := mgo.DialWithTimeout(connstr, time.Second*5)
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	err = session.Ping()
	if err != nil {
		log.Error("fail to connect to mongo,err:%v", err)
		os.Exit(1)
	} else {
		log.Info("succeed connect to mongo")
	}
	ms = session
}
