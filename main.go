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

func Arrays() {
	var intArr [3]int32 // [0,0,0] olarak veri girildi

	// For initilize with data in it
	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	intArr3 := [3]int32{1, 2, 3} // kısa yol yazımı
	fmt.Printf("birinci %v, ikinci %v üçüncü %v ", intArr, intArr2, intArr3)
}

func Slice() {
	// add a value to slice
	var intSlice []int32 = []int32{3, 4, 6}
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	// add another slice to main slice with ...
	inySlice2 := []int32{8, 9}
	intSlice = append(intSlice, inySlice2...)
	fmt.Println(intSlice)

	// create
	var intSlice3 []int = make([]int, 3, 8) // len 3 , cap 8 olarak yaratır
	fmt.Println(intSlice3)
}

// Dictionary
func Map() {
	// Key in [], value end of it
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) // return map[]

	// initiliaze
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"]) // return 23 döndürür

	fmt.Println(myMap2["Jason"])  // return 0 because it's default value, be carryful about it
	var age, ok = myMap2["Jason"] // If it's exist ok will be True. If its not then false
	if ok {
		fmt.Println(age)
	}

	// LOOPS !!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	// like foreach
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	for i, v := range myMap2 { // index and value with loop
		fmt.Printf("Index %v, Value %v", i, v)
	}

	//while loop
	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	// clasic for loop
	for i := 0; i < 10; i++ { // İ-- dec 1, i += 10 inc 10, i *= 10 multiply 10, i /= 10 divide 10
		fmt.Println(i)
	}

	//  if collection cap is decleare it then that slice performance will be better like []int{} vs make([]int,0,10000) -> make will be 3 times faster for 1 loop

}
