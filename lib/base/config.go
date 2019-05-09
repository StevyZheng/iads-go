package base

import (
	"bufio"
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
)

/****************************************/
//config file like ifcfg-eth0
/****************************************/
type CommonConfigParser struct {
	filePath string
	buffer   map[string]string
}

func NewCommonConfigParser(filename string) *CommonConfigParser {
	return &CommonConfigParser{
		filePath: filename,
	}
}

func (e *CommonConfigParser) Read() (int, error) {
	e.buffer = make(map[string]string)
	file, err := os.Open(e.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nowRowStr := Trim(scanner.Text(), "\n")
		tmpList := strings.Split(nowRowStr, "=")
		listLen := len(tmpList)
		if listLen > 1 {
			e.buffer[tmpList[0]] = tmpList[1]
		}
	}
	return len(e.buffer), err
}

func (e *CommonConfigParser) save() (int, error) {
	var (
		err error
		fp  *os.File
	)
	mapLen := len(e.buffer)
	if e.buffer == nil || mapLen <= 0 {
		log.Fatal("buffer is nil")
	} else {
		fp, err = os.OpenFile(e.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if fp == nil {
			log.Fatal("open file failed.")
		}
		for k, v := range e.buffer {
			rowStr := fmt.Sprintf("%s=%s\n", k, v)
			_, err = fp.WriteString(rowStr)
		}
	}
	return mapLen, err
}

func (e *CommonConfigParser) GetValue(key string) (string, error) {
	var (
		ret string
		err error
	)
	if e.buffer == nil {
		err = errors.New("buffer is nil")
		log.Fatal("buffer is nil")
	} else {
		if _, ok := e.buffer[key]; ok {
			ret = e.buffer[key]
		} else {
			err = errors.New("nokey")
			ret = ""
		}
	}
	return ret, err
}

func (e *CommonConfigParser) SetValue(key string, value string) error {
	var (
		err error
	)
	_, err = e.Read()
	e.buffer[key] = value
	_, err = e.save()
	return err
}

type SectionConfigParserError struct {
	errorInfo string
}
type SectionConfigParser struct {
	confParser *ini.File
	filePath   string
}

func (e *SectionConfigParserError) Error() string { return e.errorInfo }

func NewSectionConfigParser(filename string) *SectionConfigParser {
	e := &SectionConfigParser{
		filePath: filename,
	}
	conf, err := ini.Load(e.filePath)
	if err != nil {
		e.confParser = nil
	}
	e.confParser = conf
	return e
}
