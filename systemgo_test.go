package main

import "testing"

type Task struct {
	Self     string
	Name     string
	Status   string
	Ppid     int
	Pid      int
	Filename string
}

func TestCli() {
	t := &Task{
		Self:     "systemgo",
		Name:     "bin/httpd",
		Status:   "start",
		Ppid:     12345,
		Pid:      12346,
		Filename: "httpd",
	}
}

func TestStart() {
}

func TestStop() {
}

func TestRestart() {
}
