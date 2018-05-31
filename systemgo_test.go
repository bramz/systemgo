package systemgo

import (
	"github.com/fsnotify/fsnotify"
	"os"
	"os/exec"
	"testing"
)

func TestCliStart(t *testing.T) {
	t := &Task{
		Self:     "systemgo",
		Name:     "bin/httpd",
		Status:   "start",
		Ppid:     "12345",
		Pid:      "12346",
		Filename: "httpd",
	}

	os.Args[0] = t.Self
	os.Args[1] = t.Name
	os.Args[2] = t.Status

	if os.Args[0] == "systemgo" && os.Args[1] == "bin/httpd" && os.Args[2] == "start" {
		startTask(os.Args[1])
	}
}

func TestCliStop(t *testing.T) {
	t := &Task{
		Self:     "systemgo",
		Name:     "bin/httpd",
		Status:   "stop",
		Ppid:     "12345",
		Pid:      "12346",
		Filename: "httpd",
	}

	os.Args[0] = t.Self
	os.Args[1] = t.Name
	os.Args[2] = t.Status

	if os.Args[0] == "systemgo" && os.Args[1] == "bin/httpd" && os.Args[2] == "stop" {
		stopTask(os.Args[1])
	}

}

func TestCliRestart(t *testing.T) {
	t := &Task{
		Self:     "systemgo",
		Name:     "bin/httpd",
		Status:   "restart",
		Ppid:     "12345",
		Pid:      "12346",
		Filename: "httpd",
	}

	os.Args[0] = t.Self
	os.Args[1] = t.Name
	os.Args[2] = t.Status

	if os.Args[0] == "systemgo" && os.Args[1] == "bin/httpd" && os.Args[2] == "restart" {
		restartTask(os.Args[1])
	}

}

func TestStartTask(name string, filename string) {
	cmd := exec.Command(name)
	err := cmd.Start()

	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(".systemgo/pidfiles/" + filename + ".pid")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	pid := []byte(strconv.Itoa(cmd.Process.Pid))
	file.Write(pid)
	fmt.Println("Started task", name)
	watchTask(cmd, name)
	return

}
