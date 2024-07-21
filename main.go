package main

import (
	"errors"
	"fmt"
	"strings"
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

	structFunc()

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

func Strings() {
	// string olarak bir string üzerinde gezebiliriz, char ları tek tek basar ve her birinin kapladığı bir byte miktarı bulunur. € gibi harflar fazla yer
	// kaplarlar. String  değişmez bir kez yaraılır sabit kalır, değiştirilemez. String e kelime eklenecekse string builder kullanılmalıdır
	strSlice := []string{"s", "s", "s", "a"}
	var stBuilder strings.Builder // it's not mutable
	for i := range strSlice {
		stBuilder.WriteString(strSlice[i])
	}
}

// struct
type gasEngine struct { // as a default properties will be created with default of deir values for this var 0 and 0
	mpg      uint8
	gallons  uint8
	carOwner owner
	owner    // giving a name to struct is not nessasary
	int      // giving a name to struct is not nessasary
}

type owner struct {
	name string
}

// Create a metot for the struct
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type engine interface { // with interface not genericly we can use many thing with it(share same metots with it)
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You cannot do it")
	}
}

func structFunc() {
	var structDeneme gasEngine // 0,0
	var structDeneme2 gasEngine = gasEngine{15, 20, owner{"Oktay"}, owner{"Oktay2"}, 15}

	// bilinmeyen struct !!! direkly for that purpose
	var myEngine2 = struct {
		mpg  uint8
		gall uint8
	}{15, 15}
	fmt.Println(myEngine2)

	fmt.Println(structDeneme, structDeneme2)

	canMakeIt(structDeneme, 150) // -> we gave struct for interface and it's done
}

func pointers() {
	var p *int32               // in a adress equls nil
	var p2 *int32 = new(int32) // creates a value in some address and p2 holding address of it
	var i int32                // for some address equls 0 value

	// if use * agains p again it will give us value -> *p2 = 0

	// &
	p3 := &i // Don't bring value of i, bring address of i so p3 will be pointer
}


// goroutines
func goRoutines() {
	
}
