package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{
	UserId int
	UserName string
}

var userList = []User{{101, "Saksham"}, {102, "Suhas"}}

func main(){
	r := gin.Default()
	r.GET("/getallusers", getallusers)
	r.GET("/getuser/:id", getuser)
	r.GET("/getuserqp", getuserqp)
	r.POST("/adduser", adduser)
	r.Run("localhost:2200")
}

func getallusers(g *gin.Context){
	g.IndentedJSON(http.StatusOK, userList)
}

func getuser(g *gin.Context){
	idstr := g.Param("id")
	id, _ := strconv.Atoi(idstr)
	for _, v := range userList{
		if id == v.UserId{
			g.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	g.IndentedJSON(http.StatusNotFound, gin.H{"message": "Element Not Found"})
}

func getuserqp(g *gin.Context){
	//qp := g.Request.URL.Query()
	//usrid, _ := strconv.Atoi(qp["id"][0])

	idstr := g.Request.URL.Query().Get("id")
	id, _ := strconv.Atoi(idstr)
	for _, v := range userList{
		if id == v.UserId{
			g.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	g.IndentedJSON(http.StatusNotFound, gin.H{"message":"Element Not Found"})
}

func adduser(g *gin.Context){
	var usr User
	err := g.BindJSON(&usr)
	if err != nil{
		return
	}
	userList = append(userList, usr)
	g.IndentedJSON(http.StatusOK, userList)
}