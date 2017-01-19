package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "os/exec"
    "strings"

    "github.com/rjeczalik/notify"
)

func main() {
    var dir, cmdstr string
    flag.StringVar(&dir,    "dir", "", "the directory to watch")
    flag.StringVar(&cmdstr, "cmd", "", "the command to run on change events (simply pipe to stdout if none given)")
    flag.Parse()

    if dir == "" {
      flag.Usage()
      log.Fatal("need dir argument")
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
                if cmdstr == "" {
                  fmt.Printf("%s\n", ev.Path())
                } else {
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
}
