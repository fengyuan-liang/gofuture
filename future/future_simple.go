// @Description:
// @Version: 1.0.0
// @Date: 2023/05/13 21:13
// @Author: fengyuan-liang@foxmail.com

package future

import "fmt"

// SimpleFuture a simple future func
func SimpleFuture(f func() (interface{}, error)) func() (interface{}, error) {
	var (
		result interface{}
		err    error
		c      = make(chan struct{}, 1)
	)
	go func() {
		// recover from panic
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic: %v", r)
			}
			close(c)
		}()
		result, err = f()
	}()
	return func() (interface{}, error) {
		<-c
		return result, err
	}
}
