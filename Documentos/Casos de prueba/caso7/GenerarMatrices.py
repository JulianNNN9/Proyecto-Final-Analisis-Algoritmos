import numpy as np

# Generar una matriz de 512X512 con números aleatorios de 6 dígitos
matrizA = np.random.randint(100000, 999999, size=(512, 512))
matrizB = np.random.randint(100000, 999999, size=(512, 512))

# Guardar la matrizA en un archivo de texto
np.savetxt('G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso7\\matrizA.txt', matrizA, fmt='%d')
# Guardar la matrizB en un archivo de texto
np.savetxt('G:\\Visual Studio Code - workspace\\Proyect_Final_Analisis\\Documentos\\Casos de prueba\\caso7\\matrizB.txt', matrizB, fmt='%d')
