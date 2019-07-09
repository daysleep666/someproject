package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("报告下还活着")
			time.Sleep(time.Second * 10)
		}
	}()
	for {
		c := "backup.sh"
		cmd := exec.Command("sh", c)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(out))
		}
		time.Sleep(time.Hour * 1)
	}
}
