import numpy as np
from concurrent.futures import ThreadPoolExecutor, as_completed

def III4ParallelBlock(A, B, C, bsize):
    size = len(A)
    result = np.zeros_like(C)
    
    # La tarea optimizada para calcular por bloques
    def task(i1, j1, k1):
        A_block = A[i1:i1+bsize, k1:k1+bsize]
        B_block = B[k1:k1+bsize, j1:j1+bsize]
        return i1, j1, np.dot(A_block, B_block)

    with ThreadPoolExecutor() as executor:
        futures = []
        # Enviar tareas a ejecutar
        for i1 in range(0, size, bsize):
            for j1 in range(0, size, bsize):
                for k1 in range(0, size, bsize):
                    futures.append(executor.submit(task, i1, j1, k1))

        # Recoger los resultados
        for future in as_completed(futures):
            i1, j1, local_result = future.result()
            # Insertar el bloque calculado en la matriz global 'result'
            result[i1:i1+bsize, j1:j1+bsize] += local_result
    
    return result
