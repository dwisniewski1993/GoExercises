package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Simple personalized logger")
	fmt.Println("With default level:")

	time.Sleep(1 * time.Millisecond)

	Console.Info("Info message example")
	Console.Debug("Debug message example", "Debbuger")
	Console.Warning("Warning message example")
	Console.Error("Error message example")
	Console.Critical("Critical message example")

	time.Sleep(1 * time.Millisecond)

	fmt.Println("With error level:")

	time.Sleep(1 * time.Millisecond)

	Console.SetLevel(LevelError)

	Console.Info("Info message example")
	Console.Debug("Debug message example", "Debbuger")
	Console.Warning("Warning message example")
	Console.Error("Error message example")
	Console.Critical("Critical message example")

	time.Sleep(1 * time.Millisecond)

	fmt.Println("With debug level:")

	time.Sleep(1 * time.Millisecond)
	Console.SetLevel(LevelDebug)

	Console.Info("Info message example")
	Console.Debug("Debug message example", "Debbuger")
	Console.Warning("Warning message example")
	Console.Error("Error message example")
	Console.Critical("Critical message example")
}
