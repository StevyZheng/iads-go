package lib

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"path"
	"time"
)

type Sftp struct {
	Username string
	Password string
	Host string
	Port string
	SftpClient *sftp.Client
}

func (e Sftp) Connect() error {
	var (
		auth []ssh.AuthMethod
		addr string
		clientConfig *ssh.ClientConfig
		sshClient *ssh.Client
		err error
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(e.Password))
	clientConfig = &ssh.ClientConfig{
		User:    e.Username,
		Auth:    auth,
		Timeout: 30 * time.Second,
	}
	addr = fmt.Sprintf("%s:%d", e.Host, e.Port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return err
	}
	if e.SftpClient, err = sftp.NewClient(sshClient); err != nil {
		return err
	}
	return nil
}

func (e Sftp) Download(filename string) error  {
	var(
		err error
	)
	srcFile, err := e.SftpClient.Open(filename)
	if err != nil{
		log.Fatal(err)
	}
	defer srcFile.Close()
	var localFileName = path.Base(filename)
	dstFile, err := os.Create(path.Join("./", localFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()
	if _, err = srcFile.WriteTo(dstFile); err != nil {
		log.Fatal(err)
	}
	return err
}

func (e Sftp) Upload(localFilename string, remoteFilename string) error {
	srcFile, err := os.Open(localFilename)
	if err != nil{log.Fatal(err)}
	defer srcFile.Close()
	dstFile, err := e.SftpClient.Create(remoteFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		_, _ = dstFile.Write(buf)
	}
	return err
}

