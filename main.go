package main 

import(
	//"fmt"
	//"experimentoMatriz/matriz"
	//"experimentoMatriz/determinante"
	"experimentoMatriz/benchmark"
	"experimentoMatriz/grafico"
)

func main(){
	//var ordem int 
	//var mat [][]int  
	//ordem=3
	//mat=matriz.CriarMatrizVazia(ordem)
	//matriz.IniciaMatrizRandomica(mat)
	//fmt.Println("Matriz Randomica:")
	//matriz.ImprimeMatriz(mat)
	//fmt.Println("determinante baseline: ")
	//determinante.Determinante(mat)
	//fmt.Println("esse e o valor da determinante bem mais rapido: ")
	//determinante.DeterminanteReforged(mat)
	var resultados benchmark.ResultadosBenchmark
	resultados=benchmark.TempoExecucao()
	grafico.GerarGraficos(resultados)
}
