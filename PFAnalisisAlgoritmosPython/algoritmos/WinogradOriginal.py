# N -> columnas de la matriz B
# P -> columnas de la matriz A y filas de la matriz B
# M -> columnas de la matriz B

def WinogradOriginal(A, B, Result, N, P, M):
    upsilon = P % 2
    gamma = P - upsilon

    y = [0.0] * M
    z = [0.0] * N

    for i in range(M):
        aux = 0.0
        for j in range(0, gamma, 2):
            aux += A[i][j] * A[i][j + 1]
        y[i] = aux

    for i in range(N):
        aux = 0.0
        for j in range(0, gamma, 2):
            aux += B[j][i] * B[j + 1][i]
        z[i] = aux

    if upsilon == 1:
        PP = P - 1
        for i in range(M):
            for k in range(N):
                aux = 0.0
                for j in range(0, gamma, 2):
                    aux += (A[i][j] + B[j + 1][k]) * (A[i][j + 1] + B[j][k])
                Result[i][k] = aux - y[i] - z[k] + A[i][PP] * B[PP][k]
    else:
        for i in range(M):
            for k in range(N):
                aux = 0.0
                for j in range(0, gamma, 2):
                    aux += (A[i][j] + B[j + 1][k]) * (A[i][j + 1] + B[j][k])
                Result[i][k] = aux - y[i] - z[k]

    return Result