package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type callBackChan chan struct{}

// Wywołuje kanał callBack co 'd' jednostek
func checkEvery(ctx context.Context, d time.Duration, cb callBackChan) {
	for {
		select {
		case <-ctx.Done():
			//ctx odwołane
			return
		case <-time.After(d):
			// Oczekiwanie na czas okna
			if cb != nil {
				cb <- struct{}{}
			}

		}
	}
}

func returnProcessList(os string) {
	var psCommand *exec.Cmd

	// Komenda określana w zależności od systemu - Windows i inne
	if os == "windows" {
		psCommand = exec.Command("Tasklist", "/svc")
	} else {
		psCommand = exec.Command("ps", "a")
	}

	resp, err := psCommand.CombinedOutput()
	if err != nil {
		log.Fatal("main.go [ERROR]: command execution failed")
	}

	out := string(resp)
	lines := strings.Split(out, "\n")

	for _, line := range lines {
		if line != "" {
			fmt.Println(line)
		}
	}
}

func main() {
	log.Printf("main.go [INFO]: Process Monitor Application")

	// Detekcja systemu operacyjnego
	const os string = runtime.GOOS

	// Wyświetlenie listy procesów niezwłocznie po starcie aplikacji
	ctx := context.Background()
	returnProcessList(os)

	// Kanał do porozumiewania miedzy goroutine'ami
	callBack := make(callBackChan)

	// Goroutine sprawdzania
	go checkEvery(ctx, 5*time.Second, callBack)

	// Goroutine wyświetlania
	go func() {
		for {
			select {
			case <-callBack:
				returnProcessList(os)
			}
		}
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}
