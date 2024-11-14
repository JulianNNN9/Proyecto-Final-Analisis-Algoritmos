import json
import matplotlib.pyplot as plt

import json

# Función para leer un archivo JSON y devolver el promedio del primer caso
def obtener_promedio_primero(archivo):
    try:
        # Intentar abrir el archivo JSON
        with open(archivo, "r") as file:
            try:
                # Intentar cargar el contenido del archivo como JSON
                data = json.load(file)
            except json.JSONDecodeError as e:
                # Manejar errores al decodificar el archivo JSON
                print(f"Error al leer el archivo JSON: {e}")
                return 0  # O algún otro valor predeterminado
            
            # Verificar si 'casos' está en el JSON y contiene al menos un elemento
            if "casos" in data and len(data["casos"]) > 0:
                return data["casos"][1]["promedio"]
            else:
                # Manejar el caso en que 'casos' no exista o esté vacío
                print(f"Advertencia: el archivo {archivo} no contiene la clave 'casos' o está vacío.")
                return 0  # O algún otro valor predeterminado

    except FileNotFoundError as e:
        # Manejar el error en caso de que el archivo no exista
        print(f"Error: El archivo {archivo} no se encuentra: {e}")
        return 0  # O algún otro valor predeterminado
    except PermissionError as e:
        # Manejar el error si no se tienen permisos para acceder al archivo
        print(f"Error: Permiso denegado para abrir el archivo {archivo}: {e}")
        return 0  # O algún otro valor predeterminado
    except Exception as e:
        # Manejar cualquier otro error inesperado
        print(f"Ocurrió un error inesperado: {e}")
        return 0  # O algún otro valor predeterminado


# Rutas de los archivos JSON para Go y Python
ruta_go = "Documentos\\Resultados\\go\\"
ruta_python = "Documentos\\Resultados\\python\\"

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
plt.title('Caso 2 (16x16)')
plt.xlabel('Algoritmo')
plt.ylabel('Tiempo en (s)')

# Mostrar el gráfico
plt.show()