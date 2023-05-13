# go-future

[简体中文](https://github.com/fengyuan-liang/gofuture/blob/master/README_ZH.md)

Go-future gives an implementation similar to Java/Scala Futures.

Although there are many ways to handle this behaviour in Golang.
This library is useful for people who got used to Java/Scala Future implementation.


#### Import:

```golang
go get "github.com/fengyuan-liang/gofuture"
```

#### Usage:

```golang
futureFunc := future.FutureFunc[int](func() int {
time.Sleep(5 * time.Second)
return x * 10
})
// do something else here
// get result when needed
result, err := futureFunc.Get()
```

Also it is possible to use timeouts on Get

```golang
result, err := future.GetWithTimeout(3 * time.Second)
```

Or it can be used more elegantly

```go
futureFunc := future.FutureFunc[int](Fibonacci, 10)
result, err := futureFunc.Get()

// Fibonacci returns the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
```

you can handle func error
```go
futureFunc := future.FutureFunc[int](Add, 10, 20)
result, err := futureFunc.Get()

func Add(a, b int) (int, error) {
    return a + b, errors.New("you can return error")
}
```

#### Note:

This is a very basic implementation where only Get and GetWithTimeout functions are implemented.

Future Get returns an interface so type casting should be done by user.

Every future creates a channel but it is not closed so it is better allow garbage collection of Future after usage.



Java Futures: https://docs.oracle.com/javase/8/docs/api/index.html?java/util/concurrent/Future.html

Scala Futures: https://docs.scala-lang.org/overviews/core/futures.html