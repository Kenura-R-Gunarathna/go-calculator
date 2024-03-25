package main

import "fmt"

func main() {

  var option int 

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
      fmt.Println("Exiting program.")
      return
    case 1: // Circle
      fmt.Println("Area is: ", calCircleArea())
    case 2: // Triangle
      fmt.Println("Area is: ", calTriangleArea())
    case 3: // Rectangle
      fmt.Println("Area is: ", calRectangleArea())
    case 4: // Parallelogram
      fmt.Println("Area is: ", calParallelogramArea())
    case 5: // Trapizium
      fmt.Println("Area is: ", calTrapiziumArea())
    default:
        fmt.Println("Invalid value.")
    }

    fmt.Println("-------------------")
  }
}

func calCircleArea() float64 {
  const PI float64 = 3.14
  var r, area float64

  fmt.Println("- Circle")
  fmt.Print("Enter the radius: ")
  fmt.Scanln(&r)
  
  area = PI * r * r
  return area
}

func calTriangleArea() float64 {
  var perpendicularHeight, baseLength, area float64

  fmt.Println("- Triangle")
  fmt.Print("Enter the perpendicular height: ")
  fmt.Scanln(&perpendicularHeight)
  fmt.Print("Enter the base length: ")
  fmt.Scanln(&baseLength)
  
  area = perpendicularHeight * baseLength
  return area
}

func calRectangleArea() float64 {
  var length, width, area float64

  fmt.Println("- Rectangle")
  fmt.Print("Enter the length: ")
  fmt.Scanln(&length)
  fmt.Print("Enter the width: ")
  fmt.Scanln(&width)
  
  area = length * width
  return area
}

func calParallelogramArea() float64 {
  var perpendicularHeight, baseLength, area float64

  fmt.Println("- Parallelogram")
  fmt.Print("Enter the perpendicular height: ")
  fmt.Scanln(&perpendicularHeight)
  fmt.Print("Enter the base length: ")
  fmt.Scanln(&baseLength)
  
  area = perpendicularHeight * baseLength
  return area
}

func calTrapiziumArea() float64 {
  var perpendicularHeight, baseLength_1, baseLength_2, area float64

  fmt.Println("- Parallelogram")
  fmt.Print("Enter the perpendicular height: ")
  fmt.Scanln(&perpendicularHeight)
  fmt.Print("Enter the base length 1: ")
  fmt.Scanln(&baseLength_1)
  fmt.Print("Enter the base length 2: ")
  fmt.Scanln(&baseLength_2)
  
  area = perpendicularHeight * (baseLength_1 + baseLength_2) / 2
  return area
}