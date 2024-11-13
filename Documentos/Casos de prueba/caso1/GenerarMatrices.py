import numpy as np

# Generar una matriz de 8x8 con números aleatorios de 6 dígitos
matrizA = np.random.randint(100000, 999999, size=(8, 8))
matrizB = np.random.randint(100000, 999999, size=(8, 8))

# Guardar la matrizA en un archivo de texto
np.savetxt('G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso1\\matrizA.txt', matrizA, fmt='%d')
# Guardar la matrizB en un archivo de texto
np.savetxt('G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso1\\matrizB.txt', matrizB, fmt='%d')
