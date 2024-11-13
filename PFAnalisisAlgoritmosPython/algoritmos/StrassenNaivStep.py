from algoritmos import NaivStandard

def StrassenNaivStep(A, B, Result, cantidadFilasMatrices, m):

    nuevoTamanio = 0
    if (cantidadFilasMatrices % 2 == 0) and (cantidadFilasMatrices > m):
        nuevoTamanio = cantidadFilasMatrices // 2
        
        matrizA11 = []
        matrizA12 = []
        matrizA21 = []
        matrizA22 = []
        
        matrizB11 = []
        matrizB12 = [] 
        matrizB21 = []
        matrizB22 = []
        
        matrizResultadoParte11 = []
        matrizResultadoParte12 = []
        matrizResultadoParte21 = []
        matrizResultadoParte22 = []
        
        ayudante1 = []
        ayudante2 = []
        
        auxiliar1 = []
        auxiliar2 = []
        auxiliar3 = []
        auxiliar4 = []
        auxiliar5 = []
        auxiliar6 = []
        auxiliar7 = []

        # Asignar memoria para cada fila
        for i in range(nuevoTamanio):
            matrizA11.append([0] * nuevoTamanio)
            matrizA12.append([0] * nuevoTamanio)
            matrizA21.append([0] * nuevoTamanio)
            matrizA22.append([0] * nuevoTamanio)
            
            matrizB11.append([0] * nuevoTamanio)
            matrizB12.append([0] * nuevoTamanio)
            matrizB21.append([0] * nuevoTamanio)
            matrizB22.append([0] * nuevoTamanio)
            
            matrizResultadoParte11.append([0] * nuevoTamanio)
            matrizResultadoParte12.append([0] * nuevoTamanio)
            matrizResultadoParte21.append([0] * nuevoTamanio)
            matrizResultadoParte22.append([0] * nuevoTamanio)
            
            ayudante1.append([0] * nuevoTamanio)
            ayudante2.append([0] * nuevoTamanio)
            
            auxiliar1.append([0] * nuevoTamanio)
            auxiliar2.append([0] * nuevoTamanio)
            auxiliar3.append([0] * nuevoTamanio)
            auxiliar4.append([0] * nuevoTamanio)
            auxiliar5.append([0] * nuevoTamanio)
            auxiliar6.append([0] * nuevoTamanio)
            auxiliar7.append([0] * nuevoTamanio)
            
        # Llenamos las matrices
        for i in range(nuevoTamanio):
            for j in range(nuevoTamanio):
                matrizA11[i][j] = A[i][j]
                matrizA12[i][j] = A[i][nuevoTamanio + j]
                matrizA21[i][j] = A[nuevoTamanio + i][j]
                matrizA22[i][j] = A[nuevoTamanio + i][nuevoTamanio + j]

                matrizB11[i][j] = B[i][j]
                matrizB12[i][j] = B[i][nuevoTamanio + j]
                matrizB21[i][j] = B[nuevoTamanio + i][j]
                matrizB22[i][j] = B[nuevoTamanio + i][nuevoTamanio + j]
        
        # computing the seven aux. variables
        
       # Computando las siete variables auxiliares
        ayudante1 = [[matrizA11[i][j] + matrizA22[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        ayudante2 = [[matrizB11[i][j] + matrizB22[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        auxiliar1 = StrassenNaivStep(ayudante1, ayudante2, auxiliar1, nuevoTamanio, m)


        ayudante1 = [[matrizA21[i][j] + matrizA22[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        auxiliar2 = StrassenNaivStep(ayudante1, matrizB11, auxiliar2, nuevoTamanio, m)
        
        ayudante1 = [[matrizB12[i][j] - matrizB22[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        auxiliar3 = StrassenNaivStep(matrizA11, ayudante1, auxiliar3, nuevoTamanio, m)
        
        ayudante1 = [[matrizB21[i][j] - matrizB11[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        auxiliar4 = StrassenNaivStep(matrizA22, ayudante1, auxiliar4, nuevoTamanio, m)
        
        ayudante1 = [[matrizA11[i][j] + matrizA12[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        auxiliar5 = StrassenNaivStep(ayudante1, matrizB22, auxiliar5, nuevoTamanio, m)
        
        ayudante1 = [[matrizA21[i][j] - matrizA11[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        ayudante2 = [[matrizB11[i][j] + matrizB12[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        auxiliar6 = StrassenNaivStep(ayudante1, ayudante2, auxiliar6, nuevoTamanio, m)
        
        ayudante1 = [[matrizA12[i][j] - matrizA22[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        ayudante2 = [[matrizB21[i][j] + matrizB22[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        auxiliar7 = StrassenNaivStep(ayudante1, ayudante2, auxiliar7, nuevoTamanio, m)
        
        # Computando las cuatro partes del resultado
        matrizResultadoParte11 = [[auxiliar1[i][j] + auxiliar4[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        matrizResultadoParte11 = [[matrizResultadoParte11[i][j] - auxiliar5[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        matrizResultadoParte11 = [[matrizResultadoParte11[i][j] + auxiliar7[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]

        matrizResultadoParte12 = [[auxiliar3[i][j] + auxiliar5[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]

        matrizResultadoParte21 = [[auxiliar2[i][j] + auxiliar4[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]

        matrizResultadoParte22 = [[auxiliar1[i][j] + auxiliar3[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        matrizResultadoParte22 = [[matrizResultadoParte22[i][j] - auxiliar2[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]
        matrizResultadoParte22 = [[matrizResultadoParte22[i][j] + auxiliar6[i][j] for j in range(nuevoTamanio)] for i in range(nuevoTamanio)]

        # Almacenar resultados en la matriz resultado
        for i in range(nuevoTamanio):
            for j in range(nuevoTamanio):
                Result[i][j] = matrizResultadoParte11[i][j]
        
        for i in range(nuevoTamanio):
            for j in range(nuevoTamanio):
                Result[i][nuevoTamanio + j] = matrizResultadoParte12[i][j]
        
        for i in range(nuevoTamanio):
            for j in range(nuevoTamanio):
                Result[nuevoTamanio + i][j] = matrizResultadoParte21[i][j]
        
        for i in range(nuevoTamanio):
            for j in range(nuevoTamanio):
                Result[nuevoTamanio + i][nuevoTamanio + j] = matrizResultadoParte22[i][j]
        
    else:
       # matrizResultado = naiv_standard(matrizA, matrizB, matrizResultado, len(matrizA), len(matrizB), len(matrizResultado))
        Result = NaivStandard(A, B, Result, cantidadFilasMatrices, cantidadFilasMatrices, cantidadFilasMatrices)
    
    return Result