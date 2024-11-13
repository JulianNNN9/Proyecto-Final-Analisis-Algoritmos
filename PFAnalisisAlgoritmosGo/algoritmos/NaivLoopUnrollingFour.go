package algoritmos

// NaivLoopUnrollingFour realiza la multiplicación de matrices con desenrollado de bucles de tamaño cuatro.
func NaivLoopUnrollingFour(A, B [][]float64, N, P, M int) [][]float64 {
	// Inicializamos la matriz Result con ceros
	Result := make([][]float64, N)
	for i := range Result {
		Result[i] = make([]float64, M)
	}

	if P%4 == 0 {
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				aux := 0.0
				for k := 0; k < P; k += 4 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
				}
				Result[i][j] = aux
			}
		}
	} else if P%4 == 1 {
		PP := P - 1
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				aux := 0.0
				for k := 0; k < PP; k += 4 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j]
			}
		}
	} else if P%4 == 2 {
		PP := P - 2
		PPP := P - 1
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				aux := 0.0
				for k := 0; k < PP; k += 4 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j] + A[i][PPP]*B[PPP][j]
			}
		}
	} else {
		PP := P - 3
		PPP := P - 2
		PPPP := P - 1
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				aux := 0.0
				for k := 0; k < PP; k += 4 {
					aux += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j] + A[i][k+2]*B[k+2][j] + A[i][k+3]*B[k+3][j]
				}
				Result[i][j] = aux + A[i][PP]*B[PP][j] + A[i][PPP]*B[PPP][j] + A[i][PPPP]*B[PPPP][j]
			}
		}
	}

	return Result
}
