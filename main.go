package main

import (
	"bufio"
	"io"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"

	"fmt"

	"path/filepath"

	"os"
)

var Mdomain map[string]string

func init() {
	Mdomain = make(map[string]string)
}

const (
	pathName = "./tmp"
	fileName = "domain.txt"
)

type Worker struct {
	watch      *fsnotify.Watcher
	reloadChan chan int
}

//监控目录

func (w *Worker) watchDir(dir string) {

	//通过Walk来遍历目录下的所有子目录

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		//这里判断是否为目录，只需监控目录即可

		//目录下的文件也在监控范围内，不需要我们一个一个加

		if info.IsDir() {

			path, err := filepath.Abs(path)

			if err != nil {

				return err

			}

			err = w.watch.Add(path)

			if err != nil {

				return err

			}

			fmt.Println("监控 : ", path)

		}

		return nil

	})

	go func() {

		for {

			select {

			case ev := <-w.watch.Events:
				if ev.Op&fsnotify.Write == fsnotify.Write {

					fmt.Println("写入文件 : ", ev.Name)
					w.reloadChan <- 1
				}

			case err := <-w.watch.Errors:

				{

					fmt.Println("error : ", err)

					return

				}

			}

		}

	}()

}

func ListenAndReload() {
	watch, _ := fsnotify.NewWatcher()

	w := Worker{
		watch:      watch,
		reloadChan: make(chan int),
	}

	w.watchDir(pathName)

	for {
		res, err := w.reloadFile(pathName + "/" + fileName)
		if err != nil {
			fmt.Printf("文件读取错误,请尽快修改.[%v]\n", err)
			time.Sleep(time.Second)
			continue
		}
		m, err := w.reloadMap(res)
		if err != nil {
			fmt.Printf("文件读取错误,请尽快修改.[%v]\n", err)
			time.Sleep(time.Second)
			continue
		}
		if len(m) == 0 {
			fmt.Printf("当前无数据\n")
			time.Sleep(time.Second)
			continue
		}
		Mdomain = m
		fmt.Println(Mdomain)
		<-w.reloadChan
	}
}

func (w *Worker) reloadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var result []string
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n') //以'\n'为结束符读入一行
		result = append(result, line)
		if err != nil || io.EOF == err {
			break
		}

	}
	return result, nil
}

func (w *Worker) reloadMap(res []string) (map[string]string, error) {
	m := make(map[string]string)
	for _, str := range res {
		strs := strings.Split(str, "=")
		if len(strs) != 2 {
			return nil, fmt.Errorf("%v_length_is_not_2", strs)
		}
		strs[0] = strings.TrimSpace(strs[0])
		strs[1] = strings.TrimSpace(strs[1])
		m[strs[0]] = strs[1]
	}
	return m, nil
}

func main() {

	ListenAndReload()

	select {}

}
