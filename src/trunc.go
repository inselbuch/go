package main
import "fmt"
//import "math"

func main() {
   var f float32
   fmt.Printf("Enter a floating point number: ")
   _ , _ = fmt.Scan(&f)
   fmt.Printf("%d\n",int(f))   
}
