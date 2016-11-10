package controllers

import (
	"net/http"

	"views"
	"models"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/dgrijalva/jwt-go/request"
	"fmt"
	"crypto/rsa"
)
const (
	privKeyPath = "keys/app.rsa"    // openssl genrsa -out app.rsa 1024
	pubKeyPath = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

type Response struct {
	Text string `json:"text"`
}
type Token struct {
	Token string `json:"token"`
}


// Define some custom types were going to use within our tokens
type CustomerInfo struct {
	Name string
	Role string
}

type CustomClaims struct {
	*jwt.StandardClaims
	TokenType string
	CustomerInfo
}

var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
)

func init(){
	//var err error
	signBytes,err:= ioutil.ReadFile(privKeyPath)
	if err!=nil{
		log.Fatal("Error reading public key")
		return
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err!=nil{
		log.Fatal("Error reading private key")
		return
	}
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err!=nil{
		log.Fatal("Error reading public key")
		return
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err!=nil{
		log.Fatal("Error reading public key")
		return
	}

}

func jsonResponse(response interface{},w http.ResponseWriter){
	json,err:=json.Marshal(response)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	w.Write(json)
}
// get /login handler
func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "login", "base", struct{}{})
}



func createToken(user string) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims = &CustomClaims{
		&jwt.StandardClaims{
			// set the expire time
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"level1",
		CustomerInfo{user, "member"},
	}
	log.Println(signKey)
	// Creat token string
	return t.SignedString(signKey)
}



// post /login handler
func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "post login")
	var user models.User
	r.ParseForm()
	user.Username=r.PostFormValue("username")
	user.Password=r.PostFormValue("password")
	user.CreatedAt=time.Now()
	//fmt.Fprintf(w,"%v",user)
	if user.Username!="mike" && user.Password!="123456"{
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w,"Authenticate Fail")
		return
	}

	tokenString,err:=createToken(user.Username)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"Error While Signing Token!")
		log.Printf("Token signing error:%v\n",err)
		return
	}
	response:=Token{tokenString}
	jsonResponse(response,w)
}


func RestictHandler(w http.ResponseWriter,r *http.Request){
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return  verifyKey,nil
	})
	if err!=nil{
		switch err.(type) {
		case *jwt.ValidationError:
			vErr:=err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w,"Token Expired")
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w,"Error while parsing token!")
				log.Printf("ValidationError:%+v\n",vErr.Errors)
				return
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w,"Error while parsing token!")
			log.Printf("ValidationError:%+v\n",err)
			return
		}
	}
	if token.Valid{
		response:=Response{"Authorized"}
		jsonResponse(response,w)
	}else{
		response:=Response{"Invalid token"}
		jsonResponse(response,w)
	}
}





// get /SignUp handler
func GetSignUpHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "signup", "base", struct{}{})
}

// post /SignUp handler
func PostSignUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post sing up")
}
