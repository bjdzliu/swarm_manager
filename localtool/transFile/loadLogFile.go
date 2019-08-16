package main

/*
access swarm using docker client
 */

import (
	"github.com/docker/docker/api/types"
	"context"
	"net/http"
	"github.com/docker/docker/client"
	"fmt"
	"io/ioutil"
)

func loadLogFile(containerid string,srcfile string) {

	ctx := context.Background()
	transport := new(http.Transport)
	client2 := &http.Client{Transport: transport}
	initclient := client.WithHTTPClient(client2)
	selfclient := client.Client{}

	//调用swarm的地址也可以像调用docker engine一样








	//initclienttls := client.WithTLSClientConfig("./cert-dev/ca.pem", "./cert-dev/cert.pem", "./cert-dev/key.pem")

	initclienttls := client.WithTLSClientConfig("/tmp/sushwysush/caCert", "/tmp/sushwysush/serverCert", "/tmp/sushwysush/key")

	initclienthost := client.WithHost("tcp://url:port")
	initscheme := client.WithScheme("https")
	err := initscheme(&selfclient)
	if err != nil {
		panic(err)
	}

	err = initclient(&selfclient)
	if err != nil {
		panic(err)
	}

	err = initclienthost(&selfclient)
	if err != nil {
		panic(err)
	}
	err = initclienttls(&selfclient)
	if err != nil {
		panic(err)
	}

	options := types.ContainerListOptions{All: true}
	containerslist, err := selfclient.ContainerList(ctx, options)
	for _, v := range containerslist {

		//fmt.Println(k)
		if containerid == v.ID {
			content,_,err:=selfclient.CopyFromContainer(ctx,v.ID,srcfile)
			if err != nil {
				panic(err)
			}
			body, err := ioutil.ReadAll(content)
			localfile:="."+srcfile
			if ioutil.WriteFile(localfile,body,0644) == nil {
				fmt.Println("保存文件成功")
			}
			// 缺点：2048大小。小于2048字节，自动补0
		}

	}
	if err != nil {
		panic(err)
	}

	/*	//options := types.ContainerLogsOptions{ShowStdout: true,Tail: "1"}
		// Replace this ID with a container that really exists
		out, err := cli.ContainerLogs(ctx, "0708c50a463dd94ec030cf0771eb3f7cf037e367bf4f7e4996cc6a1776bcef52", options)
		if err != nil {
			panic(err)
		}

		io.Copy(os.Stdout, out)
	*/
}
