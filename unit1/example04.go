package main
/**
 * Programa que imprime los tipos de datos simples por pantalla
 * @author: Dicxie Madrigal Brack <dicxiemadrigal@gmail.com>
 * @date: 20 september 2021
 * @licence: Licencie MIT
*/
import "fmt"

func main() {
	var cadena string
	var entero int
	var booleano bool
	cadena = "cadena de texto"
	entero = 230
	booleano = true
	fmt.Println(cadena, entero, booleano)
}
