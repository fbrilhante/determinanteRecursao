package determinante 

//import "fmt"

func VerificaQuadradaOrdem(mat [][]int) (bool, int) {
	var numLinhas int
	var numColunas int
	var ehQuadrada bool
	
	numLinhas = len(mat) 
	numColunas = len(mat[0])
	
	ehQuadrada = false
	if (numLinhas == numColunas){
		ehQuadrada = true
	}
	
	return ehQuadrada,numLinhas
}

func  CalculaSinal(indiceL int, indiceC int) int{
	var sinal int

	sinal = -1
	if ((indiceL + indiceC)% 2) == 0 {
		sinal = 1
	}

	return sinal		
}

func CopiaMatrizMaiorParaMenor(maior [][]int, menor [][]int, isqn int, jsqn int){
	var contAi,contAj,contBi,contBj,temp,numL,numC int;
	numL = len(menor) 
	numC = len(menor[0])

	contAi = 0
	for contBi = 0; contBi < numL; contBi++ {
		if(contAi == isqn){
			contAi++
		}
		contAj = 0
		for contBj = 0; contBj < numC; contBj++ {
			if(contAj == jsqn){
				contAj++ 
			}
			temp = maior[contAi][contAj]
			menor[contBi][contBj] = temp
			contAj++
		}
		contAi++
	}
}

func DetOrdem1(mat [][]int) int{
	return mat[0][0]
}

func DetOrdem2(mat [][]int) int{
	var diagonalP int
	var diagonalI int

	diagonalP = mat[0][0] * mat[1][1]		
	diagonalI = mat[1][0] * mat[0][1]		
	return (diagonalP - diagonalI)
}

