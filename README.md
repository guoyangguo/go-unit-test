# go-unit-test

## go unit-test 用来演示golang中的单元测试

## Go语言测试

Go语言的测试方法看起来比较‘简单’，这也符合Go语言的设计理念，Go语言的单元测试语法与编程Go的程序代码并无差异，并不需要添加多余的配置，只是聚焦的重点在于某个函数或者某个模块。

## Go 语言测试工具

Go语言中使用 go test 命令驱动测试程序，这与go build 命令的功能相类似，只不过go test 命令用来驱动 *_test.go 文件，而 go build 命令则会忽略_test.go 文件，所以Go语言编译后的二进制文件中并不包含测试文件； go test 命令会扫描_test.go文件，并且查找Test函数，最后生成一个临时的main包来调用它们，然后进行编译、运行、返回结果，最后清空临时文件。

### 编写Test函数

Go语言中每个测试文件必须是*_test.go文件，每个_test.go文件中的测试函数签名如下：

``` golang
func TestXxx(t *testing.T) {}
```
