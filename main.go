package main
import "./runtime"
import "fmt"

func main() {
    o := runtime.NewObject()
    o.Set("hi", runtime.Fstring("hey"))
    fmt.Println(o)
}
