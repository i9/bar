# BetterAnyResolver

Proper handle unknown Any messages and output valid json string.

## Usage
```go
import "github.com/i9/bar"

m := jsonpb.Marshaler{
  AnyResolver: bar.BetterAnyResolver,
}

m.MarshalToString(someAnyMsg)
```
