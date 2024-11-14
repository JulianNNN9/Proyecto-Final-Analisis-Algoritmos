import numpy as np
import os

# Generar una matriz de 16x16 con números aleatorios de 6 dígitos
matrizA = np.random.randint(100000, 999999, size=(16, 16))
matrizB = np.random.randint(100000, 999999, size=(16, 16))

try:

    # Definir las rutas de los archivos
    ruta_matrizA = 'Documentos\\Casos de prueba\\caso1\\matrizA.txt'
    ruta_matrizB = 'Documentos\\Casos de prueba\\caso1\\matrizB.txt'

    # Verificar si el directorio existe, si no, crear uno nuevo
    if not os.path.exists(os.path.dirname(ruta_matrizA)):
        os.makedirs(os.path.dirname(ruta_matrizA))

    # Guardar la matrizA en un archivo de texto
    np.savetxt(ruta_matrizA, matrizA, fmt='%d')
    # Guardar la matrizB en un archivo de texto
    np.savetxt(ruta_matrizB, matrizB, fmt='%d')

    print("Las matrices se guardaron correctamente.")

except FileNotFoundError as e:
    print(f"Error: No se encontró el archivo o el directorio especificado: {e}")
except PermissionError as e:
    print(f"Error: No se tienen los permisos necesarios para guardar los archivos: {e}")
except Exception as e:
    print(f"Ocurrió un error inesperado: {e}")