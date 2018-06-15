package lib

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type Task struct {
	Self     string
	Name     string
	State    string
	Ppid     int
	Pid      int
	Filename string
	Proc     *os.Process
	Cmd      *exec.Cmd
}

func StartTask(t *Task) {
	cmd := exec.Command(t.Name)
	err := cmd.Start()

	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(".systemgo/pidfiles/" + t.Filename + ".pid")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	pid := []byte(strconv.Itoa(cmd.Process.Pid))
	file.Write(pid)
	t.State = "running"
	watchTask(cmd, t.Name, t)
	return
}

func watchTask(cmd *exec.Cmd, name string, t *Task) {
	watchState(t)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
	}

	done := make(chan bool)

	// watch for file change events
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
					run := exec.Command("./systemgo", name, "start")
					run.Start()
					os.Exit(1)
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-done
	watcher.Close()

}

func watchState(t *Task) {
	cstate := make(chan *os.ProcessState)

	go func() {
		state, err := t.Proc.Wait()
		if err != nil {
			fmt.Println(err)
			return
		}
		cstate <- state
	}()
	select {
	case s := <-state:
		if t.State == "stopped" {
			fmt.Println("Process is not running")
			return
		}
	}
}

func StopTask(t *Task) {
	fmt.Println("Stopping task")
	pid, err := ioutil.ReadFile(".systemgo/pidfiles/" + t.Filename + ".pid")
	if err != nil {
		fmt.Println(err)
	}
	out := exec.Command("kill", "-9", string(pid))
	out.Run()
	t.State = "stopped"
	return
}

func RestartTask(pid int, t *Task) {
	fmt.Println("Restarting task")
	kill := exec.Command("kill", "-9", strconv.Itoa(pid))
	kill.Start()
	kill.Process.Kill()
	kill.Run()
	run := exec.Command("./systemgo", t.Name, "start")
	run.Start()
	os.Exit(1)
	return
}
