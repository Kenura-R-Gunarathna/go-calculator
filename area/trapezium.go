package area

import (
	"fmt"
)

func CalculateTrapeziumArea() float64 {
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