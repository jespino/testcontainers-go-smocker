# Testcontainers Smocker module

## Introduction

The Testcontainers module for Smocker.

## Adding this module to your project dependencies

Please run the following command to add the Smocker module to your Go dependencies:

```
go get github.com/jespino/testcontainers-go-smocker
```

## Usage example

```go
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
```

## Module reference

The Smocker module exposes one entrypoint function to create the Smocker container, and this function receives two parameters:

```golang
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*SmockerContainer, error)
```

- `context.Context`, the Go context.
- `testcontainers.ContainerCustomizer`, a variadic argument for passing options.

### Container Options

When starting the Smocker container, you can pass options in a variadic way to configure it.

#### Image

If you need to set a different Smocker Docker image, you can use `testcontainers.WithImage` with a valid Docker image
for Smocker. E.g. `testcontainers.WithImage("thiht/smocker:0.18.5")`.

You can use the standar testcontainers-go functional options to define your
container, like `WithEnv`, `WithImage`, or `WithLogConsumers`.

### Container Methods

The Smocker container exposes the following methods:

#### MockURL

This method returns the Mock URL to connect to the Smocker container, using the default `8080` port.
This URL is the one that you will use to connect your application instead of the real service.

```go
mockUrl, err := container.MockURL(ctx)
```

#### ApiURL

This method returns the API URL to connect to the Smocker container, using the default `8081` port.
This URL is the one that you will use to configure your request/response mocks.

```go
apiURL, err := container.ApiURL(ctx)
```
