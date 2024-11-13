class ResultadoMet:
    def __init__(self, nombre, casos, lenguaje):
        self.nombre = nombre
        self.casos = casos
        self.lenguaje = lenguaje
    def to_json(self):
        return {
            "nombre": self.nombre,
            "casos": [caso.to_json() for caso in self.casos],
            "lenguaje": self.lenguaje
        }