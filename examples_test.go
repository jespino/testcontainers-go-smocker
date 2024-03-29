package smocker_test

import (
	"context"
	"fmt"
	"log"

	smocker "github.com/jespino/testcontainers-go-smocker"
	"github.com/testcontainers/testcontainers-go"
)

func ExampleRunContainer() {
	ctx := context.Background()

	smockerContainer, err := smocker.RunContainer(ctx, testcontainers.WithImage("thiht/smocker:0.18.5"))
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	// Clean up the container
	defer func() {
		if err := smockerContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	state, err := smockerContainer.State(ctx)
	if err != nil {
		log.Fatalf("failed to get container state: %s", err) // nolint:gocritic
	}

	fmt.Println(state.Running)

	// Output:
	// true
}
