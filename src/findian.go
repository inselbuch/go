package main
import "fmt"
import "regexp"
import "bufio"
import "os"
import "strings"

func main() {
   
   reader := bufio.NewReader(os.Stdin)

   fmt.Printf("Enter a string: ")
   text, _ := reader.ReadString('\n')
   text = strings.Replace(text, "\n", "", -1)

   sregex := "(?i)^i.*a.*n$"
   re := regexp.MustCompile(sregex)
   match := re.Match([]byte(text))

   if match {
      fmt.Printf("Found!\n")
   } else {
      fmt.Printf("Not found!\n")
   }
}
