package systemgo

import (
    "fmt"
    "os"
    "strings"
    "github.com/bramz/lib"
)

func cmdLine() {
    t := &lib.Task{
        Self:     os.Args[0],
        Name:     os.Args[1],
        State:    os.Args[2],
        Ppid:     os.Getppid(),
        Pid:      os.Getpid(),
        Filename: strings.TrimPrefix(os.Args[1], "bin/"),
    }

    switch t.State {
    case "start":
        lib.StartTask(t.Name, t.Filename)
    case "stop":
        lib.StopTask(t.Name)
    case "restart":
        lib.RestartTask(t.Pid, t.Name)
    default:
        fmt.Println("usage: <application> <start/stop/restart>")
    }
}

func main() {
    cmdLine()
}
