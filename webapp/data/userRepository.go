package data

import (
	"gopkg.in/mgo.v2"

	"webapp/models"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"webapp/common"
	"net/http"
)

type UserRepository struct {
	C *mgo.Collection
}


func(u *UserRepository)CreateUser(user *models.User) error{

	exist,err:=u.C.Find(bson.M{"username":user.Username}).Count()
	if err!=nil{
		return common.ServerSideError{Code:http.StatusInternalServerError}
	}
	if exist!=0{
		return common.UserExistError{Code:http.StatusNotAcceptable}
	}
	user.Id=bson.NewObjectId()
	hpass,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err!=nil{
		panic(err)
	}
	user.HashPassword=hpass
	user.Password=""

	err = u.C.Insert(&user)
	return common.ServerSideError{Code:http.StatusInternalServerError}
}

//login process.
func(u *UserRepository) Login(user models.User)(loginUser models.User,err error){
	err=u.C.Find(bson.M{"username":user.Username}).One(&loginUser)
	if err!= nil{
		return loginUser,common.UserNotExistError{Code:http.StatusInternalServerError}
	}
	err = bcrypt.CompareHashAndPassword(loginUser.HashPassword,[]byte(user.Password))
	if err!=nil{
		return loginUser,common.AuthenticationError{Code:http.StatusForbidden}
	}
	return
}

