package area

import (
	"fmt"
)

func CalculateCircleArea() float64 {
    const PI float64 = 3.14
    var r, area float64

    fmt.Println("- Circle")
    fmt.Print("Enter the radius: ")
    fmt.Scanln(&r)

    area = PI * r * r
    return area
}
