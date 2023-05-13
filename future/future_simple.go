// @Description:
// @Version: 1.0.0
// @Date: 2023/05/13 21:13
// @Author: fengyuan-liang@foxmail.com

package future

// SimpleFuture a simple future func
func SimpleFuture[T any](f func() (T, error)) func() (interface{}, error) {
	var (
		result T
		err    error
		c      = make(chan struct{}, 1)
	)
	go func() {
		defer close(c)
		result, err = f()
	}()
	return func() (interface{}, error) {
		<-c
		return result, err
	}
}
