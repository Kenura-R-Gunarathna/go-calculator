package area

import (
	"fmt"
)

func CalculateRectangleArea() float64 {
  var length, width, area float64

  fmt.Println("- Rectangle")
  fmt.Print("Enter the length: ")
  fmt.Scanln(&length)
  fmt.Print("Enter the width: ")
  fmt.Scanln(&width)
  
  area = length * width
  return area
}
