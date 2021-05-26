package main

import "fmt"

// https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go-pt#:~:text=Um%20int8%20%C3%A9%20um%20n%C3%BAmero,positivos%20entre%200%20a%20255.
func main() {
	fmt.Println("Inteiros positivos >= 0 ")
	var uNumInt8 uint8 = 8
	var uNumByte byte = 8
	var uNumInt16 uint16 = 16
	var uNumInt32 uint32 = 32
	var uNumInt64 uint64 = 64
	var uNumInt uint = 3264

	fmt.Printf("\tVariável uNumInt8 : Tipo %T	Valor %v \n", uNumInt8, uNumInt8)
	fmt.Printf("\tVariável uNumByte : Tipo %T	Valor %v \n", uNumByte, uNumByte)
	fmt.Printf("\tVariável uNumInt16: Tipo %T	Valor %v \n", uNumInt16, uNumInt16)
	fmt.Printf("\tVariável uNumInt32: Tipo %T	Valor %v \n", uNumInt32, uNumInt32)
	fmt.Printf("\tVariável uNumInt64: Tipo %T	Valor %v \n", uNumInt64, uNumInt64)
	fmt.Printf("\tVariável uNumInt  : Tipo %T	Valor %v \n", uNumInt, uNumInt)

	fmt.Println("\nInteiros negativos/positivos <= 0 e >= 0")
	var numInt8 int8 = 8
	var numInt16 int16 = 16
	var numInt32 int32 = 32
	var numRune rune = 32
	var numInt64 int64 = 64
	var numInt int = 3264

	fmt.Printf("\tVariável numInt8 : Tipo %T	Valor %v \n", numInt8, numInt8)
	fmt.Printf("\tVariável numInt16: Tipo %T	Valor %v \n", numInt16, numInt16)
	fmt.Printf("\tVariável numInt32: Tipo %T	Valor %v \n", numInt32, numInt32)
	fmt.Printf("\tVariável numRune : Tipo %T	Valor %v \n", numRune, numRune)
	fmt.Printf("\tVariável numInt64: Tipo %T	Valor %v \n", numInt64, numInt64)
	fmt.Printf("\tVariável numInt  : Tipo %T	Valor %v \n", numInt, numInt)
}
