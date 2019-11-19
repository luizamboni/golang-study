package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {

	// cria um contexto a partir do contexto pai (Background)
	// este contexto expirará após 1 segundo (mesmo que não seja cancelado)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	// na função main, este defer é redundante, pois o processo se encerrará
	// após a sua execução
	defer cancel()

	go func(ctx context.Context, cancel context.CancelFunc) {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			fmt.Println("char:", scanner.Text())
		}
		cancel()
	}(ctx, cancel)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
