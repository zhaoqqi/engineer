##### Type and Interface
反射是建立在类型系统基础之上的，要想学习反射，需从类型系统开始。   
Go是静态类型语言，每个变量必须有其静态类型。   
变量静态类型为其声明时的类型，哪怕两个变量底层类型一样，在不做类型转换的情况下也不能相互赋值，比如：   
```golang
type MyInt int
var i int
var j MyInt
```   
interface 类型，代表了固定的方法集合。关于 interface 类型经典的例子就是 io.Reader和io.Writer。
```golang
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Write interface {
    Writer(p []byte) (n int, err error)
}
```   
所有实现了Read() 或 Write()方法的类型，都是一个 Reader 或 Writer。
```golang
var r io.Reader
r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)
// and so on
```   
   
极端的例子是空接口 interface{}，代表了空的方法集合，所以其满足所有类型。
   
##### The first law of reflection
##### 1. Reflection goes from interface value to reflection object.
在基础层面上，reflection只是用来检测存储在 interface变量中的 <类型和值 对>的一种机制。   
在reflect包中有两种类型，reflect.type和reflect.value。   
reflect.TypeOf 返回类型 reflect.Type   
reflect.ValueOf 返回类型 reflect.Value   
先看看 reflect.Type
```golang
func main() {
    var x float64 = 3.4
    fmt.Println("type:", reflect.TypeOf(x))
}
输出：
type: float64
```   
   
TypeOf方法实现如下：   
func TypeOf(i interface{}) Type    
其先将参数 x 转换为 空接口 interface{}，再传入作为方法参数。   
   

reflect.Type 和 reflect.Value 均有很多方法去做变量的检查和操作（examine and manipulate）。   
比如 reflect.Value 的 Type()方法，Kind()方法，Float()方法等，见下面代码：   
```golang
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type: ", v.Type())
fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
fmt.Println("value: ", v.Value())
输出：
type: float64
kind is float64: true
value: 3.4
```    
reflect.Value 还有 setInt() 和 setFloat() 方法，但只有操作对象是 settability 才可以。   
Kind() 方法返回的是变量底层的类型，比如返回 int64 而不是自定义的 MyInt。   
   
##### The second law of reflection
##### 2. Reflection goes from reflection object to interface value
记住一点，Interface() 方法与 ValueOf() 方法为护逆方法。   
func Interface(v Value) interface{}    
```golang
y := v.Interface().(float64) // y will have type float64
fmt.Println(y)
```    
    
fmt系列的 Print方法，可以把传递进去的 Interface 类型转换为真实的运行时类型打印出来：   
fmt.Println(v.Interface())   
比如，fmt.Println("value is %7.1e", v.Interface())   
输出：3.4e+00   
   
##### The third law of reflection
##### 3. To modify a reflection object, the value must be settlable.
记住一点，什么叫做 settlable？即按照指针传递的就是settlable，按照值传递的就是 unsettlable。比如：   
```golang
var x float64 = 3.4
v := reflect.ValueOf(&x)
fmt.Println("type of p:", p.Type())
fmt.Println("settlability of p:", p.CanSet())
输出：
type of p: *float64
settlability of p: false
```
   
反射对象 p 不能 set，但我们要 set 的是 *p 而不是 p。我们可以使用 Elem() 方法获取 p 指向的对象。   
```golang
v := p.Elem()
fmt.Println("settlability of v", v.CanSet())
输出：
settlability of v: true
```
