package utilidades

// ResultadoMet define la estructura para los resultados
type ResultadoMet struct {
	Nombre   string
	Casos    []CasoMet
	Lenguaje string
}

// ToJSON convierte el resultado a un formato JSON
func (resultado *ResultadoMet) ToJSON() map[string]interface{} {
	var casosJSON []map[string]interface{}
	for _, caso := range resultado.Casos {
		casosJSON = append(casosJSON, caso.ToJSON())
	}
	return map[string]interface{}{
		"nombre":   resultado.Nombre,
		"casos":    casosJSON,
		"lenguaje": resultado.Lenguaje,
	}
}
