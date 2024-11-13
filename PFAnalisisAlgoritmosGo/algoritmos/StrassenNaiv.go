package algoritmos

import (
	"math"
)

// StrassenNaiv realiza la multiplicación de matrices usando el algoritmo de Strassen con un paso base de tamaño "m".
func StrassenNaiv(A, B, Result [][]float64, N, P, M int) [][]float64 {
	tamanioMaximo := int(math.Max(float64(N), math.Max(float64(P), float64(M))))

	if tamanioMaximo < 16 {
		tamanioMaximo = 16
	}

	k := int(math.Floor(math.Log2(float64(tamanioMaximo))) - 4)
	m := int(math.Floor(float64(tamanioMaximo)*math.Pow(2, float64(-k)))) + 1
	nuevoTamanio := int(float64(m) * math.Pow(2, float64(k)))

	nuevaMatrizA := crearMatrizCuadrada(nuevoTamanio)
	nuevaMatrizB := crearMatrizCuadrada(nuevoTamanio)
	matrizResultadoAuxiliar := crearMatrizCuadrada(nuevoTamanio)

	// Asignamos en cada posición i,j de las nuevas matrices A y B los valores que están en las matrices A y B respectivamente
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

	// Calculamos la multiplicación con el método StrassenNaivStep
	matrizResultadoAuxiliar = StrassenNaivStep(nuevaMatrizA, nuevaMatrizB, nuevoTamanio, m)

	// Guardamos el resultado en la matriz Result
	for i := 0; i < N; i++ {
		for j := 0; j < P; j++ {
			Result[i][j] = matrizResultadoAuxiliar[i][j]
		}
	}

	return Result
}

// crearMatrizCuadrada crea una matriz cuadrada de tamaño `tamanio` inicializada con ceros.
func crearMatrizCuadrada(tamanio int) [][]float64 {
	matriz := make([][]float64, tamanio)
	for i := range matriz {
		matriz[i] = make([]float64, tamanio)
	}
	return matriz
}
