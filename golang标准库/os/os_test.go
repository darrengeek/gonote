package os__test

import (
	"testing"
	"fmt"
	"os"
	"github.com/Echosong/beego_blog/models"
)

// Hostname返回内核提供的主机名。
func Test_Os(t *testing.T)  {
	fmt.Println(os.Hostname())
}

// Getpagesize返回底层的系统内存页的尺
func Test_GetPageSize(t *testing.T)  {
	fmt.Println(os.Getpagesize())
}

// Environ返回表示环境变量的格式为"key=value"的字符串的切片拷贝。
func Test_Environ(t *testing.T)  {
	fmt.Println(os.Environ())
}

// Getenv检索并返回名为key的环境变量的值。如果不存在该环境变量会返回空字符串。
func Test_Getenv(t *testing.T)  {
	fmt.Println(os.Getenv("GOROOT"))
}

// Setenv设置名为key的环境变量。如果出错会返回该错误。
func Test_Setenv(t *testing.T)  {
	fmt.Println(os.Setenv("my_name","xmge"))
	fmt.Println(os.Getenv("my_name"))
}

// Clearenv删除所有环境变量。
func Test_Clarenv(t *testing.T)  {
	os.Clearenv()
	fmt.Println(os.Environ())
}

// Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行。
func Test_Exit(t *testing.T)  {
	go func() {
		os.Exit(1)
	}()
	go func() {
		os.Exit(0)
	}()
}

// Expand函数替换s中的${var}或$var为mapping(var)。例如，os.ExpandEnv(s)等价于os.Expand(s, os.Getenv)。
func Test_Expand(t *testing.T)  {
	os.Expand("PATH", func(s string) string {
		fmt.Println(s)
		return s
	})
}