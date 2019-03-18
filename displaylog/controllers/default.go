package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"fmt"
	//        "strings"
	// "bufio"
	"log"
	"time"
	//         "io/ioutil"
	"github.com/gorilla/websocket"
)

type MainController struct {
	beego.Controller
}

type WsController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type InstanceInfo struct {
	clusterid   string
	containerid string
}

type ShowlogController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}


type Result struct{ lines []byte }

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 8192

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on connection.
	closeGracePeriod = 10 * time.Second
)

func (s *ShowlogController) Display() {

	s.TplName = "showlogs.tpl"
}

var ioreader io.ReadCloser

func (s *WsController) Wsreturn() {

	var msgchan = make(chan *Result)
	var stopSignal = make(chan bool)
	var containerid string

	ws, err := upgrader.Upgrade(s.Ctx.ResponseWriter, s.Ctx.Request, nil)
	if err != nil {
		log.Println("write:", err)
	}

	//循环等待网页端的输入
for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if messageType==1{
			containerid=string(p)
			fmt.Printf("containerid is %s\n",containerid)
		}

	go test(msgchan,stopSignal,containerid)
	go sendmsg(ws, msgchan)
	s.TplName = "showlogs.tpl"
}

}

func monitorSignal(stopSignal  chan  bool){
	if <- stopSignal{
		fmt.Println("right i read it")
	}
	return
}


func sendmsg(ws *websocket.Conn, dstchan <-chan *Result) {
	var b []byte = make([]byte, 4096)
	for {
		//fmt.Println(i)
		select {
		case output, ok := <-dstchan:

			b = output.lines
//通道被关闭
			if !ok {
				fmt.Println("no data")
				ws.WriteMessage(websocket.CloseMessage, []byte{})
			}

			ws.SetWriteDeadline(time.Now().Add(writeWait))
			fmt.Println("Ain default.go", string(b))
			err := ws.WriteMessage(websocket.TextMessage, b)
			//w, err := ws.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println("somethins wrong after WriteMessage", err)
				ws.Close()
				break
			}

			//w.Write(b)
			//w.Write([]byte("1231"))
			//			if err := w.Close(); err != nil {
			//				return
			//			}



		default:
			time.Sleep(1 * time.Second)

		} //end select
	} // end for

} //end func
