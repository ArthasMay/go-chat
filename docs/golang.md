### Golang的一些语法

#### 1. sync.RWMutex 不需要初始化
The mutexes are generally designed to work without any type of initialization. That is, a zero valued mutex is all you need to use them. To fix your program, declare foo as a value. Not a pointer to a mutex:

互斥锁通常被设计为不需要任何类型的初始化就可以工作。也就是说，只需要使用一个零值互斥对象即可。要修复程序，请将 foo 声明为一个值。不是互斥对象的指针: