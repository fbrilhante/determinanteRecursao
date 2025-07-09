package determinante 

import "fmt"

func DetOrdemN(mat [][]int) int{
	var sinal,cofator,detTemp,resposta,contL,contC,numL,numC,cont int
	var matMenor [][]int
	numL = len(mat) 
	numC = len(mat[0])
	
	resposta = 0;
	contL = 0;
	for contC = 0; contC < numC; contC++ {
		cofator = mat[contL][contC]
		sinal = CalculaSinal(contL,contC)
		//criando a matriz menor
		matMenor = make([][]int, numL-1)
		for cont = 0; cont < (numL-1); cont++ {
    			matMenor[cont] = make([]int, numC-1)
		}
		
		CopiaMatrizMaiorParaMenor(mat,matMenor,contL,contC)
		detTemp = Determinante(matMenor);
		//fmt.Println("DetTemp ",detTemp)
		//fmt.Println("resposta ",resposta)
		//fmt.Println("cofator ",cofator)
		//fmt.Println("sinal ",sinal)
		resposta = resposta + (cofator * sinal * detTemp)
		//fmt.Println("resposta dps",resposta)
	}
	
	return resposta
}

func Determinante(mat [][]int) int{
	var ordem int
	var ehQuadrada bool
	var det int

	ehQuadrada, ordem = VerificaQuadradaOrdem(mat)
	det = 0
	if(ehQuadrada){		
		switch (ordem) {
		    case 1:
		       // fmt.Println("Ordem 1")
		    	det = DetOrdem1(mat)
		    case 2:
		    	//fmt.Println("Ordem 2")
		    	det = DetOrdem2(mat)
		    default: 
		        //fmt.Println("Ordem ", ordem)
		    	det = DetOrdemN(mat)
			
		}
		//matriz.IniciaMatrizRandomica(mat)
		//fmt.Println("Det ", det)
		
	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}
	return det
}

	
