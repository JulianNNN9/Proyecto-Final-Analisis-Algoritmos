package algoritmos

// IV3SequentialBlock realiza la multiplicación de matrices en bloques secuencialmente.
func IV3SequentialBlock(A, B [][]int, tamanioBloque int) [][]int {
	size := len(A)
	Result := make([][]int, size)
	for i := range Result {
		Result[i] = make([]int, size)
	}

	for i1 := 0; i1 < size; i1 += tamanioBloque {
		for j1 := 0; j1 < size; j1 += tamanioBloque {
			for k1 := 0; k1 < size; k1 += tamanioBloque {
				for i := i1; i < min(i1+tamanioBloque, size); i++ {
					for j := j1; j < min(j1+tamanioBloque, size); j++ {
						for k := k1; k < min(k1+tamanioBloque, size); k++ {
							Result[i][k] += A[i][j] * B[j][k]
						}
					}
				}
			}
		}
	}

	return Result
}
