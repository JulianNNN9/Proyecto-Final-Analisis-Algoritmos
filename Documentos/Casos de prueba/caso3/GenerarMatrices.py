import numpy as np

# Generar una matriz de 32x32 con números aleatorios de 6 dígitos
matrizA = np.random.randint(100000, 999999, size=(32, 32))
matrizB = np.random.randint(100000, 999999, size=(32, 32))

# Guardar la matrizA en un archivo de texto
np.savetxt('G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso3\\matrizA.txt', matrizA, fmt='%d')
# Guardar la matrizB en un archivo de texto
np.savetxt('G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso3\\matrizB.txt', matrizB, fmt='%d')
