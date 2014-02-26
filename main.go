package main

import (
    "fmt"
    "io/ioutil"
    "os"
    // I'm not using text/scanner because this is an exercise
    "github.com/uberj/bindparse/scanner"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("File argument is required")
        os.Exit(1)
    }
    content, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Printf("couldn't read file: %s\n%s\n", os.Args[1], err)
        os.Exit(1)
    }
    var n scanner.Token
    //s := scanner.Scanner{Source: string(content)}
    s := &scanner.Scanner{Source: string(content)}
    i := 1
    for {
        if n = <-s.Next(); n.End {
            break
        }
        fmt.Printf("(%d) Token: '%s'\n", i,  n.Value)
        p := <-s.Peek()
        fmt.Printf("(%d) Peek Token is: '%s'\n", i + 1, p.Value)
        i++
    }
}
