package controllers

import (
	//	"io"
	////	"os"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"context"
	"bufio"
	"strconv"

)

//var r,w,_=os.Pipe()
func test(src chan *Result,stopSignal <- chan  bool,containerid string) {

	ctx, cancel := context.WithCancel(context.Background())
	//	cli, err := client.NewEnvClient()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true, Follow: true, Tail: "10"}
	// Replace this ID with a container that really exists
	containerid="8ab1da7806e4"
	out, err := cli.ContainerLogs(ctx, containerid, options)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(out)
	i := 0
	for s.Scan() {
		if len(s.Bytes()) > 0 {
			//msgchan <- s.Bytes()
			//fmt.Println("in getlogs",s.Bytes())
			//fmt.Println("in getlogs",strconv.Itoa(i))
			line := new(Result)
			//fmt.Println([]byte(s.Text()))
			//gg:=`ڲ019-03-02T00:30:28.618+0000 I NETWORK  [conn8] Error receiving request from client: ProtocolError: Client sent an HTTP request over a native MongoDB connection. Ending connection from 127.0.0.1:54600 (connection id: 8)5`

			gg1 := strconv.Quote(s.Text())

			//line.lines =  []byte(s.Text()+ strconv.Itoa(i) )
			line.lines = []byte(gg1 + strconv.Itoa(i) )
			src <- line



			//fmt.Println(" after src <- line" + strconv.Itoa(i))
			i = i + 1
			select{
			case flag := <-stopSignal:
			if flag == true{
				cancel()
			}
			default:
				continue
			}
		}
	}

}

//func main(){
//str1:=output()
//fmt.Println(str1)

//}
