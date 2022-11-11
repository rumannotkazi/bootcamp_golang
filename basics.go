// package main

import (
	"fmt"
	"math"
	"strconv"
)

func for_test() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func calc_sqrt(x float64, threshold float64) (string, int) {
	z := 1.0
	temp_z := 1.0
	i := 0
	y := math.Abs(x)
	for ; i < 1000; i++ {
		z -= (z*z - y) / (2 * z)
		if math.Abs(temp_z-z) < threshold {
			if x < 0 {
				return fmt.Sprint(z) + "i", i
			}
			return fmt.Sprint(z), i
		}
		temp_z = z
		// fmt.Println("Value of z = ", z, " iteration ", i)
	}
	return fmt.Sprint(z), i
}

func calc_sqrtv2(x float64, threshold float64) (float64, int) {
	z := 1.0
	// temp_z := 1.0
	counter := 0
	for math.Abs((z*z-x)/(2*z)) > threshold {
		z -= (z*z - x) / (2 * z)
		counter++
	}
	return z, counter
}

func call_sqrt_func() {
	// % Square root of number
	number := 14.0
	fmt.Println("Input Number")
	fmt.Scanln(&number)

	threshold := 0.001
	fmt.Println("Input Threshold/tolerance")
	fmt.Scanln(&threshold)

	sqrt_val, niter := calc_sqrt(number, threshold)
	fmt.Println("Square Root of", number, " is =", sqrt_val, " After iterations : ", niter)
}

func pointers_slices() {
	var p *int
	i := 2
	p = &i
	q := &p

	fmt.Println(*p)
	*p = *p + 3
	fmt.Println(*q)

	type Coord struct {
		x int
		y float64
	}
	cord := Coord{89, 4}
	fmt.Println(cord)

	cord_add := &cord

	fmt.Println(cord_add, *cord_add, cord_add.x, cord_add.y, &cord_add.x, &cord_add.y)
	cord.x = 3
	fmt.Println(cord_add, *cord_add, cord_add.x, cord_add.y, &cord_add.x, &cord_add.y)

	a := make([]int, 4, 10)
	// a[3] = 21
	fmt.Println(a)

	b := append(a, 9, 10, 11, 12, 13, 13, 17)

	fmt.Println(b, cap(b))

	pw := [4]int{1, 2, 3, 4}
	for range pw {
		fmt.Println("Hey!")
	}
}

type Coord struct {
	x, y float64
}

// type ret_fun func f

func ret_fun() int {
	return 3
}
func (c *Coord) DistFromOrigin(scale float64) float64 {
	// c.x = 4
	return math.Sqrt(c.x*c.x+c.y*c.y) * scale
}
func methods_interface() {
	cordinate := Coord{3.0, 4.0}
	fmt.Println(cordinate.DistFromOrigin(3))
}

func error_handling(x string) (int, error) {
	i, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Couldn't convert")
		return i, fmt.Errorf("Hello")
	}

	fmt.Println(i)
	return i, fmt.Errorf("Hello")
}

type InfiniteReader struct{}

func (r InfiniteReader) blabla(b []byte) (int, error) {
	i := 0
	for ; i < len(b); i++ {
		b[i] = 65
	}
	return len(b), nil
}

func readers() {
	buf := make([]byte, 5)
	r := InfiniteReader{}
	for x := 0; x < 5; x++ {
		n, err := r.blabla(buf)
		// fmt.Printf("n = %v err = %v b = %v\n", n, err, buf)

		if err != nil {
			break
		}
		// fmt.Printf("b[:n] = %q\n", buf[:n])
		fmt.Printf("%s\n", buf[:n])
	}
}

//func main() {
//	// fmt.Println("Hello, 世界")
//	// fmt.Println("my lucky number is " + fmt.Sprint(rand.Intn(6)))
//	// fmt.Println(for_test())
//
//	// call_sqrt_func()
//	// methods_interface()
//	// error_handling("4")
//	readers()
//
//}
