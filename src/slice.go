/*

Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

	June 16, 2020 - Inselbuch
		new module

*/
package main

import "fmt"
import "strconv"
import "sort"


func main() {

   var text string

   mySlice := [] int {}

   for {

      fmt.Printf("Enter an integer or X to quit: ")
      fmt.Scan(&text)

      if text == "X" || text == "x" {
         break
      }

      i,err := strconv.Atoi(text)
      if err != nil {
         fmt.Println("invalid number")
      } else {
         mySlice = append(mySlice,i)
         sort.Ints(mySlice)
      }
      fmt.Println(mySlice)
   }
}


   
