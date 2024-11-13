from math import floor,log

from algoritmos import StrassenWinogradStep

def StrassenWinograd (A, B, Result, N, P, M):

    tamanioMaximo = max(N, M)
    tamanioMaximo = max(tamanioMaximo, N)
    
    if tamanioMaximo < 16:
        tamanioMaximo = 16
        
    k = floor(log(tamanioMaximo)/log(2)) - 4
    m = floor(tamanioMaximo * pow(2,-k)) + 1
    nuevoTamanio = m * pow(2,k)
    
    nuevaMatrizA = []
    nuevaMatrizB = [] 
    matrizResultadoAuxiliar = []

    for i in range(nuevoTamanio):
         nuevaMatrizA.append([0 for i in range(nuevoTamanio)] )
         nuevaMatrizB.append([0 for i in range(nuevoTamanio)] )
         matrizResultadoAuxiliar.append([0 for i in range(nuevoTamanio)] )

    # Asignamos en cada posición i,j de las nuevas matrices A y B los valores que están en las matrices A y B respectivamente
    for i in range(N):
        for j in range(M):
            nuevaMatrizA[i][j] = A[i][j]
            
    for i in range(M):
        for j in range(P):
            nuevaMatrizB[i][j] = B[i][j]
            
    matrizResultadoAuxiliar = StrassenWinogradStep(nuevaMatrizA, nuevaMatrizB, matrizResultadoAuxiliar, nuevoTamanio, m)
    
    for i in range(N):
        for j in range(P):
            Result[i][j] = matrizResultadoAuxiliar[i][j]
            
    return Result