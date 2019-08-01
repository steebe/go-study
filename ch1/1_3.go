package main
import (
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    startIterative := time.Now()
    fmt.Printf("%.9f elapsed\n", time.Since(startIterative).Seconds())

    startStringJoin := time.Now()
    fmt.Printf("%.9f elapsed\n", time.Since(startStringJoin).Seconds())
}

func iterative() {
    for i := 0; i < len(os.Args); i++ {}
}

func stringJoin() {
    strings.Join(os.Args[1:], " ")
}
