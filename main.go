package goarea

import "math"

//Pi é uma proporção numérica definida pela relação entre o perímetro de uma circunferência e seu diametro
const Pi = 3.1416

//Circulo calcula área da circunferência
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é responsável por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visível!
func _TrianguloEquilatero(base, altura float64) float64 {
	return (base * altura) / 2
}
