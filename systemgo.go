package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Proc struct {
	Name string
	Pid  int
}

func (p *Proc) SetName(Name string) {
	p.Name = Name
}

func (p *Proc) GetName() string {
	return p.Name
}

func (p *Proc) SetPid(Pid int) {
	p.Pid = Pid
}

func (p *Proc) GetPid() int {
	return int(p.Pid)
}

func main() {
	pn := os.Args[1]
	use := os.Args[2]
	p := new(Proc)
	p.SetName(pn)
	filename := strings.TrimPrefix(pn, "bin/")

	switch use {
	case "start":
		cmd := exec.Command(pn)
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Started", pn)
		fmt.Println(cmd.Process.Pid)

		file, err := os.Create("pidfiles/" + filename + ".pid")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		pid := []byte(strconv.Itoa(cmd.Process.Pid))
		file.Write(pid)
	case "stop":
		pid, err := ioutil.ReadFile("pidfiles/" + filename + ".pid")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(pid))
        run := "kill -9 " + string(pid)
        cmd := exec.Command(run)
        err = cmd.Run()
        if err != nil {
            fmt.Println(err)
        }
		//        cmd := exec.Command("kill -9 " + pid)
		//		err := cmd.Start()
		//		if err != nil {
		//			fmt.Println(err)
		//		}
	case "restart":
	default:
		fmt.Println("usage: <application> <start/stop/restart>")
	}
}
