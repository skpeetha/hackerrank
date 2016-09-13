package main
import (
    "fmt"
    "bufio"
    "os"
)
func main(){
    str := bufio.NewReader(os.Stdin)
    fmt.Println("Hello, World.")
    // str.ReadString method returns multiple values we need only first
    txt, _ := str.ReadString('\n')
    fmt.Println(txt)
    
}
