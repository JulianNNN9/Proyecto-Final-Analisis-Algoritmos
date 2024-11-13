# N -> filas de la matriz A
# P -> columnas de la matriz A y filas de la matriz B
# M -> columnas de la matriz B
def NaivOnArray(A, B, Result, N, P, M):
    for i in range(N):
        for j in range(M):
            Result[i][j] = 0.0
            for k in range(P):
                Result[i][j] += A[i][k] * B[k][j]

    return Result
