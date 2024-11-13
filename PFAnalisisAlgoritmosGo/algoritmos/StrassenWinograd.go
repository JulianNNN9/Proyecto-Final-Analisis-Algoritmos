package algoritmos

import (
	"math"
)

// StrassenWinograd realiza la multiplicación de matrices utilizando el algoritmo de Strassen-Winograd.
func StrassenWinograd(A, B, Result [][]float64, N, P, M int) [][]float64 {

	// Obtener el tamaño máximo entre las dimensiones de las matrices
	tamanioMaximo := int(math.Max(float64(N), math.Max(float64(M), float64(P))))
	if tamanioMaximo < 16 {
		tamanioMaximo = 16
	}

	// Calcular los valores de k, m y nuevoTamanio
	k := int(math.Floor(math.Log2(float64(tamanioMaximo)))) - 4
	m := int(math.Floor(float64(tamanioMaximo)*math.Pow(2, -float64(k)))) + 1
	nuevoTamanio := int(float64(m) * math.Pow(2, float64(k)))

	// Inicializar las nuevas matrices A y B y la matriz de resultado auxiliar
	nuevaMatrizA := make([][]float64, nuevoTamanio)
	nuevaMatrizB := make([][]float64, nuevoTamanio)
	matrizResultadoAuxiliar := make([][]float64, nuevoTamanio)

	for i := 0; i < nuevoTamanio; i++ {
		nuevaMatrizA[i] = make([]float64, nuevoTamanio)
		nuevaMatrizB[i] = make([]float64, nuevoTamanio)
		matrizResultadoAuxiliar[i] = make([]float64, nuevoTamanio)
	}

	// Asignar valores de las matrices originales a las nuevas matrices
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			nuevaMatrizA[i][j] = A[i][j]
		}
	}
	for i := 0; i < M; i++ {
		for j := 0; j < P; j++ {
			nuevaMatrizB[i][j] = B[i][j]
		}
	}

	// Realizar la multiplicación utilizando el paso de Strassen-Winograd
	matrizResultadoAuxiliar = StrassenWinogradStep(nuevaMatrizA, nuevaMatrizB, matrizResultadoAuxiliar, nuevoTamanio, m)

	// Asignar el resultado calculado a la matriz Result
	for i := 0; i < N; i++ {
		for j := 0; j < P; j++ {
			Result[i][j] = matrizResultadoAuxiliar[i][j]
		}
	}

	return Result
}
