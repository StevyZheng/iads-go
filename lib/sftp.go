package lib

import (
	"bytes"
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime/debug"
	"time"
)

type Ssh struct {
	Username   string
	Password   string
	Host       string
	Port       string
	session    *ssh.Session
	sftpClient *sftp.Client
}

func NewSsh(host string, username string, password string) *Ssh {
	ret := new(Ssh)
	ret.Host = host
	ret.Username = username
	ret.Password = password
	ret.Port = "22"
	ret.session = nil
	ret.sftpClient = nil
	return ret
}

func (e *Ssh) SshConnect() error {
	var (
		auth []ssh.AuthMethod
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(e.Password))
	clientConfig := &ssh.ClientConfig{
		User:            e.Username,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //ssh.FixedHostKey(hostKey),
	}
	addr := fmt.Sprintf("%s:%d", e.Host, e.Port)
	sshClient, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		return err
	}
	if e.session, err = sshClient.NewSession(); err != nil {
		return err
	}
	return err
}

func (e *Ssh) execCommand(cmd string) (string, string, error) {
	if err := e.SshConnect(); err != nil {
		log.Fatal(err)
	}
	defer e.session.Close()
	var outbt, errbt bytes.Buffer
	e.session.Stdout = &outbt
	e.session.Stderr = &errbt
	_ = e.session.Run(cmd)
	outStr := outbt.String()
	errStr := errbt.String()
	return outStr, errStr, err
}

func (e *Ssh) execCommandList(cmdList arraylist.List) (arraylist.List, arraylist.List, error) {
	if err := e.SshConnect(); err != nil {
		log.Fatal(err)
	}
	var (
		outList arraylist.List
		errList arraylist.List
	)
	defer e.session.Close()
	var outbt, errbt bytes.Buffer
	e.session.Stdout = &outbt
	e.session.Stderr = &errbt
	cit := cmdList.Iterator()
	for cit.Next() {
		tmpCmd := cit.Value().(string)
		_ = e.session.Run(tmpCmd)
		outList.Add(outbt.String())
		errList.Add(errbt.String())
	}
	return outList, errList, err
}

func (e *Ssh) CloseSftp() {
	e.sftpClient.Close()
}

func (e *Ssh) SetPort(port int) {
	e.Port = string(port)
}

func (e *Ssh) SftpConnect() error {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(e.Password))
	clientConfig = &ssh.ClientConfig{
		User:            e.Username,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //ssh.FixedHostKey(hostKey),
	}
	addr = fmt.Sprintf("%s:%d", e.Host, e.Port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		e.sftpClient = nil
		return err
	}
	if e.sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return err
	}
	return nil
}

func (e *Ssh) DownloadFile(remoteFile string, localFile string) error {
	if nil != e.sftpClient {
		srcFile, err := e.sftpClient.Open(remoteFile)
		if err != nil {
			log.Fatal(err)
		}
		defer srcFile.Close()
		dstFile, err := os.Create(localFile)
		if err != nil {
			log.Fatal(err)
		}
		defer dstFile.Close()
		if _, err = srcFile.WriteTo(dstFile); err != nil {
			log.Fatal(err)
		}
	} else {
		if panicHandle := recover(); panicHandle != nil {
			fmt.Println("panic")
			debug.PrintStack()
			return err
		}
		panic("sftpclient is nil, func Ssh.connect not run")
	}
	return err
}

func (e *Ssh) UploadFile(localFile string, remoteFile string) error {
	if e.sftpClient != nil {
		srcFile, err := os.Open(localFile)
		if err != nil {
			log.Fatal(err)
		}
		defer srcFile.Close()
		dstFile, err := e.sftpClient.Create(remoteFile)
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

	} else {
		if panicHandle := recover(); panicHandle != nil {
			fmt.Println("panic")
			debug.PrintStack()
			return err
		}
		panic("sftpclient is nil, func Ssh.connect not run")
	}
	return err
}

func (e *Ssh) UploadDir(localDir string, remoteDir string) error {
	if e.sftpClient != nil {
		localFiles, err := ioutil.ReadDir(localDir)
		if err != nil {
			log.Fatal("read local dir list fail ", err)
		}
		for _, backupDir := range localFiles {
			localFilePath := path.Join(localDir, backupDir.Name())
			remoteFilePath := path.Join(remoteDir, backupDir.Name())
			if backupDir.IsDir() {
				_ = e.sftpClient.Mkdir(remoteFilePath)
				_ = e.UploadDir(localFilePath, remoteFilePath)
			} else {
				_ = e.UploadFile(path.Join(localDir, backupDir.Name()), remoteDir)
			}
		}
	} else {
		if panicHandle := recover(); panicHandle != nil {
			fmt.Println("panic")
			debug.PrintStack()
			return err
		}
		panic("sftpclient is nil, func Ssh.connect not run")
	}
	return err
}

func (e *Ssh) DownloadDir(remoteDir string, localDir string) error {
	if e.sftpClient != nil {
		remoteFiles, err := e.sftpClient.ReadDir(remoteDir)
		if err != nil {
			log.Fatal("read remote dir list fail ", err)
		}
		for _, backupDir := range remoteFiles {
			localFilePath := path.Join(localDir, backupDir.Name())
			remoteFilePath := path.Join(remoteDir, backupDir.Name())
			if backupDir.IsDir() {
				_ = os.MkdirAll(localFilePath, os.ModePerm)
				_ = e.DownloadDir(remoteFilePath, localFilePath)
			} else {
				_ = e.DownloadFile(remoteDir, path.Join(localDir, backupDir.Name()))
			}
		}
	} else {
		if panicHandle := recover(); panicHandle != nil {
			fmt.Println("panic")
			debug.PrintStack()
			return err
		}
		panic("sftpclient is nil, func Ssh.connect not run")
	}
	return err
}
