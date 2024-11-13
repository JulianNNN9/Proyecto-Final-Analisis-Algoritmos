package algoritmos

// StrassenNaivStep realiza la multiplicación de matrices utilizando el algoritmo de Strassen
func StrassenNaivStep(A, B [][]float64, cantidadFilasMatrices, m int) [][]float64 {
	Result := make([][]float64, cantidadFilasMatrices)
	for i := range Result {
		Result[i] = make([]float64, cantidadFilasMatrices)
	}

	// Caso base
	if cantidadFilasMatrices <= m {
		return NaivStandard(A, B, cantidadFilasMatrices, cantidadFilasMatrices, cantidadFilasMatrices)
	}

	nuevoTamanio := cantidadFilasMatrices / 2

	// Inicializamos submatrices para A y B
	A11, A12, A21, A22 := crearSubmatrices(nuevoTamanio)
	B11, B12, B21, B22 := crearSubmatrices(nuevoTamanio)
	Result11, Result12, Result21, Result22 := crearSubmatrices(nuevoTamanio)

	// Inicializamos matrices auxiliares y ayudantes
	ayudante1, _, _, _ := crearSubmatrices(nuevoTamanio)
	ayudante2, _, _, _ := crearSubmatrices(nuevoTamanio)

	auxiliar1, _, _, _ := crearSubmatrices(nuevoTamanio)
	auxiliar2, _, _, _ := crearSubmatrices(nuevoTamanio)
	auxiliar3, _, _, _ := crearSubmatrices(nuevoTamanio)
	auxiliar4, _, _, _ := crearSubmatrices(nuevoTamanio)
	auxiliar5, _, _, _ := crearSubmatrices(nuevoTamanio)
	auxiliar6, _, _, _ := crearSubmatrices(nuevoTamanio)
	auxiliar7, _, _, _ := crearSubmatrices(nuevoTamanio)

	// Llenamos las submatrices de A y B
	llenarSubmatrices(A, B, A11, A12, A21, A22, B11, B12, B21, B22, nuevoTamanio)

	// Computamos las siete multiplicaciones necesarias
	ayudante1 = sumarMatrices(A11, A22)
	ayudante2 = sumarMatrices(B11, B22)
	auxiliar1 = StrassenNaivStep(ayudante1, ayudante2, nuevoTamanio, m)

	ayudante1 = sumarMatrices(A21, A22)
	auxiliar2 = StrassenNaivStep(ayudante1, B11, nuevoTamanio, m)

	ayudante1 = restarMatrices(B12, B22)
	auxiliar3 = StrassenNaivStep(A11, ayudante1, nuevoTamanio, m)

	ayudante1 = restarMatrices(B21, B11)
	auxiliar4 = StrassenNaivStep(A22, ayudante1, nuevoTamanio, m)

	ayudante1 = sumarMatrices(A11, A12)
	auxiliar5 = StrassenNaivStep(ayudante1, B22, nuevoTamanio, m)

	ayudante1 = restarMatrices(A21, A11)
	ayudante2 = sumarMatrices(B11, B12)
	auxiliar6 = StrassenNaivStep(ayudante1, ayudante2, nuevoTamanio, m)

	ayudante1 = restarMatrices(A12, A22)
	ayudante2 = sumarMatrices(B21, B22)
	auxiliar7 = StrassenNaivStep(ayudante1, ayudante2, nuevoTamanio, m)

	// Calculamos las submatrices de Result
	Result11 = sumarMatrices(sumarMatrices(restarMatrices(auxiliar1, auxiliar5), auxiliar7), auxiliar4)
	Result12 = sumarMatrices(auxiliar3, auxiliar5)
	Result21 = sumarMatrices(auxiliar2, auxiliar4)
	Result22 = sumarMatrices(sumarMatrices(restarMatrices(auxiliar1, auxiliar2), auxiliar3), auxiliar6)

	// Combinar los resultados en la matriz final
	combinarResultados(Result, Result11, Result12, Result21, Result22, nuevoTamanio)

	return Result
}

// Funciones de utilidad para crear, sumar, restar y combinar submatrices
func crearSubmatrices(tamanio int) ([][]float64, [][]float64, [][]float64, [][]float64) {
	a, b, c, d := make([][]float64, tamanio), make([][]float64, tamanio), make([][]float64, tamanio), make([][]float64, tamanio)
	for i := range a {
		a[i], b[i], c[i], d[i] = make([]float64, tamanio), make([]float64, tamanio), make([]float64, tamanio), make([]float64, tamanio)
	}
	return a, b, c, d
}

func llenarSubmatrices(A, B, A11, A12, A21, A22, B11, B12, B21, B22 [][]float64, tamanio int) {
	for i := 0; i < tamanio; i++ {
		for j := 0; j < tamanio; j++ {
			A11[i][j] = A[i][j]
			A12[i][j] = A[i][j+tamanio]
			A21[i][j] = A[i+tamanio][j]
			A22[i][j] = A[i+tamanio][j+tamanio]

			B11[i][j] = B[i][j]
			B12[i][j] = B[i][j+tamanio]
			B21[i][j] = B[i+tamanio][j]
			B22[i][j] = B[i+tamanio][j+tamanio]
		}
	}
}

func sumarMatrices(A, B [][]float64) [][]float64 {
	N := len(A)
	C := make([][]float64, N)
	for i := range C {
		C[i] = make([]float64, N)
		for j := 0; j < N; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

func restarMatrices(A, B [][]float64) [][]float64 {
	N := len(A)
	C := make([][]float64, N)
	for i := range C {
		C[i] = make([]float64, N)
		for j := 0; j < N; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
	return C
}

func combinarResultados(Result, C11, C12, C21, C22 [][]float64, tamanio int) {
	for i := 0; i < tamanio; i++ {
		for j := 0; j < tamanio; j++ {
			Result[i][j] = C11[i][j]
			Result[i][j+tamanio] = C12[i][j]
			Result[i+tamanio][j] = C21[i][j]
			Result[i+tamanio][j+tamanio] = C22[i][j]
		}
	}
}
