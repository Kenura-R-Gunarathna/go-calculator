package area

import (
	"fmt"
)

func CalculateTriangleArea() float64 {
    var perpendicularHeight, baseLength, area float64

    fmt.Println("- Triangle")
    fmt.Print("Enter the perpendicular height: ")
    fmt.Scanln(&perpendicularHeight)
    fmt.Print("Enter the base length: ")
    fmt.Scanln(&baseLength)

    area = perpendicularHeight * baseLength / 2
    return area
}
