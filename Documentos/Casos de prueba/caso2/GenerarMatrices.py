import numpy as np

# Generar una matriz de 16x16 con números aleatorios de 6 dígitos
matrizA = np.random.randint(100000, 999999, size=(16, 16))
matrizB = np.random.randint(100000, 999999, size=(16, 16))

# Guardar la matrizA en un archivo de texto
np.savetxt('G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso2\\matrizA.txt', matrizA, fmt='%d')
# Guardar la matrizB en un archivo de texto
np.savetxt('G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso2\\matrizB.txt', matrizB, fmt='%d')