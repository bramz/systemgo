package systemgo

import (
	"github.com/fsnotify/fsnotify"
	"os"
	"os/exec"
	"testing"
)

func TestCliStart(test *testing.T) {
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
		TestStartTask(os.Args[1])
	}
}

func TestCliStop() {
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
		TestStopTask(os.Args[1])
	}

}

func TestCliRestart() {
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
		TestRestartTask(os.Args[1])
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
	TestwatchTask(cmd, name)
	return

}

func TestWatchTask(cmd *exec.Cmd, name string) {
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

func TestStopTask(name string) {
    fmt.Println("Stopping task", name)
    os.Exit(1)
    return
}

func TestRestartTask(name string) {
}
