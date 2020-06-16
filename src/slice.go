/*

Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

	June 16, 2020 - Inselbuch
		new module

*/
package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"
import "sort"


func main() {

   var done bool

   mySlice := make([] int,3)
   reader := bufio.NewReader(os.Stdin)

   for done==false {

      fmt.Printf("Enter an integer or X to quit: ")
      c,_ := reader.ReadString('\n')
      c = strings.Replace(c,"\n","",-1)

      if c=="X" {
         done=true
      } else {
         i,_ := strconv.Atoi(c)
         mySlice = append(mySlice,i)
         sort.Ints(mySlice)
         fmt.Print(mySlice)
         fmt.Printf("\n")
      }
   }
}


   
