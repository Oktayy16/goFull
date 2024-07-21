package main

import (
	"errors"
	"fmt"
)

func main() {
	var intNum int16 = 32 //
	fmt.Println(intNum)

	var deneme float32 = 5
	fmt.Println(deneme) // float32 ve 64 olmak üzere 2 tip decimal bulunmakta

	var pozitif uint = 5
	println(pozitif)

	//stringlerde direk işlediğin gibi çıkması için `` kullanılıyor
	var string1 string = `Hello
						  World` // çıktısı Hello altta World olarak aynı çıkar

	var number = len(string1) // burada uzunluğu kaç byte tutuğu olarak işler harf sayısı değil`!!
	fmt.Println(number)

	// convert işlemi
	var int1 int = 5
	var float1 float32 = float32(int1)
	fmt.Println(float1)

	const pi float32 = 3.1415
	fmt.Println(pi)

	var result, remainder, err = intDivision(11, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Check the remainder")
	}
	fmt.Printf("The result is %v and remainder is %v", result, remainder)

	// switchcase types
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Println("result 0")
	default:
		fmt.Printf("The result is %v and remainder is %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("Cannot perform")
	case 1, 2:
		fmt.Println("Is it 1 or 2")
	default:
		fmt.Println("not close")
	}

}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
