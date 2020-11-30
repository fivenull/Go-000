# Go进阶笔记-关于error



很多人对于Go的error比较吐槽，说代码中总是会有大量的如下代码：

```
if err != nil {
    ...
}
```

其实很多时候是使用的姿势不对，或者说，对于error的用法没有完全理解，这里整理一下关于Go中的error

## 关于源码中的error

先看一下go源码中`go/src/builtin/builtin.go`对于error的定义：


```
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
        Error() string
}
```

我们使用的时候经常会通过errors.New() 来返回一个error对象，这里可以看一下我们调用errors.New()的这段源码文件`go/src/errors/errors.go`,可以看到errorString实现了error解接口，而errors.New()其实返回的是一个 `&errorString{text}` 即errorString对象的指针。


```
package errors

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

如果之前看过一些优秀源码或者go源码的，会发现代码中通常会定义很多自定义的error，并且都是包级别的变量，即变量名首字母大写:

```
// https://golang.org/pkg/bufio


var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
    ErrBufferFull        = errors.New("bufio: buffer full")
    ErrNegativeCount     = errors.New("bufio: negative count")
)
```

注意：自己之后在代码中关于这种自定义错误的定义，也要参照这种格式规范定义。
**"当前的包名：错误信息"**


```
package main

import (
	"errors"
	"fmt"
)

type errorString string

// 实现 error 接口
func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	return errorString(text)
}

var errNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

func main() {
	//  这里其实就是两个结构体值的比较
	if errNamedType == New("EOF") {
		fmt.Println("Named Type Error")   // 这行打印会输出
	}
	// 标准库中errors.New() 返回的是一个地址，每次调用都会返回一个新的内存地址
	// 标准库这样设计也是为了避免碰巧如果两个结构体值相同了，而引发一些不期望的问题
	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")  // 这行打印不会输出
	}
}
```

关于结构体值的比较：

如果两个结构体值的类型均为可比较类型，则它们仅在它们的类型相同或者它们的底层类型相同（要考虑字段标签）并且其中至少有一个结构体值的类型为非定义类型时才可以互相比较。
如果两个结构体值可以相互比较，则它们的比较结果等同于逐个比较它们的相应字段


**注意：关于Go中函数支持多参数返回，如果函数有error的通常把返回值的最后一个参数作为error**

如果一个函数返回（value, error）这个时候必须先判定error
Go中的panic 意味着程序挂了不能继续运行了，不能假设调用者来解决panic。

对于刚学习go的时候经常用如下代码开启一个goroutine执行任务：
```
go func() {
    ...
}
```
这种情况也叫野生goroutine,并且这个时候recover是不能解决的。

可以定义一个包,通过调用该包中的Go() 方法来开goroutine，来避免野生goroutine。

```
package sync

func Go(x func()) {
    
    if err := recover(); err != nil {
        ....
    }
    go x()
}
```

关于代码的panic 通常在代码中是很少使用的，只有在极少情况下，我们需要panic，如我们项目的初始化地方连接数据库连接不上，并且这个时候，数据库是我们程序的强依赖，那么这个时候是可以panic。

下面通过一个例子来演示error的使用姿势：


```
package main

import (
	"errors"
	"fmt"
)

// 判断正负数
func Positivie(n int) (bool, error) {
	if n == 0 {
		return false, errors.New("undefined")
	}
	return true, nil
}

func Check(n int) {
	pos, err := Positivie(n)
	if err != nil {
		fmt.Println(n, err)
		return
	}
	if pos {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

func main() {
	Check(1)
	Check(0)
	Check(-1)
}

```

上面是一种非常正确的姿势，我们通过返回`(value, error)` 这种方式来解决，也是非常go 的一种写法，只有`err!=nil` 的时候我们的`value`才有意义

那么在实际中可能有很多各种姿势来解决上述的问题，如下：


```
package main

import "fmt"

func Positive(n int) *bool {
	if n == 0 {
		return nil
	}
	r := n > -1
	return &r
}

func Check(n int) {
	pos := Positive(n)
	if pos == nil {
		fmt.Println(n, "is neither")
		return
	}
	if *pos {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

func main() {
	Check(1)
	Check(0)
	Check(-1)
}
```

另外一种姿势：

```
package main

import "fmt"

func Positive(n int) bool {
	if n == 0 {
		panic("undefined")
	}
	return n > -1
}

func Check(n int) {
	defer func() {
		if recover() != nil {
			fmt.Println("is neither")
		}
	}()
	
	if Positive(n) {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

func main() {
	Check(1)
	Check(0)
	Check(-1)
}

```

上面这两种姿势虽然也可以实现这个功能，但是非常的不好，也不推荐使用。在代码中尽可能还是使用`(value, error)` 这种返回值来解决error的情况。

对于真正意外的情况，那些不可恢复的程序错误，例如索引越界，不可恢复的环境问题，栈溢出等才会使用panic,对于其他的情况我们应该还是期望使用error来进行判定。 

## error 处理套路

### Sentinel Error 预定义error

通常我们把代码包中如下的这种error叫预定义error.

```
// https://golang.org/pkg/bufio


var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
    ErrBufferFull        = errors.New("bufio: buffer full")
    ErrNegativeCount     = errors.New("bufio: negative count")
)
```

这种姿势的缺点：

- 对于这种错误，在实际中的使用中我们通常会使用 `if err == ErrSomething {....}` 这种姿势来进行判断。但是也不得不说，这种姿势是最不灵活的错误处理策略，并且不能对于错误提供有用的上下文。
- Sentinel errors 成为API的公共部分。如果你的公共函数或方法返回一个特定值的错误，那么该错误就必须是公共的，当然要有文档记录，这最终会增加API的表面积。
- Sentinel errors 在两个包之间创建了依赖。对于使用者不得不导入这些错误，这样就在两个包之间建立了依赖关系，当项目中有许多类似的导出错误值时，存在耦合，项目中的其他包必须导入这些错误值才能检查特定的错误条件。


### Error types

Error type 是实现了error接口的自定义类型，例如MyError类型记录了文件和行号以展示发生了什么


```
type MyError struct {
    Msg string
    File string
    Line int
}

func (e *MyError) Error() string {
    return fmt.Sprintf("%s:%d:%s", e.File,e.Line, e.Msg)
}

func test() error {
    return &MyError("something happened", "server.go", 11)
}

func main() {
    err := test()
    switch err := err.(type){
    case nil:
        // ....
    case *MyError:
        fmt.Println("error occurred on line:", err.Line)
    default:
        // ....
    }
}

```

这种方式其实在标准库中也有使用如os.PathError


```
// https://golang.org/pkg/os/#PathError

type PathError struct {
    Op   string
    Path string
    Err  error
}
```

调用者要使用类型断言和类型switch，就要让自定义的error变成public，这种模型会导致和调用者产生强耦合，从而导致API变得脆弱。

### Opaque errors

这种方式也称为不透明处理，这也是相对来说比较优雅的处理方式,如下

```
func fn() error {
    
    x, err := bar.Foo()
    if err != nil {
        return err
    }
    // use x
}
```
这种不透明的实现方式，一种比较好的用法，这里以net库的代码来看：

```
// https://golang.org/pkg/net/#Error

type Error interface {
    error
    Timeout() bool   // Is the error a timeout?
    Temporary() bool // Is the error temporary?
}
```

这里是定义了一个Error接口，而让其他需要用到error的来实现这个接口，如net中的下面这个错误

```
// https://golang.org/pkg/net/#DNSConfigError

type DNSConfigError
    func (e *DNSConfigError) Error() string
    func (e *DNSConfigError) Temporary() bool
    func (e *DNSConfigError) Timeout() bool
    func (e *DNSConfigError) Unwrap() error
```

按照这个方式实现我们使用net时的异常处理可能就是如下情况：

```
if neerr, ok := err.(net.err); ok && nerr.Temporary() {
    time.Sleep(time.Second * 10)
    continue
}
if err != nil {
    log.Fatal(err)
}
```

其实这样还是不够优雅，好的方式是我们卡一定义temporary的接口，然后取实现这个接口，这样整体代码就看着非常简洁清楚，对外我们就只需要暴露IsTemporary方法即可，而不用外部再进行断言。

```
Type temporary interface {
    Temporary() bool
}

func IsTemporary(err error) bool {
    te, ok := err.(temporary)
    return ok && te.Temporary()
}

```

以上这几种姿势，其实各有各的用处，不同的场景，选择可能也不同，需要根据实际场景实际分析。