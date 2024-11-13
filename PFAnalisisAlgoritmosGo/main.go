package main

import (
	"PFAnalisisAlgoritmosGo/algoritmos" // Paquete de algoritmos
	"PFAnalisisAlgoritmosGo/utilidades" // Paquete de utilidades
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// LeerMatriz lee una matriz desde un archivo de texto
func LeerMatriz(ruta string, tam int) [][]int {
	// Leer el archivo
	data, err := ioutil.ReadFile(ruta)
	if err != nil {
		fmt.Println("Error al leer archivo:", err)
		return nil
	}

	// Convertir el contenido del archivo en una matriz
	lineas := strings.Split(string(data), "\n")
	matriz := make([][]int, tam)

	for i, linea := range lineas {
		if i >= tam {
			break
		}
		valores := strings.Fields(linea)
		matriz[i] = make([]int, tam)
		for j, valor := range valores {
			num, err := strconv.Atoi(valor)
			if err != nil {
				fmt.Println("Error al convertir el valor:", err)
				return nil
			}
			matriz[i][j] = num
		}
	}

	return matriz
}


func ejecutar_WinogradOriginal() {
	// Casos de prueba
	casosPrueba := []map[string]interface{}{
		{"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques": 0},
		{"num": 2, "tam": 16, "numMuestras": 2, "tamanioBloques": 0},
		{"num": 3, "tam": 32, "numMuestras": 2, "tamanioBloques": 0},
		{"num": 4, "tam": 64, "numMuestras": 2, "tamanioBloques": 0},
		{"num": 5, "tam": 128, "numMuestras": 2, "tamanioBloques": 0},
		{"num": 6, "tam": 256, "numMuestras": 2, "tamanioBloques": 0},
		{"num": 7, "tam": 512, "numMuestras": 1, "tamanioBloques": 0},
		{"num": 8, "tam": 1024, "numMuestras": 1, "tamanioBloques": 0},
	}

	// Resultado
	resultado := utilidades.ResultadoMet{
		Nombre:   "WinogradOriginal",
		Casos:    []utilidades.CasoMet{},
		Lenguaje: "go",
	}

	// Procesar los casos de prueba
	for _, caso := range casosPrueba {
		objeto := utilidades.CasoMet{
			Tam:            caso["tam"].(int),
			Muestras:       []float64{},
			Promedio:       0,
			TamanioBloques: 0, // No se usa en este caso
		}

		// Leer matrices desde los archivos correspondientes
		matrizA := LeerMatriz(fmt.Sprintf("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Casos de prueba/caso%d/matrizA.txt", caso["num"].(int)), objeto.Tam)
		matrizB := LeerMatriz(fmt.Sprintf("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Casos de prueba/caso%d/matrizB.txt", caso["num"].(int)), objeto.Tam)

		// Convertir las matrices de [][]int a [][]float64
		matrizAFloat := convertirAMatrizFloat64(matrizA)
		matrizBFloat := convertirAMatrizFloat64(matrizB)

		// Realizar el cálculo y medir el tiempo
		for i := 0; i < caso["numMuestras"].(int); i++ {
			N := len(matrizAFloat)
			P := len(matrizAFloat[0])
			M := len(matrizBFloat[0])

			// Inicializar la matriz de resultados
			result := make([][]float64, N)
			for i := range result {
				result[i] = make([]float64, M)
			}

			// Medir el tiempo de ejecución
			tiempoInicio := time.Now()
			algoritmos.WinogradOriginal(matrizAFloat, matrizBFloat, result, N, P, M) // Algoritmo de WinogradOriginal
			tiempoFinalizacion := time.Since(tiempoInicio)

			// Agregar el tiempo de ejecución a las muestras
			objeto.Muestras = append(objeto.Muestras, tiempoFinalizacion.Seconds())
		}

		// Calcular el promedio de las muestras
		objeto.Promedio = objeto.CalcularPromedio()

		// Agregar el objeto al resultado
		resultado.Casos = append(resultado.Casos, objeto)
	}

	// Guardar el resultado en un archivo JSON
	resultadoJSON, err := json.MarshalIndent(resultado.ToJSON(), "", "  ")
	if err != nil {
		fmt.Println("Error al generar JSON:", err)
		return
	}

	// Guardar en archivo
	err = os.WriteFile("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Resultados/go/WinogradOriginalResultadoGo.json", resultadoJSON, 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
	}
}

func convertirAMatrizFloat64(matriz [][]int) [][]float64 {
	N := len(matriz)
	M := len(matriz[0])
	result := make([][]float64, N)
	for i := range matriz {
		result[i] = make([]float64, M)
		for j := range matriz[i] {
			result[i][j] = float64(matriz[i][j])
		}
	}
	return result
}

func ejecutar_V3SequentialBlock() {
	// Casos de prueba
	casosPrueba := []map[string]interface{}{
		{"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques": 4},
		{"num": 2, "tam": 16, "numMuestras": 2, "tamanioBloques": 8},
		{"num": 3, "tam": 32, "numMuestras": 2, "tamanioBloques": 16},
		{"num": 4, "tam": 64, "numMuestras": 1, "tamanioBloques": 32},
		{"num": 5, "tam": 128, "numMuestras": 1, "tamanioBloques": 64},
		{"num": 6, "tam": 256, "numMuestras": 1, "tamanioBloques": 64},
		{"num": 7, "tam": 512, "numMuestras": 1, "tamanioBloques": 128},
		{"num": 8, "tam": 1024, "numMuestras": 1, "tamanioBloques": 128},
	}

	// Resultado
	resultado := utilidades.ResultadoMet{
		Nombre:   "V3SequentialBlock",
		Casos:    []utilidades.CasoMet{},
		Lenguaje: "go",
	}

	// Procesar los casos de prueba
	for _, caso := range casosPrueba {
		objeto := utilidades.CasoMet{
			Tam:            caso["tam"].(int),
			Muestras:       []float64{},
			Promedio:       0,
			TamanioBloques: caso["tamanioBloques"].(int),
		}

		// Leer matrices desde los archivos correspondientes
		matrizA := LeerMatriz(fmt.Sprintf("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Casos de prueba/caso%d/matrizA.txt", caso["num"].(int)), objeto.Tam)
		matrizB := LeerMatriz(fmt.Sprintf("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Casos de prueba/caso%d/matrizB.txt", caso["num"].(int)), objeto.Tam)

		// Convertir las matrices de [][]int a [][]float64
		matrizAFloat := convertirAMatrizFloat64(matrizA)
		matrizBFloat := convertirAMatrizFloat64(matrizB)

		// Realizar el cálculo y medir el tiempo
		for i := 0; i < caso["numMuestras"].(int); i++ {
			N := len(matrizAFloat)
			M := len(matrizBFloat[0])

			// Inicializar la matriz de resultados
			result := make([][]float64, N)
			for i := range result {
				result[i] = make([]float64, M)
			}

			// Medir el tiempo de ejecución
			tiempoInicio := time.Now()
			algoritmos.V3SequentialBlock(matrizAFloat, matrizBFloat, result, objeto.TamanioBloques)
			tiempoFinalizacion := time.Since(tiempoInicio)

			// Agregar el tiempo de ejecución a las muestras
			objeto.Muestras = append(objeto.Muestras, tiempoFinalizacion.Seconds())
		}

		// Calcular el promedio de las muestras
		objeto.Promedio = objeto.CalcularPromedio()

		// Agregar el objeto al resultado
		resultado.Casos = append(resultado.Casos, objeto)
	}

	// Guardar el resultado en un archivo JSON
	resultadoJSON, err := json.MarshalIndent(resultado.ToJSON(), "", "  ")
	if err != nil {
		fmt.Println("Error al generar JSON:", err)
		return
	}

	// Guardar en archivo
	err = os.WriteFile("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Resultados/go/V3SequentialBlockResultadoGo.json", resultadoJSON, 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
	}
}

func ejecutar_StrassenWinograd() {
	// Casos de prueba
	casosPrueba := []map[string]interface{}{
		{"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques": 0},
		{"num": 2, "tam": 16, "numMuestras": 2, "tamanioBloques": 0},
		{"num": 3, "tam": 32, "numMuestras": 2, "tamanioBloques": 0},
		{"num": 4, "tam": 64, "numMuestras": 2, "tamanioBloques": 0},
		{"num": 5, "tam": 128, "numMuestras": 1, "tamanioBloques": 0},
		{"num": 6, "tam": 256, "numMuestras": 1, "tamanioBloques": 0},
		{"num": 7, "tam": 512, "numMuestras": 1, "tamanioBloques": 0},
		{"num": 8, "tam": 1024, "numMuestras": 1, "tamanioBloques": 0},
	}

	// Resultado
	resultado := utilidades.ResultadoMet{
		Nombre:   "StrassenWinograd",
		Casos:    []utilidades.CasoMet{},
		Lenguaje: "go",
	}

	// Procesar los casos de prueba
	for _, caso := range casosPrueba {
		objeto := utilidades.CasoMet{
			Tam:            caso["tam"].(int),
			Muestras:       []float64{},
			Promedio:       0,
			TamanioBloques: caso["tamanioBloques"].(int),
		}

		// Leer matrices desde los archivos correspondientes
		matrizA := LeerMatriz(fmt.Sprintf("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Casos de prueba/caso%d/matrizA.txt", caso["num"].(int)), objeto.Tam)
		matrizB := LeerMatriz(fmt.Sprintf("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Casos de prueba/caso%d/matrizB.txt", caso["num"].(int)), objeto.Tam)

		// Convertir las matrices de [][]int a [][]float64
		matrizAFloat := convertirAMatrizFloat64(matrizA)
		matrizBFloat := convertirAMatrizFloat64(matrizB)

		// Realizar el cálculo y medir el tiempo
		for sample := 0; sample < caso["numMuestras"].(int); sample++ {
			// Obtener el tamaño de las matrices
			N := len(matrizAFloat)
			M := len(matrizBFloat[0])

			// Inicializar la matriz de resultados
			result := make([][]float64, N)
			for row := range result {
				result[row] = make([]float64, M)
			}

			// Medir el tiempo de ejecución
			tiempoInicio := time.Now()
			// Asumiendo que el parámetro adicional es algo como 'bloques'
			algoritmos.StrassenWinograd(matrizAFloat, matrizBFloat, result, N, M, objeto.TamanioBloques)

			tiempoFinalizacion := time.Since(tiempoInicio)

			// Agregar el tiempo de ejecución a las muestras
			objeto.Muestras = append(objeto.Muestras, tiempoFinalizacion.Seconds())
		}

		// Calcular el promedio de las muestras
		objeto.Promedio = objeto.CalcularPromedio()

		// Agregar el objeto al resultado
		resultado.Casos = append(resultado.Casos, objeto)
	}

	// Guardar el resultado en un archivo JSON
	resultadoJSON, err := json.MarshalIndent(resultado, "", "  ")
	if err != nil {
		fmt.Println("Error al generar JSON:", err)
		return
	}

	// Guardar en archivo
	err = os.WriteFile("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Resultados/go/StrassenWinogradResultadoGo.json", resultadoJSON, 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
	}
}


func ejecutar_StrassenNaiv() {

}

func ejecutar_NaivOnArray() {

}

func ejecutar_NaivLoopUnrollingTwo() {

}

func ejecutar_NaivLoopUnrollingFour() {

}

func ejecutar_IV3SequentialBlock() {
	casosPrueba := []map[string]interface{}{
		{"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques": 4},
		{"num": 2, "tam": 16, "numMuestras": 2, "tamanioBloques": 8},
		{"num": 3, "tam": 32, "numMuestras": 2, "tamanioBloques": 16},
		{"num": 4, "tam": 64, "numMuestras": 1, "tamanioBloques": 32},
		{"num": 5, "tam": 128, "numMuestras": 1, "tamanioBloques": 64},
		{"num": 6, "tam": 256, "numMuestras": 1, "tamanioBloques": 64},
		{"num": 7, "tam": 512, "numMuestras": 1, "tamanioBloques": 128},
		{"num": 8, "tam": 1024, "numMuestras": 1, "tamanioBloques": 128},
	}

	// Resultado
	resultado := utilidades.ResultadoMet{
		Nombre:   "IV3SequentialBlock",
		Casos:    []utilidades.CasoMet{},
		Lenguaje: "go",
	}

	// Procesar los casos de prueba
	for _, caso := range casosPrueba {
		objeto := utilidades.CasoMet{
			Tam:            caso["tam"].(int),
			Muestras:       []float64{},
			Promedio:       0,
			TamanioBloques: caso["tamanioBloques"].(int),
		}

		// Leer matrices desde los archivos correspondientes
		matrizA := LeerMatriz(fmt.Sprintf("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Casos de prueba/caso%d/matrizA.txt", caso["num"].(int)), objeto.Tam)
		matrizB := LeerMatriz(fmt.Sprintf("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Casos de prueba/caso%d/matrizB.txt", caso["num"].(int)), objeto.Tam)

		// Realizar el cálculo y medir el tiempo
		for i := 0; i < caso["numMuestras"].(int); i++ {
			// Inicializar la matriz de resultados
			result := make([][]int, len(matrizA))
			for i := range result {
				result[i] = make([]int, len(matrizB[0]))
			}

			// Medir el tiempo de ejecución
			tiempoInicio := time.Now()
			// Llamar a la función de multiplicación de matrices
			algoritmos.IV3SequentialBlock(matrizA, matrizB, objeto.TamanioBloques)
			tiempoFinalizacion := time.Since(tiempoInicio)

			// Agregar el tiempo de ejecución a las muestras
			objeto.Muestras = append(objeto.Muestras, tiempoFinalizacion.Seconds())
		}

		// Calcular el promedio de las muestras
		objeto.Promedio = objeto.CalcularPromedio()

		// Agregar el objeto al resultado
		resultado.Casos = append(resultado.Casos, objeto)
	}

	// Guardar el resultado en un archivo JSON
	resultadoJSON, err := json.MarshalIndent(resultado.ToJSON(), "", "  ")
	if err != nil {
		fmt.Println("Error al generar JSON:", err)
		return
	}

	// Guardar en archivo
	err = os.WriteFile("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Resultados/go/IV3SequentialBlockResultadoGo.json", resultadoJSON, 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
	}
}

func ejecutar_III4ParallelBlock() {
	// Casos de prueba
	casosPrueba := []map[string]interface{}{
		{"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques": 4},
		{"num": 2, "tam": 16, "numMuestras": 2, "tamanioBloques": 8},
		{"num": 3, "tam": 32, "numMuestras": 2, "tamanioBloques": 16},
		{"num": 4, "tam": 64, "numMuestras": 1, "tamanioBloques": 32},
		{"num": 5, "tam": 128, "numMuestras": 1, "tamanioBloques": 64},
		{"num": 6, "tam": 256, "numMuestras": 1, "tamanioBloques": 64},
		{"num": 7, "tam": 512, "numMuestras": 1, "tamanioBloques": 128},
		{"num": 8, "tam": 1024, "numMuestras": 1, "tamanioBloques": 128},
	}

	// Resultado
	resultado := utilidades.ResultadoMet{
		Nombre:   "III4ParallelBlock",
		Casos:    []utilidades.CasoMet{},
		Lenguaje: "go",
	}

	// Procesar los casos de prueba
	for _, caso := range casosPrueba {
		objeto := utilidades.CasoMet{
			Tam:            caso["tam"].(int),
			Muestras:       []float64{},
			Promedio:       0,
			TamanioBloques: caso["tamanioBloques"].(int),
		}

		// Crear matrices de ejemplo A y B
		N := objeto.Tam
		P := N // Suponiendo una matriz cuadrada
		M := N
		matrizA := make([][]int, N)
		matrizB := make([][]int, N)

		for i := 0; i < N; i++ {
			matrizA[i] = make([]int, P)
			matrizB[i] = make([]int, M)
		}

		// Leer matrices desde los archivos correspondientes
		matrizA = LeerMatriz(fmt.Sprintf("G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso%d\\matrizA.txt", caso["num"].(int)), N)
		matrizB = LeerMatriz(fmt.Sprintf("G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso%d\\matrizB.txt", caso["num"].(int)), N)

		// Realizar el cálculo y medir el tiempo
		for i := 0; i < caso["numMuestras"].(int); i++ {
			tiempoInicio := time.Now()
			// Llamar a la función del paquete algoritmos
			algoritmos.III4ParallelBlock(matrizA, matrizB, objeto.TamanioBloques)
			tiempoFinalizacion := time.Since(tiempoInicio)

			// Agregar el tiempo de ejecución a las muestras
			objeto.Muestras = append(objeto.Muestras, tiempoFinalizacion.Seconds())
		}

		// Calcular el promedio de las muestras
		objeto.Promedio = objeto.CalcularPromedio()

		// Agregar el objeto al resultado
		resultado.Casos = append(resultado.Casos, objeto)
	}

	// Guardar el resultado en un archivo JSON
	resultadoJSON, err := json.MarshalIndent(resultado.ToJSON(), "", "  ")
	if err != nil {
		fmt.Println("Error al generar JSON:", err)
		return
	}

	// Guardar en archivo
	err = os.WriteFile("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Resultados/go/III4ParallelBlockResultadoGo.json", resultadoJSON, 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
	}
}

func ejecutar_III3SequentialBlock() {
	// Casos de prueba
	casosPrueba := []map[string]interface{}{
		{"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques": 4},
		{"num": 2, "tam": 16, "numMuestras": 2, "tamanioBloques": 8},
		{"num": 3, "tam": 32, "numMuestras": 2, "tamanioBloques": 16},
		{"num": 4, "tam": 64, "numMuestras": 1, "tamanioBloques": 32},
		{"num": 5, "tam": 128, "numMuestras": 1, "tamanioBloques": 64},
		{"num": 6, "tam": 256, "numMuestras": 1, "tamanioBloques": 64},
		{"num": 7, "tam": 512, "numMuestras": 1, "tamanioBloques": 128},
		{"num": 8, "tam": 1024, "numMuestras": 1, "tamanioBloques": 128},
	}

	// Resultado
	resultado := utilidades.ResultadoMet{
		Nombre:   "III3SequentialBlock",
		Casos:    []utilidades.CasoMet{},
		Lenguaje: "go",
	}

	// Procesar los casos de prueba
	for _, caso := range casosPrueba {
		objeto := utilidades.CasoMet{
			Tam:            caso["tam"].(int),
			Muestras:       []float64{},
			Promedio:       0,
			TamanioBloques: caso["tamanioBloques"].(int),
		}

		// Crear matrices de ejemplo A y B
		N := objeto.Tam
		P := N // Suponiendo una matriz cuadrada
		M := N
		matrizA := make([][]int, N)
		matrizB := make([][]int, N)
		result := make([][]int, N)

		for i := 0; i < N; i++ {
			matrizA[i] = make([]int, P)
			matrizB[i] = make([]int, M)
			result[i] = make([]int, M)
		}

		// Leer matrices desde los archivos correspondientes
		matrizA = LeerMatriz(fmt.Sprintf("G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso%d\\matrizA.txt", caso["num"].(int)), N)
		matrizB = LeerMatriz(fmt.Sprintf("G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso%d\\matrizB.txt", caso["num"].(int)), N)

		// Realizar el cálculo y medir el tiempo
		for i := 0; i < caso["numMuestras"].(int); i++ {
			tiempoInicio := time.Now()
			// Llamar a la función del paquete algoritmos
			algoritmos.III3SequentialBlock(matrizA, matrizB, result, objeto.TamanioBloques)
			tiempoFinalizacion := time.Since(tiempoInicio)

			// Agregar el tiempo de ejecución a las muestras
			objeto.Muestras = append(objeto.Muestras, tiempoFinalizacion.Seconds())
		}

		// Calcular el promedio de las muestras
		objeto.Promedio = objeto.CalcularPromedio()

		// Agregar el objeto al resultado
		resultado.Casos = append(resultado.Casos, objeto)
	}

	// Guardar el resultado en un archivo JSON
	resultadoJSON, err := json.MarshalIndent(resultado.ToJSON(), "", "  ")
	if err != nil {
		fmt.Println("Error al generar JSON:", err)
		return
	}

	// Guardar en archivo
	err = os.WriteFile("G:/Visual Studio Code - workspace/Proyect_Final_Analisis/Documentos/Resultados/go/III3SequentialBlockResultadoGo.json", resultadoJSON, 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
	}
}

func main() {
	ejecutar_III3SequentialBlock()
	ejecutar_III4ParallelBlock()
	ejecutar_IV3SequentialBlock()
	ejecutar_NaivLoopUnrollingFour()
	ejecutar_NaivLoopUnrollingTwo()
	ejecutar_NaivOnArray()
	ejecutar_StrassenNaiv()
	ejecutar_StrassenWinograd()
	ejecutar_V3SequentialBlock()
	ejecutar_WinogradOriginal()
}
