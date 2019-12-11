package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Cswapi/Web-Server/database/database"
	jwt "github.com/dgrijalva/jwt-go"
)

const secret = "Cswapi"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

// 注册处理器
func registerHandler(res http.ResponseWriter, req *http.Request) {
	var user User
	// 请求解析
	err := req.ParseForm()
	// 判断参数格式是否正确
	if err != nil && req.PostForm["username"] != nil && req.PostForm["password"] != nil {
		res.WriteHeader(http.StatusForbidden)
		fmt.Println("Register failed!")
		res.Write([]byte("Parameters format is wrong!\n"))
		return
	}

	user.Username = req.PostForm["username"][0]
	user.Password = req.PostForm["password"][0]
	// 调用database中的CheckKey函数查看新建用户名是否存在
	if database.CheckKey([]byte("users"), []byte(user.Username)) {
		res.WriteHeader(http.StatusForbidden)
		fmt.Println("Register failed!")
		res.Write([]byte("Username already exists.\n"))
		return
	}
	// 更新数据库
	database.Update([]byte("users"), []byte(user.Username), []byte(user.Password))
	// 将处理结果写到http.Header中
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("User created successfully!\n"))
}

// 登陆处理器：jwt产生token对用户进行认证
func loginHandler(res http.ResponseWriter, req *http.Request) {
	var user User
	err := req.ParseForm()
	// 参数格式解析
	if err != nil && req.PostForm["username"] != nil && req.PostForm["password"] != nil {
		res.WriteHeader(http.StatusForbidden)
		fmt.Println("Login failed!")
		res.Write([]byte("Parameters format is wrong.\n"))
		return
	}

	user.Username = req.PostForm["username"][0]
	user.Password = req.PostForm["password"][0]
	// 调用database中的CheckKey函数查看用户名和密码是否匹配
	if !database.CheckKey([]byte("users"), []byte(user.Username)) || user.Password != database.GetValue([]byte("users"), []byte(user.Username)) {
		res.WriteHeader(http.StatusForbidden)
		fmt.Println("Login failed!")
		res.Write([]byte("The user does not exist.\n"))
		return
	}
	// 生成Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(res, "Token is not valid!")
		log.Fatal(err)
	}
	// 以Token的形式生成response
	response := Token{tokenStr}
	writeResponse(response, res)
}

// 生成JSON格式的response
func writeResponse(response interface{}, res http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write(json)
}
