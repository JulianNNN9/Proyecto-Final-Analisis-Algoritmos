import numpy as np
from concurrent.futures import ThreadPoolExecutor, as_completed


def III4ParallelBlock(A, B, C, bsize):
    size = len(A)
    def task(i1, j1, k1):
        result = np.zeros_like(C)
        for i in range(i1, min(i1 + bsize, size)):
            for j in range(j1, min(j1 + bsize, size)):
                for k in range(k1, min(k1 + bsize, size)):
                    result[i][j] += A[i][k] * B[k][j]
        return result

    with ThreadPoolExecutor() as executor:
        futures = []
        for i1 in range(0, size, bsize):
            for j1 in range(0, size, bsize):
                for k1 in range(0, size, bsize):
                    futures.append(executor.submit(task, i1, j1, k1))

        result = np.zeros_like(C)
        for future in as_completed(futures):
            result += future.result()

    return result
