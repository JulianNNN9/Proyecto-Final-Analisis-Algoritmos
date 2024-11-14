package algoritmos

import (
	"sync"
)

// III4ParallelBlock realiza la multiplicación de matrices en bloques.
func III4ParallelBlock(A, B [][]int, bsize int) [][]int {
	size := len(A)
	C := make([][]int, size)
	for i := range C {
		C[i] = make([]int, size)
	}

	var wg sync.WaitGroup
	mutex := sync.Mutex{}

	// Función para realizar la multiplicación en bloques
	task := func(i1, j1, k1 int) {
		defer wg.Done()
		localResult := make([][]int, size)
		for i := range localResult {
			localResult[i] = make([]int, size)
		}

		for i := i1; i < min(i1+bsize, size); i++ {
			for j := j1; j < min(j1+bsize, size); j++ {
				for k := k1; k < min(k1+bsize, size); k++ {
					localResult[i][j] += A[i][k] * B[k][j]
				}
			}
		}

		// Bloquea el acceso a C para agregar los resultados del bloque
		mutex.Lock()
		for i := range localResult {
			for j := range localResult[i] {
				C[i][j] += localResult[i][j]
			}
		}
		mutex.Unlock()
	}

	// Crear go rutinas para cada bloque
	for i1 := 0; i1 < size; i1 += bsize {
		for j1 := 0; j1 < size; j1 += bsize {
			for k1 := 0; k1 < size; k1 += bsize {
				wg.Add(1)
				go task(i1, j1, k1)
			}
		}
	}

	// Espera a que todas las gorutinas terminen
	wg.Wait()

	return C
}
