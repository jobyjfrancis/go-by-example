package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")

	switch(i) {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
	fmt.Println("Its a weekend")
default:
	fmt.Println("Its a weekday")
}

t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("Its Morning")
case t.Hour() > 12 && t.Hour() < 4:
	fmt.Println("Its Afternnon")
default:
	fmt.Println("Evening")
}

whatAmI := func (i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("Boolean")
	case int:
		fmt.Println("Integer")
	default:
		fmt.Printf("Data type is: %T\n", t)
	}
}
whatAmI(true)
whatAmI(25)
whatAmI("Hello")

}
