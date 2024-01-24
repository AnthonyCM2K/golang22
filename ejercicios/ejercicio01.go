package ejercicios

/* Ejercicio 01
1. Crear un paquete nuevo llamado 'ejercicios'
2. Crear un archivo GO en ese paquete llamado 'ejercicio01.go'
3. Crear una función Pública, que devuelva 2 valores (int y string) y que reciba de parámetro un valor 'string'
4. El parámetro de tipo String recibido deberá ser convertido a entero y si el entero es > a 100, el String retornado debe decir "Es mayor a 100" de lo contrario, devolver el mensaje "Es menor a 100"
5. El valor numérico entero retornado debe corresponder al string convertido
6. En Main, llamar a dicha función asignándola a 2 variables y luego mostrar dichas variables por consola.
7. Grabar, ejecutar y luego subir todo a Github */
import (
	"strconv"
)

func ConversorNumerico(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}

}
