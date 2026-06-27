package main

import (
    "flag"
    "log"
    "os"
    "path/filepath"
)

func main() {
    c:= flag.Bool("c", false, "option: c")
    l:= flag.Bool("l", false, "option: l")
    m:= flag.Bool("m", false, "option: m")
    w:= flag.Bool("w", false, "option: w")

    flag.Parse()

    command := filepath.Base(os.Args[0])
    args := flag.Args()
    if command == "zerobox" {
        command = args[0]
        args = args[1:]
    } 

    switch command {
    case "wc":
        Wc(args, c, l, m, w)
    default:
        log.Fatalf("invalid command %v", command)
    }

}

