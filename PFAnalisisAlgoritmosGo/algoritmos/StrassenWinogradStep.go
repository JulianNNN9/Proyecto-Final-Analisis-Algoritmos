package algoritmos

// StrassenWinogradStep realiza la multiplicación de matrices con el método de Strassen Winograd.
func StrassenWinogradStep(A, B, Result [][]float64, cantidadFilasMatrices, m int) [][]float64 {
	nuevoTamanio := 0
	if cantidadFilasMatrices%2 == 0 && cantidadFilasMatrices > m {
		nuevoTamanio = cantidadFilasMatrices / 2

		// Inicialización de matrices
		var matrizA11, matrizA12, matrizA21, matrizA22 [][]float64
		var matrizB11, matrizB12, matrizB21, matrizB22 [][]float64
		var matrizA1, matrizA2, matrizB1, matrizB2 [][]float64
		var matrizResultadoParte11, matrizResultadoParte12, matrizResultadoParte21, matrizResultadoParte22 [][]float64
		var ayudante1, ayudante2 [][]float64
		var auxiliar1, auxiliar2, auxiliar3, auxiliar4, auxiliar5, auxiliar6, auxiliar7 [][]float64
		var auxiliar8, auxiliar9 [][]float64

		// Asignar memoria para cada fila
		for i := 0; i < nuevoTamanio; i++ {
			matrizA11 = append(matrizA11, make([]float64, nuevoTamanio))
			matrizA12 = append(matrizA12, make([]float64, nuevoTamanio))
			matrizA21 = append(matrizA21, make([]float64, nuevoTamanio))
			matrizA22 = append(matrizA22, make([]float64, nuevoTamanio))

			matrizB11 = append(matrizB11, make([]float64, nuevoTamanio))
			matrizB12 = append(matrizB12, make([]float64, nuevoTamanio))
			matrizB21 = append(matrizB21, make([]float64, nuevoTamanio))
			matrizB22 = append(matrizB22, make([]float64, nuevoTamanio))

			matrizA1 = append(matrizA1, make([]float64, nuevoTamanio))
			matrizA2 = append(matrizA2, make([]float64, nuevoTamanio))
			matrizB1 = append(matrizB1, make([]float64, nuevoTamanio))
			matrizB2 = append(matrizB2, make([]float64, nuevoTamanio))

			matrizResultadoParte11 = append(matrizResultadoParte11, make([]float64, nuevoTamanio))
			matrizResultadoParte12 = append(matrizResultadoParte12, make([]float64, nuevoTamanio))
			matrizResultadoParte21 = append(matrizResultadoParte21, make([]float64, nuevoTamanio))
			matrizResultadoParte22 = append(matrizResultadoParte22, make([]float64, nuevoTamanio))

			ayudante1 = append(ayudante1, make([]float64, nuevoTamanio))
			ayudante2 = append(ayudante2, make([]float64, nuevoTamanio))

			auxiliar1 = append(auxiliar1, make([]float64, nuevoTamanio))
			auxiliar2 = append(auxiliar2, make([]float64, nuevoTamanio))
			auxiliar3 = append(auxiliar3, make([]float64, nuevoTamanio))
			auxiliar4 = append(auxiliar4, make([]float64, nuevoTamanio))
			auxiliar5 = append(auxiliar5, make([]float64, nuevoTamanio))
			auxiliar6 = append(auxiliar6, make([]float64, nuevoTamanio))
			auxiliar7 = append(auxiliar7, make([]float64, nuevoTamanio))

			auxiliar8 = append(auxiliar8, make([]float64, nuevoTamanio))
			auxiliar9 = append(auxiliar9, make([]float64, nuevoTamanio))
		}

		// Llenamos las matrices
		for i := 0; i < nuevoTamanio; i++ {
			for j := 0; j < nuevoTamanio; j++ {
				matrizA11[i][j] = A[i][j]
				matrizA12[i][j] = A[i][nuevoTamanio+j]
				matrizA21[i][j] = A[nuevoTamanio+i][j]
				matrizA22[i][j] = A[nuevoTamanio+i][nuevoTamanio+j]

				matrizB11[i][j] = B[i][j]
				matrizB12[i][j] = B[i][nuevoTamanio+j]
				matrizB21[i][j] = B[nuevoTamanio+i][j]
				matrizB22[i][j] = B[nuevoTamanio+i][nuevoTamanio+j]
			}
		}

		// Computing the seven aux. variables
		matrizA1 = subtractMatrices(matrizA11, matrizA21, nuevoTamanio)
		matrizA2 = subtractMatrices(matrizA22, matrizA1, nuevoTamanio)

		matrizB1 = subtractMatrices(matrizB22, matrizB12, nuevoTamanio)
		matrizB2 = addMatrices(matrizB1, matrizB11, nuevoTamanio)

		auxiliar1 = StrassenWinogradStep(matrizA11, matrizB11, auxiliar1, nuevoTamanio, m)
		auxiliar2 = StrassenWinogradStep(matrizA12, matrizB21, auxiliar2, nuevoTamanio, m)
		auxiliar3 = StrassenWinogradStep(matrizA2, matrizB2, auxiliar3, nuevoTamanio, m)

		ayudante1 = addMatrices(matrizA21, matrizA22, nuevoTamanio)
		ayudante2 = subtractMatrices(matrizB12, matrizB11, nuevoTamanio)

		auxiliar4 = StrassenWinogradStep(ayudante1, ayudante2, auxiliar4, nuevoTamanio, m)

		auxiliar5 = StrassenWinogradStep(matrizA1, matrizB1, auxiliar5, nuevoTamanio, m)

		ayudante1 = subtractMatrices(matrizA12, matrizA2, nuevoTamanio)

		auxiliar6 = StrassenWinogradStep(ayudante1, matrizB22, auxiliar6, nuevoTamanio, m)

		ayudante1 = subtractMatrices(matrizB21, matrizB2, nuevoTamanio)

		auxiliar7 = StrassenWinogradStep(matrizA22, ayudante1, auxiliar7, nuevoTamanio, m)

		auxiliar8 = addMatrices(auxiliar1, auxiliar3, nuevoTamanio)
		auxiliar9 = addMatrices(auxiliar8, auxiliar4, nuevoTamanio)

		// Calcular partes de la matriz resultado
		matrizResultadoParte11 = addMatrices(auxiliar1, auxiliar2, nuevoTamanio)
		matrizResultadoParte12 = addMatrices(auxiliar9, auxiliar6, nuevoTamanio)
		ayudante1 = addMatrices(auxiliar8, auxiliar5, nuevoTamanio)
		matrizResultadoParte21 = addMatrices(ayudante1, auxiliar7, nuevoTamanio)
		matrizResultadoParte22 = addMatrices(auxiliar9, auxiliar5, nuevoTamanio)

		// Almacenar resultados en la matriz resultado
		for i := 0; i < nuevoTamanio; i++ {
			for j := 0; j < nuevoTamanio; j++ {
				Result[i][j] = matrizResultadoParte11[i][j]
			}
		}

		for i := 0; i < nuevoTamanio; i++ {
			for j := 0; j < nuevoTamanio; j++ {
				Result[i][nuevoTamanio+j] = matrizResultadoParte12[i][j]
			}
		}

		for i := 0; i < nuevoTamanio; i++ {
			for j := 0; j < nuevoTamanio; j++ {
				Result[nuevoTamanio+i][j] = matrizResultadoParte21[i][j]
			}
		}

		for i := 0; i < nuevoTamanio; i++ {
			for j := 0; j < nuevoTamanio; j++ {
				Result[nuevoTamanio+i][nuevoTamanio+j] = matrizResultadoParte22[i][j]
			}
		}
	} else {
		// Usar el algoritmo naiv
		Result = NaivStandard(A, B, len(A), len(B[0]), len(Result[0]))
	}

	return Result
}

// Funciones auxiliares para operaciones de matrices
func addMatrices(A, B [][]float64, size int) [][]float64 {
	result := make([][]float64, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			result[i][j] = A[i][j] + B[i][j]
		}
	}
	return result
}

func subtractMatrices(A, B [][]float64, size int) [][]float64 {
	result := make([][]float64, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			result[i][j] = A[i][j] - B[i][j]
		}
	}
	return result
}
