package benchmark

import (
	"fmt"
	"time"
	"experimentoMatriz/matriz"
	"experimentoMatriz/determinante"
)

func TempoExecucao(){
	var ordens []int 
	var matrizOriginal [][]int
	var copiaBaseline [][]int
	var copiaReforged [][]int 
	var tempoBaseline []int64
	var tempoReforged []int64
	var tempoExperimento int64
	var contOrdem int 
	var contRepeticoes int 
	var repeticoesTotais int
	var ordem int 
	var inicio time.Time
	var duracao time.Duration
	
	
	
	//ordens=[]int{3,5,7,9,11}
	ordens=[]int{3}
	repeticoesTotais=3
	tempoBaseline=make([]int64,len(ordens))
	tempoReforged=make([]int64,len(ordens))
	
	for contOrdem=0;contOrdem<len(ordens);contOrdem++{
		ordem=ordens[contOrdem]
		for contRepeticoes=0;contRepeticoes<repeticoesTotais;contRepeticoes++{
			matrizOriginal=matriz.CriarMatrizVazia(ordem)
			matriz.IniciaMatrizRandomica(matrizOriginal)
			fmt.Printf("\nMatriz Randomica(Repeticao %d):\n",contRepeticoes+1)
			matriz.ImprimeMatriz(matrizOriginal)
			copiaBaseline=matriz.CopiarMatriz(matrizOriginal)
			copiaReforged=matriz.CopiarMatriz(matrizOriginal)
			
			//baseline
			inicio=time.Now()
			determinante.Determinante(copiaBaseline)
			duracao=time.Since(inicio)
			tempoExperimento=duracao.Nanoseconds()
			tempoBaseline[contOrdem] += tempoExperimento
			
			//prime
			inicio=time.Now()
			determinante.DeterminanteReforged(copiaReforged)
			duracao=time.Since(inicio)
			tempoExperimento=duracao.Nanoseconds()
			tempoReforged[contOrdem]+=tempoExperimento
		}
		//media
		tempoBaseline[contOrdem]=tempoBaseline[contOrdem]/int64(repeticoesTotais)
		tempoReforged[contOrdem]=tempoReforged[contOrdem]/int64(repeticoesTotais)
	}
	fmt.Println("Resultados médios em nanosegundos:")
	for contOrdem=0;contOrdem<len(ordens);contOrdem++{
	fmt.Printf("Ordem: %d | Baseline: %d ns | Reforged: %d ns\n" ,ordens[contOrdem],tempoBaseline[contOrdem],tempoReforged[contOrdem])
	}
}
			
