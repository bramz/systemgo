package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"os"
	"os/exec"
	//	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	pn := os.Args[1]
	use := os.Args[2]
	filename := strings.TrimPrefix(pn, "bin/")

	switch use {
	case "start":
		cmd := exec.Command(pn)
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Started", pn)

		file, err := os.Create(".systemgo/pidfiles/" + filename + ".pid")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		pid := []byte(strconv.Itoa(cmd.Process.Pid))
		file.Write(pid)

		fpid, err := ioutil.ReadFile(".systemgo/pidfiles/" + filename + ".pid")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("watching " + pn + " with pid: " + string(fpid))
		fmt.Printf("%v", time.Now())
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			fmt.Println(err)
		}

		done := make(chan bool)

		// handle events
		go func() {
			for {
				select {
				case event := <-watcher.Events:
					fmt.Println("\nevent:", event)
					if event.Op&fsnotify.Chmod == fsnotify.Chmod {
						fmt.Println("Rebuild issued, restarting", cmd.Process.Pid)
						kill := exec.Command("kill", "-9", strconv.Itoa(cmd.Process.Pid))
						kill.Start()
						cmd.Process.Kill()
						cmd.Run()
						run := exec.Command("./systemgo", pn, "start")
						run.Start()
						os.Exit(1)
					}
				case err := <-watcher.Errors:
					fmt.Println("error:", err)
				}
			}
		}()

		err = watcher.Add(pn)
		if err != nil {
			fmt.Println(err)
		}

		<-done

		watcher.Close()
	case "stop":
		pid, err := ioutil.ReadFile(".systemgo/pidfiles/" + filename + ".pid")
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
	default:
		fmt.Println("usage: <application> <start/stop>")
	}
}
