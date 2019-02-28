package main

import (
	"fmt"
	"math/cmplx"
)

func main(){
	noDataType()
	newWayOfAssigningData(12)
	diffTypes()
	initValues()
	typeConversions()
	inferType()
	constantDeclaration()
	attendNumericConstants()
}


// not mensioning the type of the data

func noDataType(){
	var school, college, passed = "GPSSHS", "GCELT",  true
	fmt.Println(school, college, passed)
}

// inside the function we can use := contruct

func newWayOfAssigningData(value int){
	// outside the function we can not use this type of operation
	a := value
	fmt.Println(a)
}

// basic types

// bool string 
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64
// byte --- alias for uint8
// rune --- alias for uint32
// float32 float64
// complex64 complex128

var(
	MaxInt 	uint64 		= 1<<64 - 1
	z		complex128	= cmplx.Sqrt(-5 + 12i)

)

func diffTypes(){
	fmt.Printf("Type: %T Value %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value %v \n", z, z)
}


// Initial value of variables

func initValues(){


	var i int
	var f float64
	var s string
	var b bool

	fmt.Printf("%v %v %v %v \n", i, f, s, b)

}


// type conversions

func typeConversions(){
	i := 42
	f := float64(i)
	u := uint(f)

	fmt.Println(i, f, u)
}


// infer the type by value

func inferType(){
	i := 42;
	j := 42.2;
	k := 0.6827 + 7373.9i

	fmt.Printf("%T ~ %T ~ %T \n", i, j, k)
}

// constants

const Pi = 3.14
func constantDeclaration(){

	const World = "世界"
	fmt.Println("Happy ", World, Pi)

}

// numeric constants

const(
	Big = 1<< 100
	Small = Big >> 99
)

func attendNumericConstants(){
	//fmt.Println(needInt(Big))
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
}

func needInt(a int) int { return a * 10 + 1 }

func needFloat(a float64) float64{return a * 0.1 }