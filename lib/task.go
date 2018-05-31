package task

import (
)

type Task struct {
    Self     string
    Name     string
    State   string
    Ppid     int
    Pid      int
    Filename string
    Cmd      string
}

func startTask(name string, filename string) {
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


func watchTask(cmd *exec.Cmd, name string) {
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

func stopTask(name string) {
    fmt.Println("Stopping task", name)
    pid, err := ioutil.ReadFile(".systemgo/pidfiles/" + name + ".pid")
    spid := string(pid)
    out, err := exec.Command("kill", "-9", spid).CombinedOutput()
    if err != nil {
        fmt.Println(err)
        return
    } else {
        fmt.Println("Stopping", name, out)
//      out.Run()
        return
    }
    return
}

func restartTask(pid int, name string) {
    fmt.Println("Restarting task", name)
    kill := exec.Command("kill", "-9", strconv.Itoa(pid))
    kill.Start()
    kill.Process.Kill()
    kill.Run()
    run := exec.Command("./systemgo", name, "start")
    run.Start()
    os.Exit(1)
    return
}

