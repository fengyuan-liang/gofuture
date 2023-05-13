// @Description: future func
// @Version: 1.0.0
// @Date: 2023/05/13 21:12
// @Author: fengyuan-liang@foxmail.com

package future

import (
	"errors"
	"reflect"
	"time"
)

// Future type holds Result and state
type Future[T any] struct {
	Success          bool
	Done             bool
	Result           T
	InterfaceChannel <-chan T
	err              error
}

// Get return the result when available. This is a blocking call
func (f *Future[T]) Get() (T, error) {
	if f.Done {
		return f.Result, f.err
	}
	f.Result = <-f.InterfaceChannel
	f.Success = true
	f.Done = true
	return f.Result, f.err
}

// GetWithTimeout return the result until timeout.
func (f *Future[T]) GetWithTimeout(timeout time.Duration) (T, error) {
	if f.Done {
		return f.Result, f.err
	}
	timeoutChannel := time.After(timeout)
	select {
	case res := <-f.InterfaceChannel:
		f.Result = res
		f.Success = true
		f.Done = true
	case <-timeoutChannel:
		f.Result = reflect.Zero(reflect.TypeOf(f.Result)).Interface().(T)
		f.Done = true
		f.Success = false
		f.err = errors.New("timed out")
	}
	return f.Result, f.err
}

// FutureFunc creates a function that returns its response in future
func FutureFunc[T any](implem interface{}, args ...interface{}) *Future[T] {
	valIn := make([]reflect.Value, len(args), len(args))

	fnVal := reflect.ValueOf(implem)

	for idx, elt := range args {
		valIn[idx] = reflect.ValueOf(elt)
	}
	interfaceChannel := make(chan T, 1)

	go func() {
		res := fnVal.Call(valIn)
		// Only one result is supported
		if len(res) > 0 {
			interfaceChannel <- res[0].Interface().(T)
		} else {
			interfaceChannel <- reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
		}
	}()

	return &Future[T]{
		Success:          false,
		Done:             false,
		Result:           reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T),
		InterfaceChannel: interfaceChannel,
	}
}
