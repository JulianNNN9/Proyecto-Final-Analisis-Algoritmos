package algoritmos

// WinogradOriginal realiza la multiplicación de matrices usando el algoritmo de Winograd.
func WinogradOriginal(A, B, Result [][]float64, N, P, M int) [][]float64 {
	upsilon := P % 2
	gamma := P - upsilon

	y := make([]float64, M)
	z := make([]float64, N)

	// Calcular el vector y
	for i := 0; i < M; i++ {
		aux := 0.0
		for j := 0; j < gamma; j += 2 {
			aux += A[i][j] * A[i][j+1]
		}
		y[i] = aux
	}

	// Calcular el vector z
	for i := 0; i < N; i++ {
		aux := 0.0
		for j := 0; j < gamma; j += 2 {
			aux += B[j][i] * B[j+1][i]
		}
		z[i] = aux
	}

	// Realizar la multiplicación dependiendo del valor de upsilon
	if upsilon == 1 {
		PP := P - 1
		for i := 0; i < M; i++ {
			for k := 0; k < N; k++ {
				aux := 0.0
				for j := 0; j < gamma; j += 2 {
					aux += (A[i][j] + B[j+1][k]) * (A[i][j+1] + B[j][k])
				}
				Result[i][k] = aux - y[i] - z[k] + A[i][PP]*B[PP][k]
			}
		}
	} else {
		for i := 0; i < M; i++ {
			for k := 0; k < N; k++ {
				aux := 0.0
				for j := 0; j < gamma; j += 2 {
					aux += (A[i][j] + B[j+1][k]) * (A[i][j+1] + B[j][k])
				}
				Result[i][k] = aux - y[i] - z[k]
			}
		}
	}

	return Result
}
