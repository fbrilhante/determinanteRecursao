package benchmark

import (
	"fmt"
	"time"
	"experimentoMatriz/matriz"
	"experimentoMatriz/determinante"
)


type ResultadosBenchmark struct {
	Ordens        []int
	TempoBaseline []int
	TempoReforged []int
}

func TempoExecucao()ResultadosBenchmark{
	var ordens []int 
	var matrizOriginal [][]int
	var copiaBaseline [][]int
	var copiaReforged [][]int 
	var tempoBaseline []int
	var tempoReforged []int
	var tempoExperimento int
	var detBaseline int 
	var detOtimizada int 
	var contOrdem int 
	var contRepeticoes int 
	var repeticoesTotais int
	var ordem int 
	var inicio time.Time
	var duracao time.Duration
	
	
	
	ordens=[]int{3,5,7,9,11}
	//ordens=[]int{3}
	repeticoesTotais=3
	tempoBaseline=make([]int,len(ordens))
	tempoReforged=make([]int,len(ordens))
	
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
			detBaseline=determinante.Determinante(copiaBaseline)
			duracao=time.Since(inicio)
			tempoExperimento=int(duracao.Nanoseconds())
			tempoBaseline[contOrdem] += tempoExperimento
			fmt.Printf("Determinante (baseline): %d\n",detBaseline)
			//prime
			inicio=time.Now()
			detOtimizada=determinante.DeterminanteReforged(copiaReforged)
			duracao=time.Since(inicio)
			tempoExperimento=int(duracao.Nanoseconds())
			tempoReforged[contOrdem]+=tempoExperimento
			fmt.Printf("Determinante (otimizada): %d\n",detOtimizada)
			fmt.Println()
		}
		//media
		tempoBaseline[contOrdem]=tempoBaseline[contOrdem]/repeticoesTotais
		tempoReforged[contOrdem]=tempoReforged[contOrdem]/repeticoesTotais
	}
	fmt.Println("Resultados médios em nanosegundos:")
	for contOrdem=0;contOrdem<len(ordens);contOrdem++{
	fmt.Printf("Ordem: %d | Baseline: %d ns | Reforged: %d ns\n" ,ordens[contOrdem],tempoBaseline[contOrdem],tempoReforged[contOrdem])
	}

	return ResultadosBenchmark{
		Ordens: ordens,
		TempoBaseline: tempoBaseline,
		TempoReforged: tempoReforged,
	}
}
			
