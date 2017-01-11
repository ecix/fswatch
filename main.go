package main

import (
    "flag"
    "log"
    "os"
    "os/exec"
    "strings"

    "github.com/rjeczalik/notify"
)

func main() {
    var dir, cmdstr string
    flag.StringVar(&dir,    "dir", "", "the directory to watch")
    flag.StringVar(&cmdstr, "cmd", "", "the command to run on change events")
    flag.Parse()

    if dir == "" || cmdstr == "" {
      flag.Usage()
      log.Fatal("need both dir and cmd flags")
    }

    watcher := make(chan notify.EventInfo, 1)

    dir = strings.Join([]string{dir, "..."}, "/")
    err := notify.Watch(dir, watcher, notify.Write)

    if err != nil {
        log.Fatal(err)
    }

    defer notify.Stop(watcher)

    for {
        select {
            case ev := <-watcher:
                cmd := exec.Command(cmdstr, ev.Path())
                cmd.Stdout = os.Stdout
                cmd.Stderr = os.Stderr
                err := cmd.Run()


                if err != nil {
                  log.Println("Error executing command:", err)
                }
        }
    }
}
