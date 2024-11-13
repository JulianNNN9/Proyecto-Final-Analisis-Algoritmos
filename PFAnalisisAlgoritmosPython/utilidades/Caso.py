class CasoMet:
    def __init__(self, tam, muestras, promedio, tamanioBloques):
        self.tam = tam
        self.muestras = muestras
        self.promedio = promedio
        self.tamanioBloques = tamanioBloques

    def calcular_promedio(self):
        suma = 0
        for muestra in self.muestras:
            suma += muestra
        return suma / len(self.muestras)

    def to_json(self):
        return {
            "tam": self.tam,
            "muestras": self.muestras,
            "promedio": self.promedio,
            "tamanioBloques": self.tamanioBloques
        }
