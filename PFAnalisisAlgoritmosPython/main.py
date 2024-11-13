import json

import numpy as np
import time

from algoritmos import III3SequentialBlock
from algoritmos import III4ParallelBlock
from algoritmos import IV3SequentialBlock
from algoritmos import NaivLoopUnrollingFour
from algoritmos import NaivLoopUnrollingTwo
from algoritmos import NaivOnArray
from algoritmos import StrassenNaiv
from algoritmos import StrassenWinograd
from algoritmos import V3SequentialBlock
from algoritmos import WinogradOriginal
from utilidades import Caso
from utilidades import Resultado

def ejecutar_WinogradOriginal():
    casosPrueba = ({"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques":0},
               {"num": 2,  "tam": 16, "numMuestras": 2, "tamanioBloques":0},
               {"num": 3,  "tam": 32, "numMuestras": 2, "tamanioBloques":0},
               {"num": 4,  "tam": 64, "numMuestras": 2, "tamanioBloques":0},
               {"num": 5,  "tam": 128, "numMuestras": 2, "tamanioBloques":0},
               {"num": 6,  "tam": 256, "numMuestras": 2, "tamanioBloques":0},
               {"num": 7,  "tam": 512, "numMuestras": 1, "tamanioBloques":0},
               {"num": 8,  "tam": 1024, "numMuestras": 1, "tamanioBloques":0})
    resultado = Resultado.ResultadoMet("WinogradOriginal", [], "python")

    for caso in casosPrueba:
        objeto = Caso.CasoMet(caso["tam"], [], 0,0)
        matrizA = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizA.txt'.format(caso["num"]), dtype=int)
        matrizB = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizB.txt'.format(caso["num"]), dtype=int)

        for i in range(caso["numMuestras"]):
            N = len(matrizA)
            P = len(matrizA[0])
            M = len(matrizB[0])

            Result = [[0 for _ in range(M)] for _ in range(N)]

            tiempoInicio = time.time()
            matrizResultado = WinogradOriginal.WinogradOriginal(matrizA, matrizB, Result, N, P, M)
            tiempoFinalizacion = time.time() - tiempoInicio

            objeto.muestras.append(tiempoFinalizacion)

        objeto.promedio = objeto.calcular_promedio()
        resultado.casos.append(objeto)

    with open("Documentos\\Resultados\\python\\WinogradOriginalResultadoPython.json", "w") as archivo:
        json.dump(resultado.to_json(), archivo)

def ejecutar_V3SequentialBlock():
    casosPrueba = ({"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques":4},
               {"num": 2,  "tam": 16, "numMuestras": 2, "tamanioBloques":8},
               {"num": 3,  "tam": 32, "numMuestras": 2, "tamanioBloques":16},
               {"num": 4,  "tam": 64, "numMuestras": 1, "tamanioBloques":32},
               {"num": 5,  "tam": 128, "numMuestras": 1, "tamanioBloques":64},
               {"num": 6,  "tam": 256, "numMuestras": 1, "tamanioBloques":64},
               {"num": 7,  "tam": 512, "numMuestras": 1, "tamanioBloques":128},
               {"num": 8,  "tam": 1024, "numMuestras": 1,  "tamanioBloques":128})
    resultado = Resultado.ResultadoMet("V3SequentialBlock", [], "python")

    for caso in casosPrueba:
        objeto = Caso.CasoMet(caso["tam"], [], 0,caso["tamanioBloques"])
        matrizA = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizA.txt'.format(caso["num"]), dtype=int)
        matrizB = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizB.txt'.format(caso["num"]), dtype=int)

        for i in range(caso["numMuestras"]):
            N = len(matrizA)
            P = len(matrizA[0])
            M = len(matrizB[0])

            Result = [[0 for _ in range(M)] for _ in range(N)]

            tiempoInicio = time.time()
            matrizResultado = V3SequentialBlock.V3SequentialBlock(matrizA, matrizB, Result, caso["tamanioBloques"])
            tiempoFinalizacion = time.time() - tiempoInicio

            objeto.muestras.append(tiempoFinalizacion)

        objeto.promedio = objeto.calcular_promedio()
        resultado.casos.append(objeto)

    with open("Documentos\\Resultados\\python\\V3SequentialBlockResultadoPython.json", "w") as archivo:
        json.dump(resultado.to_json(), archivo)

def ejecutar_StrassenWinograd():
    casosPrueba = ({"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques":0},
               {"num": 2,  "tam": 16, "numMuestras": 2, "tamanioBloques":0},
               {"num": 3,  "tam": 32, "numMuestras": 2, "tamanioBloques":0},
               {"num": 4,  "tam": 64, "numMuestras": 2, "tamanioBloques":0},
               {"num": 5,  "tam": 128, "numMuestras": 1, "tamanioBloques":0},
               {"num": 6,  "tam": 256, "numMuestras": 1, "tamanioBloques":0},
               {"num": 7,  "tam": 512, "numMuestras": 1, "tamanioBloques":0},
               {"num": 8,  "tam": 1024, "numMuestras": 1, "tamanioBloques":0})
    resultado = Resultado.ResultadoMet("StrassenWinograd", [], "python")

    for caso in casosPrueba:
        objeto = Caso.CasoMet(caso["tam"], [], 0,0)
        matrizA = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizA.txt'.format(caso["num"]), dtype=int)
        matrizB = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizB.txt'.format(caso["num"]), dtype=int)

        for i in range(caso["numMuestras"]):
            N = len(matrizA)
            P = len(matrizA[0])
            M = len(matrizB[0])

            Result = [[0 for _ in range(M)] for _ in range(N)]

            tiempoInicio = time.time()
            matrizResultado = StrassenWinograd.StrassenWinograd(matrizA, matrizB, Result, N, P, M)
            tiempoFinalizacion = time.time() - tiempoInicio

            objeto.muestras.append(tiempoFinalizacion)

        objeto.promedio = objeto.calcular_promedio()
        resultado.casos.append(objeto)

    with open("Documentos\\Resultados\\python\\StrassenWinogradResultadoPython.json", "w") as archivo:
        json.dump(resultado.to_json(), archivo)

def ejecutar_StrassenNaiv():
    casosPrueba = ({"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques":0},
               {"num": 2,  "tam": 16, "numMuestras": 2, "tamanioBloques":0},
               {"num": 3,  "tam": 32, "numMuestras": 2, "tamanioBloques":0},
               {"num": 4,  "tam": 64, "numMuestras": 2, "tamanioBloques":0},
               {"num": 5,  "tam": 128, "numMuestras": 1, "tamanioBloques":0},
               {"num": 6,  "tam": 256, "numMuestras": 1, "tamanioBloques":0},
               {"num": 7,  "tam": 512, "numMuestras": 1, "tamanioBloques":0},
               {"num": 8,  "tam": 1024, "numMuestras": 1, "tamanioBloques":0})
    resultado = Resultado.ResultadoMet("StrassenNaiv", [], "python")

    for caso in casosPrueba:
        objeto = Caso.CasoMet(caso["tam"], [], 0,0)
        matrizA = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizA.txt'.format(caso["num"]), dtype=int)
        matrizB = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizB.txt'.format(caso["num"]), dtype=int)

        for i in range(caso["numMuestras"]):
            N = len(matrizA)
            P = len(matrizA[0])
            M = len(matrizB[0])

            Result = [[0 for _ in range(M)] for _ in range(N)]

            tiempoInicio = time.time()
            matrizResultado = StrassenNaiv.StrassenNaiv(matrizA, matrizB, Result, N, P, M)
            tiempoFinalizacion = time.time() - tiempoInicio

            objeto.muestras.append(tiempoFinalizacion)

        objeto.promedio = objeto.calcular_promedio()
        resultado.casos.append(objeto)

    with open("Documentos\\Resultados\\python\\StrassenNaivResultadoPython.json", "w") as archivo:
        json.dump(resultado.to_json(), archivo)

def ejecutar_NaivOnArray():
    casosPrueba = ({"num": 1, "tam": 8, "numMuestras": 3, "tamanioBloques":0},
               {"num": 2,  "tam": 16, "numMuestras": 3, "tamanioBloques":0},
               {"num": 3,  "tam": 32, "numMuestras": 3, "tamanioBloques":0},
               {"num": 4,  "tam": 64, "numMuestras": 3, "tamanioBloques":0},
               {"num": 5,  "tam": 128, "numMuestras": 3, "tamanioBloques":0},
               {"num": 6,  "tam": 256, "numMuestras": 3, "tamanioBloques":0},
               {"num": 7,  "tam": 512, "numMuestras": 3, "tamanioBloques":0},
               {"num": 8,  "tam": 1024, "numMuestras": 2, "tamanioBloques":0})
    resultado = Resultado.ResultadoMet("NaivOnArray", [], "python")

    for caso in casosPrueba:
        objeto = Caso.CasoMet(caso["tam"], [], 0,0)
        matrizA = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizA.txt'.format(caso["num"]), dtype=int)
        matrizB = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizB.txt'.format(caso["num"]), dtype=int)

        for i in range(caso["numMuestras"]):
            N = len(matrizA)
            P = len(matrizA[0])
            M = len(matrizB[0])

            Result = [[0 for _ in range(M)] for _ in range(N)]

            tiempoInicio = time.time()
            matrizResultado = NaivOnArray.NaivOnArray(matrizA, matrizB, Result, N, P, M)
            tiempoFinalizacion = time.time() - tiempoInicio

            objeto.muestras.append(tiempoFinalizacion)

        objeto.promedio = objeto.calcular_promedio()
        resultado.casos.append(objeto)

    with open("Documentos\\Resultados\\python\\NaivOnArrayResultadoPython.json", "w") as archivo:
        json.dump(resultado.to_json(), archivo)

def ejecutar_NaivLoopUnrollingTwo():
    casosPrueba = ({"num": 1, "tam": 8, "numMuestras": 2,"tamanioBloques":0},
               {"num": 2,  "tam": 16, "numMuestras": 2, "tamanioBloques":0},
               {"num": 3,  "tam": 32, "numMuestras": 2, "tamanioBloques":0},
               {"num": 4,  "tam": 64, "numMuestras": 1, "tamanioBloques":0},
               {"num": 5,  "tam": 128, "numMuestras": 1, "tamanioBloques":0},
               {"num": 6,  "tam": 256, "numMuestras": 1, "tamanioBloques":0},
               {"num": 7,  "tam": 512, "numMuestras": 1, "tamanioBloques":0},
               {"num": 8,  "tam": 1024, "numMuestras": 1, "tamanioBloques":0})
    resultado = Resultado.ResultadoMet("NaivLoopUnrollingTwo", [], "python")

    for caso in casosPrueba:
        objeto = Caso.CasoMet(caso["tam"], [], 0,0)
        matrizA = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizA.txt'.format(caso["num"]), dtype=int)
        matrizB = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizB.txt'.format(caso["num"]), dtype=int)

        for i in range(caso["numMuestras"]):
            N = len(matrizA)
            P = len(matrizA[0])
            M = len(matrizB[0])

            Result = [[0 for _ in range(M)] for _ in range(N)]

            tiempoInicio = time.time()
            matrizResultado = NaivLoopUnrollingTwo.NaivLoopUnrollingTwo(matrizA, matrizB, Result, N, P, M)
            tiempoFinalizacion = time.time() - tiempoInicio

            objeto.muestras.append(tiempoFinalizacion)

        objeto.promedio = objeto.calcular_promedio()
        resultado.casos.append(objeto)

    with open("Documentos\\Resultados\\python\\NaivLoopUnrollingTwoResultadoPython.json", "w") as archivo:
        json.dump(resultado.to_json(), archivo)

def ejecutar_NaivLoopUnrollingFour():
    casosPrueba = ({"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques":0},
               {"num": 2,  "tam": 16, "numMuestras": 2, "tamanioBloques":0},
               {"num": 3,  "tam": 32, "numMuestras": 2, "tamanioBloques":0},
               {"num": 4,  "tam": 64, "numMuestras": 1, "tamanioBloques":0},
               {"num": 5,  "tam": 128, "numMuestras": 1, "tamanioBloques":0},
               {"num": 6,  "tam": 256, "numMuestras": 1, "tamanioBloques":0},
               {"num": 7,  "tam": 512, "numMuestras": 1, "tamanioBloques":0},
               {"num": 8,  "tam": 1024, "numMuestras": 1, "tamanioBloques":0})
    resultado = Resultado.ResultadoMet("NaivLoopUnrollingFour", [], "python")

    for caso in casosPrueba:
        objeto = Caso.CasoMet(caso["tam"], [], 0,0)
        matrizA = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizA.txt'.format(caso["num"]), dtype=int)
        matrizB = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizB.txt'.format(caso["num"]), dtype=int)

        for i in range(caso["numMuestras"]):
            N = len(matrizA)
            P = len(matrizA[0])
            M = len(matrizB[0])

            Result = [[0 for _ in range(M)] for _ in range(N)]

            tiempoInicio = time.time()
            matrizResultado = NaivLoopUnrollingFour.NaivLoopUnrollingFour(matrizA, matrizB, Result, N, P, M)
            tiempoFinalizacion = time.time() - tiempoInicio

            objeto.muestras.append(tiempoFinalizacion)

        objeto.promedio = objeto.calcular_promedio()
        resultado.casos.append(objeto)

    with open("Documentos\\Resultados\\python\\NaivLoopUnrollingFourResultadoPython.json", "w") as archivo:
        json.dump(resultado.to_json(), archivo)

def ejecutar_IV3SequentialBlock():
    casosPrueba = ({"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques":4},
               {"num": 2,  "tam": 16, "numMuestras": 2, "tamanioBloques":8},
               {"num": 3,  "tam": 32, "numMuestras": 2, "tamanioBloques":16},
               {"num": 4,  "tam": 64, "numMuestras": 1, "tamanioBloques":32},
               {"num": 5,  "tam": 128, "numMuestras": 1, "tamanioBloques":64},
               {"num": 6,  "tam": 256, "numMuestras": 1, "tamanioBloques":64},
               {"num": 7,  "tam": 512, "numMuestras": 1, "tamanioBloques":128},
               {"num": 8,  "tam": 1024, "numMuestras": 1,  "tamanioBloques":128})
    resultado = Resultado.ResultadoMet("IV3SequentialBlock", [], "python")

    for caso in casosPrueba:
        objeto = Caso.CasoMet(caso["tam"], [], 0,caso["tamanioBloques"])
        matrizA = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizA.txt'.format(caso["num"]), dtype=int)
        matrizB = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizB.txt'.format(caso["num"]), dtype=int)

        for i in range(caso["numMuestras"]):
            N = len(matrizA)
            P = len(matrizA[0])
            M = len(matrizB[0])

            Result = [[0 for _ in range(M)] for _ in range(N)]

            tiempoInicio = time.time()
            matrizResultado = IV3SequentialBlock.IV3SequentialBlock(matrizA, matrizB, Result, caso["tamanioBloques"])
            tiempoFinalizacion = time.time() - tiempoInicio

            objeto.muestras.append(tiempoFinalizacion)

        objeto.promedio = objeto.calcular_promedio()
        resultado.casos.append(objeto)

    with open("Documentos\\Resultados\\python\\IV3SequentialBlockResultadoPython.json", "w") as archivo:
        json.dump(resultado.to_json(), archivo)

def ejecutar_III4ParallelBlock():
    casosPrueba = ({"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques":4},
               {"num": 2,  "tam": 16, "numMuestras": 2, "tamanioBloques":8},
               {"num": 3,  "tam": 32, "numMuestras": 2, "tamanioBloques":16},
               {"num": 4,  "tam": 64, "numMuestras": 1, "tamanioBloques":32},
               {"num": 5,  "tam": 128, "numMuestras": 1, "tamanioBloques":64},
               {"num": 6,  "tam": 256, "numMuestras": 1, "tamanioBloques":64},
               {"num": 7,  "tam": 512, "numMuestras": 1, "tamanioBloques":128},
               {"num": 8,  "tam": 1024, "numMuestras": 1,  "tamanioBloques":128})
    resultado = Resultado.ResultadoMet("III4ParallelBlock", [], "python")

    for caso in casosPrueba:
        objeto = Caso.CasoMet(caso["tam"], [], 0,caso["tamanioBloques"])
        matrizA = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizA.txt'.format(caso["num"]), dtype=int)
        matrizB = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizB.txt'.format(caso["num"]), dtype=int)

        for i in range(caso["numMuestras"]):
            N = len(matrizA)
            P = len(matrizA[0])
            M = len(matrizB[0])

            Result = [[0 for _ in range(M)] for _ in range(N)]

            tiempoInicio = time.time()
            matrizResultado = III4ParallelBlock.III4ParallelBlock(matrizA, matrizB, Result, caso["tamanioBloques"])
            tiempoFinalizacion = time.time() - tiempoInicio

            objeto.muestras.append(tiempoFinalizacion)

        objeto.promedio = objeto.calcular_promedio()
        resultado.casos.append(objeto)

    with open("Documentos\\Resultados\\python\\III4ParallelBlockResultadoPython.json", "w") as archivo:
        json.dump(resultado.to_json(), archivo)

def ejecutar_III3SequentialBlock():
    casosPrueba = ({"num": 1, "tam": 8, "numMuestras": 2, "tamanioBloques":4},
                {"num": 2,  "tam": 16, "numMuestras": 2, "tamanioBloques":8},
                {"num": 3,  "tam": 32, "numMuestras": 2, "tamanioBloques":16},
                {"num": 4,  "tam": 64, "numMuestras": 1, "tamanioBloques":32},
                {"num": 5,  "tam": 128, "numMuestras": 1, "tamanioBloques":64},
                {"num": 6,  "tam": 256, "numMuestras": 1, "tamanioBloques":64},
                {"num": 7,  "tam": 512, "numMuestras": 1, "tamanioBloques":128},
                {"num": 8,  "tam": 1024, "numMuestras": 1,  "tamanioBloques":128})
    resultado = Resultado.ResultadoMet("III3SequentialBlock", [], "python")

    for caso in casosPrueba:
        objeto = Caso.CasoMet(caso["tam"], [], 0,caso["tamanioBloques"])
        matrizA = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizA.txt'.format(caso["num"]), dtype=int)
        matrizB = np.loadtxt('Documentos\\Casos de prueba\\caso{}\\matrizB.txt'.format(caso["num"]), dtype=int)

        for i in range(caso["numMuestras"]):
            N = len(matrizA)
            P = len(matrizA[0])
            M = len(matrizB[0])

            Result = [[0 for _ in range(M)] for _ in range(N)]

            tiempoInicio = time.time()
            matrizResultado = III3SequentialBlock.III3SequentialBlockMet(matrizA, matrizB, Result, caso["tamanioBloques"])
            tiempoFinalizacion = time.time() - tiempoInicio

            objeto.muestras.append(tiempoFinalizacion)

        objeto.promedio = objeto.calcular_promedio()
        resultado.casos.append(objeto)

    with open("Documentos\\Resultados\\python\\III3SequentialBlockResultadoPython.json", "w") as archivo:
        json.dump(resultado.to_json(), archivo)

if __name__ == "__main__":
    ejecutar_III3SequentialBlock()
    ejecutar_III4ParallelBlock()
    ejecutar_IV3SequentialBlock()
    ejecutar_NaivLoopUnrollingFour()
    ejecutar_NaivLoopUnrollingTwo()
    ejecutar_NaivOnArray()
    ejecutar_StrassenNaiv()
    ejecutar_StrassenWinograd()
    ejecutar_V3SequentialBlock()
    ejecutar_WinogradOriginal()