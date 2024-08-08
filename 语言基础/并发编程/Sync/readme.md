# sync.Mutex
共享资源互斥访问
```azure
mutex := sync.Mutex{}

mutex.Lock()
// update shared resource
mutex.Unlock()
```


# sync.RWMutex
读写互斥锁
```azure
mutex := sync.RWMutex{}
mutex.Lock()
// update shared resource
mutex.Unlock()
    
    

mutex.RLock()
// read shared resource
mutex.RUnlock()
```

# sync.WaitGroup
在并发环境下等待多个协程完成
```azure
wg := sync.WaitGroup{}

for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
    // do something 
    wg.Done()
    }()
}
wg.Wait()
```

# sync.Map
并发版本的go map，可以安全的并发读写
```azure
m := &sync.Map{}

// 添加元素
m.Store("key", "value")


// 读取元素
value, ok := m.Load("key")

// 删除元素
m.Delete("key")

// 遍历所有元素
m.Range(func(key, value interface{}) bool {
    // do something with key and value
    return true
})
``` 

# sync.Pool
并发池，可以缓存并发对象，减少创建和销毁对象的开销
```azure
pool := &sync.Pool{}
pool.Put(obj)

obj := pool.Get()
```