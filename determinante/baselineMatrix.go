package determinante 

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
	
