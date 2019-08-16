package main

/*
access swarm using http
 */

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"flag"
	"os"
)

var appname string
var srcfilename string

func init() {
	//从参数获取应用的名字,获取值
	flag.StringVar(&appname, "appname", "", "app name in swarm cluster")

	//flag.StringVar(&projecturl, "projecturl", "", "project url,like https://domain name:port/projects/")

	//从参数文件获取
	flag.StringVar(&srcfilename, "srcfilename", "", "filename in container")

}

func Getconf(){

}

func main() {
	flag.Parse()
	client := initConnectSwarm()

	/* optional
	//从配置文件获取集群地址和证书路径
	conf, err := config.NewConfig("ini", "conf.ini")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	// 从配置文件获取url
	//projecturl:=conf.String("projecturl")
	*/


	//hard code url
	projecturl:= "https://url:port/projects/"

	url :=  projecturl + appname
	fmt.Println(url)

	resp, err2 := client.Get(url)
	if err2 != nil {
		panic(err2)
	}

	defer resp.Body.Close()
	containerResult := getContainerid(resp)

	for k, v := range containerResult {
		fmt.Println(k)
		fmt.Println(v)
		//传输传输容器内的文件
		//transferLog("root", "password", v, k)
		////传输log日志到本地
		loadLogFile(k,srcfilename)

	}


	//var P123 interface{}
	//for k, v := range result {
	//	switch vv := v.(type) {
	//	case string:
	//		fmt.Println(k, "is string", vv)
	//	case int:
	//		fmt.Println(k, "is int", vv)
	//	case []interface{}:
	//		fmt.Println(k, "is an array:",vv)
	//		P123 = reflect.ValueOf(vv)
	//		fmt.Sprintf("in vvvv %T", vv)
	//
	//	default:
	//		fmt.Println(k, "is of a type I don't know how to handle")
	//	}
	//}

	//获取到接口类型
	//	fmt.Println("123",P123)
	//	fmt.Println(reflect.TypeOf(P123).String())
	//	fmt.Sprintf("%T", P123)

	//使用simplejson,获取复杂结构json字符串中的key
	//res, err := simplejson.NewJson([]byte(string(body)))
	//fmt.Println("res:",res)
	//if err != nil {
	//	fmt.Printf(" error is %v\n", err)
	//	return
	//}
	//content := res.Get("services")
	//if err != nil {
	//	fmt.Printf("error is %v\n", err)
	//	return
	//}
	//a := content.Get("containers")
	//fmt.Println(a)
	/// end simplejson

}

//func echoMap(a interface{}){
//	c,_:=a.(map[string])
//}

func WriteWithIoutil(name string ,data []byte) {
	if ioutil.WriteFile(name,data,0600) == nil {
		fmt.Println("ok")
	}
}

func initConnectSwarm() (*http.Client) {

	/*  optional
	//从文件获取pem
	caCert, err := ioutil.ReadFile("./cert-test/ca.pem")
	if err != nil {
		log.Fatal(err)
	}*/

	//hard codepem
	caCert:= []byte(`-----BEGIN CERTIFICATE-----
pem content
-----END CERTIFICATE-----`)

	//待改进，本地留存一份，docker client 必须读取证书文件
	err := os.Mkdir("/tmp/sushwysush/",os.ModePerm)
	WriteWithIoutil("/tmp/sushwysush/caCert",caCert)

	/*  optional
	//从文件获取pem
	//cliCrt, err := tls.LoadX509KeyPair("./cert-dev/cert.pem", "./cert-dev/key.pem")
*/
	//hard codepem
	server := []byte(`-----BEGIN CERTIFICATE-----
pem content
-----END CERTIFICATE-----`)
	//待改进
	WriteWithIoutil("/tmp/sushwysush/serverCert",server)





	key := []byte(`-----BEGIN RSA PRIVATE KEY-----
pem content
-----END RSA PRIVATE KEY-----`)
	//待改进，
	WriteWithIoutil("/tmp/sushwysush/key",key)
	cliCrt, err := tls.X509KeyPair(server,key)

	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cliCrt},
			},
		},
	}
	return client

}

func getContainerid(resp *http.Response) map[string]string {
	dict := make(map[string]string)

	body, err := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	type containers struct {
		id string `json:"containers"`
	}

	type appinfo struct {
		Name   string `json:"name"`
		containers    `json:"containers"`
		Yaml_v string `json:"compose_version"`
	}

	//appinfostruct := appinfo{}

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		panic(err)
	}

	//fmt.Println("get desired",result["services"].([]interface{})[0].(map[string]interface{})["containers"])
	type containerObject struct {
		name   string
		node   string
		status string
	}

	type arrayContainer struct {
		id containerObject
	}
   fmt.Println("result content is",result)
	containerslist := result["services"].([]interface{})[0].(map[string]interface{})["containers"]
	//fmt.Printf("leixing : %v \n", containerslist)

	//jsonbody is []byte
	jsonbody, err := json.Marshal(containerslist)

	fmt.Println("jsonbody:", string(jsonbody))

	//json convert map
	var mapResult map[string]interface{}
	err = json.Unmarshal(jsonbody, &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	for k, v := range mapResult {
		//fmt.Println("container id is :", k)
		//fmt.Printf("%v \n", v)
		containerdetail_json, _ := json.Marshal(v)
		res, _ := simplejson.NewJson(containerdetail_json)
		nodejson := res.Get("node")
		nodeip, _ := nodejson.String()
		dict[k] = nodeip

	}
	return dict

}


