import json
import matplotlib.pyplot as plt

# Función para leer un archivo JSON y devolver el promedio del primer caso
def obtener_promedio_primero(archivo):
    with open(archivo, "r") as file:
        data = json.load(file)
        return data["casos"][2]["promedio"]

# Rutas de los archivos JSON para Go y Python
ruta_go = "..\\Documentos\\Resultados\\go\\"
ruta_python = "..\\Documentos\\Resultados\\python\\"

# Nombres de los algoritmos
nombres_algoritmos = ["NaivOnArray", "NaivLoopUnrollingTwo", "NaivLoopUnrollingFour", "WinogradOriginal", "StrassenNaiv",
                       "StrassenWinograd", "III3SequentialBlock", "III4ParallelBlock", "IV3SequentialBlock", "V3SequentialBlock"]

archivos_go = [ruta_go + nombre + "ResultadoGo.json" for nombre in nombres_algoritmos]
archivos_python = [ruta_python + nombre + "ResultadoPython.json" for nombre in nombres_algoritmos]

# Promedios para Go y Python
promedios_go = [obtener_promedio_primero(archivo) for archivo in archivos_go]
promedios_python = [obtener_promedio_primero(archivo) for archivo in archivos_python]


# Configuración del gráfico
bar_width = 0.1
index = range(len(nombres_algoritmos))

# Crear el gráfico de barras
plt.bar(index, promedios_go, bar_width, color='b', label='Go')
plt.bar([i + bar_width for i in index], promedios_python, bar_width, color='r', label='Python')

# Obtener el valor máximo de todos los promedios
max_promedio = max(max(promedios_go), max(promedios_python))

# Factor de desplazamiento vertical
desplazamiento_vertical = 0.01  # Puedes ajustar este valor según tus necesidades

# Agregar nombre del algoritmo encima de cada grupo de barras con desplazamiento
for i, algoritmo in enumerate(nombres_algoritmos):
    y_texto = max(promedios_go[i], promedios_python[i]) + desplazamiento_vertical * max_promedio
    plt.text(i + bar_width/2, y_texto, algoritmo, ha='center')

# Agregar leyenda
plt.legend()

# Agregar título y etiquetas a los ejes
plt.title('Caso 3 (32x32)')
plt.xlabel('Algoritmo')
plt.ylabel('Tiempo en (s)')

# Mostrar el gráfico
plt.show()