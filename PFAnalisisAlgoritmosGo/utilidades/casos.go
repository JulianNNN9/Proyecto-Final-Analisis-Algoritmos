package utilidades

// CasoMet define la estructura para los casos de prueba
type CasoMet struct {
	Tam            int
	Muestras       []float64
	Promedio       float64
	TamanioBloques int
}

// CalcularPromedio calcula el promedio de las muestras
func (caso *CasoMet) CalcularPromedio() float64 {
	var suma float64
	for _, muestra := range caso.Muestras {
		suma += muestra
	}
	return suma / float64(len(caso.Muestras))
}

// ToJSON convierte el caso a un formato JSON
func (caso *CasoMet) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"tam":            caso.Tam,
		"muestras":       caso.Muestras,
		"promedio":       caso.Promedio,
		"tamanioBloques": caso.TamanioBloques,
	}
}
