package main

import (
	"context"
	"fmt"
	"log"
)

func run() error {
	app, err := NewApp()
	if err != nil {
		return fmt.Errorf("failed to build app: %w", err)
	}

	ctx := context.Background()

	catFact, err := app.GetCatFact(ctx)
	if err != nil {
		return fmt.Errorf("failed to get cat fact: %w", err)
	}

	log.Printf("Cat Fact: %s", catFact)

	return nil
}

func main() {
	closeTracer, err := ConfigureTracer()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = closeTracer()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = run()
	if err != nil {
		log.Fatal(err)
	}
}
