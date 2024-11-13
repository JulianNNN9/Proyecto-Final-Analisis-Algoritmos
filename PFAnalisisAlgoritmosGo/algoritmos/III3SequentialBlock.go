package algoritmos

// III3SequentialBlockMet realiza la operación de multiplicación de matrices con bloques.
func III3SequentialBlock(A, B, Result [][]int, tamanioBloque int) [][]int {
	size := len(Result)

	for i1 := 0; i1 < size; i1 += tamanioBloque {
		for j1 := 0; j1 < size; j1 += tamanioBloque {
			for k1 := 0; k1 < size; k1 += tamanioBloque {
				for i := i1; i < min(i1+tamanioBloque, size); i++ {
					for j := j1; j < min(j1+tamanioBloque, size); j++ {
						for k := k1; k < min(k1+tamanioBloque, size); k++ {
							Result[i][j] += A[i][k] * B[k][j]
						}
					}
				}
			}
		}
	}
	return Result
}
