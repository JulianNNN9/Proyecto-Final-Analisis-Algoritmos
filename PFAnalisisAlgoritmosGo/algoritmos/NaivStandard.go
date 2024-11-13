package algoritmos

// NaivStandard realiza la multiplicación de matrices estándar de A y B, almacenando el resultado en Result.
func NaivStandard(A, B [][]float64, N, P, M int) [][]float64 {
	// Inicializamos la matriz Result con ceros
	Result := make([][]float64, N)
	for i := range Result {
		Result[i] = make([]float64, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			aux := 0.0
			for k := 0; k < P; k++ {
				aux += A[i][k] * B[k][j]
			}
			Result[i][j] = aux
		}
	}

	return Result
}
