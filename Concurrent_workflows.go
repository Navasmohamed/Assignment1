package main
import (
 "fmt"
 "time"
)
func myFunc(done chan string) {
   for i := 0; i < 10; i++ {
      time.Sleep(time.Millisecond * 500)
      fmt.Println(i, " Hello")
   }
   fmt.Println("Finished loop in my function")
   done <- "Goroutine finished" 
}
func main() {
   done := make(chan string)   
   go myFunc(done)             
   for i := 0; i < 5; i++ {
      time.Sleep(time.Millisecond * 500)
      fmt.Println(i, " Welcome")
   }
msg := <- done               
   fmt.Println(msg)
}