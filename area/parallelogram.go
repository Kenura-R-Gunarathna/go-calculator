package area

import (
	"fmt"
)

func CalculateParallelogramArea() float64 {
  var perpendicularHeight, baseLength, area float64

  fmt.Println("- Parallelogram")
  fmt.Print("Enter the perpendicular height: ")
  fmt.Scanln(&perpendicularHeight)
  fmt.Print("Enter the base length: ")
  fmt.Scanln(&baseLength)
  
  area = perpendicularHeight * baseLength
  return area
}
