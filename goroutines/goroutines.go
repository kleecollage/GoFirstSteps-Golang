package goroutines

import (
	"fmt"
	"net/http"
	"time"
)

func TimeLapse() {
	// canal := make(chan int) // declaramos el canal
	// canal <- 15             // enviamos datos al canal
	// valor := <- canal // recibimos datos del canal

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com", // esta no existe
		"https://graph.microsoft.com",
	}

	ch := make(chan string)
	for _, api := range apis {
		go checkAPI(api, ch) // aplicamos la concurrencia con go
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	// time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomo %v segundos!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR: ¡%s esta caido!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS ¡%s esta en funcionamiento!\n", api)
}
