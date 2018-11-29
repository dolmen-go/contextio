package contextio_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/dolmen-go/contextio"
)

func ExampleNewWriter() {
	// interrupt context after 500ms
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// interrupt context with SIGTERM (CTRL+C)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs
		cancel()
	}()

	f, err := os.OpenFile("/dev/null", os.O_WRONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Writer that fails when context is canceled
	out := contextio.NewWriter(ctx, f)

	// Infinite loop. Will only be interrupted once write fails.
	for {
		if _, err := out.Write([]byte{'a', 'b', 'c'}); err != nil {
			fmt.Println("Err:", err)
			break
		}
	}

	fmt.Println("Closing.")

	// Output:
	// Err: context deadline exceeded
	// Closing.
}
