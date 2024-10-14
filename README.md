***Nota: Este proyecto contiene un paquete para usarse ***
# Saludos en Go
Este paquete proporciona una lista de salidos que pueden usarse en Go

## Instalacion
Ejecuta el siguiente comando para instalar  el paquete:
```bash
go get -u github.com/kleecollage/GoFirstSteps-Golang/greetings
```

## Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu codigo

```go
import (
	"fmt"
	"github.com/kleecollage/GoFirstSteps-Golang/greetings"
)

func main() {
	message, err := greetings.Hello("Jane Smith")
  if err != nil {
    fmt.Println("Ocurrio un error: ", err)
    return
  }
	fmt.Println(message)
}
```
Este ejemplo iporta el paquete y llama a la funcion Hello que contiene un saludo personalizado. En caso de que ocurra un error, se imprime el mensaje de error
