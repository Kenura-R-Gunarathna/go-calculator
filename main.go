package main

import (
	"fmt"

	"github.com/Kenura-R-Gunarathna/go-calculator/area"
	"github.com/fatih/color"
)

func main() {

  var option int 

  red := color.New(color.FgRed).SprintFunc()
  green := color.New(color.FgGreen).SprintFunc()
  
  fmt.Println("-------------------")
  fmt.Println("Go KRAG Calculator")
  fmt.Println("-------------------")

  for {

    fmt.Println("Area equations -")

    fmt.Println("[1] Circle")
    fmt.Println("[2] Triangle")
    fmt.Println("[3] Rectangle")
    fmt.Println("[4] Parallelogram")
    fmt.Println("[5] Trapizium")
    fmt.Println("[0] Quit program")

    fmt.Print("Select the equation (by entering number): ")

    fmt.Scanln(&option)

    switch option {
    case 0: // Exit
      fmt.Println(red("Exiting program."))
      return
    case 1: // Circle
      fmt.Println(green("Area is: ", area.CalculateCircleArea()))
    case 2: // Triangle
      fmt.Println(green("Area is: ", area.CalculateTriangleArea()))
    case 3: // Rectangle
      fmt.Println(green("Area is: ", area.CalculateRectangleArea()))
    case 4: // Parallelogram
      fmt.Println(green("Area is: ", area.CalculateParallelogramArea()))
    case 5: // Trapizium
      fmt.Println(green("Area is: ", area.CalculateTrapeziumArea()))
    default:
        fmt.Println(red("Invalid value."))
    }

    fmt.Println("-------------------")
  }
}

