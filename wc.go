package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "unicode"
)

func Wc(args []string, cOpt *bool, lOpt *bool, mOpt *bool, wOpt *bool) {
    if !(*cOpt || *lOpt || *mOpt || *wOpt) {
        // no options is the default c, l, w option
        *cOpt, *lOpt, *wOpt = true, true, true
    }

    if len(args) > 0 {
        if *mOpt {
            // multibyte char count
            file, err := os.Open(args[0])
            if err != nil {
                log.Fatalf("open fail: %v", err)
            }
            defer file.Close()

            reader := bufio.NewReader(file)
            count := 0
            for {
                _, _, err := reader.ReadRune()
                if err != nil {
                    if err == io.EOF {
                        // finish reading
                        break
                    }
                    log.Fatalf("read error: %v", err)
                    return
                }
                count++
            }

            fmt.Printf("%d ", count)
            fmt.Println(args[0])
            return
        } else {
            // count byte, line, word or all of them
            file, err := os.Open(args[0])
            if err != nil {
                log.Fatalf("open fail: %v", err)
                return
            }
            defer file.Close()

            reader := bufio.NewReader(file)
            nchar, nline, nword := 0, 0, 0      
            onWord := false

            for {
                b, err := reader.ReadByte()
                if err != nil {
                    if err == io.EOF {
                        // finish reading
                        break
                    }
                    log.Fatalf("read error: %v", err)
                    return
                }

                if *cOpt {
                    nchar++
                }

                if *lOpt && b == '\n' {
                    nline++
                }

                if *wOpt {
                    if !onWord && !unicode.IsSpace(rune(b)) {
                        // head of the word
                        onWord = true
                        nword++
                    } else if onWord && unicode.IsSpace(rune(b)) {
                        // tail of the word
                        onWord = false
                    } 
                }
            }

            // If the original wc with l, w and c option, 
            // the output format is starting white space and its different
            // from the case only single option has selected.
            // However it must be useless.
            
            if *lOpt {
                fmt.Printf("%d ", nline)
            }
            if *wOpt {
                fmt.Printf("%d ", nword)
            }
            if *cOpt {
                fmt.Printf("%d ", nchar)
            }
            fmt.Println(args[0])
            return
        }  // c, l, w option
    } else {
        // no input file
        log.Fatalf("need a input file.")
    }        
}

