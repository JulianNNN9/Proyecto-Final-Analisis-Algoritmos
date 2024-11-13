import numpy as np

# Generar una matriz de 1024X1024 con números aleatorios de 6 dígitos
matrizA = np.random.randint(100000, 999999, size=(1024, 1024))
matrizB = np.random.randint(100000, 999999, size=(1024, 1024))

# Guardar la matrizA en un archivo de texto
np.savetxt('G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso8\\matrizA.txt', matrizA, fmt='%d')
# Guardar la matrizB en un archivo de texto
np.savetxt('G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso8\\matrizB.txt', matrizB, fmt='%d')
