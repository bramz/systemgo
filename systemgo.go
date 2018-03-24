package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
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

		file, err := os.Create("pidfiles/" + filename + ".pid")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		pid := []byte(strconv.Itoa(cmd.Process.Pid))
		file.Write(pid)

		fpid, err := ioutil.ReadFile("pidfiles/" + filename + ".pid")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("watching " + pn + " with pid: " + string(fpid))
		fmt.Printf("%v", time.Now())
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
        // uncomment for parent loop
		//        for {
		//            time.Sleep(time.Second)
		//        }
	case "stop":
		pid, err := ioutil.ReadFile("pidfiles/" + filename + ".pid")
		spid := string(pid)
		out, err := exec.Command("kill", "-9", spid).CombinedOutput()
		if err != nil {
			fmt.Println(err)
		} else {
			if string(out) == "" {
			}
			fmt.Println("stopping", pn)
			fmt.Println(out)

		}
	case "restart":
	default:
		fmt.Println("usage: <application> <start/stop/restart>")
	}
}
