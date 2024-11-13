# N -> filas de la matriz A
# P -> columnas de la matriz A y filas de la matriz B
# M -> columnas de la matriz B
def NaivLoopUnrollingTwo(A, B, Result, N, P, M):
    if P % 2 == 0:
        for i in range(N):
            for j in range(M):
                aux = 0.0
                for k in range(0, P, 2):
                    aux += A[i][k] * B[k][j] + A[i][k+1] * B[k+1][j]

                Result[i][j] = aux
    else:
        PP = P - 1
        for i in range(N):
            for j in range(M):
                aux = 0.0

                for k in range(0, PP, 2):
                    aux += A[i][k] * B[k][j] + A[i][k+1] * B[k+1][j]

                Result[i][j] = aux + A[i][PP] * B[PP][j]

    return Result