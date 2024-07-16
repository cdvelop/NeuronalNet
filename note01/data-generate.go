package note01

import (
	"math"
	"math/rand"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// HousePriceSamples genera un dataset sintético de precios de casas con valores redondeados a un decimal
func HousePriceSamples(total int) dataframe.DataFrame {
	rooms := make([]float64, total)
	prices := make([]float64, total)
	isException := make([]bool, total)

	for i := 0; i < total; i++ {
		// Generar número de habitaciones como valor continuo entre 3 y 9
		rooms[i] = roundToTwoDecimals(3 + rand.Float64()*6)

		// Calcular precio base con una relación no lineal con el número de habitaciones
		basePrice := 15000 + math.Pow(rooms[i], 2)*1500

		// Añadir variación aleatoria al precio (±20%)
		randomFactor := 0.8 + rand.Float64()*0.4
		prices[i] = roundToTwoDecimals(basePrice * randomFactor)

		// Añadir excepciones (5% de probabilidad)
		if rand.Float64() < 0.05 {
			isException[i] = true
			// Para excepciones, podemos tener casas caras con pocas habitaciones o baratas con muchas
			if rand.Float64() < 0.5 {
				// Casa cara con pocas habitaciones (ej. apartamento de lujo en el centro)
				rooms[i] = roundToTwoDecimals(3 + rand.Float64()*2)              // Entre 3 y 5 habitaciones
				prices[i] = roundToTwoDecimals(prices[i] * (2 + rand.Float64())) // 2 a 3 veces más caro
			} else {
				// Casa barata con muchas habitaciones (ej. casa antigua en las afueras)
				rooms[i] = roundToTwoDecimals(7 + rand.Float64()*2)                    // Entre 7 y 9 habitaciones
				prices[i] = roundToTwoDecimals(prices[i] * (0.5 + rand.Float64()*0.3)) // 50% a 80% más barato
			}
		}
	}

	return dataframe.New(
		series.New(rooms, series.Float, "rooms"),
		series.New(prices, series.Float, "price"),
		series.New(isException, series.Bool, "is_exception"),
	)
}

func roundToTwoDecimals(num float64) float64 {
	return math.Round(num*100) / 100
}
