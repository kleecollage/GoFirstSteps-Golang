package handleErrors

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		// return 0, errors.New("no es posible dividir entre 0")
		return 0, fmt.Errorf("no es posible dividir entre 0") // error personalizado
	}
	return dividendo / divisor, nil
}

func handleError() {
	result, errorResponse := divide(10, 0)
	if errorResponse != nil {
		fmt.Println("Error: ", errorResponse)
		return
	}
	fmt.Println("Resultado", result)

	str := "12321@"
	// con Atoi convertimos string a int
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Nuemero: ", num)
}

func myDefer() {
	// defer prolonga la ejecucion
	defer fmt.Println(3)
	defer fmt.Println(2)
	// con una pila de defers, el primero en ser delcarado sera el ultimo en ser ejecutado
	defer fmt.Println(1)

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // prolongamos la ejecucion para cerrar el flujo
	_, err = file.Write([]byte("Hola Alex, Roel"))
	if err != nil {
		fmt.Println(err)
		// file.Close() // ahorramos estas lineas con defer
		return
	}
	// file.Close()
}

func divide2(dividendo, divisor int) {
	// Funcion anonima
	defer func() {
		// capturamos el panico con recover sin interrupir el programa
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividir por cero")
	}
}

func panicRecover() {
	divide2(100, 10)
	divide2(200, 25)
	divide2(34, 0)
	divide2(100, 4)
}

func myLogs() {
	log.Print("¿Puedes verme?") // log nos da un registro del tiempo
	log.SetPrefix("myLogs(): ")
	log.Fatal("¡Oye, soy un regitro de errores!") // aqui la aplicacion se detiene
	log.Panic("¡Oye, soy un regitro de errores!") // aqui la aplicacion se detiene y nos da detalles
}

func myLogsFile() {
	log.SetPrefix("myLogsFile(): ")
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("¡Ojo, soy un Log!")
}
