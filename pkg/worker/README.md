# example usage

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/money-stats/go-lib/worker"
)

func main() {
	exec := func(ctx context.Context) {
		fmt.Println("Work!")
	}
	workers := worker.New("example", 1, time.Second, exec)

	workers.Start(context.Background())
	workers.Trigger()
}
```