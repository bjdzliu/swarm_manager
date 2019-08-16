package main

/*
通过ssh的方式，将server上的文件，传输到本地
 */
import (
	"time"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"path"
	"os"
	"net"
)

func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

func transferLog(user string,password string,host string,containerid string) {

	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err := connect(user, password, host, 22)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	// 用来测试的远程文件路径 和 本地文件夹
	var remoteFilePath = "/var/lib/docker/containers/"+containerid+"/"+containerid+"-json.log"
	fmt.Println(remoteFilePath)
	var localDir = "/home/devpay"

	srcFile, err := sftpClient.Open(remoteFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	var localFileName = path.Base(remoteFilePath)
	dstFile, err := os.Create(path.Join(localDir, localFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	if _, err = srcFile.WriteTo(dstFile); err != nil {
		log.Fatal(err)
	}

	fmt.Println("copy log file from remote server finished!")

}

