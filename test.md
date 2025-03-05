## 简介
Go语言的标准库非常的丰富（数据结构除外）功能也十分强大，API对开发人员而言十分友好，能够熟练运用标准库对日常开发将会有非常大的帮助，我们可以不用任何外部依赖，仅使用`net/http`，`html/template`，`database/sql`这三个库就可以开发一个独立完整的Web项目。

下面列出了所有的标准库：
- `html`，html模板操作包。
- `net`，提供一系列网络操作的包，相当的强大。
- `reflect`，反射操作包。
- `sort`，排序操作包。
- `regex`，正则表达式包。
- `image`，图像操作的包。
- `unicode`，unicode字符集操作包。
- `unsafe`，提供没有类型限制的指针操作。
- `syscall`，操作系统提供的系统调用包。
- `testing`，测试包，主要是写测试的时候用。
- `sync` ，并发操作包。
- `context`，上下文包。
- `strconv`，字符串转换包。
- `strings`，字符串操作包。
- `bytes`，字节切片操作包。
- `maps`，map操作包。
- `slices`，切片操作包。
- `io`，定义了一系列基础的IO接口，例如常用的`Closer`，`Writer`等，还有一些基础的IO操作函数。
- `io/fs`，定义了一系列文件系统的接口，算是对文件系统的抽象。
- `os`，操作系统进行交互的包，提供的是一系列具体的函数，可以简单理解为是上面两个的具体实现。
- `os/signal`，操作系统信号包，用于监听操作系统的信号，主要用于实现程序的优雅关闭。
- `os/exec`，操作系统命令包，用于执行操作系统命令。
- `os/user`，操作系统用户包，用于获取操作系统的用户信息。
- `bufio`，有缓冲IO读写包。
- `math`，数学操作包。
- `math/bits`，位运算操作包。
- `math/cmplx`，复数操作包，也提供了一些三角函数的工具函数。
- `math/rand`，伪随机包。
- `math/big`，大数高精度计算包。
- `container/heap`，最小堆的实现。
- `container/list`，双线链表的实现。
- `container/ring`，环的实现。
- `archive/zip`，zip压缩归档的包。
- `archive/tar`，tar归档文件的包。
- `compress`，压缩算法实现的包，比如gzip，flate。
- `database/sql`，与数据库进行交互的包。
- `encoding`，处理编码的包，其下有很多子包，比如`json`，`xml`之类的，还有`base64`这种。
- `crypto`，处理加密的包，其下有很多子包，比如sha1，rsa等。
- `go/ast`，go源代码抽象语法树映射的包。
- `go/parser` ，将go源代码解析成抽象语法树的包。
- `go/importer`，提供了对导入器的访问。
- `go/format`，go语言格式化包。
- `areana`，提供可手动分配和释放内存的功能，处于实验阶段。
- `runtime`，运行时操作包，go中有许多操作都是这个包实现的。

## fmt库
`fmt`库是用于处理格式化输入输出的包。

`Print`函数直接输入内容到终端，是最简单的输出方式：
```go
fmt.Print("hello")
```

`Prinft`函数用于输出格式化的字符串：
```go
fmt.Printf("%%v-pi :%v \n", pi)
```

`Println`会在输出的内容后加换行符：
```go
fmt.Println("hello")
```

`Fprint`系列的函数会将内容输出指定的流中，通常我们是输出到文件中：
```go
package main
 
import (
	"fmt"
	"os"
)
 
func main() {
	fmt.Println("开始向标准输出写入内容...")
	fileObj, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    // O_CREATE（不存在则创建）O_WRONLY（只写）O_APPEND（追加内容）
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	content := "chengzi子"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写入信息：%s\n", content)
	fmt.Fprintln(fileObj, "又一行信息。")
	fmt.Fprint(fileObj, "再一行信息。")
	fmt.Fprint(fileObj, "最后的信息。")
	fmt.Println("写入完成！")
}
```

`Sprint`系列的函数将传入的数据以字符串的形式输出：
```go
package main
 
import (
	"fmt"
)
 
func main() {
	s1 := fmt.Sprint("chengzi家")
	fmt.Println(s1)
	s2 := fmt.Sprintln("chengzi家")
	fmt.Println(s2)
	name := "chengzi家"
	age := 18
	s3 := fmt.Sprintf("name:%s,age:%d", name, age)
	fmt.Println(s3)
}
```

`Errorf`函数返回一个包装好的错误，错误的内容包含了传入的字符串：
```go
package main
 
import (
	"errors"
	"fmt"
)
 
func main() {
	original_err := errors.New("original error")
	new_err := fmt.Errorf("an error occurred: %w", original_err)
	fmt.Println(new_err)
}
```

`Scan`函数从标准输入中读取文本，读取由空白符分隔（换行符或空格）的值保存到传递给本函数的参数中：
```go
fmt.Scan(&name, &age, &married)
fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
```

`Scanf`函数可以定义输入参数的格式，只有按照指定格式输入才会识别，更加标准化：
```go
fmt.Scanf("1:%v 2:%v 3:%v", &name, &age, &married)
fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
```

`Scanln`函数不允许输入多个参数间回车，只能通过空格来分隔多个参数，最后回车确认输入：
```go
fmt.Scanln(&name, &age, &married)
fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
```

`Fscan`系列的函数从指定的输入流中读取数据：
```go
reader := strings.NewReader("chengzijia 18 false")
intt, _ := fmt.Fscan(reader, &name, &age, &married)
fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
fmt.Printf("intt:%v\n", intt)
```

`Sscan`系列的函数是从指定的字符串中读取数据：
```go
var name string
var alphabet_count int
n, err := fmt.Sscan("GFG 3", &name, &alphabet_count) // 根据空格取值
if err != nil {
	panic(err)
}
fmt.Printf("%d:%s, %d\n", n, name, alphabet_count)
```

## flag库
flag库包含了命令行参数解析的相关函数，该库使得命令行工具开发更为简单。

`Bool`函数表示定义一个布尔类型的命令行参数：
```go
/**
 * 参数一：表示命令行参数名称
 * 参数二：表示命令行参数的默认值
 * 参数三：表示命令行参数的帮助信息
 */
var isRegister = flag.Bool("isRegister",false,"是否已经注册")
```

`BoolVar`函数表示定义一个布尔类型的命令行参数，其值存到指定的变量中：
```go
var isRegister bool
/**
 * 参数一：表示存储值的指针
 * 参数二：表示命令行参数名称
 * 参数三：表示命令行参数的默认值
 * 参数四：表示命令行参数的帮助信息
 */
flag.Bool(&bool,"isRegister",false,"是否已经注册")
```

`String`函数表示定义一个字符类型的命令行参数：
```go
/**
 * 参数一：表示命令行参数名称
 * 参数二：表示命令行参数的默认值
 * 参数三：表示命令行参数的帮助信息
 */
var ps = flag.String("ps","-a","查看所有")
```

`StringVar`函数表示定义一个字符类型的命令行参数，其值存到指定的变量中：
```go
var ps string
/**
 * 参数一：表示存储值的指针
 * 参数二：表示命令行参数名称
 * 参数三：表示命令行参数的默认值
 * 参数四：表示命令行参数的帮助信息
 */
flag.StringVar(&ps,"ps","-a","查看所有")
```

`Int`函数表示定义一个整数类型的命令行参数：
```go
/**
 * 参数一：表示命令行参数名称
 * 参数二：表示命令行参数的默认值
 * 参数三：表示命令行参数的帮助信息
 */
var count = lag.Int("count",10,"执行次数")
```

`IntVar`函数表示定义一个整数类型的命令行参数，其值存到指定的变量中：
```go
var count int
/**
 * 参数一：表示命令行参数名称
 * 参数二：表示命令行参数的默认值
 * 参数三：表示命令行参数的帮助信息
 */
lag.IntVar(&count,"count",10,"执行次数")
```

`Int64`函数表示定义个64位整数类型的命令行参数：
```go
/**
 * 参数一：表示命令行参数名称
 * 参数二：表示命令行参数的默认值
 * 参数三：表示命令行参数的帮助信息
 */
var count = flag.Int64("count",10,"执行次数")
```

`Int64Var`函数表示定义个64位整数类型的命令行参数，其值存到指定的变量中：
```go
var count int64
/**
 * 参数一：表示存储值的指针
 * 参数二：表示命令行参数名称
 * 参数三：表示命令行参数的默认值
 * 参数四：表示命令行参数的帮助信息
 */
flag.Int64Var(&count,count",10,"执行次数")
```

`Float64`函数表示定义一个浮点数类型的命令行参数：
```go
/**
 * 参数一：表示命令行参数名称
 * 参数二：表示命令行参数的默认值
 * 参数三：表示命令行参数的帮助信息
 */
var score = flag.Float64("score",19.2,"分数")
```

`Float64Var`函数表示定义一个浮点数类型的命令行参数，其值存到指定的变量中：
```go
var score floa64
/**
 * 参数一：表示存储值的指针
 * 参数二：表示命令行参数名称
 * 参数三：表示命令行参数的默认值
 * 参数四：表示命令行参数的帮助信息
 */
flag.Float64Var(&score"score",19.2,"分数")
```

`Duration`函数表示定义一个时间格式的命令行参数：
```go
varr d = flag.Duration("d",0,"时间间隔")
```

`DurationVar`函数表示定义一个时间格式的命令行参数：
```go
var d time.Duration
/**
 * 参数一：表示存储值的指针
 * 参数二：表示命令行参数名称
 * 参数三：表示命令行参数的默认值
 * 参数四：表示命令行参数的帮助信息
 */
varr d = flag.DurationVar(&d,"d",0,"时间间隔")
```

`Parse`函数可以解析传入的命令行参数：
```go
stringflag string
flag.StringVar(&stringflag, "stringflag", "default", "string flag value")
flag.Parse()
fmt.Println("string flag:", stringflag)
```

`NFlag`可以获取到参数的个数：
```go

```

`Arg`函数可以获取到某个索引处的参数：
```go

```

`Args`函数在通过`Parse`函数解析完后如果还有未被解析的参数则返回这些参数的切片：
```go

```

`NArg`函数获取到`Parse`函数解析完后还有未被解析的参数数量：
```go

```

当选项不是从命令行传递的而是从配置表等地方获取，这个时候可以通过`FlagSet`结构的相关方法来解析这些选项：
```go

```

## log库
Log跨定义了Logger类型，该类型提供了一些格式化输出的方法。

`Print`函数用于输出日志信息：
```go
log.Print("这是一条普通日志")
```

`Println`函数用于输出日志信息并在末尾添加换行符：
```go
log.Println("这是一条带换行的日志")
```

`Printf`函数按照指定格式输出日志信息：
```go
log.Printf("这是一条格式化日志：变量值为 %d", 42)
```

`Fatal`类型函数输出日志信息后调用`os.Exit(1)`来终止程序：
```go
log.Fatal("这是一条普通日志")
log.Fatalln("这是一条带换行的日志")
log.Fatalf("这是一条格式化日志：变量值为 %d", 42)
```

`Panic`类型函数输出日志信息后引发`panic`：
```go
log.Panic("这是一条普通日志")
log.Panicln("这是一条带换行的日志")
log.Panicf("这是一条格式化日志：变量值为 %d", 42)
```

`Prefix`函数获取日志信息的前缀：
```go
var prefix = log.Prefix()
```

`SetPrefix`函数设置日志信息的前缀：
```go
log.SetPrefix("ERROR: ")
```

`Flags`函数查看日志格式的标志信息：
```go
var logFlags = log.Flags()
```

`SetFlags`函数设置日志格式：
```go
/**
 * 可选的标志：
 * Ldate：输出当地时区的日期，如2020/02/07；
 * Ltime：输出当地时区的时间，如11:45:45；
 * Lmicroseconds：输出的时间精确到微秒，设置了该选项就不用设置Ltime了。如11:45:45.123123；
 * Llongfile：输出长文件名+行号，含包名，如github.com/go-quiz/go-daily-lib/log/flag/main.go:50；
 * Lshortfile：输出短文件名+行号，不含包名，如main.go:50；
 * LUTC：如果设置了Ldate或Ltime，将输出 UTC 时间，而非当地时区。
 */
log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
```

`SetOutput`函数设置日志的输出位置：
```go
log.SetOutput(os.Stdout)
```

`通过New`函数可以创建一个新的logger对象：
```go
// 打开日志文件
file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
    log.Fatalf("无法打开日志文件: %v", err)
}
defer file.Close()

// 创建自定义 Logger
logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

// 记录日志
logger.Println("应用程序启动")
logger.Printf("当前状态：%s", "运行中")

// 错误日志示例
logger.SetPrefix("ERROR: ")
logger.Println("发生错误，无法连接数据库")
```

### log/slog
slog库允许以结构化的方式记录日志，结构化日志时将日志信息以键值对的心事的组织，方便后续日志分许和处理。

日志级别分为了`debug`、`info`、`warn`、`error`四个级别，分别提供了输出不同级别的日志：
```go
// 记录一条简单的信息日志
slog.Info("这是一条信息日志", "key", "value")
// 记录一条警告日志
slog.Warn("这是一条警告日志", "error", "some warning")
// 记录一条错误日志
slog.Error("这是一条错误日志", "err", "an error occurred")
```

默认情况下slog会记录所有级别的日志，我们可以通过设置`Handler`来控制日志的级别：
```go
// 创建一个新的文本处理程序，并设置日志级别为 Info
handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
// 设置全局的日志处理程序
slog.SetDefault(slog.New(handler))
 
// 记录 Debug 级别的日志，由于级别低于 Info，不会输出
slog.Debug("这是一条调试日志", "debugKey", "debugValue")
// 记录 Info 级别的日志，会输出
slog.Info("这是一条信息日志", "infoKey", "infoValue")
```

默认情况下slog使用文本格式输出日志，我们也可以通过创建文本处理程序来输出日志：
```go
// 创建一个文本处理程序
textHandler := slog.NewTextHandler(os.Stdout, nil)
// 创建一个使用文本处理程序的 Logger
logger := slog.New(textHandler)
 
// 使用自定义的 Logger 记录日志
logger.Info("文本格式日志", "type", "text")
```

除了文本格式以外slog还支持JSON格式的日志输出：
```go
// 创建一个 JSON 处理程序
jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
// 创建一个使用 JSON 处理程序的 Logger
logger := slog.New(jsonHandler)
 
// 使用自定义的 Logger 记录日志
logger.Info("JSON 格式日志", "type", "json")
```

我们也可以自定义日志处理程序来满足特定的需求：
```go
package main
 
import (
    "log/slog"
    "os"
    "strings"
)
 
// PrefixHandler 是一个自定义的日志处理程序，用于在每条日志前添加前缀
type PrefixHandler struct {
    handler slog.Handler
    prefix  string
}
 
// NewPrefixHandler 创建一个新的 PrefixHandler
func NewPrefixHandler(handler slog.Handler, prefix string) *PrefixHandler {
    return &PrefixHandler{
        handler: handler,
        prefix:  prefix,
    }
}
 
// Enabled 检查指定级别的日志是否启用
func (h *PrefixHandler) Enabled(_ []slog.Attr, level slog.Level) bool {
    return h.handler.Enabled(nil, level)
}
 
// Handle 处理日志记录
func (h *PrefixHandler) Handle(_ []slog.Attr, r slog.Record) error {
    r.Message = h.prefix + r.Message
    return h.handler.Handle(nil, r)
}
 
// WithAttrs 返回一个带有额外属性的新处理程序
func (h *PrefixHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
    return &PrefixHandler{
        handler: h.handler.WithAttrs(attrs),
        prefix:  h.prefix,
    }
}
 
// WithGroup 返回一个带有指定组的新处理程序
func (h *PrefixHandler) WithGroup(name string) slog.Handler {
    return &PrefixHandler{
        handler: h.handler.WithGroup(name),
        prefix:  h.prefix,
    }
}
 
func main() {
    // 创建一个文本处理程序
    textHandler := slog.NewTextHandler(os.Stdout, nil)
    // 创建一个带有前缀的自定义处理程序
    prefixHandler := NewPrefixHandler(textHandler, "[CUSTOM] ")
    // 创建一个使用自定义处理程序的 Logger
    logger := slog.New(prefixHandler)
 
    // 使用自定义的 Logger 记录日志
    logger.Info("自定义处理程序日志")
}
```

`HandlerOptions`结构用于配置自定义日志处理器的行为，可以实现日志的源位置跟踪、日志级别动态调整以及对日志属性的定制化处理：
```go
type HandlerOptions struct {
	// 是否记录源代码的位置
	AddSource bool
	// 日志的级别
	Level Leveler
	// 修改、删除或过滤属性，支持灵活的日志控制
	ReplaceAttr func(groups []string, a Attr) Attr
}
```

对于需要频繁输出日志的场景，通过`LogAttrs`函数和`Attr`类型来输出日志效率更高，因为它们减少了类型解析的过程：
```go
jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
/**
 * 参数一表示上下文类型
 * 参数二表示日志的级别
 * 参数三表示日志消息
 * 参数四表示其它键值对信息
 */
jsonLogger.logAttrs(context.Background(),slog.LevelInfo,"日志信息",slog.String("name","Jack"),slog.String("age",18))
```

如果所有的日志中都要包含相同的一个键值对可以通过`With`方法设置一个或多个统一的属性并返回一个新的`Logger`实例：
```go
jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger := jsonLogger.With("systemID","s1")

logger.logAttrs(context.Background(),slog.LevelInfo,"json-log",slog.String("name","Jack"))
logger.logAttrs(context.Background(),slog.LevelInfo,"json-log",slog.String("age",18))
```

### log/syslog
syslog库可以将不同程序的日志信息发送到统一的管理系统中，有助于日志数据的集中处理和监控。

```go
package main

import (
    "log/syslog"
    "log"
)

func main() {
    // LOG_INFO表示日志级别、LOG_LOCAL0表示日志消息的来源类别，MyApp表示日志消息的标识符
    logger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, "MyApp")
    if err != nil {
        log.Fatal(err)
    }

    logger.Info("这是通过 syslog 发送的信息")
}
```

## time库
time库提供了时间的显示和测量用的函数。

`Time`类表示时间，我们可以通过`now`函数来获取当前时间对象，然后根据该对象获取到年月日时分秒等信息：
```go
now := time.Now() //获取当前时间
fmt.Printf("current time:%v\n", now)

year := now.Year()     //年
month := now.Month()   //月
day := now.Day()       //日
hour := now.Hour()     //小时
minute := now.Minute() //分钟
second := now.Second() //秒
timestamp1 := now.Unix()     //时间戳
timestamp2 := now.UnixNano() //纳秒时间戳
timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式

// 将当前时间加上或减去一定的时间得到一个新的时间
time72hour := time.Now().Add( -72 * time.Hour )

/**
 * 将当前日期减去一定的天数得到一个新的日期
 * 参数一：表示年
 * 参数二：表示月
 * 参数三：表示日
 */
time1year := now.AddDate(-1,0, 0)

// 根据传入的值创建一个时间
start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2023, 7, 11, 12, 0, 0, 0, time.UTC)

/**
 * 计算两个时间相差的值
 */
sub := end.Sub(start)
fmt.Printf("sub = %v\n", sub)
// sub = 4596h0m0s

// 查看两个日期是否相等
fmt.Println(start.Equal(end))
// 查看当前实例时间是否在某个时间之前
fmt.Println(start.Before(end))
// 查看当前实例时间是否在某个时间之后
fmt.Println(start.Eed(end))

// 可以将一个字符串时间格式化
time.Now().Format("2006-01-02 15:04:05")

// 将时间字符串转为Time实例
timeParse,err := time.Parse(timerFormat,"2023-07-10 10:10:10")
if err == nil {
	fmt.Printf("time string:%s\n", timeParse)
}
// 	打印结果：time string:2023-07-10 10:10:10 +0000 UTC

timeParseLocation,err1 := time.ParseInLocation(timerFormat,"2023-07-10 10:10:10",time.Local)
if err1 == nil {
	fmt.Printf("time string:%s\n", timeParseLocation.String())
}
// 打印结果：time string:2023-07-10 10:10:10 +0800 CST
```

`Duration`是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位：
```go
// time包中定义的时间间隔类型的常量
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
```

`UnixNano`函数将事件转为时间戳：
```go
startTime = time.Now().UnixNano()
```

## path库
path实现了对斜杠分割的路径相关的操作函数。

`IsAbs`函数查看一个路径是否是绝对路径：
```go
var result = path.IsAbs("/home/test")
```

`Split`函数将路径从最后一个斜杠后面位置分割为两部分并返回，一部分表示文件路径，一部分表示文件名：
```go
var filePath,fileName = path.Split("/home/test/nginx.log")
```

`Join`函数将传入的所有字符串合并为一个新的路径：
```go
var result = path.Join("home","test","/nginx.log")
```

`Dir`函数返回路径除去最后一个路径元素的部分，表示获取文件所在的目录：
```go
var result = path.Dir("/home/test/nginx.log")
```

`Base`函数返回路径的最后一个元素，如果路径为空则返回"."，如果路径只有一个斜杠组成则会返回"/"：
```go
var result = path.Base("/home/test/")
```

`Ext`函数获取到文件的扩展名：
```go
var result = path.Ext(/home/test/nginx.log)
```

## path/filepath库
`path/filepath`库实现了兼容各种操作系统的文件路径的操作函数，path库中提供的函数该库中都有提供。

`IsAbs`函数查看一个路径是否是绝对路径：
```go
var result = filepath.IsAbs("/home/test")
```

`Abs`函数会返回指定路径的绝对路径，如果传入的参数不是一个绝对路径会在其基础上添加当前工作目录来形成一个绝对路径：
```go
filepath.Asb("go.log")
```

`Split`函数将路径从最后一个斜杠后面位置分割为两部分并返回，一部分表示文件路径，一部分表示文件名：
```go
var filePath,fileName = filepath.Split("/home/test/nginx.log")
```

`SplitList`函数

`Join`函数将传入的所有字符串合并为一个新的路径：
```go
var result = filepath.Join("home","test","/nginx.log")
```

`Dir`函数返回路径除去最后一个路径元素的部分，表示获取文件所在的目录：
```go
var result = filepath.Dir("/home/test/nginx.log")
```

`Base`函数返回路径的最后一个元素，如果路径为空则返回"."，如果路径只有一个斜杠组成则会返回"/"：
```go
var result = filepath.Base("/home/test/")
```

`Ext`函数获取到文件的扩展名：
```go
var result = filepath.Ext(/home/test/nginx.log)
```

`Rel`函数可以返回一个路径相对于基础路径的相对路径：
```go
// 这里返回nginx.log相对于/home/test的相对路径，如果后者无法表示为相对前者的路径将返回错误
var result,err = filepath.Rel("/home/test","nginx.log")
```

`FromSlash`函数将路径中的“/”替换为路径分割符，多个“/”会被替换为多个路径分隔符：
```go
var reuslt = filepath.FromSlash("/home/test/nginx.log")
```

`ToSlash`函数将路径职工的分隔符替换为“/”，多个路径分隔符会被替换为多个“/”：
```go
var reuslt = filepath.ToSlash("home\\test\\nginx\\log")
```

`VolumeName`函数用于获取卷名，通常这个函数用于获取Window路径中的卷名：
```go
var reuslt = filepath.VolumeName("C:\foo\test")
```

## io库
io库提供了用于进行数据的读取和写入的接口，这些接口是所有I/O操作的基本接口。

在io库中提供了三个常量，这些常量用于在文件读写时进行定位：
```go
const (
  SeekStart = 0   // 定位到文件头 
  SeekCurrent = 1 // 定位到当前读写的位置
  SeekEnd = 2     // 定位到文件尾
)
```

`Reader`接口是用于读取数据的基本接口，内部提供了一个`Read`方法：
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

`ReaderFrom`接口用于从指定的流中读取数据：
```go
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}
```

`ReadAt`接口从流中的指定位置读取数据到缓冲区中：
```go
type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}
```

`ByteReader`接口用于进行单字节读取的基本接口：
```go
type ByteReader interface {
	ReadByte() (byte, error)
}
```

`Writer`接口是用于写出数据的基本接口，其内部提供了一个`Write`方法：
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

`WriterTo`接口可以将数据写入到指定的流中：
```go
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}
```

`WriteAt`接口将缓冲区的数据写入到指定的位置：
```go
type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}
```

`ByteWrite`接口是用于单字节写出数据的基本接口：
```go
type ByteWriter interface {
	WriteByte(c byte) error
}
```

`WriteString`接口将一个字符串写出：
```go
type StringWriter interface {
	WriteString(s string) (n int, err error)
}
```

`Closer`接口包含一个`Close`方法，这个方法用于释放资源：
```go
type Closer interface {
    Close() error
}
```

`Seeker`接口允许我们在数据流中跳转到指定的位置：
```go
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```


将上面的一些基础接口进行组合组合为一些新接口：
```go
// 同时支持ReadByte()和UnreadByte()
type ByteScanner interface {
	ByteReader
	UnreadByte() error
}
 
// 同时支持ReadRune()和UnreadRune()
type RuneScanner interface {
	RuneReader
	UnreadRune() error
}
 
// 同时有Read和Close
type ReadCloser interface {
	Reader
	Closer
}
 
// 同时有Read、Seek
type ReadSeeker interface {
	Reader
	Seeker
}
 
// 同时有Read、Seek和Close 
type ReadSeekCloser interface {
	Reader
	Seeker
	Closer
}
 
// 同时有Read和Write
type ReadWriter interface {
	Reader
	Writer
}
 
// 同时有Read、Write和Seek
type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}
 
// 同时有Read、Write和Close
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
```

拷贝函数用于将数据从源位置拷贝到新的位置：
```go
func Copy(dst Writer, src Reader) (written int64, err error)
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

io库还提供了三个用于读取的函数可以直接使用：
```go
func ReadAll(r Reader) ([]byte, error)
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
func ReadFull(r Reader, buf []byte) (n int, err error)
```

`Pipe`函数对应于Linux中的`pipe`接口，创建一个输入流和输出流，写到输出流中的数据可以从输入流中读出：
```go
func Pipe() (*PipeReader, *PipeWriter)
```

`PipeReader`接口可以读取、关闭或者是关闭并返回一个错误：
```go
type PipeReader struct {
	// contains filtered or unexported fields
}
 
// 关闭后另一端再写会返回io.ErrClosePipe
func (r *PipeReader) Close() error
// 关闭后另一端再写会返回传入的err
func (r *PipeReader) CloseWithError(err error) error
// 读取时如果另一端关闭了会返回io.EOF
func (r *PipeReader) Read(data []byte) (n int, err error)
```

`PipeWriter`接口可以写、关闭或者在关闭的时候返回错误：
```go
type PipeWriter struct {
	// contains filtered or unexported fields
}
 
func (w *PipeWriter) Close() error
func (w *PipeWriter) CloseWithError(err error) error
func (w *PipeWriter) Write(data []byte) (n int, err error)
```

## io/fs库
io/fs包在Go中提供了一套抽象的文件系统接口，通过统一的方式来访问不同来源的文件系统。

`FS`接口提供了访问文件系统的基本方法：
```go
type FS interface {
    Open(name string) (File, error)
}
```

`File`和`DirEntry`些接口描述了文件系统中的文件和目录应有的行为：
```go
type File interface {
    Stat() (FileInfo, error)
    Read([]byte) (int, error)
    Close() error
}

type DirEntry interface {
    Name() string       // 文件或目录的名字
    IsDir() bool        // 判断是否为目录
    Type() FileMode     // 文件的模式和权限
    Info() (FileInfo, error) // 获取文件的详细信息
}
```

`fs.ReadFile`是一个方便的函数，用于快速读取文件的全部内容，它屏蔽了打开文件、读取文件以及关闭文件的复杂操作，使得文件读取变得简单明了：
```go
func ReadFile(fsys fs.FS, name string) ([]byte, error) {
    file, err := fsys.Open(name)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return io.ReadAll(file)
}
```

`fs.ReadDir`提供了一个高效的方式来列出目录中的所有项，这个函数返回一个`DirEntry`切片，每个`DirEntry`代表目录中的一个文件或子目录：
```go
func ReadDir(fsys fs.FS, name string) ([]fs.DirEntry, error) {
    dir, err := fsys.Open(name)
    if err != nil {
        return nil, err
    }
    defer dir.Close()
    return dir.ReadDir(-1)
}
```

`fs.WalkDir`是一个非常有用的函数，它允许开发者定义一个回调函数，该函数会被应用到每一个遍历到的文件和目录上：
```go
func WalkDir(root string, fn fs.WalkDirFunc) error {
    return fs.WalkDir(os.DirFS(root), ".", fn)
}
```

`fs.Glob`函数提供了模式匹配功能，可以使用通配符模式来查找匹配的文件：
```go
func Glob(fsys fs.FS, pattern string) ([]string, error) {
    // 实现具体的模式匹配逻辑
}

files, err := fs.Glob(os.DirFS("path/to/search"), "*.txt")
if err != nil {
    log.Fatal(err)
}
for _, file := range files {
    fmt.Println(file)
}
```

## bufio库
该库提供了数据的缓冲读写功能，其通过包装一个`io.Reader`和`io.Writer`对象创建一个带缓冲的读写器，使得读写操作更高效。

`Reader`结构体封装了`Reader`对象并提供了多种方法来对缓冲进行操作：
```go
type Reader struct {
	buf          []byte
	rd           io.Reader // 读取流
	r, w         int       // 缓冲区读、写的位置
	err          error
	lastByte     int // last byte read for UnreadByte; -1 means invalid
	lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
}

func (b *Reader) Buffered() int
func (b *Reader) Discard(n int) (discarded int, err error)
// 预览接下来的数据而不移动读取位置
func (b *Reader) Peek(n int) ([]byte, error)
// 读取数据到数组中，返回读取的字节数
func (b *Reader) Read(p []byte) (n int, err error)
// 读取单个字节
func (b *Reader) ReadByte() (byte, error)
// 读取数据直到遇到指定的分隔符
func (b *Reader) ReadBytes(delim byte) ([]byte, error)
// 读取一行
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
// 读取一个码点
func (b *Reader) ReadRune() (r rune, size int, err error)
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
// 读取数据直到遇到指定的分隔符，返回一个字符串
func (b *Reader) ReadString(delim byte) (string, error)
func (b *Reader) Reset(r io.Reader)
func (b *Reader) Size() int
func (b *Reader) UnreadByte() error
func (b *Reader) UnreadRune() error
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
```

`NewReader`函数可以创建一个新的`Reader`对象：
```go
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		println(line)
		if err != nil {
			log.Fatal(err)
		}
		// 打印行数据
	}
}

```

`Writer`结构体封装了`Writer`对象并提供了多种方法来对缓冲进行操作：
```go
type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}
func (b *Writer) Available() int
func (b *Writer) AvailableBuffer() []byte
func (b *Writer) Buffered() int
func (b *Writer) Flush() error
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
func (b *Writer) Reset(w io.Writer)
func (b *Writer) Size() int
// 将数据写入
func (b *Writer) Write(p []byte) (nn int, err error)
// 写入单个字节
func (b *Writer) WriteByte(c byte) error
// 写入单个码点
func (b *Writer) WriteRune(r rune) (size int, err error)
// 写入一个字符串
func (b *Writer) WriteString(s string) (int, error)
```

`NewWriter`函数可以创建一个新的`Writer`对象：
```go
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Create("example_output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("Hello, World!")
	if err != nil {
		log.Fatal(err)
	}

	writer.Flush()
}
```

`ReadWriter`结构体表示既可以读缓冲也可以写缓冲结构：
```go
type ReadWriter struct {
	*Reader
	*Writer
}

func NewReadWriter(r *Reader, w *Writer) *ReadWriter {
	return &ReadWriter{r, w}
}
```

`NewReadWriter`函数可以创建一个新的`ReadWriter`实例：
```go
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Create("example_output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewReadWriter(file)

	_, err = writer.WriteString("Hello, World!")
	if err != nil {
		log.Fatal(err)
	}

	writer.Flush()
}
```

## os库
os库中定义了一些常量表示文件的打开模式以及相对文件的位置：
```go
const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)

const (
    SEEK_SET int = 0 // 相对于文件起始位置seek
    SEEK_CUR int = 1 // 相对于文件当前位置seek
    SEEK_END int = 2 // 相对于文件结尾位置seek
)

// 下面的常量是文件的模式和权限的相关位
const (
    // 单字符是被String方法用于格式化的属性缩写。
    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
    ModeAppend                                     // a: 只能写入，且只能写入到末尾
    ModeExclusive                                  // l: 用于执行
    ModeTemporary                                  // T: 临时文件（非备份文件）
    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
    ModeDevice                                     // D: 设备
    ModeNamedPipe                                  // p: 命名管道（FIFO）
    ModeSocket                                     // S: Unix域socket
    ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
    ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
    ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
    ModeSticky                                     // t: 只有root/创建者能删除/移动文件
    // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
    ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
)
```

`Hostname`函数返回当前主机名：
```go
var hostName,err = os.Hostname()
```

`Getpagesize`函数获取系统内存页的大小，单位为字节（该值为系统的分页大小不一定和硬件分页大小相同）：
```go
var pageSize = os.Getpagesize()
```

`Environ`函数获取一个字符串切片，这些字符串表示环境变量：
```go
for _, value := range os.Environ() {
	fmt.Println(value)
}
```

`LookupEnv`函数用于获取指定名名称的环境变量的值，如果该变量不存在则返回空：
```go
goPath, bgExist := os.LookupEnv("GOPATH")
fmt.Println("GOPATH =", goPath, ", bgExist =", bgExist)
luaPath, blExist := os.LookupEnv("LUAPATH")
fmt.Println("LUAPATH =", luaPath, ", blExist =", blExist)
```

`Getenv`函数可以获取一个环境变量对应的值，如果该变量不存在则返回空字符串：
```go
runHome := os.Getenv("runHome")
```

`Setenv`函数可以设置环境变量的值：
```go
err := os.Setenv("runHome","/home/test/")
```

`Unsetenv`函数可以删除单个环境变量：
```go
err := os.Unsetenv("TESTENV")
if err != nil {
	os.Exit(1)
}
testValue := os.Getenv("TESTENV")
fmt.Println("TESTENV = ", testValue)
```

`Clearenv`函数清空当前环境变量：
```go
os.Clearenv()
```

`ExpandEnv`函数可以将字符串中通过“$”符号引用的环境变量填充到对应的位置：
```go
os.Setenv("GAME_NAME", "Chronicles of Darkness")
os.Setenv("GAME_DIR", "/usr/games")
fmt.Println(os.ExpandEnv("The $GAME_NAME game is placed in the ${GAME_DIR} path."))
``

`Exit`函数可以让当前程序以给定的状态码退出，通常0表示成功，非0表示出错：
```go
os.Exit(0)
```

`IsDir`函数查看

`IsRegular`

`Perm`

`String`



func Chdir(dir string) error
func Chmod(name string, mode FileMode) error
func Chown(name string, uid, gid int) error
func Chtimes(name string, atime time.Time, mtime time.Time) error
func CopyFS(dir string, fsys fs.FS) error
func DirFS(dir string) fs.FS
func Executable() (string, error)
func Expand(s string, mapping func(string) string) string
func Getegid() int
func Geteuid() int
func Getgid() int
func Getgroups() ([]int, error)
func Getpid() int
func Getppid() int
func Getuid() int
func Getwd() (dir string, err error)
func IsExist(err error) bool
func IsNotExist(err error) bool
func IsPathSeparator(c uint8) bool
func IsPermission(err error) bool
func IsTimeout(err error) bool
func Lchown(name string, uid, gid int) error
func Link(oldname, newname string) error
func Mkdir(name string, perm FileMode) error
func MkdirAll(path string, perm FileMode) error
func MkdirTemp(dir, pattern string) (string, error)
func NewSyscallError(syscall string, err error) error
func Pipe() (r *File, w *File, err error)
func ReadFile(name string) ([]byte, error)
func Readlink(name string) (string, error)
func Remove(name string) error
func RemoveAll(path string) error
func Rename(oldpath, newpath string) error
func SameFile(fi1, fi2 FileInfo) bool
func Symlink(oldname, newname string) error
func TempDir() string
func Truncate(name string, size int64) error
func UserCacheDir() (string, error)
func UserConfigDir() (string, error)
func UserHomeDir() (string, error)
func WriteFile(name string, data []byte, perm FileMode) error
type DirEntry
func ReadDir(name string) ([]DirEntry, error)
type File
func Create(name string) (*File, error)
func CreateTemp(dir, pattern string) (*File, error)
func NewFile(fd uintptr, name string) *File
func Open(name string) (*File, error)
func OpenFile(name string, flag int, perm FileMode) (*File, error)
func OpenInRoot(dir, name string) (*File, error)
func (f *File) Chdir() error
func (f *File) Chmod(mode FileMode) error
func (f *File) Chown(uid, gid int) error
func (f *File) Close() error
func (f *File) Fd() uintptr
func (f *File) Name() string
func (f *File) Read(b []byte) (n int, err error)
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
func (f *File) ReadDir(n int) ([]DirEntry, error)
func (f *File) ReadFrom(r io.Reader) (n int64, err error)
func (f *File) Readdir(n int) ([]FileInfo, error)
func (f *File) Readdirnames(n int) (names []string, err error)
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
func (f *File) SetDeadline(t time.Time) error
func (f *File) SetReadDeadline(t time.Time) error
func (f *File) SetWriteDeadline(t time.Time) error
func (f *File) Stat() (FileInfo, error)
func (f *File) Sync() error
func (f *File) SyscallConn() (syscall.RawConn, error)
func (f *File) Truncate(size int64) error
func (f *File) Write(b []byte) (n int, err error)
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
func (f *File) WriteString(s string) (n int, err error)
func (f *File) WriteTo(w io.Writer) (n int64, err error)
type FileInfo
func Lstat(name string) (FileInfo, error)
func Stat(name string) (FileInfo, error)
type FileMode
type LinkError
func (e *LinkError) Error() string
func (e *LinkError) Unwrap() error
type PathError
type ProcAttr
type Process
func FindProcess(pid int) (*Process, error)
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
func (p *Process) Kill() error
func (p *Process) Release() error
func (p *Process) Signal(sig Signal) error
func (p *Process) Wait() (*ProcessState, error)
type ProcessState
func (p *ProcessState) ExitCode() int
func (p *ProcessState) Exited() bool
func (p *ProcessState) Pid() int
func (p *ProcessState) String() string
func (p *ProcessState) Success() bool
func (p *ProcessState) Sys() any
func (p *ProcessState) SysUsage() any
func (p *ProcessState) SystemTime() time.Duration
func (p *ProcessState) UserTime() time.Duration
type Root
func OpenRoot(name string) (*Root, error)
func (r *Root) Close() error
func (r *Root) Create(name string) (*File, error)
func (r *Root) FS() fs.FS
func (r *Root) Lstat(name string) (FileInfo, error)
func (r *Root) Mkdir(name string, perm FileMode) error
func (r *Root) Name() string
func (r *Root) Open(name string) (*File, error)
func (r *Root) OpenFile(name string, flag int, perm FileMode) (*File, error)
func (r *Root) OpenRoot(name string) (*Root, error)
func (r *Root) Remove(name string) error
func (r *Root) Stat(name string) (FileInfo, error)
type Signal
type SyscallError
func (e *SyscallError) Error() string
func (e *SyscallError) Timeout() bool
func (e *SyscallError) Unwrap() error

## os/signal库
func Ignore(sig ...os.Signal)
func Ignored(sig os.Signal) bool
func Notify(c chan<- os.Signal, sig ...os.Signal)
func NotifyContext(parent context.Context, signals ...os.Signal) (ctx context.Context, stop context.CancelFunc)
func Reset(sig ...os.Signal)
func Stop(c chan<- os.Signal)

## os/exec库
func LookPath(file string) (string, error)
type Cmd
func Command(name string, arg ...string) *Cmd
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd
func (c *Cmd) CombinedOutput() ([]byte, error)
func (c *Cmd) Environ() []string
func (c *Cmd) Output() ([]byte, error)
func (c *Cmd) Run() error
func (c *Cmd) Start() error
func (c *Cmd) StderrPipe() (io.ReadCloser, error)
func (c *Cmd) StdinPipe() (io.WriteCloser, error)
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
func (c *Cmd) String() string
func (c *Cmd) Wait() error
type Error
func (e *Error) Error() string
func (e *Error) Unwrap() error
type ExitError
func (e *ExitError) Error() string

## embed库
embed库的核心功能就是将指定的文件或目录嵌入到Go程序的可执行文件中，在编译的时候embed会将标记为嵌入的文件或目录的内容作为字节数组直接嵌入到编译后的二进制文件中，在程序运行时可以直接从二进制文件中提取这些资源。

可以直接使用`//go:embed`指令来指示编译器嵌入特定的文件或目录：
```go
// 这里将image.png文件嵌入到程序中，可以通过myImage变量来访问
//go:embed image.png
var myImage embed.FS
```

一旦资源被嵌入就可以通过embed包提供的接口来访问它们：
```go
data, err := myImage.ReadFile("image.png")
if err != nil {
    // 处理错误
}
// 使用 data 变量


// 嵌套templates整个目录
//go:embed templates/*
var templatesFS embed.FS

// 通过readDir方法来遍历目录中的文件
entries, err := templatesFS.ReadDir("templates")
if err != nil {
    // 错误处理
}
for _, entry := range entries {
    fileData, err := templatesFS.ReadFile("templates/" + entry.Name())
    if err != nil {
        // 错误处理
    }
    // 使用 fileData
}

```

**需要注意的是embed只能嵌入项目目录下的文件或目录，并且embed指令必须紧跟在变量声明之前，不能有空行或其它注释**。

## os/user库
func LookPath(file string) (string, error)
type Cmd
func Command(name string, arg ...string) *Cmd
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd
func (c *Cmd) CombinedOutput() ([]byte, error)
func (c *Cmd) Environ() []string
func (c *Cmd) Output() ([]byte, error)
func (c *Cmd) Run() error
func (c *Cmd) Start() error
func (c *Cmd) StderrPipe() (io.ReadCloser, error)
func (c *Cmd) StdinPipe() (io.WriteCloser, error)
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
func (c *Cmd) String() string
func (c *Cmd) Wait() error
type Error
func (e *Error) Error() string
func (e *Error) Unwrap() error
type ExitError
func (e *ExitError) Error() string

## strings库
`Contains`函数查看一个字符串是否在指定的字符串中出现，如果出现了返回True，否则返回False：
```go
strings.Contains("Hello world","wo")
```

`ContainsAny`函数查看一个字符串中任意一个Unicode码点是否在指定的字符串中出现，如果出现了返回True，否则返回False：
```go
strings.ContainsAny("Hello world","wo")
```

`ContainsRune`函数查看一个码点是否在指定的字符串中出现，如果出现了返回True，否则返回False：
```go
strings.ContainsRune("Hello world","w")
```

`Index`函数查看字符串在指定字符串中第一次出现的位置，如果不存在则返回-1：
```go

```

`IndexAny`函数查看字符串的任意一个Unicode码点在指定字符串中第一次出现的位置，如果不存在则返回-1：
```go

```

`IndexFunc`
```go

```

`IndexByte`函数查看单字节字符在指定字符串中第一次出现的位置，如果不存在则返回-1：
```go

```

`IndexRune`函数查看一个码点在指定的字符串中第一次出现的位置，如果不存在则返回-1：
```go

```

`LastIndex`函数查看字符串在指定字符串中最后一次出现的位置，如果不存在则返回-1：
```go

```

`LastIndexByte`函数查看单字节字符在指定字符串中最后一次出现的位置，如果不存在则返回-1：
```go

```

`LastIndexAny`函数查看字符串的任意一个Unicode码点在指定字符串中最后一次出现的位置，如果不存在则返回-1：
```go

```

`LastIndexFunc`
```go

```

`Count`函数查看字符串在指定字符串中出现的次数：
```go

```

`HasPrefix`函数查看字符串是否以指定的前缀开头：
```go

```

`HasSuffix`函数查看字符串是否以指定的后缀开头：
```go

```

`Compare`函数用于比较两个字符串大小，如果字符串相等则返回0，如果前者小于后者则返回-1，否则返回1：
```go
a := "gopher"
b := "hello world"
fmt.Println(strings.Compare(a, b))
fmt.Println(strings.Compare(a, a))
fmt.Println(strings.Compare(b, a))
```

`EqualFold`函数忽略两个字符串的大小写，比较两个字符串是否相同：
```go
fmt.Println(strings.EqualFold("GO", "go"))
fmt.Println(strings.EqualFold("壹", "一"))
```

`Title`
`ToLower`
`ToLowerSpecial`
`ToTitle`
`ToTitleSpecial`
`ToUpper`
`ToUpperSpecial`


`Replace`
`ReplaceAll`

`Trim`
`TrimFunc`
`TrimLeft`
`TrimLeftFunc`
`TrimPrefix`
`TrimRight`
`TrimRightFunc`
`TrimSpace`
`TrimSuffix`

`Fields`
`FieldsFunc`
`FieldsFuncSeq`
`FieldsSeq`

`Split`
`SplitAfter`
`SplitAfterN`
`SplitAfterSeq`
`SplitN`
`SplitSeq`

`Join`

`Repeat`


func Clone(s string) string
func ContainsFunc(s string, f func(rune) bool) bool
func Cut(s, sep string) (before, after string, found bool)
func CutPrefix(s, prefix string) (after string, found bool)
func CutSuffix(s, suffix string) (before string, found bool)
func Lines(s string) iter.Seq[string]
func Map(mapping func(rune) rune, s string) string
func ToValidUTF8(s, replacement string) string
type Builder
func (b *Builder) Cap() int
func (b *Builder) Grow(n int)
func (b *Builder) Len() int
func (b *Builder) Reset()
func (b *Builder) String() string
func (b *Builder) Write(p []byte) (int, error)
func (b *Builder) WriteByte(c byte) error
func (b *Builder) WriteRune(r rune) (int, error)
func (b *Builder) WriteString(s string) (int, error)
type Reader
func NewReader(s string) *Reader
func (r *Reader) Len() int
func (r *Reader) Read(b []byte) (n int, err error)
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
func (r *Reader) ReadByte() (byte, error)
func (r *Reader) ReadRune() (ch rune, size int, err error)
func (r *Reader) Reset(s string)
func (r *Reader) Seek(offset int64, whence int) (int64, error)
func (r *Reader) Size() int64
func (r *Reader) UnreadByte() error
func (r *Reader) UnreadRune() error
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
type Replacer
func NewReplacer(oldnew ...string) *Replacer
func (r *Replacer) Replace(s string) string
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)

## html库
func EscapeString(s string) string
func UnescapeString(s string) string

## reflect库
Constants
func Copy(dst, src Value) int
func DeepEqual(x, y any) bool
func Swapper(slice any) func(i, j int)
type ChanDir
func (d ChanDir) String() string
type Kind
func (k Kind) String() string
type MapIter
func (iter *MapIter) Key() Value
func (iter *MapIter) Next() bool
func (iter *MapIter) Reset(v Value)
func (iter *MapIter) Value() Value
type Method
func (m Method) IsExported() bool
type SelectCase
type SelectDir
type SliceHeaderdeprecated
type StringHeaderdeprecated
type StructField
func VisibleFields(t Type) []StructField
func (f StructField) IsExported() bool
type StructTag
func (tag StructTag) Get(key string) string
func (tag StructTag) Lookup(key string) (value string, ok bool)
type Type
func ArrayOf(length int, elem Type) Type
func ChanOf(dir ChanDir, t Type) Type
func FuncOf(in, out []Type, variadic bool) Type
func MapOf(key, elem Type) Type
func PointerTo(t Type) Type
func PtrTo(t Type) Typedeprecated
func SliceOf(t Type) Type
func StructOf(fields []StructField) Type
func TypeFor[T any]() Type
func TypeOf(i any) Type
type Value
func Append(s Value, x ...Value) Value
func AppendSlice(s, t Value) Value
func Indirect(v Value) Value
func MakeChan(typ Type, buffer int) Value
func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
func MakeMap(typ Type) Value
func MakeMapWithSize(typ Type, n int) Value
func MakeSlice(typ Type, len, cap int) Value
func New(typ Type) Value
func NewAt(typ Type, p unsafe.Pointer) Value
func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
func SliceAt(typ Type, p unsafe.Pointer, n int) Value
func ValueOf(i any) Value
func Zero(typ Type) Value
func (v Value) Addr() Value
func (v Value) Bool() bool
func (v Value) Bytes() []byte
func (v Value) Call(in []Value) []Value
func (v Value) CallSlice(in []Value) []Value
func (v Value) CanAddr() bool
func (v Value) CanComplex() bool
func (v Value) CanConvert(t Type) bool
func (v Value) CanFloat() bool
func (v Value) CanInt() bool
func (v Value) CanInterface() bool
func (v Value) CanSet() bool
func (v Value) CanUint() bool
func (v Value) Cap() int
func (v Value) Clear()
func (v Value) Close()
func (v Value) Comparable() bool
func (v Value) Complex() complex128
func (v Value) Convert(t Type) Value
func (v Value) Elem() Value
func (v Value) Equal(u Value) bool
func (v Value) Field(i int) Value
func (v Value) FieldByIndex(index []int) Value
func (v Value) FieldByIndexErr(index []int) (Value, error)
func (v Value) FieldByName(name string) Value
func (v Value) FieldByNameFunc(match func(string) bool) Value
func (v Value) Float() float64
func (v Value) Grow(n int)
func (v Value) Index(i int) Value
func (v Value) Int() int64
func (v Value) Interface() (i any)
func (v Value) InterfaceData() [2]uintptrdeprecated
func (v Value) IsNil() bool
func (v Value) IsValid() bool
func (v Value) IsZero() bool
func (v Value) Kind() Kind
func (v Value) Len() int
func (v Value) MapIndex(key Value) Value
func (v Value) MapKeys() []Value
func (v Value) MapRange() *MapIter
func (v Value) Method(i int) Value
func (v Value) MethodByName(name string) Value
func (v Value) NumField() int
func (v Value) NumMethod() int
func (v Value) OverflowComplex(x complex128) bool
func (v Value) OverflowFloat(x float64) bool
func (v Value) OverflowInt(x int64) bool
func (v Value) OverflowUint(x uint64) bool
func (v Value) Pointer() uintptr
func (v Value) Recv() (x Value, ok bool)
func (v Value) Send(x Value)
func (v Value) Seq() iter.Seq[Value]
func (v Value) Seq2() iter.Seq2[Value, Value]
func (v Value) Set(x Value)
func (v Value) SetBool(x bool)
func (v Value) SetBytes(x []byte)
func (v Value) SetCap(n int)
func (v Value) SetComplex(x complex128)
func (v Value) SetFloat(x float64)
func (v Value) SetInt(x int64)
func (v Value) SetIterKey(iter *MapIter)
func (v Value) SetIterValue(iter *MapIter)
func (v Value) SetLen(n int)
func (v Value) SetMapIndex(key, elem Value)
func (v Value) SetPointer(x unsafe.Pointer)
func (v Value) SetString(x string)
func (v Value) SetUint(x uint64)
func (v Value) SetZero()
func (v Value) Slice(i, j int) Value
func (v Value) Slice3(i, j, k int) Value
func (v Value) String() string
func (v Value) TryRecv() (x Value, ok bool)
func (v Value) TrySend(x Value) bool
func (v Value) Type() Type
func (v Value) Uint() uint64
func (v Value) UnsafeAddr() uintptr
func (v Value) UnsafePointer() unsafe.Pointer
type ValueError
func (e *ValueError) Error() string
Bugs



## sort库
func Find(n int, cmp func(int) int) (i int, found bool)
func Float64s(x []float64)
func Float64sAreSorted(x []float64) bool
func Ints(x []int)
func IntsAreSorted(x []int) bool
func IsSorted(data Interface) bool
func Search(n int, f func(int) bool) int
func SearchFloat64s(a []float64, x float64) int
func SearchInts(a []int, x int) int
func SearchStrings(a []string, x string) int
func Slice(x any, less func(i, j int) bool)
func SliceIsSorted(x any, less func(i, j int) bool) bool
func SliceStable(x any, less func(i, j int) bool)
func Sort(data Interface)
func Stable(data Interface)
func Strings(x []string)
func StringsAreSorted(x []string) bool
type Float64Slice
func (x Float64Slice) Len() int
func (x Float64Slice) Less(i, j int) bool
func (p Float64Slice) Search(x float64) int
func (x Float64Slice) Sort()
func (x Float64Slice) Swap(i, j int)
type IntSlice
func (x IntSlice) Len() int
func (x IntSlice) Less(i, j int) bool
func (p IntSlice) Search(x int) int
func (x IntSlice) Sort()
func (x IntSlice) Swap(i, j int)
type Interface
func Reverse(data Interface) Interface
type StringSlice
func (x StringSlice) Len() int
func (x StringSlice) Less(i, j int) bool
func (p StringSlice) Search(x string) int
func (x StringSlice) Sort()
func (x StringSlice) Swap(i, j int)

## regex库

## image库
func RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), ...)
type Alpha
func NewAlpha(r Rectangle) *Alpha
func (p *Alpha) AlphaAt(x, y int) color.Alpha
func (p *Alpha) At(x, y int) color.Color
func (p *Alpha) Bounds() Rectangle
func (p *Alpha) ColorModel() color.Model
func (p *Alpha) Opaque() bool
func (p *Alpha) PixOffset(x, y int) int
func (p *Alpha) RGBA64At(x, y int) color.RGBA64
func (p *Alpha) Set(x, y int, c color.Color)
func (p *Alpha) SetAlpha(x, y int, c color.Alpha)
func (p *Alpha) SetRGBA64(x, y int, c color.RGBA64)
func (p *Alpha) SubImage(r Rectangle) Image
type Alpha16
func NewAlpha16(r Rectangle) *Alpha16
func (p *Alpha16) Alpha16At(x, y int) color.Alpha16
func (p *Alpha16) At(x, y int) color.Color
func (p *Alpha16) Bounds() Rectangle
func (p *Alpha16) ColorModel() color.Model
func (p *Alpha16) Opaque() bool
func (p *Alpha16) PixOffset(x, y int) int
func (p *Alpha16) RGBA64At(x, y int) color.RGBA64
func (p *Alpha16) Set(x, y int, c color.Color)
func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16)
func (p *Alpha16) SetRGBA64(x, y int, c color.RGBA64)
func (p *Alpha16) SubImage(r Rectangle) Image
type CMYK
func NewCMYK(r Rectangle) *CMYK
func (p *CMYK) At(x, y int) color.Color
func (p *CMYK) Bounds() Rectangle
func (p *CMYK) CMYKAt(x, y int) color.CMYK
func (p *CMYK) ColorModel() color.Model
func (p *CMYK) Opaque() bool
func (p *CMYK) PixOffset(x, y int) int
func (p *CMYK) RGBA64At(x, y int) color.RGBA64
func (p *CMYK) Set(x, y int, c color.Color)
func (p *CMYK) SetCMYK(x, y int, c color.CMYK)
func (p *CMYK) SetRGBA64(x, y int, c color.RGBA64)
func (p *CMYK) SubImage(r Rectangle) Image
type Config
func DecodeConfig(r io.Reader) (Config, string, error)
type Gray
func NewGray(r Rectangle) *Gray
func (p *Gray) At(x, y int) color.Color
func (p *Gray) Bounds() Rectangle
func (p *Gray) ColorModel() color.Model
func (p *Gray) GrayAt(x, y int) color.Gray
func (p *Gray) Opaque() bool
func (p *Gray) PixOffset(x, y int) int
func (p *Gray) RGBA64At(x, y int) color.RGBA64
func (p *Gray) Set(x, y int, c color.Color)
func (p *Gray) SetGray(x, y int, c color.Gray)
func (p *Gray) SetRGBA64(x, y int, c color.RGBA64)
func (p *Gray) SubImage(r Rectangle) Image
type Gray16
func NewGray16(r Rectangle) *Gray16
func (p *Gray16) At(x, y int) color.Color
func (p *Gray16) Bounds() Rectangle
func (p *Gray16) ColorModel() color.Model
func (p *Gray16) Gray16At(x, y int) color.Gray16
func (p *Gray16) Opaque() bool
func (p *Gray16) PixOffset(x, y int) int
func (p *Gray16) RGBA64At(x, y int) color.RGBA64
func (p *Gray16) Set(x, y int, c color.Color)
func (p *Gray16) SetGray16(x, y int, c color.Gray16)
func (p *Gray16) SetRGBA64(x, y int, c color.RGBA64)
func (p *Gray16) SubImage(r Rectangle) Image
type Image
func Decode(r io.Reader) (Image, string, error)
type NRGBA
func NewNRGBA(r Rectangle) *NRGBA
func (p *NRGBA) At(x, y int) color.Color
func (p *NRGBA) Bounds() Rectangle
func (p *NRGBA) ColorModel() color.Model
func (p *NRGBA) NRGBAAt(x, y int) color.NRGBA
func (p *NRGBA) Opaque() bool
func (p *NRGBA) PixOffset(x, y int) int
func (p *NRGBA) RGBA64At(x, y int) color.RGBA64
func (p *NRGBA) Set(x, y int, c color.Color)
func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA)
func (p *NRGBA) SetRGBA64(x, y int, c color.RGBA64)
func (p *NRGBA) SubImage(r Rectangle) Image
type NRGBA64
func NewNRGBA64(r Rectangle) *NRGBA64
func (p *NRGBA64) At(x, y int) color.Color
func (p *NRGBA64) Bounds() Rectangle
func (p *NRGBA64) ColorModel() color.Model
func (p *NRGBA64) NRGBA64At(x, y int) color.NRGBA64
func (p *NRGBA64) Opaque() bool
func (p *NRGBA64) PixOffset(x, y int) int
func (p *NRGBA64) RGBA64At(x, y int) color.RGBA64
func (p *NRGBA64) Set(x, y int, c color.Color)
func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64)
func (p *NRGBA64) SetRGBA64(x, y int, c color.RGBA64)
func (p *NRGBA64) SubImage(r Rectangle) Image
type NYCbCrA
func NewNYCbCrA(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *NYCbCrA
func (p *NYCbCrA) AOffset(x, y int) int
func (p *NYCbCrA) At(x, y int) color.Color
func (p *NYCbCrA) ColorModel() color.Model
func (p *NYCbCrA) NYCbCrAAt(x, y int) color.NYCbCrA
func (p *NYCbCrA) Opaque() bool
func (p *NYCbCrA) RGBA64At(x, y int) color.RGBA64
func (p *NYCbCrA) SubImage(r Rectangle) Image
type Paletted
func NewPaletted(r Rectangle, p color.Palette) *Paletted
func (p *Paletted) At(x, y int) color.Color
func (p *Paletted) Bounds() Rectangle
func (p *Paletted) ColorIndexAt(x, y int) uint8
func (p *Paletted) ColorModel() color.Model
func (p *Paletted) Opaque() bool
func (p *Paletted) PixOffset(x, y int) int
func (p *Paletted) RGBA64At(x, y int) color.RGBA64
func (p *Paletted) Set(x, y int, c color.Color)
func (p *Paletted) SetColorIndex(x, y int, index uint8)
func (p *Paletted) SetRGBA64(x, y int, c color.RGBA64)
func (p *Paletted) SubImage(r Rectangle) Image
type PalettedImage
type Point
func Pt(X, Y int) Point
func (p Point) Add(q Point) Point
func (p Point) Div(k int) Point
func (p Point) Eq(q Point) bool
func (p Point) In(r Rectangle) bool
func (p Point) Mod(r Rectangle) Point
func (p Point) Mul(k int) Point
func (p Point) String() string
func (p Point) Sub(q Point) Point
type RGBA
func NewRGBA(r Rectangle) *RGBA
func (p *RGBA) At(x, y int) color.Color
func (p *RGBA) Bounds() Rectangle
func (p *RGBA) ColorModel() color.Model
func (p *RGBA) Opaque() bool
func (p *RGBA) PixOffset(x, y int) int
func (p *RGBA) RGBA64At(x, y int) color.RGBA64
func (p *RGBA) RGBAAt(x, y int) color.RGBA
func (p *RGBA) Set(x, y int, c color.Color)
func (p *RGBA) SetRGBA(x, y int, c color.RGBA)
func (p *RGBA) SetRGBA64(x, y int, c color.RGBA64)
func (p *RGBA) SubImage(r Rectangle) Image
type RGBA64
func NewRGBA64(r Rectangle) *RGBA64
func (p *RGBA64) At(x, y int) color.Color
func (p *RGBA64) Bounds() Rectangle
func (p *RGBA64) ColorModel() color.Model
func (p *RGBA64) Opaque() bool
func (p *RGBA64) PixOffset(x, y int) int
func (p *RGBA64) RGBA64At(x, y int) color.RGBA64
func (p *RGBA64) Set(x, y int, c color.Color)
func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64)
func (p *RGBA64) SubImage(r Rectangle) Image
type RGBA64Image
type Rectangle
func Rect(x0, y0, x1, y1 int) Rectangle
func (r Rectangle) Add(p Point) Rectangle
func (r Rectangle) At(x, y int) color.Color
func (r Rectangle) Bounds() Rectangle
func (r Rectangle) Canon() Rectangle
func (r Rectangle) ColorModel() color.Model
func (r Rectangle) Dx() int
func (r Rectangle) Dy() int
func (r Rectangle) Empty() bool
func (r Rectangle) Eq(s Rectangle) bool
func (r Rectangle) In(s Rectangle) bool
func (r Rectangle) Inset(n int) Rectangle
func (r Rectangle) Intersect(s Rectangle) Rectangle
func (r Rectangle) Overlaps(s Rectangle) bool
func (r Rectangle) RGBA64At(x, y int) color.RGBA64
func (r Rectangle) Size() Point
func (r Rectangle) String() string
func (r Rectangle) Sub(p Point) Rectangle
func (r Rectangle) Union(s Rectangle) Rectangle
type Uniform
func NewUniform(c color.Color) *Uniform
func (c *Uniform) At(x, y int) color.Color
func (c *Uniform) Bounds() Rectangle
func (c *Uniform) ColorModel() color.Model
func (c *Uniform) Convert(color.Color) color.Color
func (c *Uniform) Opaque() bool
func (c *Uniform) RGBA() (r, g, b, a uint32)
func (c *Uniform) RGBA64At(x, y int) color.RGBA64
type YCbCr
func NewYCbCr(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *YCbCr
func (p *YCbCr) At(x, y int) color.Color
func (p *YCbCr) Bounds() Rectangle
func (p *YCbCr) COffset(x, y int) int
func (p *YCbCr) ColorModel() color.Model
func (p *YCbCr) Opaque() bool
func (p *YCbCr) RGBA64At(x, y int) color.RGBA64
func (p *YCbCr) SubImage(r Rectangle) Image
func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr
func (p *YCbCr) YOffset(x, y int) int
type YCbCrSubsampleRatio
func (s YCbCrSubsampleRatio) String() string


## unicode库
该库包含了基本的字符判断函数，其中utf8包主要负责`rune`和`byte`之间的转换，utf16包主要负责`rune`和`uint16`数组之间的转换（`rune`表示一个unicode码点）。

`RangeTable`类型通过列出一组Unicode码点的范围来定义它，通过该结构将所有unicode涉及到的码点进行分类来表示不同类别的字符集合：
```go
type Range16 struct {
	Lo     uint16
	Hi     uint16
	Stride uint16
}

type Range32 struct {
	Lo     uint32
	Hi     uint32
	Stride uint32
}
type RangeTable struct {
	R16         []Range16
	R32         []Range32
	LatinOffset int // number of entries in R16 with Hi <= MaxLatin1
}
```

`IsControl`函数用于查看字符是否是一个控制字符：
```go
res := unicode.IsControl('\r')
```

`IsDigit`函数用于查看字符是否是一个阿拉伯数字：
```go
res := unicode.IsDigit('9')
```

`IsGraphic`函数用于查看字符是否是圆形字符：
```go
res := unicode.IsGraphic('①')
```

`IsLetter`函数用于查看是否是字母：
```go
res := unicode.IsLetter('M')
```

`IsLower`函数用于查看是否是小写字符：
```go
res := unicode.IsLower('a')
```

`IsMark`函数用于查看是否是符号字符：
```go
res := unicode.IsMark(',')
```
https://blog.csdn.net/a772304419/article/details/134355716
``
func In(r rune, ranges ...*RangeTable) bool
func Is(rangeTab *RangeTable, r rune) bool
func IsNumber(r rune) bool
func IsOneOf(ranges []*RangeTable, r rune) bool
func IsPrint(r rune) bool
func IsPunct(r rune) bool
func IsSpace(r rune) bool
func IsSymbol(r rune) bool
func IsTitle(r rune) bool
func IsUpper(r rune) bool
func SimpleFold(r rune) rune
func To(_case int, r rune) rune
func ToLower(r rune) rune
func ToTitle(r rune) rune
func ToUpper(r rune) rune
type CaseRange
type Range16
type Range32
type RangeTable
type SpecialCase
func (special SpecialCase) ToLower(r rune) rune
func (special SpecialCase) ToTitle(r rune) rune
func (special SpecialCase) ToUpper(r rune) rune
Bugs

## unsafe库
unsafe库提供了依稀诶不安全的编程操作，例如直接操作指针、修改内存等，这些操作可能会引起内存错误和安全漏洞。

`Pointer`是一个通用的指针类型，可以指向任意类型的变量
```go
type ArbitraryType int

type IntegerType int

type Pointer *ArbitraryType
```


func Alignof(x ArbitraryType) uintptr
func Offsetof(x ArbitraryType) uintptr
func Sizeof(x ArbitraryType) uintptr
func String(ptr *byte, len IntegerType) string
func StringData(str string) *byte
type ArbitraryType
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
func SliceData(slice []ArbitraryType) *ArbitraryType
type IntegerType
type Pointer
func Add(ptr Pointer, len IntegerType) Pointer

## syscall库
func Access(path string, mode uint32) (err error)
func Acct(path string) (err error)
func Adjtimex(buf *Timex) (state int, err error)
func AttachLsf(fd int, i []SockFilter) errordeprecated
func Bind(fd int, sa Sockaddr) (err error)
func BindToDevice(fd int, device string) (err error)
func BytePtrFromString(s string) (*byte, error)
func ByteSliceFromString(s string) ([]byte, error)
func Chdir(path string) (err error)
func Chmod(path string, mode uint32) (err error)
func Chown(path string, uid int, gid int) (err error)
func Chroot(path string) (err error)
func Clearenv()
func Close(fd int) (err error)
func CloseOnExec(fd int)
func CmsgLen(datalen int) int
func CmsgSpace(datalen int) int
func Connect(fd int, sa Sockaddr) (err error)
func Creat(path string, mode uint32) (fd int, err error)
func DetachLsf(fd int) errordeprecated
func Dup(oldfd int) (fd int, err error)
func Dup2(oldfd int, newfd int) (err error)
func Dup3(oldfd int, newfd int, flags int) (err error)
func Environ() []string
func EpollCreate(size int) (fd int, err error)
func EpollCreate1(flag int) (fd int, err error)
func EpollCtl(epfd int, op int, fd int, event *EpollEvent) (err error)
func EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
func Exec(argv0 string, argv []string, envv []string) (err error)
func Exit(code int)
func Faccessat(dirfd int, path string, mode uint32, flags int) (err error)
func Fallocate(fd int, mode uint32, off int64, len int64) (err error)
func Fchdir(fd int) (err error)
func Fchmod(fd int, mode uint32) (err error)
func Fchmodat(dirfd int, path string, mode uint32, flags int) error
func Fchown(fd int, uid int, gid int) (err error)
func Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error)
func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) error
func Fdatasync(fd int) (err error)
func Flock(fd int, how int) (err error)
func ForkExec(argv0 string, argv []string, attr *ProcAttr) (pid int, err error)
func Fstat(fd int, stat *Stat_t) (err error)
func Fstatfs(fd int, buf *Statfs_t) (err error)
func Fsync(fd int) (err error)
func Ftruncate(fd int, length int64) (err error)
func Futimes(fd int, tv []Timeval) (err error)
func Futimesat(dirfd int, path string, tv []Timeval) (err error)
func Getcwd(buf []byte) (n int, err error)
func Getdents(fd int, buf []byte) (n int, err error)
func Getegid() (egid int)
func Getenv(key string) (value string, found bool)
func Geteuid() (euid int)
func Getgid() (gid int)
func Getgroups() (gids []int, err error)
func Getpagesize() int
func Getpgid(pid int) (pgid int, err error)
func Getpgrp() (pid int)
func Getpid() (pid int)
func Getppid() (ppid int)
func Getpriority(which int, who int) (prio int, err error)
func Getrlimit(resource int, rlim *Rlimit) (err error)
func Getrusage(who int, rusage *Rusage) (err error)
func GetsockoptInet4Addr(fd, level, opt int) (value [4]byte, err error)
func GetsockoptInt(fd, level, opt int) (value int, err error)
func Gettid() (tid int)
func Gettimeofday(tv *Timeval) (err error)
func Getuid() (uid int)
func Getwd() (wd string, err error)
func Getxattr(path string, attr string, dest []byte) (sz int, err error)
func InotifyAddWatch(fd int, pathname string, mask uint32) (watchdesc int, err error)
func InotifyInit() (fd int, err error)
func InotifyInit1(flags int) (fd int, err error)
func InotifyRmWatch(fd int, watchdesc uint32) (success int, err error)
func Ioperm(from int, num int, on int) (err error)
func Iopl(level int) (err error)
func Kill(pid int, sig Signal) (err error)
func Klogctl(typ int, buf []byte) (n int, err error)
func Lchown(path string, uid int, gid int) (err error)
func Link(oldpath string, newpath string) (err error)
func Listen(s int, n int) (err error)
func Listxattr(path string, dest []byte) (sz int, err error)
func LsfSocket(ifindex, proto int) (int, error)deprecated
func Lstat(path string, stat *Stat_t) (err error)
func Madvise(b []byte, advice int) (err error)
func Mkdir(path string, mode uint32) (err error)
func Mkdirat(dirfd int, path string, mode uint32) (err error)
func Mkfifo(path string, mode uint32) (err error)
func Mknod(path string, mode uint32, dev int) (err error)
func Mknodat(dirfd int, path string, mode uint32, dev int) (err error)
func Mlock(b []byte) (err error)
func Mlockall(flags int) (err error)
func Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error)
func Mount(source string, target string, fstype string, flags uintptr, data string) (err error)
func Mprotect(b []byte, prot int) (err error)
func Munlock(b []byte) (err error)
func Munlockall() (err error)
func Munmap(b []byte) (err error)
func Nanosleep(time *Timespec, leftover *Timespec) (err error)
func NetlinkRIB(proto, family int) ([]byte, error)
func Open(path string, mode int, perm uint32) (fd int, err error)
func Openat(dirfd int, path string, flags int, mode uint32) (fd int, err error)
func ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string)
func ParseUnixRights(m *SocketControlMessage) ([]int, error)
func Pause() (err error)
func Pipe(p []int) error
func Pipe2(p []int, flags int) error
func PivotRoot(newroot string, putold string) (err error)
func Pread(fd int, p []byte, offset int64) (n int, err error)
func PtraceAttach(pid int) (err error)
func PtraceCont(pid int, signal int) (err error)
func PtraceDetach(pid int) (err error)
func PtraceGetEventMsg(pid int) (msg uint, err error)
func PtraceGetRegs(pid int, regsout *PtraceRegs) (err error)
func PtracePeekData(pid int, addr uintptr, out []byte) (count int, err error)
func PtracePeekText(pid int, addr uintptr, out []byte) (count int, err error)
func PtracePokeData(pid int, addr uintptr, data []byte) (count int, err error)
func PtracePokeText(pid int, addr uintptr, data []byte) (count int, err error)
func PtraceSetOptions(pid int, options int) (err error)
func PtraceSetRegs(pid int, regs *PtraceRegs) (err error)
func PtraceSingleStep(pid int) (err error)
func PtraceSyscall(pid int, signal int) (err error)
func Pwrite(fd int, p []byte, offset int64) (n int, err error)
func Read(fd int, p []byte) (n int, err error)
func ReadDirent(fd int, buf []byte) (n int, err error)
func Readlink(path string, buf []byte) (n int, err error)
func Reboot(cmd int) (err error)
func Removexattr(path string, attr string) (err error)
func Rename(oldpath string, newpath string) (err error)
func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
func Rmdir(path string) error
func Seek(fd int, offset int64, whence int) (off int64, err error)
func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error)
func Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int) (err error)
func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags int) (n int, err error)
func Sendto(fd int, p []byte, flags int, to Sockaddr) (err error)
func SetLsfPromisc(name string, m bool) errordeprecated
func SetNonblock(fd int, nonblocking bool) (err error)
func Setdomainname(p []byte) (err error)
func Setegid(egid int) (err error)
func Setenv(key, value string) error
func Seteuid(euid int) (err error)
func Setfsgid(gid int) (err error)
func Setfsuid(uid int) (err error)
func Setgid(gid int) (err error)
func Setgroups(gids []int) (err error)
func Sethostname(p []byte) (err error)
func Setpgid(pid int, pgid int) (err error)
func Setpriority(which int, who int, prio int) (err error)
func Setregid(rgid, egid int) (err error)
func Setresgid(rgid, egid, sgid int) (err error)
func Setresuid(ruid, euid, suid int) (err error)
func Setreuid(ruid, euid int) (err error)
func Setrlimit(resource int, rlim *Rlimit) error
func Setsid() (pid int, err error)
func SetsockoptByte(fd, level, opt int, value byte) (err error)
func SetsockoptICMPv6Filter(fd, level, opt int, filter *ICMPv6Filter) error
func SetsockoptIPMreq(fd, level, opt int, mreq *IPMreq) (err error)
func SetsockoptIPMreqn(fd, level, opt int, mreq *IPMreqn) (err error)
func SetsockoptIPv6Mreq(fd, level, opt int, mreq *IPv6Mreq) (err error)
func SetsockoptInet4Addr(fd, level, opt int, value [4]byte) (err error)
func SetsockoptInt(fd, level, opt int, value int) (err error)
func SetsockoptLinger(fd, level, opt int, l *Linger) (err error)
func SetsockoptString(fd, level, opt int, s string) (err error)
func SetsockoptTimeval(fd, level, opt int, tv *Timeval) (err error)
func Settimeofday(tv *Timeval) (err error)
func Setuid(uid int) (err error)
func Setxattr(path string, attr string, data []byte, flags int) (err error)
func Shutdown(fd int, how int) (err error)
func SlicePtrFromStrings(ss []string) ([]*byte, error)
func Socket(domain, typ, proto int) (fd int, err error)
func Socketpair(domain, typ, proto int) (fd [2]int, err error)
func Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)
func StartProcess(argv0 string, argv []string, attr *ProcAttr) (pid int, handle uintptr, err error)
func Stat(path string, stat *Stat_t) (err error)
func Statfs(path string, buf *Statfs_t) (err error)
func StringBytePtr(s string) *bytedeprecated
func StringByteSlice(s string) []bytedeprecated
func StringSlicePtr(ss []string) []*bytedeprecated
func Symlink(oldpath string, newpath string) (err error)
func Sync()
func SyncFileRange(fd int, off int64, n int64, flags int) (err error)
func Sysinfo(info *Sysinfo_t) (err error)
func Tee(rfd int, wfd int, len int, flags int) (n int64, err error)
func Tgkill(tgid int, tid int, sig Signal) (err error)
func Times(tms *Tms) (ticks uintptr, err error)
func TimespecToNsec(ts Timespec) int64
func TimevalToNsec(tv Timeval) int64
func Truncate(path string, length int64) (err error)
func Umask(mask int) (oldmask int)
func Uname(buf *Utsname) (err error)
func UnixCredentials(ucred *Ucred) []byte
func UnixRights(fds ...int) []byte
func Unlink(path string) error
func Unlinkat(dirfd int, path string) error
func Unmount(target string, flags int) (err error)
func Unsetenv(key string) error
func Unshare(flags int) (err error)
func Ustat(dev int, ubuf *Ustat_t) (err error)
func Utime(path string, buf *Utimbuf) (err error)
func Utimes(path string, tv []Timeval) (err error)
func UtimesNano(path string, ts []Timespec) (err error)
func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (wpid int, err error)
func Write(fd int, p []byte) (n int, err error)
type Cmsghdr
func (cmsg *Cmsghdr) SetLen(length int)
type Conn
type Credential
type Dirent
type EpollEvent
type Errno
func AllThreadsSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func AllThreadsSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func (e Errno) Error() string
func (e Errno) Is(target error) bool
func (e Errno) Temporary() bool
func (e Errno) Timeout() bool
type FdSet
type Flock_t
type Fsid
type ICMPv6Filter
func GetsockoptICMPv6Filter(fd, level, opt int) (*ICMPv6Filter, error)
type IPMreq
func GetsockoptIPMreq(fd, level, opt int) (*IPMreq, error)
type IPMreqn
func GetsockoptIPMreqn(fd, level, opt int) (*IPMreqn, error)
type IPv6MTUInfo
func GetsockoptIPv6MTUInfo(fd, level, opt int) (*IPv6MTUInfo, error)
type IPv6Mreq
func GetsockoptIPv6Mreq(fd, level, opt int) (*IPv6Mreq, error)
type IfAddrmsg
type IfInfomsg
type Inet4Pktinfo
type Inet6Pktinfo
type InotifyEvent
type Iovec
func (iov *Iovec) SetLen(length int)
type Linger
type Msghdr
func (msghdr *Msghdr) SetControllen(length int)
type NetlinkMessage
func ParseNetlinkMessage(b []byte) ([]NetlinkMessage, error)
type NetlinkRouteAttr
func ParseNetlinkRouteAttr(m *NetlinkMessage) ([]NetlinkRouteAttr, error)
type NetlinkRouteRequest
type NlAttr
type NlMsgerr
type NlMsghdr
type ProcAttr
type PtraceRegs
func (r *PtraceRegs) PC() uint64
func (r *PtraceRegs) SetPC(pc uint64)
type RawConn
type RawSockaddr
type RawSockaddrAny
type RawSockaddrInet4
type RawSockaddrInet6
type RawSockaddrLinklayer
type RawSockaddrNetlink
type RawSockaddrUnix
type Rlimit
type RtAttr
type RtGenmsg
type RtMsg
type RtNexthop
type Rusage
type Signal
func (s Signal) Signal()
func (s Signal) String() string
type SockFilter
func LsfJump(code, k, jt, jf int) *SockFilterdeprecated
func LsfStmt(code, k int) *SockFilterdeprecated
type SockFprog
type Sockaddr
func Accept(fd int) (nfd int, sa Sockaddr, err error)
func Accept4(fd int, flags int) (nfd int, sa Sockaddr, err error)
func Getpeername(fd int) (sa Sockaddr, err error)
func Getsockname(fd int) (sa Sockaddr, err error)
func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, err error)
func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int, from Sockaddr, err error)
type SockaddrInet4
type SockaddrInet6
type SockaddrLinklayer
type SockaddrNetlink
type SockaddrUnix
type SocketControlMessage
func ParseSocketControlMessage(b []byte) ([]SocketControlMessage, error)
type Stat_t
type Statfs_t
type SysProcAttr
type SysProcIDMap
type Sysinfo_t
type TCPInfo
type Termios
type Time_t
func Time(t *Time_t) (tt Time_t, err error)
type Timespec
func NsecToTimespec(nsec int64) Timespec
func (ts *Timespec) Nano() int64
func (ts *Timespec) Unix() (sec int64, nsec int64)
type Timeval
func NsecToTimeval(nsec int64) Timeval
func (tv *Timeval) Nano() int64
func (tv *Timeval) Unix() (sec int64, nsec int64)
type Timex
type Tms
type Ucred
func GetsockoptUcred(fd, level, opt int) (*Ucred, error)
func ParseUnixCredentials(m *SocketControlMessage) (*Ucred, error)
type Ustat_t
type Utimbuf
type Utsname
type WaitStatus
func (w WaitStatus) Continued() bool
func (w WaitStatus) CoreDump() bool
func (w WaitStatus) ExitStatus() int
func (w WaitStatus) Exited() bool
func (w WaitStatus) Signal() Signal
func (w WaitStatus) Signaled() bool
func (w WaitStatus) StopSignal() Signal
func (w WaitStatus) Stopped() bool
func (w WaitStatus) TrapCause() int

## testing库
func AllocsPerRun(runs int, f func()) (avg float64)
func CoverMode() string
func Coverage() float64
func Init()
func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, ...)
func RegisterCover(c Cover)
func RunBenchmarks(matchString func(pat, str string) (bool, error), ...)
func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)
func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)
func Short() bool
func Testing() bool
func Verbose() bool
type B
func (c *B) Chdir(dir string)
func (c *B) Cleanup(f func())
func (c *B) Context() context.Context
func (b *B) Elapsed() time.Duration
func (c *B) Error(args ...any)
func (c *B) Errorf(format string, args ...any)
func (c *B) Fail()
func (c *B) FailNow()
func (c *B) Failed() bool
func (c *B) Fatal(args ...any)
func (c *B) Fatalf(format string, args ...any)
func (c *B) Helper()
func (c *B) Log(args ...any)
func (c *B) Logf(format string, args ...any)
func (b *B) Loop() bool
func (c *B) Name() string
func (b *B) ReportAllocs()
func (b *B) ReportMetric(n float64, unit string)
func (b *B) ResetTimer()
func (b *B) Run(name string, f func(b *B)) bool
func (b *B) RunParallel(body func(*PB))
func (b *B) SetBytes(n int64)
func (b *B) SetParallelism(p int)
func (c *B) Setenv(key, value string)
func (c *B) Skip(args ...any)
func (c *B) SkipNow()
func (c *B) Skipf(format string, args ...any)
func (c *B) Skipped() bool
func (b *B) StartTimer()
func (b *B) StopTimer()
func (c *B) TempDir() string
type BenchmarkResult
func Benchmark(f func(b *B)) BenchmarkResult
func (r BenchmarkResult) AllocedBytesPerOp() int64
func (r BenchmarkResult) AllocsPerOp() int64
func (r BenchmarkResult) MemString() string
func (r BenchmarkResult) NsPerOp() int64
func (r BenchmarkResult) String() string
type Cover
type CoverBlock
type F
func (f *F) Add(args ...any)
func (c *F) Chdir(dir string)
func (c *F) Cleanup(f func())
func (c *F) Context() context.Context
func (c *F) Error(args ...any)
func (c *F) Errorf(format string, args ...any)
func (f *F) Fail()
func (c *F) FailNow()
func (c *F) Failed() bool
func (c *F) Fatal(args ...any)
func (c *F) Fatalf(format string, args ...any)
func (f *F) Fuzz(ff any)
func (f *F) Helper()
func (c *F) Log(args ...any)
func (c *F) Logf(format string, args ...any)
func (c *F) Name() string
func (c *F) Setenv(key, value string)
func (c *F) Skip(args ...any)
func (c *F) SkipNow()
func (c *F) Skipf(format string, args ...any)
func (f *F) Skipped() bool
func (c *F) TempDir() string
type InternalBenchmark
type InternalExample
type InternalFuzzTarget
type InternalTest
type M
func MainStart(deps testDeps, tests []InternalTest, benchmarks []InternalBenchmark, ...) *M
func (m *M) Run() (code int)
type PB
func (pb *PB) Next() bool
type T
func (t *T) Chdir(dir string)
func (c *T) Cleanup(f func())
func (c *T) Context() context.Context
func (t *T) Deadline() (deadline time.Time, ok bool)
func (c *T) Error(args ...any)
func (c *T) Errorf(format string, args ...any)
func (c *T) Fail()
func (c *T) FailNow()
func (c *T) Failed() bool
func (c *T) Fatal(args ...any)
func (c *T) Fatalf(format string, args ...any)
func (c *T) Helper()
func (c *T) Log(args ...any)
func (c *T) Logf(format string, args ...any)
func (c *T) Name() string
func (t *T) Parallel()
func (t *T) Run(name string, f func(t *T)) bool
func (t *T) Setenv(key, value string)
func (c *T) Skip(args ...any)
func (c *T) SkipNow()
func (c *T) Skipf(format string, args ...any)
func (c *T) Skipped() bool
func (c *T) TempDir() string
type TB



## strconv库
该库提供了将基本数据类型转为字符串的操作。

`ParseBool`函数将传入的字符串转为一个布尔值：
```go
result := strconv.ParseBool("1")
```

`ParseComplex`

`ParseFloat`函数将一个字符串解析为浮点数：
```go
result := strconv.ParseBool("BA",16,strconv.Int64)
```

`ParseInt`函数将传入的字符串转为一个整数：
```go
/**
 * 参数二表示字符串要转换进制类型，可以是2~32，如果值为0则根据字符串前缀判断字符串的进制类型
 * 参数三表示预期的数值bit大小
 */
result := strconv.ParseBool("BA",16,strconv.Int64)
```

`ParseUint`函数将传入的字符串转为一个正整数：
```go
res, err := strconv.ParseUint("12", 10, strconv.IntSize)
```

`FormatBool`函数将传入的布尔值转为字符串返回：
```go
res := strconv.FormatBool(true)
```

`FormatComplex`
```go

```

`FormatFloat`函数将传入的浮点数转为字符串返回：
```go
res := strconv.FormatFloat(12.9)
```

`FormatInt`函数将传入的整数转为字符串返回：
```go
res := strconv.FormatInt(12)
```

`FormatUint`函数将传入的正整数转为字符串返回：
```go
res := strconv.FormatUint(12)
```

`Atoid`函数是`ParseInt(s, 10, 0)`的简写：
```go
res := strconv.Atoid("11")
```

`Itoa`函数是`FormateInt(i,10)`的简写：
```go
res := strconv.Itoa(12)
```

`AppendBool`函数将一个布尔类型的值追加到原始的字符串切片末尾：
```go
res := strconv.AppendBool([]byte("Hello"), true)
```

`AppendFloat`函数将一个浮点数的值追加到原始的字符串切片末尾：
```go
res := strconv.AppendFloat([]byte("Hello"), 12.8)
```

`AppendInt`
`AppendUint`

`CanBackquote`函数用于检查字符串是否是单行的不包含除了空格和制表符之外的特殊字符：
```go
res := strconv.CanBackquote("Hello \r\n world")
``

`IsPrint`函数判断传入的内容是否可以被打印：
```go

```

```go
// 返回Go双引号引用的 Go字面量表示的 字符串s。对于控制字符和不可打印的字符(!IsPrint)，
//  使用Go转义序列进行表示（\t, \n, \xFF, \u0100）。
func Quote(s string) string
// 类似Quote，但对于非ASCII符号表示为Go转义序列。
func QuoteToASCII(s string) string
// 类似Quote，但对于 控制字符 和 非图片字符(!IsGraphic) 使用Go转义序列表示
func QuoteToGraphic(s string) string
// 返回Go单引号引用的 Go字面量表示的 字符串s。对于控制字符和不可打印的字符(!IsPrint)，
//  使用Go转义序列进行表示
func QuoteRune(r rune) string
// 类似QuoteRune，但对于非ASCII符号表示为Go转义序列。
func QuoteRuneToASCII(r rune) string
// 类似QuoteRune，但对于控制字符和非图片字符使用Go转义序列表示
func QuoteRuneToGraphic(r rune) string
// 返回被引号（单引号【引用单个字符】、双引号【引用字符串】、反引号【引用字符串】）引用的前缀，如果没有被引号引用的前缀则返回错误。
func QuotedPrefix(s string) (string, error)
// 将s解释为引号引用的字符串，返回去掉该引号后的字符串。
//  注意，如果引号为单引号，则被引用的应该只有一个unicode字符。
func Unquote(s string) (string, error)
// 比较复杂，我将英文原文写在下面，可自行解读
//  大致意思就是将s进行解码，返回解码后的字符、该字符是否是多字节的、后续的字符串。
//  quote表示解码方式，可选为单引号、双引号和0
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
```

## strings库
func Clone(s string) string
func Compare(a, b string) int
func Contains(s, substr string) bool
func ContainsAny(s, chars string) bool
func ContainsFunc(s string, f func(rune) bool) bool
func ContainsRune(s string, r rune) bool
func Count(s, substr string) int
func Cut(s, sep string) (before, after string, found bool)
func CutPrefix(s, prefix string) (after string, found bool)
func CutSuffix(s, suffix string) (before string, found bool)
func EqualFold(s, t string) bool
func Fields(s string) []string
func FieldsFunc(s string, f func(rune) bool) []string
func FieldsFuncSeq(s string, f func(rune) bool) iter.Seq[string]
func FieldsSeq(s string) iter.Seq[string]
func HasPrefix(s, prefix string) bool
func HasSuffix(s, suffix string) bool
func Index(s, substr string) int
func IndexAny(s, chars string) int
func IndexByte(s string, c byte) int
func IndexFunc(s string, f func(rune) bool) int
func IndexRune(s string, r rune) int
func Join(elems []string, sep string) string
func LastIndex(s, substr string) int
func LastIndexAny(s, chars string) int
func LastIndexByte(s string, c byte) int
func LastIndexFunc(s string, f func(rune) bool) int
func Lines(s string) iter.Seq[string]
func Map(mapping func(rune) rune, s string) string
func Repeat(s string, count int) string
func Replace(s, old, new string, n int) string
func ReplaceAll(s, old, new string) string
func Split(s, sep string) []string
func SplitAfter(s, sep string) []string
func SplitAfterN(s, sep string, n int) []string
func SplitAfterSeq(s, sep string) iter.Seq[string]
func SplitN(s, sep string, n int) []string
func SplitSeq(s, sep string) iter.Seq[string]
func Title(s string) stringdeprecated
func ToLower(s string) string
func ToLowerSpecial(c unicode.SpecialCase, s string) string
func ToTitle(s string) string
func ToTitleSpecial(c unicode.SpecialCase, s string) string
func ToUpper(s string) string
func ToUpperSpecial(c unicode.SpecialCase, s string) string
func ToValidUTF8(s, replacement string) string
func Trim(s, cutset string) string
func TrimFunc(s string, f func(rune) bool) string
func TrimLeft(s, cutset string) string
func TrimLeftFunc(s string, f func(rune) bool) string
func TrimPrefix(s, prefix string) string
func TrimRight(s, cutset string) string
func TrimRightFunc(s string, f func(rune) bool) string
func TrimSpace(s string) string
func TrimSuffix(s, suffix string) string
type Builder
func (b *Builder) Cap() int
func (b *Builder) Grow(n int)
func (b *Builder) Len() int
func (b *Builder) Reset()
func (b *Builder) String() string
func (b *Builder) Write(p []byte) (int, error)
func (b *Builder) WriteByte(c byte) error
func (b *Builder) WriteRune(r rune) (int, error)
func (b *Builder) WriteString(s string) (int, error)
type Reader
func NewReader(s string) *Reader
func (r *Reader) Len() int
func (r *Reader) Read(b []byte) (n int, err error)
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
func (r *Reader) ReadByte() (byte, error)
func (r *Reader) ReadRune() (ch rune, size int, err error)
func (r *Reader) Reset(s string)
func (r *Reader) Seek(offset int64, whence int) (int64, error)
func (r *Reader) Size() int64
func (r *Reader) UnreadByte() error
func (r *Reader) UnreadRune() error
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
type Replacer
func NewReplacer(oldnew ...string) *Replacer
func (r *Replacer) Replace(s string) string
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)

## bytes库
`bytes`库提供了高效的处理字节切片的功能。

`Compare`函数用于比较两个切片的大小，如果相等返回0，如果参数一小于参数二返回-1，否则返回1：
```go
func main() {
   fmt.Println(bytes.Compare([]byte{},[]byte{})) // 0
   fmt.Println(bytes.Compare([]byte{1},[]byte{2})) // -1
   fmt.Println(bytes.Compare([]byte{2},[]byte{1})) // 1
   fmt.Println(bytes.Compare([]byte{},nil)) //0
   fmt.Println([]byte{} == nil)  // false
}
```

`Equal`函数用于判断两个切片是否完全相等（**需要注意nil被视为空切片**）：
```go
func main() {
   fmt.Println(bytes.Equal([]byte{},[]byte{})) // true
   fmt.Println(bytes.Equal([]byte{'A', 'B'},[]byte{'a'})) // false
   fmt.Println(bytes.Equal([]byte{'a'},[]byte{'a'})) // true
   fmt.Println(bytes.Equal([]byte{},nil)) // true
   fmt.Println([]byte{} == nil)  // false
}
```

`EqualFold`函数比较两个字符串是否相等并且不区分大小写：
```go
func main() {
    fmt.Println(bytes.EqualFold([]byte{},[]byte{})) // true
    fmt.Println(bytes.EqualFold([]byte{'A'},[]byte{'a'})) // true
    fmt.Println(bytes.EqualFold([]byte{'B'},[]byte{'a'})) // false
    fmt.Println(bytes.EqualFold([]byte{},nil)) // true
}
```

`Runes`函数将字符串转为对应的码点：
```go
func main() {
   fmt.Println([]byte("你好world")) // [228 189 160 229 165 189 119 111 114 108 100]
   fmt.Println(bytes.Runes([]byte("你好world"))) // [20320 22909 119 111 114 108 100]
   fmt.Println([]rune("你好world"))  // [20320 22909 119 111 114 108 100]            
}
```

`HasPrefix`函数查看字符串是否以指定字符开头：
```go

```

`HasSuffix`函数查看字符串是否以指定字符结尾：
```go

```

`Contains`函数查看字符串是否包含指定字符串：
```go

```

`Count`函数查看字符串中包含多少个不同的指定字符：
```go

```

`Index`函数查看字符串中指定字符的第一次出现的位置：
```go

```

`IndexByte`函数查看字节在指定字符串中第一次出现的位置：
```go

```

`IndexRune`函数查看字符在指定字符串中第一次出现的位置：
```go

```

`LastIndex`函数查看字符串在指定字符串中最后一次出现的位置：
```go

```

`LastIndexByte`函数查看字节在指定字符串中最后一次出现的位置：
```go

```

`Title`函数将每个单词的首字符大写并返回：
```go
fmt.Printf("%s\n", bytes.Title([]byte("her royal highness")))  // Her Royal Highness
fmt.Printf("%s\n", bytes.Title([]byte("hEr royal highness")))  // HEr Royal Highness
```

`ToLower`函数将所有字符转为小写并返回：
```go
fmt.Printf("%s", bytes.ToLower([]byte("Gopher")))  // gopher
```

`ToUpper`函数将所有字符转为大写并返回：
```go
fmt.Printf("%s", bytes.ToUpper([]byte("Gopher")))  // GOPHER
```

`Repeat`函数将指定的字符串重复n次后返回：
```go
res := bytes.Repeat([]byte("0"),10)
```

`Trim`函数将字符串前后的指定字符删除：
```go
res := bytes.Trim([]byte("1122  0 2211"),"1")
```

`TrimSpace`函数将字符串前后的空格删除：
```go
res := bytes.TrimSpace([]byte("  0  "))
```

`Split`函数将字符串按照指定字符进行分割：
```go
res := bytes.Split([]byte("1,2,3,4"))
```

`Join`函数将切片中的字符按照指定字符进行拼接：
```go

```

## maps库
func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V]
func Clone[M ~map[K]V, K comparable, V any](m M) M
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2)
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool)
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool
func Insert[Map ~map[K]V, K comparable, V any](m Map, seq iter.Seq2[K, V])
func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K]
func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V]

## slices库
func All[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]
func AppendSeq[Slice ~[]E, E any](s Slice, seq iter.Seq[E]) Slice
func Backward[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]
func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool)
func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool)
func Chunk[Slice ~[]E, E any](s Slice, n int) iter.Seq[Slice]
func Clip[S ~[]E, E any](s S) S
func Clone[S ~[]E, E any](s S) S
func Collect[E any](seq iter.Seq[E]) []E
func Compact[S ~[]E, E comparable](s S) S
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S
func Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int
func CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int
func Concat[S ~[]E, E any](slices ...S) S
func Contains[S ~[]E, E comparable](s S, v E) bool
func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool
func Delete[S ~[]E, E any](s S, i, j int) S
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S
func Equal[S ~[]E, E comparable](s1, s2 S) bool
func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool
func Grow[S ~[]E, E any](s S, n int) S
func Index[S ~[]E, E comparable](s S, v E) int
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int
func Insert[S ~[]E, E any](s S, i int, v ...E) S
func IsSorted[S ~[]E, E cmp.Ordered](x S) bool
func IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool
func Max[S ~[]E, E cmp.Ordered](x S) E
func MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E
func Min[S ~[]E, E cmp.Ordered](x S) E
func MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E
func Repeat[S ~[]E, E any](x S, count int) S
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S
func Reverse[S ~[]E, E any](s S)
func Sort[S ~[]E, E cmp.Ordered](x S)
func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int)
func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int)
func Sorted[E cmp.Ordered](seq iter.Seq[E]) []E
func SortedFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E
func SortedStableFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E
func Values[Slice ~[]E, E any](s Slice) iter.Seq[E]

## math库
func Abs(x float64) float64
func Acos(x float64) float64
func Acosh(x float64) float64
func Asin(x float64) float64
func Asinh(x float64) float64
func Atan(x float64) float64
func Atan2(y, x float64) float64
func Atanh(x float64) float64
func Cbrt(x float64) float64
func Ceil(x float64) float64
func Copysign(f, sign float64) float64
func Cos(x float64) float64
func Cosh(x float64) float64
func Dim(x, y float64) float64
func Erf(x float64) float64
func Erfc(x float64) float64
func Erfcinv(x float64) float64
func Erfinv(x float64) float64
func Exp(x float64) float64
func Exp2(x float64) float64
func Expm1(x float64) float64
func FMA(x, y, z float64) float64
func Float32bits(f float32) uint32
func Float32frombits(b uint32) float32
func Float64bits(f float64) uint64
func Float64frombits(b uint64) float64
func Floor(x float64) float64
func Frexp(f float64) (frac float64, exp int)
func Gamma(x float64) float64
func Hypot(p, q float64) float64
func Ilogb(x float64) int
func Inf(sign int) float64
func IsInf(f float64, sign int) bool
func IsNaN(f float64) (is bool)
func J0(x float64) float64
func J1(x float64) float64
func Jn(n int, x float64) float64
func Ldexp(frac float64, exp int) float64
func Lgamma(x float64) (lgamma float64, sign int)
func Log(x float64) float64
func Log10(x float64) float64
func Log1p(x float64) float64
func Log2(x float64) float64
func Logb(x float64) float64
func Max(x, y float64) float64
func Min(x, y float64) float64
func Mod(x, y float64) float64
func Modf(f float64) (int float64, frac float64)
func NaN() float64
func Nextafter(x, y float64) (r float64)
func Nextafter32(x, y float32) (r float32)
func Pow(x, y float64) float64
func Pow10(n int) float64
func Remainder(x, y float64) float64
func Round(x float64) float64
func RoundToEven(x float64) float64
func Signbit(x float64) bool
func Sin(x float64) float64
func Sincos(x float64) (sin, cos float64)
func Sinh(x float64) float64
func Sqrt(x float64) float64
func Tan(x float64) float64
func Tanh(x float64) float64
func Trunc(x float64) float64
func Y0(x float64) float64
func Y1(x float64) float64
func Yn(n int, x float64) float64

## math/bits库
func Add(x, y, carry uint) (sum, carryOut uint)
func Add32(x, y, carry uint32) (sum, carryOut uint32)
func Add64(x, y, carry uint64) (sum, carryOut uint64)
func Div(hi, lo, y uint) (quo, rem uint)
func Div32(hi, lo, y uint32) (quo, rem uint32)
func Div64(hi, lo, y uint64) (quo, rem uint64)
func LeadingZeros(x uint) int
func LeadingZeros16(x uint16) int
func LeadingZeros32(x uint32) int
func LeadingZeros64(x uint64) int
func LeadingZeros8(x uint8) int
func Len(x uint) int
func Len16(x uint16) (n int)
func Len32(x uint32) (n int)
func Len64(x uint64) (n int)
func Len8(x uint8) int
func Mul(x, y uint) (hi, lo uint)
func Mul32(x, y uint32) (hi, lo uint32)
func Mul64(x, y uint64) (hi, lo uint64)
func OnesCount(x uint) int
func OnesCount16(x uint16) int
func OnesCount32(x uint32) int
func OnesCount64(x uint64) int
func OnesCount8(x uint8) int
func Rem(hi, lo, y uint) uint
func Rem32(hi, lo, y uint32) uint32
func Rem64(hi, lo, y uint64) uint64
func Reverse(x uint) uint
func Reverse16(x uint16) uint16
func Reverse32(x uint32) uint32
func Reverse64(x uint64) uint64
func Reverse8(x uint8) uint8
func ReverseBytes(x uint) uint
func ReverseBytes16(x uint16) uint16
func ReverseBytes32(x uint32) uint32
func ReverseBytes64(x uint64) uint64
func RotateLeft(x uint, k int) uint
func RotateLeft16(x uint16, k int) uint16
func RotateLeft32(x uint32, k int) uint32
func RotateLeft64(x uint64, k int) uint64
func RotateLeft8(x uint8, k int) uint8
func Sub(x, y, borrow uint) (diff, borrowOut uint)
func Sub32(x, y, borrow uint32) (diff, borrowOut uint32)
func Sub64(x, y, borrow uint64) (diff, borrowOut uint64)
func TrailingZeros(x uint) int
func TrailingZeros16(x uint16) int
func TrailingZeros32(x uint32) int
func TrailingZeros64(x uint64) int
func TrailingZeros8(x uint8) int

## math/cmplx库
func Abs(x complex128) float64
func Acos(x complex128) complex128
func Acosh(x complex128) complex128
func Asin(x complex128) complex128
func Asinh(x complex128) complex128
func Atan(x complex128) complex128
func Atanh(x complex128) complex128
func Conj(x complex128) complex128
func Cos(x complex128) complex128
func Cosh(x complex128) complex128
func Cot(x complex128) complex128
func Exp(x complex128) complex128
func Inf() complex128
func IsInf(x complex128) bool
func IsNaN(x complex128) bool
func Log(x complex128) complex128
func Log10(x complex128) complex128
func NaN() complex128
func Phase(x complex128) float64
func Polar(x complex128) (r, θ float64)
func Pow(x, y complex128) complex128
func Rect(r, θ float64) complex128
func Sin(x complex128) complex128
func Sinh(x complex128) complex128
func Sqrt(x complex128) complex128
func Tan(x complex128) complex128
func Tanh(x complex128) complex128

## math/rand库
func ExpFloat64() float64
func Float32() float32
func Float64() float64
func Int() int
func Int31() int32
func Int31n(n int32) int32
func Int63() int64
func Int63n(n int64) int64
func Intn(n int) int
func NormFloat64() float64
func Perm(n int) []int
func Read(p []byte) (n int, err error)deprecated
func Seed(seed int64)deprecated
func Shuffle(n int, swap func(i, j int))
func Uint32() uint32
func Uint64() uint64
type Rand
func New(src Source) *Rand
func (r *Rand) ExpFloat64() float64
func (r *Rand) Float32() float32
func (r *Rand) Float64() float64
func (r *Rand) Int() int
func (r *Rand) Int31() int32
func (r *Rand) Int31n(n int32) int32
func (r *Rand) Int63() int64
func (r *Rand) Int63n(n int64) int64
func (r *Rand) Intn(n int) int
func (r *Rand) NormFloat64() float64
func (r *Rand) Perm(n int) []int
func (r *Rand) Read(p []byte) (n int, err error)
func (r *Rand) Seed(seed int64)
func (r *Rand) Shuffle(n int, swap func(i, j int))
func (r *Rand) Uint32() uint32
func (r *Rand) Uint64() uint64
type Source
func NewSource(seed int64) Source
type Source64
type Zipf
func NewZipf(r *Rand, s float64, v float64, imax uint64) *Zipf
func (z *Zipf) Uint64() uint64

## math/big库
func Jacobi(x, y *Int) int
type Accuracy
func (i Accuracy) String() string
type ErrNaN
func (err ErrNaN) Error() string
type Float
func NewFloat(x float64) *Float
func ParseFloat(s string, base int, prec uint, mode RoundingMode) (f *Float, b int, err error)
func (z *Float) Abs(x *Float) *Float
func (x *Float) Acc() Accuracy
func (z *Float) Add(x, y *Float) *Float
func (x *Float) Append(buf []byte, fmt byte, prec int) []byte
func (x *Float) AppendText(b []byte) ([]byte, error)
func (x *Float) Cmp(y *Float) int
func (z *Float) Copy(x *Float) *Float
func (x *Float) Float32() (float32, Accuracy)
func (x *Float) Float64() (float64, Accuracy)
func (x *Float) Format(s fmt.State, format rune)
func (z *Float) GobDecode(buf []byte) error
func (x *Float) GobEncode() ([]byte, error)
func (x *Float) Int(z *Int) (*Int, Accuracy)
func (x *Float) Int64() (int64, Accuracy)
func (x *Float) IsInf() bool
func (x *Float) IsInt() bool
func (x *Float) MantExp(mant *Float) (exp int)
func (x *Float) MarshalText() (text []byte, err error)
func (x *Float) MinPrec() uint
func (x *Float) Mode() RoundingMode
func (z *Float) Mul(x, y *Float) *Float
func (z *Float) Neg(x *Float) *Float
func (z *Float) Parse(s string, base int) (f *Float, b int, err error)
func (x *Float) Prec() uint
func (z *Float) Quo(x, y *Float) *Float
func (x *Float) Rat(z *Rat) (*Rat, Accuracy)
func (z *Float) Scan(s fmt.ScanState, ch rune) error
func (z *Float) Set(x *Float) *Float
func (z *Float) SetFloat64(x float64) *Float
func (z *Float) SetInf(signbit bool) *Float
func (z *Float) SetInt(x *Int) *Float
func (z *Float) SetInt64(x int64) *Float
func (z *Float) SetMantExp(mant *Float, exp int) *Float
func (z *Float) SetMode(mode RoundingMode) *Float
func (z *Float) SetPrec(prec uint) *Float
func (z *Float) SetRat(x *Rat) *Float
func (z *Float) SetString(s string) (*Float, bool)
func (z *Float) SetUint64(x uint64) *Float
func (x *Float) Sign() int
func (x *Float) Signbit() bool
func (z *Float) Sqrt(x *Float) *Float
func (x *Float) String() string
func (z *Float) Sub(x, y *Float) *Float
func (x *Float) Text(format byte, prec int) string
func (x *Float) Uint64() (uint64, Accuracy)
func (z *Float) UnmarshalText(text []byte) error
type Int
func NewInt(x int64) *Int
func (z *Int) Abs(x *Int) *Int
func (z *Int) Add(x, y *Int) *Int
func (z *Int) And(x, y *Int) *Int
func (z *Int) AndNot(x, y *Int) *Int
func (x *Int) Append(buf []byte, base int) []byte
func (x *Int) AppendText(b []byte) (text []byte, err error)
func (z *Int) Binomial(n, k int64) *Int
func (x *Int) Bit(i int) uint
func (x *Int) BitLen() int
func (x *Int) Bits() []Word
func (x *Int) Bytes() []byte
func (x *Int) Cmp(y *Int) (r int)
func (x *Int) CmpAbs(y *Int) int
func (z *Int) Div(x, y *Int) *Int
func (z *Int) DivMod(x, y, m *Int) (*Int, *Int)
func (z *Int) Exp(x, y, m *Int) *Int
func (x *Int) FillBytes(buf []byte) []byte
func (x *Int) Float64() (float64, Accuracy)
func (x *Int) Format(s fmt.State, ch rune)
func (z *Int) GCD(x, y, a, b *Int) *Int
func (z *Int) GobDecode(buf []byte) error
func (x *Int) GobEncode() ([]byte, error)
func (x *Int) Int64() int64
func (x *Int) IsInt64() bool
func (x *Int) IsUint64() bool
func (z *Int) Lsh(x *Int, n uint) *Int
func (x *Int) MarshalJSON() ([]byte, error)
func (x *Int) MarshalText() (text []byte, err error)
func (z *Int) Mod(x, y *Int) *Int
func (z *Int) ModInverse(g, n *Int) *Int
func (z *Int) ModSqrt(x, p *Int) *Int
func (z *Int) Mul(x, y *Int) *Int
func (z *Int) MulRange(a, b int64) *Int
func (z *Int) Neg(x *Int) *Int
func (z *Int) Not(x *Int) *Int
func (z *Int) Or(x, y *Int) *Int
func (x *Int) ProbablyPrime(n int) bool
func (z *Int) Quo(x, y *Int) *Int
func (z *Int) QuoRem(x, y, r *Int) (*Int, *Int)
func (z *Int) Rand(rnd *rand.Rand, n *Int) *Int
func (z *Int) Rem(x, y *Int) *Int
func (z *Int) Rsh(x *Int, n uint) *Int
func (z *Int) Scan(s fmt.ScanState, ch rune) error
func (z *Int) Set(x *Int) *Int
func (z *Int) SetBit(x *Int, i int, b uint) *Int
func (z *Int) SetBits(abs []Word) *Int
func (z *Int) SetBytes(buf []byte) *Int
func (z *Int) SetInt64(x int64) *Int
func (z *Int) SetString(s string, base int) (*Int, bool)
func (z *Int) SetUint64(x uint64) *Int
func (x *Int) Sign() int
func (z *Int) Sqrt(x *Int) *Int
func (x *Int) String() string
func (z *Int) Sub(x, y *Int) *Int
func (x *Int) Text(base int) string
func (x *Int) TrailingZeroBits() uint
func (x *Int) Uint64() uint64
func (z *Int) UnmarshalJSON(text []byte) error
func (z *Int) UnmarshalText(text []byte) error
func (z *Int) Xor(x, y *Int) *Int
type Rat
func NewRat(a, b int64) *Rat
func (z *Rat) Abs(x *Rat) *Rat
func (z *Rat) Add(x, y *Rat) *Rat
func (x *Rat) AppendText(b []byte) ([]byte, error)
func (x *Rat) Cmp(y *Rat) int
func (x *Rat) Denom() *Int
func (x *Rat) Float32() (f float32, exact bool)
func (x *Rat) Float64() (f float64, exact bool)
func (x *Rat) FloatPrec() (n int, exact bool)
func (x *Rat) FloatString(prec int) string
func (z *Rat) GobDecode(buf []byte) error
func (x *Rat) GobEncode() ([]byte, error)
func (z *Rat) Inv(x *Rat) *Rat
func (x *Rat) IsInt() bool
func (x *Rat) MarshalText() (text []byte, err error)
func (z *Rat) Mul(x, y *Rat) *Rat
func (z *Rat) Neg(x *Rat) *Rat
func (x *Rat) Num() *Int
func (z *Rat) Quo(x, y *Rat) *Rat
func (x *Rat) RatString() string
func (z *Rat) Scan(s fmt.ScanState, ch rune) error
func (z *Rat) Set(x *Rat) *Rat
func (z *Rat) SetFloat64(f float64) *Rat
func (z *Rat) SetFrac(a, b *Int) *Rat
func (z *Rat) SetFrac64(a, b int64) *Rat
func (z *Rat) SetInt(x *Int) *Rat
func (z *Rat) SetInt64(x int64) *Rat
func (z *Rat) SetString(s string) (*Rat, bool)
func (z *Rat) SetUint64(x uint64) *Rat
func (x *Rat) Sign() int
func (x *Rat) String() string
func (z *Rat) Sub(x, y *Rat) *Rat
func (z *Rat) UnmarshalText(text []byte) error
type RoundingMode
func (i RoundingMode) String() string
type Word

## container/heap库
func Fix(h Interface, i int)
func Init(h Interface)
func Pop(h Interface) any
func Push(h Interface, x any)
func Remove(h Interface, i int) any
type Interface

## container/list库
type Element
func (e *Element) Next() *Element
func (e *Element) Prev() *Element
type List
func New() *List
func (l *List) Back() *Element
func (l *List) Front() *Element
func (l *List) Init() *List
func (l *List) InsertAfter(v any, mark *Element) *Element
func (l *List) InsertBefore(v any, mark *Element) *Element
func (l *List) Len() int
func (l *List) MoveAfter(e, mark *Element)
func (l *List) MoveBefore(e, mark *Element)
func (l *List) MoveToBack(e *Element)
func (l *List) MoveToFront(e *Element)
func (l *List) PushBack(v any) *Element
func (l *List) PushBackList(other *List)
func (l *List) PushFront(v any) *Element
func (l *List) PushFrontList(other *List)
func (l *List) Remove(e *Element) any

## container/ring库
type Ring
func New(n int) *Ring
func (r *Ring) Do(f func(any))
func (r *Ring) Len() int
func (r *Ring) Link(s *Ring) *Ring
func (r *Ring) Move(n int) *Ring
func (r *Ring) Next() *Ring
func (r *Ring) Prev() *Ring
func (r *Ring) Unlink(n int) *Ring

## archive/zip库
func RegisterCompressor(method uint16, comp Compressor)
func RegisterDecompressor(method uint16, dcomp Decompressor)
type Compressor
type Decompressor
type File
func (f *File) DataOffset() (offset int64, err error)
func (f *File) Open() (io.ReadCloser, error)
func (f *File) OpenRaw() (io.Reader, error)
type FileHeader
func FileInfoHeader(fi fs.FileInfo) (*FileHeader, error)
func (h *FileHeader) FileInfo() fs.FileInfo
func (h *FileHeader) ModTime() time.Timedeprecated
func (h *FileHeader) Mode() (mode fs.FileMode)
func (h *FileHeader) SetModTime(t time.Time)deprecated
func (h *FileHeader) SetMode(mode fs.FileMode)
type ReadCloser
func OpenReader(name string) (*ReadCloser, error)
func (rc *ReadCloser) Close() error
type Reader
func NewReader(r io.ReaderAt, size int64) (*Reader, error)
func (r *Reader) Open(name string) (fs.File, error)
func (r *Reader) RegisterDecompressor(method uint16, dcomp Decompressor)
type Writer
func NewWriter(w io.Writer) *Writer
func (w *Writer) AddFS(fsys fs.FS) error
func (w *Writer) Close() error
func (w *Writer) Copy(f *File) error
func (w *Writer) Create(name string) (io.Writer, error)
func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)
func (w *Writer) CreateRaw(fh *FileHeader) (io.Writer, error)
func (w *Writer) Flush() error
func (w *Writer) RegisterCompressor(method uint16, comp Compressor)
func (w *Writer) SetComment(comment string) error
func (w *Writer) SetOffset(n int64)

## archive/tar库
type FileInfoNames
type Format
func (f Format) String() string
type Header
func FileInfoHeader(fi fs.FileInfo, link string) (*Header, error)
func (h *Header) FileInfo() fs.FileInfo
type Reader
func NewReader(r io.Reader) *Reader
func (tr *Reader) Next() (*Header, error)
func (tr *Reader) Read(b []byte) (int, error)
type Writer
func NewWriter(w io.Writer) *Writer
func (tw *Writer) AddFS(fsys fs.FS) error
func (tw *Writer) Close() error
func (tw *Writer) Flush() error
func (tw *Writer) Write(b []byte) (int, error)
func (tw *Writer) WriteHeader(hdr *Header) error

## compress库

## database/sql库
func Drivers() []string
func Register(name string, driver driver.Driver)
type ColumnType
func (ci *ColumnType) DatabaseTypeName() string
func (ci *ColumnType) DecimalSize() (precision, scale int64, ok bool)
func (ci *ColumnType) Length() (length int64, ok bool)
func (ci *ColumnType) Name() string
func (ci *ColumnType) Nullable() (nullable, ok bool)
func (ci *ColumnType) ScanType() reflect.Type
type Conn
func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)
func (c *Conn) Close() error
func (c *Conn) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
func (c *Conn) PingContext(ctx context.Context) error
func (c *Conn) PrepareContext(ctx context.Context, query string) (*Stmt, error)
func (c *Conn) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...any) *Row
func (c *Conn) Raw(f func(driverConn any) error) (err error)
type DB
func Open(driverName, dataSourceName string) (*DB, error)
func OpenDB(c driver.Connector) *DB
func (db *DB) Begin() (*Tx, error)
func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)
func (db *DB) Close() error
func (db *DB) Conn(ctx context.Context) (*Conn, error)
func (db *DB) Driver() driver.Driver
func (db *DB) Exec(query string, args ...any) (Result, error)
func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
func (db *DB) Ping() error
func (db *DB) PingContext(ctx context.Context) error
func (db *DB) Prepare(query string) (*Stmt, error)
func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error)
func (db *DB) Query(query string, args ...any) (*Rows, error)
func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
func (db *DB) QueryRow(query string, args ...any) *Row
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row
func (db *DB) SetConnMaxIdleTime(d time.Duration)
func (db *DB) SetConnMaxLifetime(d time.Duration)
func (db *DB) SetMaxIdleConns(n int)
func (db *DB) SetMaxOpenConns(n int)
func (db *DB) Stats() DBStats
type DBStats
type IsolationLevel
func (i IsolationLevel) String() string
type NamedArg
func Named(name string, value any) NamedArg
type Null
func (n *Null[T]) Scan(value any) error
func (n Null[T]) Value() (driver.Value, error)
type NullBool
func (n *NullBool) Scan(value any) error
func (n NullBool) Value() (driver.Value, error)
type NullByte
func (n *NullByte) Scan(value any) error
func (n NullByte) Value() (driver.Value, error)
type NullFloat64
func (n *NullFloat64) Scan(value any) error
func (n NullFloat64) Value() (driver.Value, error)
type NullInt16
func (n *NullInt16) Scan(value any) error
func (n NullInt16) Value() (driver.Value, error)
type NullInt32
func (n *NullInt32) Scan(value any) error
func (n NullInt32) Value() (driver.Value, error)
type NullInt64
func (n *NullInt64) Scan(value any) error
func (n NullInt64) Value() (driver.Value, error)
type NullString
func (ns *NullString) Scan(value any) error
func (ns NullString) Value() (driver.Value, error)
type NullTime
func (n *NullTime) Scan(value any) error
func (n NullTime) Value() (driver.Value, error)
type Out
type RawBytes
type Result
type Row
func (r *Row) Err() error
func (r *Row) Scan(dest ...any) error
type Rows
func (rs *Rows) Close() error
func (rs *Rows) ColumnTypes() ([]*ColumnType, error)
func (rs *Rows) Columns() ([]string, error)
func (rs *Rows) Err() error
func (rs *Rows) Next() bool
func (rs *Rows) NextResultSet() bool
func (rs *Rows) Scan(dest ...any) error
type Scanner
type Stmt
func (s *Stmt) Close() error
func (s *Stmt) Exec(args ...any) (Result, error)
func (s *Stmt) ExecContext(ctx context.Context, args ...any) (Result, error)
func (s *Stmt) Query(args ...any) (*Rows, error)
func (s *Stmt) QueryContext(ctx context.Context, args ...any) (*Rows, error)
func (s *Stmt) QueryRow(args ...any) *Row
func (s *Stmt) QueryRowContext(ctx context.Context, args ...any) *Row
type Tx
func (tx *Tx) Commit() error
func (tx *Tx) Exec(query string, args ...any) (Result, error)
func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
func (tx *Tx) Prepare(query string) (*Stmt, error)
func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error)
func (tx *Tx) Query(query string, args ...any) (*Rows, error)
func (tx *Tx) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
func (tx *Tx) QueryRow(query string, args ...any) *Row
func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *Row
func (tx *Tx) Rollback() error
func (tx *Tx) Stmt(stmt *Stmt) *Stmt
func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt
type TxOptions

## encoding库
type BinaryAppender
type BinaryMarshaler
type BinaryUnmarshaler
type TextAppender
type TextMarshaler
type TextUnmarshaler

## crypto库
func RegisterHash(h Hash, f func() hash.Hash)
type Decrypter
type DecrypterOpts
type Hash
func (h Hash) Available() bool
func (h Hash) HashFunc() Hash
func (h Hash) New() hash.Hash
func (h Hash) Size() int
func (h Hash) String() string
type PrivateKey
type PublicKey
type Signer
type SignerOpts

## go/ast库
func FileExports(src *File) bool
func FilterDecl(decl Decl, f Filter) bool
func FilterFile(src *File, f Filter) bool
func FilterPackage(pkg *Package, f Filter) bool
func Fprint(w io.Writer, fset *token.FileSet, x any, f FieldFilter) error
func Inspect(node Node, f func(Node) bool)
func IsExported(name string) bool
func IsGenerated(file *File) bool
func NotNilFilter(_ string, v reflect.Value) bool
func PackageExports(pkg *Package) bool
func Preorder(root Node) iter.Seq[Node]
func Print(fset *token.FileSet, x any) error
func SortImports(fset *token.FileSet, f *File)
func Walk(v Visitor, node Node)
type ArrayType
func (x *ArrayType) End() token.Pos
func (x *ArrayType) Pos() token.Pos
type AssignStmt
func (s *AssignStmt) End() token.Pos
func (s *AssignStmt) Pos() token.Pos
type BadDecl
func (d *BadDecl) End() token.Pos
func (d *BadDecl) Pos() token.Pos
type BadExpr
func (x *BadExpr) End() token.Pos
func (x *BadExpr) Pos() token.Pos
type BadStmt
func (s *BadStmt) End() token.Pos
func (s *BadStmt) Pos() token.Pos
type BasicLit
func (x *BasicLit) End() token.Pos
func (x *BasicLit) Pos() token.Pos
type BinaryExpr
func (x *BinaryExpr) End() token.Pos
func (x *BinaryExpr) Pos() token.Pos
type BlockStmt
func (s *BlockStmt) End() token.Pos
func (s *BlockStmt) Pos() token.Pos
type BranchStmt
func (s *BranchStmt) End() token.Pos
func (s *BranchStmt) Pos() token.Pos
type CallExpr
func (x *CallExpr) End() token.Pos
func (x *CallExpr) Pos() token.Pos
type CaseClause
func (s *CaseClause) End() token.Pos
func (s *CaseClause) Pos() token.Pos
type ChanDir
type ChanType
func (x *ChanType) End() token.Pos
func (x *ChanType) Pos() token.Pos
type CommClause
func (s *CommClause) End() token.Pos
func (s *CommClause) Pos() token.Pos
type Comment
func (c *Comment) End() token.Pos
func (c *Comment) Pos() token.Pos
type CommentGroup
func (g *CommentGroup) End() token.Pos
func (g *CommentGroup) Pos() token.Pos
func (g *CommentGroup) Text() string
type CommentMap
func NewCommentMap(fset *token.FileSet, node Node, comments []*CommentGroup) CommentMap
func (cmap CommentMap) Comments() []*CommentGroup
func (cmap CommentMap) Filter(node Node) CommentMap
func (cmap CommentMap) String() string
func (cmap CommentMap) Update(old, new Node) Node
type CompositeLit
func (x *CompositeLit) End() token.Pos
func (x *CompositeLit) Pos() token.Pos
type Decl
type DeclStmt
func (s *DeclStmt) End() token.Pos
func (s *DeclStmt) Pos() token.Pos
type DeferStmt
func (s *DeferStmt) End() token.Pos
func (s *DeferStmt) Pos() token.Pos
type Ellipsis
func (x *Ellipsis) End() token.Pos
func (x *Ellipsis) Pos() token.Pos
type EmptyStmt
func (s *EmptyStmt) End() token.Pos
func (s *EmptyStmt) Pos() token.Pos
type Expr
func Unparen(e Expr) Expr
type ExprStmt
func (s *ExprStmt) End() token.Pos
func (s *ExprStmt) Pos() token.Pos
type Field
func (f *Field) End() token.Pos
func (f *Field) Pos() token.Pos
type FieldFilter
type FieldList
func (f *FieldList) End() token.Pos
func (f *FieldList) NumFields() int
func (f *FieldList) Pos() token.Pos
type File
func MergePackageFiles(pkg *Package, mode MergeMode) *File
func (f *File) End() token.Pos
func (f *File) Pos() token.Pos
type Filter
type ForStmt
func (s *ForStmt) End() token.Pos
func (s *ForStmt) Pos() token.Pos
type FuncDecl
func (d *FuncDecl) End() token.Pos
func (d *FuncDecl) Pos() token.Pos
type FuncLit
func (x *FuncLit) End() token.Pos
func (x *FuncLit) Pos() token.Pos
type FuncType
func (x *FuncType) End() token.Pos
func (x *FuncType) Pos() token.Pos
type GenDecl
func (d *GenDecl) End() token.Pos
func (d *GenDecl) Pos() token.Pos
type GoStmt
func (s *GoStmt) End() token.Pos
func (s *GoStmt) Pos() token.Pos
type Ident
func NewIdent(name string) *Ident
func (x *Ident) End() token.Pos
func (id *Ident) IsExported() bool
func (x *Ident) Pos() token.Pos
func (id *Ident) String() string
type IfStmt
func (s *IfStmt) End() token.Pos
func (s *IfStmt) Pos() token.Pos
type ImportSpec
func (s *ImportSpec) End() token.Pos
func (s *ImportSpec) Pos() token.Pos
type Importerdeprecated
type IncDecStmt
func (s *IncDecStmt) End() token.Pos
func (s *IncDecStmt) Pos() token.Pos
type IndexExpr
func (x *IndexExpr) End() token.Pos
func (x *IndexExpr) Pos() token.Pos
type IndexListExpr
func (x *IndexListExpr) End() token.Pos
func (x *IndexListExpr) Pos() token.Pos
type InterfaceType
func (x *InterfaceType) End() token.Pos
func (x *InterfaceType) Pos() token.Pos
type KeyValueExpr
func (x *KeyValueExpr) End() token.Pos
func (x *KeyValueExpr) Pos() token.Pos
type LabeledStmt
func (s *LabeledStmt) End() token.Pos
func (s *LabeledStmt) Pos() token.Pos
type MapType
func (x *MapType) End() token.Pos
func (x *MapType) Pos() token.Pos
type MergeMode
type Node
type ObjKind
func (kind ObjKind) String() string
type Objectdeprecated
func NewObj(kind ObjKind, name string) *Object
func (obj *Object) Pos() token.Pos
type Packagedeprecated
func NewPackage(fset *token.FileSet, files map[string]*File, importer Importer, ...) (*Package, error)deprecated
func (p *Package) End() token.Pos
func (p *Package) Pos() token.Pos
type ParenExpr
func (x *ParenExpr) End() token.Pos
func (x *ParenExpr) Pos() token.Pos
type RangeStmt
func (s *RangeStmt) End() token.Pos
func (s *RangeStmt) Pos() token.Pos
type ReturnStmt
func (s *ReturnStmt) End() token.Pos
func (s *ReturnStmt) Pos() token.Pos
type Scopedeprecated
func NewScope(outer *Scope) *Scope
func (s *Scope) Insert(obj *Object) (alt *Object)
func (s *Scope) Lookup(name string) *Object
func (s *Scope) String() string
type SelectStmt
func (s *SelectStmt) End() token.Pos
func (s *SelectStmt) Pos() token.Pos
type SelectorExpr
func (x *SelectorExpr) End() token.Pos
func (x *SelectorExpr) Pos() token.Pos
type SendStmt
func (s *SendStmt) End() token.Pos
func (s *SendStmt) Pos() token.Pos
type SliceExpr
func (x *SliceExpr) End() token.Pos
func (x *SliceExpr) Pos() token.Pos
type Spec
type StarExpr
func (x *StarExpr) End() token.Pos
func (x *StarExpr) Pos() token.Pos
type Stmt
type StructType
func (x *StructType) End() token.Pos
func (x *StructType) Pos() token.Pos
type SwitchStmt
func (s *SwitchStmt) End() token.Pos
func (s *SwitchStmt) Pos() token.Pos
type TypeAssertExpr
func (x *TypeAssertExpr) End() token.Pos
func (x *TypeAssertExpr) Pos() token.Pos
type TypeSpec
func (s *TypeSpec) End() token.Pos
func (s *TypeSpec) Pos() token.Pos
type TypeSwitchStmt
func (s *TypeSwitchStmt) End() token.Pos
func (s *TypeSwitchStmt) Pos() token.Pos
type UnaryExpr
func (x *UnaryExpr) End() token.Pos
func (x *UnaryExpr) Pos() token.Pos
type ValueSpec
func (s *ValueSpec) End() token.Pos
func (s *ValueSpec) Pos() token.Pos
type Visitor

## go/parser库
func ParseDir(fset *token.FileSet, path string, filter func(fs.FileInfo) bool, mode Mode) (pkgs map[string]*ast.Package, first error)
func ParseExpr(x string) (ast.Expr, error)
func ParseExprFrom(fset *token.FileSet, filename string, src any, mode Mode) (expr ast.Expr, err error)
func ParseFile(fset *token.FileSet, filename string, src any, mode Mode) (f *ast.File, err error)
type Mode

## go/importer库
func Default() types.Importer
func For(compiler string, lookup Lookup) types.Importerdeprecated
func ForCompiler(fset *token.FileSet, compiler string, lookup Lookup) types.Importer
type Lookup

## go/format库
func Node(dst io.Writer, fset *token.FileSet, node any) error
func Source(src []byte) ([]byte, error)

## areana库

## runtime库
runtime包提供了一组函数和变量可以控制和监控程序的执行状态，包括Goroutine管理、内存管理和系统信息管理。

`Goexit`函数用于立即终止当前goroutine，它不会影响其它正在运行的goruntine，并且不会执行当前goroutine的defer语句：
```go
go func() {
    defer fmt.Println("This will not be printed.")
    fmt.Println("Exiting goroutine.")
    // 这里执行后不会执行上面的defer语句
    runtime.Goexit()
    fmt.Println("This will not be printed either.")
}()
```

`Gosched`函数用于让出当前goroutine的执行权，允许其它goroutine运行，它不会挂起也不会结束当前goroutine，只是简单的将其放回队列等待下次调度：
```go
// 下面的程序会使得goroutine和主程序交替输出
go func() {
    for i := 0; i < 5; i++ {
        fmt.Println("Goroutine iteration:", i)
        runtime.Gosched()
    }
}()
for i := 0; i < 5; i++ {
    fmt.Println("Main iteration:", i)
    runtime.Gosched()
```

`NumGoroutine`函数返回当前正在运行的goroutine数量：
```go
fmt.Println("Number of goroutines:", runtime.NumGoroutine())
go func() {
    fmt.Println("Number of goroutines inside goroutine:", runtime.NumGoroutine())
}()
runtime.Gosched() // 让子goroutine有机会运行
fmt.Println("Number of goroutines after launch:", runtime.NumGoroutine())
```

`MemStats`结构体用于存储内存统计信息，通过`ReadMemStats`函数可以获取当前的内存使用情况并填充到该结构体中：
```go
var memStats runtime.MemStats
runtime.ReadMemStats(&memStats)
fmt.Printf("Alloc = %v MiB\n", memStats.Alloc / 1024 / 1024)
fmt.Printf("TotalAlloc = %v MiB\n", memStats.TotalAlloc / 1024 / 1024)
fmt.Printf("Sys = %v MiB\n", memStats.Sys / 1024 / 1024)
fmt.Printf("NumGC = %v\n", memStats.NumGC)
```

`GC`函数用于触发一次垃圾回收，通常我们不需要单独来触发垃圾回收：
```go
runtime.GC()
```

`GOMAXPROCS`函数用于设置和获取可同时执行的最大CPU数：
```go
// 获取当前系统的CPU数
numCPU := runtime.NumCPU()
fmt.Printf("Number of CPUs: %d\n", numCPU)

// 设置最大可用CPU数
runtime.GOMAXPROCS(numCPU)
fmt.Printf("GOMAXPROCS set to: %d\n", runtime.GOMAXPROCS(0))
```

`Version`函数可以返回当前使用的Go版本：
```go
var version = runtime.Version()
```

runtime包提供了一些调试和诊断功能，通过`Callers`和`FuncForPC`函数来获取当前goroutine的调用栈信息，使用`NumGoroutine`函数可以获取当前正在运行的goroutine数量：
```go
func main() {
    pc := make([]uintptr, 10)
    n := runtime.Callers(0, pc)
    frames := runtime.CallersFrames(pc[:n])

    for {
        frame, more := frames.Next()
        fmt.Printf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
        if !more {
            break
        }
    }
}
```

`GOOS`属性可以当前所在的操作系统名称：
```go
var goos = runtime.GOOS
```

`GOARCH`属性可以查看都给你钱系统的架构：
```go
var goArch = runtime.GOARCH()
```

`GOROOT`函数可以获取到Go的根目录：
```go
var goRoot = runtime.GOROOT()
```

## runtime/debug库
该库主要提供了控制垃圾回收行为、获取堆栈信息、内存管理相关的操作。

`SetGCPercent`函数可以设置GC的触发频率。
```go
// 当内存的使用情况超过总内存的80%时执行GC
debug.SetGCPrecent(80)
```

`ReadGCStats`函数可以获取GC的统计信息，包括GC的次数、暂停时间等信息。
```go
var stats debug.GCStats
debug.ReadGCStats(&stats)
fmt.Printf("GC次数: %v\n", stats.NumGC)
fmt.Printf("总暂停时间: %v\n", stats.PauseTotal)
fmt.Printf("最后一次GC暂停时间: %v\n", stats.Pause[0])
```

`PrintStack`函数会将当前的堆栈信息打印到标准错误输出中：
```go
func main() {
    // 这里当触发panic错误时defer语句就会打印当前堆栈信息，方便定位错误
    defer debug.PrintStack()
    panic("something went wrong")
}
```

`Stack`获取一个包含当前堆栈信息的字节切片，我们可以灵活的处理堆栈信息：
```go
func captureStack() []byte {
    return debug.Stack()
}

func main() {
    stackInfo := captureStack()
    fmt.Printf("%s\n", stackInfo)
}
```

`FreeOSMemory`函数会触发一次垃圾回收，并尝试将未使用的内存返回给操作系统：
```go
debug.FreeOSMemory() // 手动释放内存
```

`SetMaxStack`函数可以设置一个goroutine的最大堆栈大小，如果超过了这个大小会触发panic错误：
```go
debug.SetMaxStack(10 * 1024 * 1024) // 设置最大堆栈大小为10MB
```

`SetMaxThreads`函数可以设置程序运行时的最大线程数，超过这个数目时程序会触发panic错误：
```go
debug.SetMaxThreads(100) // 设置最大线程数为100
```

`BuildInfo`类型用于存储各个模块的信息
```go
type BuildInfo struct {
	GoVersion string
	Path string
	Main Module
	Deps []*Module
	Settings []BuildSetting
}

type Module struct {
	Path    string  // 模块路径
	Version string  // 模块版本
	Sum     string  // 校验和
	Replace *Module // 子模块
}

type BuildSetting struct {
	Key, Value string
}
```

`ReadBuildInfo`函数可以读取程序的构建详情：
```go
package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("无法读取构建信息")
		return
	}

	fmt.Printf("路径: %s\n", info.Path)
	fmt.Printf("主模块版本: %s\n", info.Main.Version)
	for _, dep := range info.Deps {
		fmt.Printf("依赖项: %s@%s\n", dep.Path, dep.Version)
	}
}
```

## text/scanner库
该库实现了对源码文本的扫描器，其将[]byte作为一个源，然后通过重复的调用Scan方法来进行标记。

```go
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	const src = `
// This is scanned code.
if a > 10 {
	someParsable = text
}`
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}
}
```

下面列出了预定义的模式位
```go
const (
        ScanIdents     = 1 << -Ident
        ScanInts       = 1 << -Int
        ScanFloats     = 1 << -Float // includes Ints
        ScanChars      = 1 << -Char
        ScanStrings    = 1 << -String
        ScanRawStrings = 1 << -RawString
        ScanComments   = 1 << -Comment
        SkipComments   = 1 << -skipComment // if set with ScanComments, comments become white space
        GoTokens       = ScanIdents | ScanFloats | ScanChars | ScanStrings | ScanRawStrings | ScanComments | SkipComments
)
```

`Scan`的结果可以是下面这些标志或Unicode字符之一：
```go
const (
        EOF = -(iota + 1)
        Ident
        Int
        Float
        Char
        String
        RawString
        Comment
)
```

`GoWhitespace`属性表示扫描仪空白字段的默认值：
```go
const GoWhitespace = 1<<'\t' | 1<<'\n' | 1<<'\r' | 1<<' '
```

`TokenString`函数将标志符或Unicode字符转为可以打印的字符：
```go
var efo = scanner.TokenString(scanner.EOF)
```

`Position`类型用于表示源位置的值：
```go
type Position struct {
	Filename string // 文件名称
	Offset   int    // 当前位置到起始位置的偏移量
	Line     int    // 当前行号，初始位1
	Column   int    // 当前所在列号
}

// 查看传入的位置是否是有效的
func (pos *Position) IsValid() bool

// 将传入的位置转为字符串
func (pos Position) String() string
```

`Scanner`类型用于从指定的输入流中读取字符：
```go
type Scanner struct {

        // 遇到错误时指定的回调函数
        Error func(s *Scanner, msg string)

        // 每次遇到错误该值+1
        ErrorCount int

        // 操作的标识位
        Mode uint

        // 识别那些字符作为空白内容
        Whitespace uint64

        IsIdentRune func(ch rune, i int) bool

        // 最近扫描的起始位置信息
        Position
}

// 初始化扫描仪
func (s *Scanner) Init(src io.Reader) *Scanner

// 读取并返回下一个Unicode字符，如果读取到了末尾则返回EOF
func (s *Scanner) Next() rune

// 读取并返回下一个Unicode字符，但是不会向前推进扫描程序
func (s *Scanner) Peek() rune

// 获取最后一次调用Next或Scan时返回的字符或标记符之后的字符位置
func (s *Scanner) Pos() (pos Position)

// 读取下一个标记或Unicode字符并返回
func (s *Scanner) Scan() rune

// 返回上次扫描的标记或字符
func (s *Scanner) TokenText() string
```

## text/tabwriter库
tabwriter库实现了一个写入过滤器，它将输入中的选项卡式列转换为正确对其的文本。
```go
  // 将内容写到标准输出流中，
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t") // trailing tab
	fmt.Fprintln(w, "aaaa\tdddd\teeee")
	w.Flush()
```

tabwriter库提供了一些用于控制格式化的常量：
```go
const (
    // 忽略html标签，并将字符实体（以'&'开始，以';'结束）视为单字符
    FilterHTML uint = 1 << iota
    // 将转义后文本两端的转义字符去掉
    StripEscape
    // 强制单元格右对齐，默认是左对齐的
    AlignRight
    // 剔除空行
    DiscardEmptyColumns
    // 使用tab而不是空格进行缩进
    TabIndent
    // 在格式化后在每一列两侧添加'|'并忽略空行
    Debug
)
```

`Init`函数用于初始化写入过滤器：
```go
/**
 * minwidth表示最小单元长度
 * tabwidth表示tab字符的宽度
 * padding表示计算单元宽度时额外加上该值
 * padchar用于填充的ASCII字符，如果是"\t"则过写过滤器会假设tabwidth作为输出中tab的宽度，并且单元必然是左对齐
 * flags表示格式化的控制标识
 */
func (b *Writer) Init(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
```

`NewWriter`函数创建并初始化一个过滤器：
```go
func (b *Writer) Init(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
```

`Writer`类型标识一个过滤器，会在输入的tab划分的列进行填充，在输出中对齐它们：
```go
type Writer struct {

}

// 将传入的内容写入到过滤器中
func (b *Writer) Write(buf []byte) (n int, err error)

// 最后一次进行写操作后需要通过该方法将缓冲区的数据刷新出去
func (b *Writer) Flush() error
```

## text/template库
`text/tepmlate`库用于生成任何基于文本的格式，其通过`{{}}`来定义模板的动态部分，其余部分称为静态部分。
```go
package main

import (
	"os"
	"text/template"
)

func main() {
	// 模板字符串
	const tpl = `Hi, {{.Name}}! Welcome to {{.Website}}.`

	// 准备模板数据
	data := struct {
		Name    string
		Website string
	}{
		Name:    "Jack",
		Website: "https://xxx.com",
	}

	// 创建模板对象并解析模板字符串
	tmpl, err := template.New("test").Parse(tpl)
	if err != nil {
		panic(err)
	}

	// 使用数据渲染模板，并将结果输出到标准输出
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
```

默认情况下执行模板时，动作（Action）之间的所有文本都会逐字复制，为了帮助格式化模板源代码，模板渲染会按照下面的规则进行格式化：
- 如果操作的左分隔符后面紧跟一个负号和空白（`{{-`）则前一个文本中的所有尾部空白都会被删除。
  
- 如果右分隔符前面有空白和减号（`-}}`）则紧跟其后的文本所有前倒空白都会被删除。
  
- 如果需要做代码模板格式化，空白符号必须存在`{{- 3}}`与`{{3 -}}`，如果没有空白符号则会被解析为值，`{{-3}}则解析为包含数字`-3`。

在模板中可以定义很多动作来执行逻辑操作。
- `注释动作`：该动作会被忽略。
```go
{{/* test data *//}}
```

- `if动作`：进行条件判断，如果条件值为非0则渲染对应的内容，否则不输出对应内容。
```go
{{if true}}<p>标题一</p>{{end}}
```

- `if-else`和`if-else-if`动作：进行多分支条件判断，输出第一个符合条件的内容。
```go
{{if pipeline}}<Hello World>{{else}}<Hi World>{{end}}

{{if pipeline1}}<Hello World>{{else if pipeline2}}<Hi World>{{end}}
```
- `range`动作：当传入的表达式不为0时模板渲染后根据循环次数循环输出。
```go
{{range 5}}<Hello World>{{end}}
```

- `break`动作：可用于结束range动作中的循环。

- `continue`动作：可用于跳过range动作中的本次循环。

- `define`动作：可用于在模板中定义一个字模板。
```go
{{define "foo"}}{{end}}
```

- `template`动作可以引用已经定义好的模板。
```go
{{template "foo"}}
```

`HTMLEscape`函数

`HTMLEscapeString`

`HTMLEscaper`

`IsTrue`

`JSEscape`

`JSEscapeString`

`JSEscaper`

`URLQueryEscaper`

`FuncMap`

`New`函数用于创建一个新的Template：
```go
var login = template.New("login.yml")
```

`Must`函数是一个帮助程序，它包装对返回`（*Template， error）`的函数的调用，如果错误为非`nil`，则`panic`错误：
```go
var t = template.Must(template.New("version").Parse(versionInfoTmpl))
```

`Template`类型表示一个模板
```go
type Template struct {
	name string
	*parse.Tree
	*common
	leftDelim  string
	rightDelim string
}

// 设置模板的名称
func (t *Template) Name()

// 创建一个新的模板
func (t *Template) New(name string) *Template

// 初始化模板
func (t *Template) init()

// 将传入的字符串解析为模板
func (t *Template) Parse(text string) (*Template, error)

// 解析传入的路径中的模板
func (t *Template) ParseFiles(filenames ...string) (*Template, error)

// 当有多个模板时通过该函数指定需要解析的模板
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) 

// 将解析好的模板应用到data，并将结果写到指定位置
func (t *Template) Execute(wr io.Writer, data any)

func (t *Template) DefinedTemplates()

func (t *Template) Clone() (*Template, error)

func (t *Template) Clone() (*Template, error)

func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)

func (t *Template) Templates() []*Template

func (t *Template) Funcs(funcMap FuncMap) *Template

func (t *Template) Delims(left, right string) *Template 

func (t *Template) Lookup(name string) *Template

func (t *Template) associate(new *Template, tree *parse.Tree) bool
```








## net库
IP地址分为`IPv4`和`IPv6`，net库提供了两个常量记录IP地址的长度：
```go
const (
    IPv4len = 4
    IPv6len = 16
)
```

net库提供了一些常用的IP地址和标识：
```go
var (
    IPv4bcast     = IPv4(255, 255, 255, 255) // 广播地址
    IPv4allsys    = IPv4(224, 0, 0, 1)       // 所有主机和路由器
    IPv4allrouter = IPv4(224, 0, 0, 2)       // 所有路由器
    IPv4zero      = IPv4(0, 0, 0, 0)         // 本地地址，只能作为源地址（曾用作广播地址）
)

var (
    IPv6zero                   = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    IPv6unspecified            = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    IPv6loopback               = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
    IPv6interfacelocalallnodes = IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
    IPv6linklocalallnodes      = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
    IPv6linklocalallrouters    = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)
```

`JoinHostPort`函数将IP地址和端口号通过“:”合并在一起：
```go
net.JoinHostPort("192.168.1.22","8080")
```

`SplitHostPort`函数将传入的注解地址分割为IP地址和端口号：
```go
host, port, err := net.SplitHostPort("192.168.0.1:2232")
```

`LookupIP`函数将解析传入的域名返回对应该域名的IP地址信息：
```go
addrs, err := net.LookupIP("www.baidu.com")
```

`LookupAddr`函数对给定的地址进行反向查找，返回一个列表，该列表记录所有映射到地址的名称。
```go
addrs,err := net.LookupAddr("192.168.1.22")
```

`LookupCNAME`函数获取给定主机的规范名称：
```go
addrs, err := net.LookupCNAME("www.baidu.com")
```

`LookupPort`函数查看给定网络和服务的端口：
```go
net.LookupPort("www.baidu.com")
```

`IP`类型表示一个IP地址信息，以字节的形式存储：
```go
type IP []byte
```

`IPv4`函数可以将传入的4个字节转换为一个IP地址：
```go
ip := net.IPv4(192, 168, 100, 10)
```

`ParseIP`函数将传入的参数解析为IP地址并返回，如果传入的参数不是合法的IP地址将返回nil：
```go
func ParseIP(s string) IP
```

`DefaultMask`函数可以获取到IP实例的默认网络掩码：
```go
mask := ip.DefaultMask()
```

`Addr`接口定义类一个地址需要的基本函数：
```go
type Addr interface {
	Network() string // 获取到网络类型名称
	String() string  // 获取到地址信息
}
```

`IPAddr`：表示一个IP地址：
```go
type IPAddr struct {
    IP   IP
    Zone string // IPv6范围寻址域
}
```

`ResolveIPAddr`函数将一个域名解析为对应的IP地址：
```go
// 可选的网络类型为ip、ipv6
ip, err := net.ResolveIPAddr("ip", "www.baidu.com")
```

`Network`函数获取到IPAddr实例的网络类型：
```go
net := ip.Network()
```

`TCPAddr`表示TCP的地址信息：
```go
type TCPAddr struct {
    IP   IP
    Port int
    Zone string // IPv6范围寻址域
}
```

`UDPAddr`结构表示UDP的地址信息：
```go
type UDPAddr struct {
    IP   IP
    Port int
    Zone string // IPv6范围寻址域
}
```

`UnixAddr`结构表示Unix地址信息：
```go
type UnixAddr struct {
    Name string
    Net  string
}
```

`InterfaceAddrs`函数获取到本地的所有IP地址信息：
```go
result, err := net.InterfaceAddrs()
```

`Buffers`类型：
```go
func (v *Buffers) Read(p []byte) (n int, err error)
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error)
```

`Dial`函数用于创建一个TCP连接，它需要指定服务器的地址和端口：
```go
conn, err := net.Dial("tcp", "127.0.0.1:8080")
```

`DialTimeout`函数类似`Dial`但采用了超时机制：
```go
conn, err := net.Dial("tcp", "127.0.0.1:8080",time.Second)
```

`Dialer`
```go
type Dialer struct {
	Timeout time.Duration

	Deadline time.Time

	LocalAddr Addr

	DualStack bool

	FallbackDelay time.Duration

	KeepAlive time.Duration

	KeepAliveConfig KeepAliveConfig

	Resolver *Resolver

	Cancel <-chan struct{}

	Control func(network, address string, c syscall.RawConn) error

	ControlContext func(ctx context.Context, network, address string, c syscall.RawConn) error
}
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)
func (d *Dialer) MultipathTCP() bool
func (d *Dialer) SetMultipathTCP(use bool)
```

通过一些常量来标记接口支持的功能：
```go
const (
	FlagUp           Flags = 1 << iota // interface is administratively up
	FlagBroadcast                      // 标识接口支持广播功能
	FlagLoopback                       // 标识接口是回环接口
	FlagPointToPoint                   // 标识接口属于点对点链路
	FlagMulticast                      // 标识接口支持组播能力
	FlagRunning                        // 标识接口是否处于运行状态
)
```

`IPMask`表示一个IP地址的掩码：
```go
type IPMask []byte
```

`IPv4Mask`函数将传入的4个字节转为一个IPv4掩码：
```go
newMask := net.IPv4Mask(255, 255, 255, 0)
```

`CIDRMask`函数可以指定掩码的总长度，以及掩码中作为网段位的长度，从而计算出一个网络掩码：
```go
// 得到的结果和255.255.255.0是一样的
cidMaks := net.CIDRMask(24, 32)
```

`Size`函数和`CIDRMask`相反，用于计算传入的掩码地址的总长度，以及网段位的长度：
```go
one, len := cidMaks.Size()
```

`Pipe`函数创建一个内存中的同步、全双工网络连接，连接的两端都实现了`Conn`接口，一端的读取对应另一端的写入，直接将数据在两端之间作拷贝，没有内部缓冲。
```go
func Pipe() (Conn, Conn)
```

`Conn`接口表示通用的面向流的网络连接：
```go
type Conn interface {
    // Read从连接中读取数据
    // Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    Read(b []byte) (n int, err error)
    // Write从连接中写入数据
    // Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    Write(b []byte) (n int, err error)
    // Close方法关闭该连接
    // 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
    Close() error
    // 返回本地网络地址
    LocalAddr() Addr
    // 返回远端网络地址
    RemoteAddr() Addr
    // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
    // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
    // deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
    // 参数t为零值表示不设置期限
    SetDeadline(t time.Time) error
    // 设定该连接的读操作deadline，参数t为零值表示不设置期限
    SetReadDeadline(t time.Time) error
    // 设定该连接的写操作deadline，参数t为零值表示不设置期限
    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
    SetWriteDeadline(t time.Time) error
}
```

`PacketConn`接口代表通用的面向数据包的网络连接。
```go
type PacketConn interface {
    // ReadFrom方法从连接读取一个数据包，并将有效信息写入b
    // ReadFrom方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    // 返回写入的字节数和该数据包的来源地址
    ReadFrom(b []byte) (n int, addr Addr, err error)
    // WriteTo方法将有效数据b写入一个数据包发送给addr
    // WriteTo方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    // 在面向数据包的连接中，写入超时非常罕见
    WriteTo(b []byte, addr Addr) (n int, err error)
    // Close方法关闭该连接
    // 会导致任何阻塞中的ReadFrom或WriteTo方法不再阻塞并返回错误
    Close() error
    // 返回本地网络地址
    LocalAddr() Addr
    // 设定该连接的读写deadline
    SetDeadline(t time.Time) error
    // 设定该连接的读操作deadline，参数t为零值表示不设置期限
    // 如果时间到达deadline，读操作就会直接因超时失败返回而不会阻塞
    SetReadDeadline(t time.Time) error
    // 设定该连接的写操作deadline，参数t为零值表示不设置期限
    // 如果时间到达deadline，写操作就会直接因超时失败返回而不会阻塞
    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
    SetWriteDeadline(t time.Time) error
}
```

`Listen`函数用于在指定的IP地址和端口上监听，可用的网络类型为tcp、tcp4、tcp6、unix或unixpacket：
```go
ln, err := net.Listen("tcp", "127.0.0.1:8080")
```

`ListenPacket`函数监听本地网络地址，网络类型必须是面向数据包的网络类型ip、ip4、ip6、udp、udp4、udp6、或unixgram。
```go
ln,err := net.ListenPacket(udp,"127.0.0.1:8888")
```

`Listen`接口是用于面相流的网络协议公用的网络监听接口：
```go
type Listener interface {
    // Addr返回该接口的网络地址
    Addr() Addr
    // Accept等待并返回下一个连接到该接口的连接
    Accept() (c Conn, err error)
    // Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。
    Close() error
}
```

`FileSystem`接口用于将文件系统抽象为一个可供HTTP服务器处理的接口，我们可以让HTTP服务器直接从文件系统中读取文件并返回给客户端。
```go
type FileSystem interface {
  // 该方法将给定的文件名打开文件并返回一个File接口
	Open(name string) (File, error)
}

// 该接口定义了文件的基本操作
type File interface {
  // 关闭文件
	io.Closer
  // 从文件中读取数据到字节切片中
	io.Reader
  // 设置文件指针的位置
	io.Seeker
  // 读取目录中的文件信息
	Readdir(count int) ([]fs.FileInfo, error)
	Stat() (fs.FileInfo, error)
}
```

## net/url库
该库包含了用于URL解析、构建和查询的功能，能够处理URL并从中提取出各个部分信息。

`URL`类型表示解析到的URL：
```go
type URL struct {
    Scheme   string      //具体指访问服务器上的资源使用的哪种协议
    Opaque   string      // 编码后的不透明数据
    User     *Userinfo  // 用户名和密码信息,有些协议需要传入明文用户名和密码来获取资源，比如 FTP
    Host     string     // host或host:port，服务器地址，可以是 IP 地址，也可以是域名信息
    Path     string     //路径，使用"/"分隔
    RawPath    string   // 已编码的路径提示(参见EscapedPath方法)
	  ForceQuery bool     // 添加一个查询('?')，即使RawQuery为空
    RawQuery string     // 编码后的查询字符串，没有'?'
    Fragment string     // 引用的片段（文档位置），没有'#'
}
```

`Parse`函数可以将传入的URL解析为一个URL实例：
```go
var url = url.Parse("https://www.baidu.com")
```

`ParseRequestURI`函数解析rawurl为一个URL结构体：
```go
var urlString := "https://admin:passwd@www.baidu.com:80/search?mq=test#12345"
var u, err := url.ParseRequestURI(urlString)
```

`URL`结构提供了一些获取或设置URL相关信息的函数：
```go
// 查看URL是否是一个绝对路径
func (u *URL) IsAbs() bool

// 获取到解析到的所有查询参数
func (u *URL) Query() Values

// 获取到请求的主机名
func (u *URL) Hostname() string

// 获取到请求的端口号
func (u *URL) Port() string

```

`Userinfo`结构表示URL中用于认证的信息
```go
type Userinfo struct {
	username    string    //用户名
	password    string    // 密码
	passwordSet bool
}
```

`User`函数可以创建一个Userinfo实例并设置用户名：
```go
var userInfo = url.User("test")
```

`UserPassword`函数可以创建一个Userinfo实例并设置用户名和密码：
```go
var userInfo = url.UserPassword("test","887766")
```

`Password`函数可以获取到认证信息中的用户名：
```go
var password = userInfo.Password()
```

`Username`函数可以获取到认证信息中的密码：
```go
var user = userInfo.Username()
```

`Values`结构用于表示查询参数信息
```go
type Values map[string][]string
```

`ParseQuery`函数可以将传入的查询参数解析为一个`Values`实例：
```go
var values = url.ParseQuery("name=jack&password=1111")
```

`Values`结构还提供了一些获取和设置查询参数的函数：
```Go
// 添加一个新的查询参数
values.Add("type",0)

// 删除指定的查询参数
values.Del("type")

// 对所有的查询参数进行编码
var result = values.Encode()

// 获取到对应的查询参数的值
var type = values.Get("type")

// 查看是否设置了对应的查询参数
var isName = values.Has("name")

// 修改对应的查询参数的值
values.Set("isShow",false)
```

`JoinPath`函数会将多个路径合并为一个路径并处理`../`和`./`，将多个`//`合并为一个`/`：
```go
j1, _ := url.JoinPath("http://test.com/", "../..//findAll?tag=read")
```

`PathEscape`函数会将字符串转移出来，以便将其安全的放置在URL路径中：
```go
package main

import (
	"fmt"
	"net/url"
)
func main() {
	path := "path with?reserved+characters"
	// path%20with%3Freserved+characters
	fmt.Println(url.PathEscape(path))
}
```

`PathUnescape`对`PathEscape`的逆转换，将`%AB`转换为字节`0xAB`，如果`%`后没有十六进制数字则返回一个错误。

`QueryEscape`函数将字符串转义，以便将其安全的放置在URL查询中：
```go
package main

import (
	"fmt"
	"net/url"
)
func main() {

	query := "Hellö Wörld@Golang"

	// Hell%C3%B6+W%C3%B6rld%40Golang
	fmt.Println(url.QueryEscape(query))
}
```

`QueryUnescape`对`QueryEscape`的逆转换，将`%AB`转换为字节`0xAB`，将`+`转换为空格，如果`%`后没有十六进制数字则返回一个错误。

## net/http库
该库提供了一套完整的http请求和响应服务。

`Client`结构表示客户端，其内部主要的属性是`Transport`，它负责客户端请求发起到返回响应的实现：
```go
type Client struct {
	Transport RoundTripper  //http请求的具体实现
 
	Jar CookieJar         // cookie
 
	Timeout time.Duration   //超时
}
```

`RoundTripper`接口实现一个`RoundTrip`方法，该方法负责实现传入请求，返回请求对应的响应，是客户端请求的核心：
```go
type RoundTripper interface {
	RoundTrip(*Request) (*Response, error)
}
```

`Request`表示请求的结果，其包含了请求方法、请求体、请求地址、协议等信息：
```go
type Request struct {
  // 请求的地址信息
	URL *url.URL
  // 使用的协议信息
	Proto      string
  // 协议主版本号
	ProtoMajor int
  // 协议子版本号
	ProtoMinor int

  // 请求头信息
	Header Header

  // 请求体信息
	Body io.ReadCloser

  // 获取到请求体的函数
	GetBody func() (io.ReadCloser, error)

  // 正文长度
	ContentLength int64

  // 编码格式
	TransferEncoding []string

	Close bool

  // 服务器主机
	Host string

    // 表单信息
	Form url.Values

	PostForm url.Values

    // 上传文件表单信息
	MultipartForm *multipart.Form

	Trailer Header

  // 服务器地址
	RemoteAddr string

  // 请求的URI
	RequestURI string

	TLS *tls.ConnectionState

	Cancel <-chan struct{}

  // 响应体
	Response *Response

	Pattern string

	ctx context.Context

	pat         *pattern          // the pattern that matched
	matches     []string          // values for the matching wildcards in pat
	otherValues map[string]string // for calls to SetPathValue that don't match a wildcard
}
```

`Response`表示响应体结构，包含请求状态、协议、响应头、响应体、请求体等信息：
```go
type Response struct {
  // 响应状态
	Status     string // e.g. "200 OK"
  // 响应状态码
	StatusCode int    // e.g. 200
  // 使用的协议信息
	Proto      string
  // 协议主版本号
	ProtoMajor int
  // 协议子版本号
	ProtoMinor int
  // 响应头
	Header Header
  // 响应体
	Body io.ReadCloser
  // 正文长度
	ContentLength int64
  // 编码格式
	TransferEncoding []string
	Close bool

	Uncompressed bool

	Trailer Header
  // 对应的请求体
	Request *Request

	TLS *tls.ConnectionState
}
```

`Handler`接口定义了ServeHTTP方法，我们可以直接定义一个类型并实现该方法来处理请求：
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// 该类型实现了Handler接口
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

`HandleFunc`函数可以注册一个请求路径对应的处理函数：
```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
})
```

`HandlerFunc`函数可以将普通的函数转为实现了`Handler`接口的处理器对象：
```go
func demo(h http.Handler) {}

func handler(w http.ResponseWriter, r *http.Request) {}

func main() {
    // 这个方法实际上就是对handler函数的类型转换将其转为Handler类型
    demo(http.Handler(handler))
}
```

`ListenAndServe`函数可以开启一个HTTP服务端：
```go
http.ListenAndServe(":8080", nil)
```

通过`Get`函数可以向服务端发起get请求：
```go
resp, err := http.Get("http://example.com/")
if err != nil {
    fmt.Println(err)
    return
}
defer resp.Body.Close()
body, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(body))
```

`Server`结构表示服务端
```go
type Server struct {

	Addr string

	Handler Handler // handler to invoke, http.DefaultServeMux if nil

	DisableGeneralOptionsHandler bool

	TLSConfig *tls.Config

	ReadTimeout time.Duration

	ReadHeaderTimeout time.Duration

	WriteTimeout time.Duration

	IdleTimeout time.Duration

	MaxHeaderBytes int

	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

	ConnState func(net.Conn, ConnState)

	ErrorLog *log.Logger

	BaseContext func(net.Listener) context.Context

	ConnContext func(ctx context.Context, c net.Conn) context.Context
	HTTP2 *HTTP2Config

	Protocols *Protocols

	inShutdown atomic.Bool // true when server is in shutdown

	disableKeepAlives atomic.Bool
	nextProtoOnce     sync.Once // guards setupHTTP2_* init
	nextProtoErr      error     // result of http2.ConfigureServer if used

	mu         sync.Mutex
	listeners  map[*net.Listener]struct{}
	activeConn map[*conn]struct{}
	onShutdown []func()

	listenerGroup sync.WaitGroup
}

// http请求的多路复用器
type ServeMux struct {
	mu     sync.RWMutex
	tree   routingNode
	index  routingIndex
	mux121 serveMux121 // used only when GODEBUG=httpmuxgo121=1
}
```

`FileServer`函数根据传入的文件路径构建一个`Handler`实例返回：
```go
http.FileServer(http.Dir("/home/test/"))
```

## context库
context库提供了一种在函数之间传递请求作用域的方法，通常用于跨API边界换地取消信号、超时值、截止时间以及请求范围的数据。

`Context`接口用于管理请求作用域的重要接口：
```go
type Context interface {
  // 返回上下文的截止时间，如果没有设置截止时间则ok为false
	Deadline() (deadline time.Time, ok bool)

  // 返回一个通道Channel，当上下文被取消或超时的时候该通道会被关闭，可以通过监听这个通道来接收取消信号
	Done() <-chan struct{}

  // 返回上下文取消的原因，如果上下文没有被取消则返回nil
	Err() error

  // 返回和键管理的值，如果键不存在则返回nil
	Value(key any) any
}
```

`Background`函数用于创建一个默认的上下文，其它的上下文都应该从它衍生出来：
```go

```

`TODO`函数用于创建一个不确定类型的上下文：
```go
```

`WithCancel`函数用于创建一个带有取消功能的上下文，当调用其返回的CancelFunc时所有通过该上下文启动的goroutine都会收到一个取消信号：
```go
// 模拟一个长时间运行的任务
func longRunningTask(ctx context.Context, name string) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done(): // 检查上下文是否被取消
			return // 如果被取消，立即返回
		case <-time.After(1 * time.Second): // 每秒输出一次进度
			fmt.Printf("%s: %d
", name, i)
		}
	}
}

func main() {
	// 创建一个新的带取消功能的上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 使用新创建的上下文启动任务
	go longRunningTask(ctx, "task")

	// 主线程等待5秒后取消上下文
	time.Sleep(5 * time.Second)
	cancel() // 发送取消信号

	// 为了演示，我们再等2秒钟以确保任务已经接收到了取消信号并且退出
	time.Sleep(2 * time.Second)
}
```

`WithDeadline`函数用于创建一个带有截止时间的上下文，这个截止时间可以用来限制操作的执行时间，当超过截止时间时可以自动取消上下文：
```go
func main() {
	parentCtx := context.Background()
	deadline := time.Now().Add(time.Millisecond * 50)
	// 创建一个带有截止时间的新context，设置截止时间为当前时间的50毫秒之后
	ctx, cancel := context.WithDeadline(parentCtx, deadline)
	// 确保在main函数返回前取消context以释放资源
	defer cancel()

	// 模拟一个需要执行的操作，比如等待100毫秒
	go func() {
		select {
		case <-time.After(100 * time.Millisecond):
			fmt.Println("操作完成")
		case <-ctx.Done():
			fmt.Println("操作超时")
		}
	}()

	// 等待一段时间，以便观察操作是否在截止时间之前完成
	time.Sleep(time.Millisecond * 200)
}
```

`WithTimeout`函数用于创建一种可以设置操作执行的超时时间的上下文，允许我们指定一个时间段，在这个时间段内如果操作没有完成机会抛出一个超时异常或执行一个指定的回调函数：
```go
func myFunction(ctx context.Context, duration time.Duration) {
	select {
	case <-time.After(duration):
		fmt.Println("操作成功")
	case <- ctx.Done():
		fmt.Println("操作失败")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	myFunction(ctx, 2*time.Second)
}
```

`WithValue`函数用于创建一个带有额外值的新上下文，这个值可以再上下文中进行传递：
```go
// 定义一个自定义类型作为键
type contextKey string

func processRequest(ctx context.Context) {
	// 从上下文中获取请求ID
	reqID := ctx.Value(contextKey("requestID")).(int)

	// 模拟处理请求的操作
	fmt.Printf("Processing request with ID %d\n", reqID)
}

func main() {
	// 创建一个带有请求ID的上下文
	ctx := context.WithValue(context.Background(), contextKey("requestID"), 12345)

	// 处理请求
	processRequest(ctx)
}
```















## sync库
该库是Go提供的并发同步包，提供了一些列工具来处理并发过程中的同步问题。

`Mutex`结构表示互斥锁，用于保护共享资源，保证同一时刻只有一个goroutine能访问邻接代码：
```go
package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()   // 获取锁
		counter++      // 临界区：访问共享变量
		mutex.Unlock() // 释放锁
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	fmt.Println("Counter:", counter) // 输出：Counter: 5000
}
```

`RWMutex`结构表示读写锁，区分读和写的操作，多个goroutine可以同时进行读取，但是一旦存在写操作其它goroutine就会阻塞，直到执行写操作的goroutine执行完毕：
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var data int
var rwMutex sync.RWMutex

func readData(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.RLock()
	fmt.Printf("Reader %d: Read data %d\n", id, data)
	time.Sleep(time.Millisecond * 10)
	rwMutex.RUnlock()
}

func writeData(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.Lock()
	data += 1
	fmt.Printf("Writer %d: Wrote data %d\n", id, data)
	time.Sleep(time.Millisecond * 10)
	rwMutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go writeData(i, &wg) // 写操作
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go readData(i, &wg) // 读操作
	}

	wg.Wait()
}
```

`WaitGroup`结构表示等待组，用于等待一组goroutine完成：
```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
  // 完成一个goroutine计数
	defer wg.Done()
	fmt.Printf("Worker %d is working\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
    // 增加一个等待的goroutine计数
		wg.Add(1)
		go worker(i, &wg)
	}

  // 阻塞直到计数为0
	wg.Wait()
	fmt.Println("All workers finished")
}
```

`Once`结构表示执行一次，确保某个操作只执行一次：
```go
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initializing...")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(initialize) // 确保initialize只执行一次
	fmt.Printf("Worker %d is working\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
```

`Cond`结构表示条件变量，在goroutine之间通过条件变量进行信号通信：
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var ready bool
// 创建一个新的条件变量
var cond = sync.NewCond(&sync.Mutex{})

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	for !ready {
		cond.Wait() // 等待条件满足
	}
	cond.L.Unlock()
	fmt.Printf("Worker %d is working\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	time.Sleep(time.Second)
	cond.L.Lock()
	ready = true
	cond.Broadcast() // 唤醒所有等待的 goroutine
	cond.L.Unlock()

  // 等待条件满足
	wg.Wait()
}
```

`Pool`结构表示对象复用池，用于缓存临时对象，减少内存分配的次数来提提高性能：
```go
package main

import (
	"fmt"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return "new object"
	},
}

func main() {
  // 获取一个对象，如果存了多个对象无法确定会获取到那个对象
	obj := pool.Get()
	fmt.Println(obj) // 输出：new object

  // 放回对象
	pool.Put("reused object")
	fmt.Println(pool.Get()) // 输出：reused object
}
```

`Map`结构是Go 1.9新增的，用于对携程环境下读写共享数据，适用于读取多、写入少的场景：
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// 存储值
	m.Store("key1", "value1")
	m.Store("key2", "value2")

	// 加载值
	if val, ok := m.Load("key1"); ok {
		fmt.Println("key1:", val) // 输出: key1: value1
	}

	// 遍历所有键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

	// 删除键
	m.Delete("key1")
	if _, ok := m.Load("key1"); !ok {
		fmt.Println("key1 已删除")
	}
}

```

## sync/atomic
该模块提供了一组底层的原子操作，用于对整数、指针和其他变量进行原子级别的读写操作，下面列出了常用的操作：
- `AddInt32/AddInt64`：对整数执行加法操作。

- `LoadInt32/LoadInt64`：读取整数的值。

- `StoreInt32/StoreInt64`：设置整数的值。

- `CompareAndSwapInt32`：比较并交换值，CAS操作。

- `SwapInt32/SwapInt64`：交换值。