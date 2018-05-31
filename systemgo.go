package systemgo

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "io/ioutil"
    "os"
    "os/exec"
    "strconv"
    "strings"
    "github.com/bramz/lib/task"
)

func cmdLine() {
    t := &Task{
        Self:     os.Args[0],
        Name:     os.Args[1],
        State:    os.Args[2],
        Ppid:     os.Getppid(),
        Pid:      os.Getpid(),
        Filename: strings.TrimPrefix(os.Args[1], "bin/"),
    }

    switch t.State {
    case "start":
        startTask(t.Name, t.Filename)
    case "stop":
        stopTask(t.Name)
    case "restart":
        restartTask(t.Pid, t.Name)
    default:
        fmt.Println("usage: <application> <start/stop/restart>")
    }
}

func main() {
    cmdLine()
}
