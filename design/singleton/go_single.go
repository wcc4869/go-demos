package singleton

import "sync"

var (
	instance *Instance
	lock     sync.Mutex
)

type Instance struct {
	Name string
}

// 常规的获取单例
func GetInstance(name string) *Instance {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Instance{Name: name}
		}
	}
	return instance
}


// go 语言风格 单例
var (
	goInstance *Instance
	once       sync.Once
)

func GoInstance(name string) *Instance {
	if goInstance == nil {
		once.Do(func() {
			goInstance = &Instance{Name: name}
		})
	}
	return goInstance
}
