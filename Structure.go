package main
import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {

    emp1 := Employee{
        firstName: "NAVAS",
        age:       25,
        salary:    15000,
        lastName:  "MOHAMED",
    }

    emp2 := Employee{"NANDHA", "KUMAR", 25, 35000}
    emp3 := Employee{"SYED", "MOHAMED", 28, 45000}
    emp4 := Employee{"NOOR", "MOHAMED", 30, 55000}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
    fmt.Println("Employee 3", emp3)
    fmt.Println("Employee 4", emp4)
}