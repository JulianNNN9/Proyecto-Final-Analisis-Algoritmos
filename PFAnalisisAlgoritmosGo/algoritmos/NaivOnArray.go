package algoritmos

// NaivOnArray realiza la multiplicación de matrices básica A * B y almacena el resultado en Result.
func NaivOnArray(A, B [][]float64, N, P, M int) [][]float64 {
	// Inicializamos la matriz Result con ceros
	Result := make([][]float64, N)
	for i := range Result {
		Result[i] = make([]float64, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			Result[i][j] = 0.0
			for k := 0; k < P; k++ {
				Result[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return Result
}
