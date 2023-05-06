# Simple package for golang web errors with stacktrace

### Getting started
```bash
go get -u github.com/brmatvey/weberrors
```
### Using package
For creating error without message just call constructor
```go
err := weberrors.New(http.StatusNotFound)
```
For creating error with message use constructor with args
```go
err := weberrors.NewWithArgs(http.StatusNotFound, "must %d %d %d", 1, 2, 3)
```
For casting error to weberrors type use general golang cast. You can implement it, for instance, in general wrapping middleware.
```go
castedErr, ok := err.(*weberrors.Error)
if !ok {
    t.Fatal("must be casted")
}
```