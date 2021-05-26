package main

import "fmt"

func main() {
	fmt.Println("Float negativos/positivos <= 0 e >= 0")
	var numFloat32 float32 = 32
	var numFloat64 float64 = 64

	fmt.Printf("\tVariável numFloat32    : Tipo %T    Valor %v \n", numFloat32, numFloat32)
	fmt.Printf("\tVariável numFloat64    : Tipo %T    Valor %v \n", numFloat64, numFloat64)

	fmt.Println("Complex negativos/positivos <= 0 e >= 0")
	var numComplex64 complex64 = 64
	var numComplex128 complex128 = 128

	fmt.Printf("\tVariável numComplex64  : Tipo %T    Valor %v \n", numComplex64, numComplex64)
	fmt.Printf("\tVariável numComplex128 : Tipo %T    Valor %v \n", numComplex128, numComplex128)
}
