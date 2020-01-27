package p1_base

import (
	"sync/atomic"
)

var config atomic.Value

/*
初始化配置
config.Store(loadConfig())
*/

/*
后台goroutine定时更新配置
go func() {
for {
time.Sleep(time.Second)
config.Store(loadConfig())
}
}()
*/

/*前台多个工作 者线程获取最新的配置信息
for i := 0; i < 10; i++ {
go func() {
for r := range requests() {
c := config.Load()
// ...
}
}()
}
*/
