package contextio_test

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/dolmen-go/contextio"
)

func Example_copy() {
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

	fIn, err := os.Open("/dev/zero")
	if err != nil {
		log.Fatal(err)
	}
	defer fIn.Close()

	fOut, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer fOut.Close()

	// Reader that fails when context is canceled
	in := contextio.NewReader(ctx, fIn)
	// Writer that fails when context is canceled
	out := contextio.NewWriter(ctx, fOut)

	n, err := io.Copy(out, in)
	log.Println(n, "bytes copied.")
	if err != nil {
		fmt.Println("Err:", err)
	}

	fmt.Println("Closing.")

	// Output:
	// Err: context deadline exceeded
	// Closing.
}

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

	f, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
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
