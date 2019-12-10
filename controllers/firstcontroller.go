package controllers

import "github.com/astaxie/beego"

type FirstController struct {
	beego.Controller
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Users []User

var users []User

func init() {
	users = Users{
		User{ID: 1, Name: "hota1024"},
	}
}

func (this *FirstController) GetUsers() {
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = users
	this.ServeJSON()
}
