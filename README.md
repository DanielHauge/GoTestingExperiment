Golang Testing Experiments
=========================
This repository is to learn how to do unit testing in golang

## usefull links
- [5 Good tips for testing](https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742)
- [Go vet, Code analysis looking at correctness](https://golang.org/cmd/vet/)
- [Go lint, Coding standards, used by google](https://github.com/golang/lint)
- [Go Import, importing dependices](https://godoc.org/golang.org/x/tools/cmd/goimports)
- [Go testing documentation](https://golang.org/pkg/testing/)
- [Go Test Coverage](https://golang.org/cmd/cover/)

## golint results

```
TestHelperFunction.go:1:1: don't use MixedCaps in package name; GoTestingExperiment should be gotestingexperiment
Triangle.go:1:1: don't use MixedCaps in package name; GoTestingExperiment should be gotestingexperiment
Triangle.go:11:6: exported type Triangle should have comment or be unexported
Triangle.go:18:1: error should be the last type when returning multiple items
Triangle.go:18:1: exported function MakeTriangle should have comment or be unexported
Triangle.go:30:1: error should be the last type when returning multiple items
Triangle.go:30:1: exported function CalculateTriangleType should have comment or be unexported
Triangle.go:50:1: exported function CalculateAngles should have comment or be unexported
Triangle.go:79:1: exported function CalculateArea should have comment or be unexported
Triangle.go:90:1: exported method Triangle.PrintInfo should have comment or be unexported
Triangle_test.go:1:1: don't use MixedCaps in package name; GoTestingExperiment should be gotestingexperiment
```
