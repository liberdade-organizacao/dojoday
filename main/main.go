package main

import (
    "fmt"
    "os"
    "bufio"
    "github.com/liberdade-organizacao/dojoday/view"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    v := view.NewView()

    for v.Working {
        fmt.Print(v.Write())
        line, _ := reader.ReadString('\n')
        v.Read(line)
    }
}

func pass() {

}
