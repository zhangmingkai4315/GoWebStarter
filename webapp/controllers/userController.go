package controllers

import (
	"net/http"

	"webapp/views"
	"webapp/models"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"github.com/dgrijalva/jwt-go/request"
	"webapp/common"
	"webapp/data"
)
const (
	privKeyPath = "keys/app.rsa"    // openssl genrsa -out app.rsa 1024
	pubKeyPath = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// get /logout handler
func GetLogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w,&http.Cookie{Name:"token",Value:"",Expires:time.Now().Add(time.Hour*24)})
	http.Redirect(w,r,"/login",http.StatusFound)
}

// get /login handler
func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "login", "base", struct{}{})
}




// post /login handler
func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	r.ParseForm()
	user.Username=r.PostFormValue("username")
	user.Password=r.PostFormValue("password")
	user.CreatedAt=time.Now()
	//fmt.Fprintf(w,"%v",user)

	context:=NewContext()
	defer context.Close()
	c:=context.DbCollection("users")
	repo:=&data.UserRepository{c}
	loginUser,err:=repo.Login(user)
	if err!=nil{
		common.DisplayAppError(w,err,"Login Fail",http.StatusForbidden)
		return
	}

	tokenString,err:=common.CreateToken(loginUser.Username)
	if err!=nil{
		//w.WriteHeader(http.StatusInternalServerError)
		//fmt.Fprintf(w,"Error While Signing Token!")
		http.Redirect(w,r,"/login",http.StatusInternalServerError)
		return
	}
	http.SetCookie(w,&http.Cookie{Name:"token",Value:tokenString,Expires:time.Now().Add(time.Hour*24)})
	response:=common.Token{tokenString}
	common.DisplayAppData(w,response)
	return
}


func RestictHandler(w http.ResponseWriter,r *http.Request){
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &common.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return  common.GetVerifyKey(),nil
	})
	if err!=nil{
		switch err.(type) {
		case *jwt.ValidationError:
			vErr:=err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:

				common.DisplayAppError(w,err,"Token Expired",http.StatusUnauthorized)
				return
			default:

				common.DisplayAppError(w,err,"Error while parsing token",http.StatusInternalServerError)
				return
			}
		default:
			common.DisplayAppError(w,err,"Error while parsing token",http.StatusInternalServerError)
			log.Printf("ValidationError:%+v\n",err)
			return
		}
	}
	if token.Valid{
		response:=common.Response{"Authorized"}
		common.DisplayAppData(w,response)

	}else{
		common.DisplayAppError(w,nil,"Invalid token",http.StatusForbidden)
	}
}

// get /SignUp handler
func GetSignUpHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "signup", "base", struct{}{})
}

// post /SignUp handler
func PostSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var useRegister models.RegisterUser
	r.ParseForm()
	useRegister.Username=r.PostFormValue("username")
	useRegister.Password=r.PostFormValue("password")
	useRegister.Password=r.PostFormValue("password2")
	useRegister.Email=r.PostFormValue("email")
	//You should do server side validation!

	//if validate==false{
	//	log.Println(err)
	//	common.DisplayAppError(w,err,"Post user data error",http.StatusBadRequest)
	//	return
	//}




	user:=models.User{Username:useRegister.Username,Password:useRegister.Password,Email:useRegister.Email,CreatedAt:time.Now()}
	context:=NewContext()
	defer context.Close()
	c:=context.DbCollection("users")
	repo:=&data.UserRepository{c}
	err:=repo.CreateUser(&user)
	if err!=nil{
		switch v:=err.(type) {
		case common.AuthenticationError:
			common.DisplayAppError(w,err,err.Error(),v.Code)
			break
		case common.UserExistError:
			common.DisplayAppError(w,err,err.Error(),v.Code)
			break
		case common.ServerSideError:
			common.DisplayAppError(w,err,err.Error(),v.Code)
			break
		default:
			common.DisplayAppError(w,err,err.Error(),http.StatusInternalServerError)
		}
		return
	}
	user.HashPassword=nil
	common.DisplayAppData(w,user)
	return
}
