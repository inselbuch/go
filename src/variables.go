# variables must start with a letter
# case sensitive
#
# every variable must be declared and have a type
#
package main

import "fmt"

func main() {
   var x int
   var float1 float32
   var str1 string
 
   # you can define an alternate name for a type (alias)

   type Celcius float64
   type IDnum int

   var temp Celcius
   var pid IDnum

   # initialize
   var x int = 100

   # infers the type from the right-hand side
   var x = 100

   # initialize after the declaration
   x = 110

   # untialized values get the zero value for that type
   # int = 0
   # string = empty string

   # short variable declaration, declares and initializes
   # can only be done within a function
   x := 100

   

}


