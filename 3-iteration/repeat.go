package iteration

import "strings"

const repeatCount = 5

// en GO los strings son inmutables, cada concatenación implica crear una nueva cadena
func Repeat(character string) string {
	var repeated string

	for i := 0 ; i < repeatCount ; i++ {
		repeated += character
	}

	return repeated
}

// esta función utiliza una librería que optimiza la operación de concatenación
func RepeatOptimized(character string) string {
	var repeated strings.Builder

	for i := 0 ; i < repeatCount ; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}