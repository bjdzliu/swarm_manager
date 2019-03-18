package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type SpecController struct{
	beego.Controller
}

//type Object struct {
//	Title string
//	Content int
//}

//全局变量，两个客户端都会来修改
//var a int = 100


func (t *SpecController) Specify(){
	//ob := &Object{}
	//json.Unmarshal(s.Ctx.Input.RequestBody, ob)

	contentvalue,_:=t.GetInt32("content")
	f:=make(chan bool,1)
	fmt.Println("1")
	f<-true
	fmt.Println("2")
	monitorSignal(f)
	fmt.Println("3")
	fmt.Println("in request get contentvalue is",contentvalue)
	t.Ctx.WriteString("successfuly terminate")


}


