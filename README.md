# cpdirall

Create a copy of all source folders and files to destination. It is called recursively and will respect nested folder structures.

```go
err := CpdirAll("src", "dst")
```

_The files mode and permission bits are preset and not configurable. This is a basic example for demonstration. There are many edge cases when considering full support._
