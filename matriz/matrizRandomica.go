package matriz 

import(
	"math/rand"
	"time"
)

func CriarMatrizVazia(ordem int)[][]int{
	var cont int 
	var mat [][]int 
	mat=make([][]int,ordem)
	for cont=0;cont<ordem;cont++{
		mat[cont]=make([]int,ordem)
	}
	return mat
}
func IniciaMatrizRandomica(mat [][]int){
	var contI int 
	var contJ int
	var limite int 
	rand.Seed(time.Now().UnixNano())
	limite = (len(mat)*len(mat))/2
	for contI=0;contI<len(mat);contI++{
		for contJ=0;contJ<len(mat);contJ++{
			mat[contI][contJ]=rand.Intn(limite+1)
		}
	}
}


