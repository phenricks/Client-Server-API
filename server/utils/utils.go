package utils

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func PrintBanner() {
	fmt.Printf(`
   ___                               ___           _           _   ___ ___ 
  / __|  _ _ _ _ _ ___ _ _  __ _  _ / _ \ _  _ ___| |_ ___    /_\ | _ \_ _|
 | (_| || | '_| '_/ -_) ' \/ _| || | (_) | || / _ \  _/ -_)  / _ \|  _/| | 
  \___\_,_|_| |_| \___|_||_\__|\_, |\__\_\\_,_\___/\__\___| /_/ \_\_| |___|
                               |__/`)
}

func CreateContext(number int) (context.Context, context.CancelFunc) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(number))
	go func() {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				fmt.Println("Tempo limite atingido. A operação demorou muito para ser concluída.")
			} else {
				fmt.Printf("Erro inesperado: %v\n", ctx.Err())
			}
		}
	}()

	return ctx, cancel
}
