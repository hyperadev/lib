# hypera.dev/lib/util/retry

[![Go Reference](https://pkg.go.dev/badge/hypera.dev/lib/util/retry.svg)](https://pkg.go.dev/hypera.dev/lib/util/retry#section-documentation)

An easy-to-use retry utility.

```shell
go get -u hypera.dev/lib/util/retry
```

## Usage

```go
// Retry with exponential backoff (and default values).
err := retry.Exponential(ctx, func(ctx context.Context) error {
	if err := doSomething(ctx); err != nil {
		if errors.Is(err, io.EOF) {
			// This will not be retried.
			return retry.PermanentError(err)
		}
		// Retry the operation.
		return err
	}

	// Success!
	return nil
})
```

## Backoffs

### Constant

`1s -> 1s -> 1s -> 1s -> 1s -> 1s -> 1s`

### Exponential

Example: `1s -> 2s -> 4s -> 8s -> 16s -> 32s -> 64s`

Default: `500ms -> 750ms -> 1.125s -> 1.6875s -> 2.53125s -> 3.796875s -> 5.6953125s`

#### Jitter

## Extras

### Maximum retries

You can limit the number of retries with `retry.WithMaxRetries(Backoff, uint64)`.<br/>
*Note: This is **retries**, not attempts. The attempts would be `maxRetries + 1`.*

```go
b := retry.DefaultExponentialBackoff()

// Stop after 5 retries. In this example, the maximum elapsed time would be
// 500ms -> 750ms -> 1.125s -> 1.6875s -> 2.53125s (19.2525312s),
// excluding function execution time.
b = retry.WithMaxRetries(b, 5)
```

## Acknowledgements

