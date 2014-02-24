package debug

import (
    "os"
    "log"
    "syscall"
)

func Debugger() {
    pid := os.Getpid()
    proc, err := os.FindProcess(pid)
    if err != nil {
        log.Fatal(err)
    }
    proc.Signal(syscall.SIGINT)
}
