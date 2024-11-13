package algoritmos

// NaivLoopUnrollingTwo realiza la multiplicación de matrices con desenrollado de bucles de tamaño dos.
func NaivLoopUnrollingTwo(A, B [][]float64, N, P, M int) [][]float64 {
	// Inicializamos la matriz Result con ceros
	Result := make([][]float64, N)
	for i := range Result {
		Result[i] = make([]float64, M)
	}

	if P%2 == 0 {
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				aux := 0.0
				for k := 0; k < P; k += 2 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j]
				}
				Result[i][j] = aux
			}
		}
	} else {
		PP := P - 1
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				aux := 0.0
				for k := 0; k < PP; k += 2 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j]
			}
		}
	}

	return Result
}
