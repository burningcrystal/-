package middleware

import (
	"calloflife/utils"
	"calloflife/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte(utils.JwtKey)

type myClaims struct {
	UserName string 	`json:"username"`
	PassWord string	`json:"password"`
	jwt.StandardClaims //创建时间、创建人等一些必要的信息

}
//认证的步骤：
//生成token
func SetToken(username string,password string)(string,int) {
	expireTime := time.Now().Add(10*time.Hour)
	setClaim := myClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),  //加一个时间戳，关于time的用法还得好好学学
			Issuer:"am",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodES256,setClaim)  //签发jwt
	token,err := reqClaim.SignedString(jwtKey)
	if err != nil {
		return "",errmsg.ERROR
	}
	log.Println("token:",token)
	return token,errmsg.SUCCESS
}
//验证token
func CheckToken(token string)(*myClaims,int)  {
	settoken,_ := jwt.ParseWithClaims(token,&myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey,nil
	})
	if key,ok := settoken.Claims.(*myClaims);ok&&settoken.Valid {
		return key,errmsg.SUCCESS
	}
	return nil,errmsg.ERROR
}
//jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := errmsg.SUCCESS
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
		}

		checkToken := strings.SplitN(tokenHeader," ",2)
		if len(checkToken) != 2 && checkToken[0]!="Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.Abort()
		}
		key,Tcode := CheckToken(checkToken[1])
		if Tcode == errmsg.ERROR {
			code =errmsg.ERROR_TOKEN_WRONG
		}
		if time.Now().Unix()>key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_LONG_TIME
			c.Abort()
		}
		c.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":errmsg.GetErrMsg(code),
		})
		c.Set("username",key.UserName)
		c.Next()
	}
}