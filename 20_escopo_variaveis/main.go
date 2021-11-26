package main

import "fmt"

var nome string = "fulano"

func main() {
	nome += " da silva"

	var nome string

	fmt.Println(nome)
}

/* var numPkg int = 53453
var floatPkg float32 = 4324324.044
var choveuPkg bool = false
var nomePkg string = "fulano"

func main() {
	fmt.Println("Variaveis com escopo de pacote")
	fmt.Printf("\tVariável numPkg    : Tipo %T	Valor %v \n", numPkg, numPkg)
	fmt.Printf("\tVariável floatPkg  : Tipo %T	Valor %v \n", numPkg, numPkg)
	fmt.Printf("\tVariável choveuPkg : Tipo %T	Valor %v \n", numPkg, numPkg)
	fmt.Printf("\tVariável nomePkg   : Tipo %T	Valor %v \n", numPkg, numPkg)

	fmt.Println("\nVariaveis com escopo de funcao")
	var num int = 243
	var flo float32 = 243.002
	var choveu bool = true
	var nome string = "fulana"
	fmt.Printf("\tVariável num    : Tipo %T	Valor %v \n", num, num)
	fmt.Printf("\tVariável float  : Tipo %T	Valor %v \n", flo, flo)
	fmt.Printf("\tVariável choveu : Tipo %T	Valor %v \n", choveu, choveu)
	fmt.Printf("\tVariável nome   : Tipo %T	Valor %v \n", nome, nome)
} */
