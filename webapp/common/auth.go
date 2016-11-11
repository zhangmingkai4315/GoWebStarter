package common

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"github.com/dgrijalva/jwt-go/request"
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

func GetVerifyKey() *rsa.PublicKey{
	return verifyKey;
}

func initKeys(){
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
	log.Println("Loading private and public RSA keys done.")

}

func CreateToken(user string) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = &CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		"level1",
		CustomerInfo{user, "member"},
	}
	return t.SignedString(signKey)
}

func AdminMiddleware(w http.ResponseWriter,r *http.Request,next http.HandlerFunc){
	var tokenString string
	var err error
	var token *jwt.Token
	for _,cookie:=range r.Cookies(){
		if cookie.Name=="token"{
			tokenString=cookie.Value;
		}
	}

	if tokenString =="" {
		//parse with authorized header
		token, err = request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return verifyKey,nil
		})
	}else{
		token, err = jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
	}

	if err!=nil{
		switch err.(type) {
		case *jwt.ValidationError:
			vErr:=err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				DisplayAppError(w,err,"Token Expired",403)
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				DisplayAppError(w,err,"Error while parsing token!",500)
				return
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			DisplayAppError(w,err,"Error while parsing token!",500)
			return
		}
	}
	if token.Valid{
		next(w,r)
	}else{
		//response:=Response{"Invalid token"}
		DisplayAppError(w,nil,"Invalid token",403)
	}
}

