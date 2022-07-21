# contextio - Context-aware I/O streams for Go

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/dolmen-go/contextio)
[![stability: stable](https://img.shields.io/badge/stability-stable-green.svg)](https://github.com/emersion/stability-badges#stable)
[![codecov](https://codecov.io/gh/dolmen-go/contextio/branch/master/graph/badge.svg)](https://codecov.io/gh/dolmen-go/contextio)
[![Travis-CI](https://api.travis-ci.org/dolmen-go/contextio.svg?branch=master)](https://travis-ci.org/dolmen-go/contextio)
[![Go Report Card](https://goreportcard.com/badge/github.com/dolmen-go/contextio)](https://goreportcard.com/report/github.com/dolmen-go/contextio)

Author: [@dolmen](https://github.com/dolmen)  (Olivier Mengué).

# Example

`go test -run ExampleWriter`

```go
func main() {
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
}
```

# See Also

* [github.com/jbenet/go-context/io](https://godoc.org/github.com/jbenet/go-context/io) Context-aware reader and writer
* [github.com/northbright/ctx/ctxcopy](https://godoc.org/github.com/northbright/ctx/ctxcopy) Context-aware io.Copy
* [gitlab.com/streamy/concon](https://godoc.org/gitlab.com/streamy/concon) Context-aware net.Conn
