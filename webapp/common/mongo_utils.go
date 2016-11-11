package common

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

var session *mgo.Session

func GetSession() *mgo.Session{
	if session==nil{
		var err error
		session,err=mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:[]string{AppConfig.Database.Host},
			Username:AppConfig.Database.User,
			Password:AppConfig.Database.Password,
			Timeout:60*time.Second,
		})
		if err!=nil{
			log.Fatal("Session created fail:%s\n",err)
		}

	}
	return session
}

func createDbSession(){
	var err error
	session,err=mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:[]string{AppConfig.Database.Host},
		Username:AppConfig.Database.User,
		Password:AppConfig.Database.Password,
		Timeout:60*time.Second,
	})
	if err!=nil{
		log.Fatal("Session created fail:%s\n",err)
	}
	log.Println("Create mongo session pool done.")
}