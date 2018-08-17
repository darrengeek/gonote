package builtin_test
import (
	"testing"
	"fmt"
)


// 获取复数的实数部分
func TestReal(t *testing.T)  {
	var f = complex(1,2)
	fmt.Println(real(f))
}

// 获取复数的虚数部分
// 翻译： 虚数[imaginary number][iˈmædʒinəri ˈnʌmbə]
func TestImag(t *testing.T)  {
	var f = complex(1,2)
	fmt.Println(imag(f))
}

// 生成一个复数
func TestComplex(t *testing.T)  {
	var f complex64 = complex(1,2)
	fmt.Print(f)
}


// 创建结构体对象 初始化并返回对象指针
func TestNew(t *testing.T)  {
	var f *complex64 = new(complex64) 		// 这种定义方式可以提高代码可读性
	fmt.Print(f)
}

// 定义 slice、map、channel
func TestMake(t *testing.T)  {
	var s = make([]string,10,1)
}

// slice的容量
func TestCap(t *testing.T)  {

}

//
func TestLen(t *testing.T)  {

}

func TestAppend(t *testing.T)  {

}

func TestCopy(t *testing.T)  {

}

func TestDelete(t *testing.T)  {

}

func TestClose(t *testing.T)  {

}

func TestRecover(t *testing.T)  {

}

func TestPrint(t *testing.T)  {

}

func TestPrintln(t *testing.T)  {

}