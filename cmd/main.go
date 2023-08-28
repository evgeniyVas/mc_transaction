package main

import (
	"context"
	"fmt"
	"github.com/mc_transaction/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	a, err := app.New()
	if err != nil {
		panic(fmt.Sprintf("не удалось инициализировать приложение error: (%v)", err.Error()))
	}

	go a.Start()
	defer a.Stop(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
