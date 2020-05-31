package main


import  (
            "fmt"
)

type myowntype int
var x myowntype

func main() {

fmt.Println(x)
fmt.Printf("%T\n", x)
x = 42
fmt.Println(x)

}
