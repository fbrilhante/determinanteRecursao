package determinante 

import "fmt"

func ProcuradorDeZeros(mat [][]int)(int,int){
	var ZeroPorLinha []int 
	var ZeroPorColuna []int 
	var contI int 
	var contJ int 
	var maxLinha int 
	var maxColuna int 
	
	ZeroPorLinha=make([]int,len(mat))
	ZeroPorColuna=make([]int,len(mat[0]))
	//contar zeros, linha e coluna 
	for contI=0;contI<len(mat);contI++{	
		for contJ=0;contJ<len(mat[0]);contJ++{
			if mat[contI][contJ]==0{
				ZeroPorLinha[contI]++
				ZeroPorColuna[contJ]++
			}	
		}
	}
	//encontrando linha com mais zero 
	maxLinha=0
	for contI=1;contI<len(mat);contI++{
		if ZeroPorLinha[contI]>ZeroPorLinha[maxLinha]{
			maxLinha=contI
		}
	}
	//encontrando coluna com mais zero 
	maxColuna=0
		for contJ=1;contJ<len(mat[0]);contJ++{
			if ZeroPorColuna[contJ]>ZeroPorColuna[maxColuna]{
				maxColuna=contJ 
			}
		}
		return maxLinha,maxColuna	
}


func DetOrdemNReforged(mat [][]int) int{
	var sinal,cofator,detTemp,resposta,contL,contC,numL,numC,cont,maisZeroL,maisZeroC int
	var matMenor [][]int
	maisZeroL,maisZeroC = ProcuradorDeZeros(mat)
	numL = len(mat) 
	numC = len(mat[0])

	if maisZeroL>=maisZeroC {
		fmt.Printf("O indice da linha com mais zeros é %d \n",maisZeroL)
		for contC=0;contC<numC;contC++{
			cofator = mat[maisZeroL][contC]
			if cofator != 0 {
				sinal= CalculaSinal(maisZeroL,contC)

				//criar matriz menor para o det - se cofator for 0 o det vai ser 0 entao nao precisa 
				matMenor=make([][]int,numL-1)
				for cont=0;cont<numL-1;cont++{
					matMenor[cont]=make([]int,numC-1)
				}

				CopiaMatrizMaiorParaMenor(mat,matMenor,maisZeroL,contC)
				detTemp=Determinante(matMenor)
				resposta=resposta+(cofator*sinal*detTemp)
			}
		}
	}else{
		fmt.Printf("o Indice da colunas com mais zeros é %d \n",maisZeroC)
		for contL=0;contL<numL;contL++{
			cofator=mat[contL][maisZeroC]
			if cofator !=0{
				sinal = CalculaSinal(contL,maisZeroC)

				//criar matriz menor 
				matMenor=make([][]int,numL-1)
				for cont=0;cont<numL-1;cont++{
					matMenor[cont]=make([]int,numC-1)
				}
				CopiaMatrizMaiorParaMenor(mat,matMenor,contL,maisZeroC)
				detTemp=DeterminanteReforged(matMenor)
				resposta=resposta+(cofator*sinal*detTemp)
			}
		}
	}
	return resposta
}

func DeterminanteReforged(mat [][]int) int{
	var ordem int
	var ehQuadrada bool
	var det int

	ehQuadrada, ordem = VerificaQuadradaOrdem(mat)
	det = 0
	if(ehQuadrada){		
		switch (ordem) {
		    case 1:
		        fmt.Println("Ordem 1")
		    	det = DetOrdem1(mat)
		    case 2:
		    	fmt.Println("Ordem 2")
		    	det = DetOrdem2(mat)
		    default: 
		        fmt.Println("Ordem ", ordem)
		    	det = DetOrdemNReforged(mat)
			
		}
		//matriz.IniciaMatrizRandomica(mat)
		fmt.Println("Det ", det)
		
	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}
	return det
}










		//sinal = CalculaSinal(maisZeroL,contC)
	//criar matriz menor sem linha maisZeroL e coluna contC
		/* matMenor=make([][]int,numL-1)
		for cont=0;cont<numL-1;cont++{
			matMenor[cont]=make([]int,numC-1)
		}
	CopaMatrizMaiorParaMenor(mat,matMenor,MaisZeroL,contJ)
	detTemp=Determinante(matMenor)
	resposta=resposta + (cofator*sinal*detTemp)
	}
	return resposta*/

	

