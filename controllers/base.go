package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type BaseController struct {
	beego.Controller
}

type NestPreparer interface {
	NestPreparer()
}

func (ctx *BaseController) Prepare() {
	log.Println("BaseControll")
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPreparer()
	}
}
