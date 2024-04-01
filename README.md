# goBtreeGeneric
goBtreeGeneric

Simple example of using the Google Btree library with generics
https://github.com/google/btree/
https://pkg.go.dev/github.com/google/btree

This example allows sorting of items by time.Time
```bash
type Descriptions struct {
	Description string
	Time        time.Time
}
```

Another example of a double nested sync map
