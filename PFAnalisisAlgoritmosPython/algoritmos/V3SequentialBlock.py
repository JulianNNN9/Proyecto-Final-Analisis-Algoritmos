def V3SequentialBlock(A, B, Result, tamanioBloque):

    size = len(Result)

    for i1 in range(0, size, tamanioBloque):
        for j1 in range(0, size, tamanioBloque):
            for k1 in range(0, size, tamanioBloque):
                for i in range(i1, min(i1 + tamanioBloque, size)):
                    for j in range(j1, min(j1 + tamanioBloque, size)):
                        for k in range(k1, min(k1 + tamanioBloque, size)):
                            Result[k][i] += A[k][j] * B[j][i]

    return Result
