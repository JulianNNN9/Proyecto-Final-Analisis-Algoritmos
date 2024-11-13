def NaivStandard(A, B, Result, N, P, M):
    
    aux = 0.0
    for i in range(N):
        for j in range(M):
            aux = 0.0
            for k in range(P):
                aux += A[i][k] * B[k][j]
            Result[i][j] = aux
            
    return Result