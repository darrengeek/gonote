# Go 协程调度

## 一、进程线程和协程

1. 进程是在浏览器中启动一个应用时创建的。它实际是一个操作系统的级别的线程。这个操作级别的线程可以创将很多子线程，子线程又可以创建很多字线程。
2. 线程是在进程（操作系统级别的线程）中创建的执行路径。
3. 协程是线程中一部分。协程在线程中按照一定的顺序执行。
4. 操作系统会在茫茫线程中通过算法选择当前cpu要执行的线程。（进程只是对线程的分类）

如图：

## 二、协程调度方式

### 1. 代码测试

#### （1）无睡眠

代码：

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println("Starting Go Routines")


	go func() {
		defer wg.Done()
		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()


	go func() {
		defer wg.Done()
		for char:='A';char<'A'+26; char++{
			fmt.Printf("%c ", char)
		}
	}()


	go func() {
		defer wg.Done()
		for char:='a';char<'a'+26; char++{
			fmt.Printf("%c ", char)
		}
	}()



	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
```

结果：

```
Starting Go Routines
Waiting To Finish
a b c d e f g h i j k l m n o p q r s t u v w x y z 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
Terminating Program
```

### （2）睡眠1ms

代码：

```
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println("Starting Go Routines")


	go func() {
		defer wg.Done()
		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()


	go func() {
		defer wg.Done()
		for char:='A';char<'A'+26; char++{
			fmt.Printf("%c ", char)
		}
	}()


	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond)
		for char:='a';char<'a'+26; char++{
			fmt.Printf("%c ", char)
		}
	}()



	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
```

结果：

```
Starting Go Routines
Waiting To Finish
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 A B C D E F G H I J K L M N O P Q R S T U V W X Y Z a b c d e f g h i j k l m n o p q r s t u v w x y z
Terminating Program
```

### （3）睡眠1ns

代码：

```
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println("Starting Go Routines")


	go func() {
		defer wg.Done()
		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()


	go func() {
		defer wg.Done()
		for char:='A';char<'A'+26; char++{
			fmt.Printf("%c ", char)
		}
	}()


	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Nanosecond)
		for char:='a';char<'a'+26; char++{
			fmt.Printf("%c ", char)
		}
	}()



	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
```

结果：

```
Waiting To Finish
a b c d e f g h i j k l m n o p q r s t u v w x y z 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
Terminating Program
```

### 2. 结论

如图：

