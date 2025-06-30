package main 

import(
	"fmt"
	"experimentoMatriz/matriz"
	"experimentoMatriz/determinante"
)

func main(){
	var ordem int 
	var mat [][]int 
	ordem=3
	mat=matriz.CriarMatrizVazia(ordem)
	matriz.IniciaMatrizRandomica(mat)
	fmt.Println("Matriz Randomica:")
	for _,linha:=range mat{
		fmt.Println(linha)
	determinante.ProcuradorDeZeros(mat)
	}
	determinante.Determinante(mat) 
}
