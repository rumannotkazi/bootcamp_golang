package main

import (
	"fmt"
    "math"
)

func calc_amount(nums []int) int {
// Calculating the amount of times a number is higher than the previous
    counter := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			counter++
		}
        if nums[i]==nums[i+1]{
            counter = counter+3
        }
	}
	return counter
}

type circle struct{
    r float64
}

type rectangle struct{
    l float64
    b float64
}



func calc_area_circle(circle circle) float64{
    return math.Pi*circle.r*circle.r
}

func calc_area_rectangle(rectangle rectangle) float64 {
    return rectangle.l*rectangle.b
}

func (c circle) Area() float64{
    return math.Pi*c.r*c.r
}


func (rect rectangle) Area() float64 {
    return rect.l*rect.b
}

type shape interface{
    Area() float64
    shape_name() string
}

func(circ circle) shape_name() string{
    return "circle"
}

func print_area(s shape){
    fmt.Println("Printing through the interface.")
    fmt.Println("Area of", s.shape_name())
    fmt.Println(s.Area())
}

func main() {
//	numbers := []int{1, 5, 7, 7, 10, 15, 6, 19, 6, 18}
//	value := calc_amount(numbers)
//	fmt.Println(value)

    r:=2.0
    circ:=circle{r:2.0}
    rect:=rectangle{l:4.0, b:5.0}
    fmt.Println(calc_area_circle(circle{r})) // functions
    fmt.Println(calc_area_rectangle(rect)) // functions

    //methods
    fmt.Println(rect.Area())
    fmt.Println(circ.Area())

    //interface
//    calc_area(rect)
    print_area(circ)

}
