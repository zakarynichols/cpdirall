# cpdirall

![Workflow](https://github.com/zakarynichols/cpdirall/actions/workflows/ci.yml/badge.svg?branch=master)

Create a copy of all source folders and files to destination.

```
$ go get github.com/zakarynichols/cpdirall
```

```go
import (
    "github.com/zakarynichols/cpdirall/dir"
)

err := dir.Cp("src", "dst")
```
