Description
pflag是Go的flag包的直接替代品，实现了POSIX/GNU风格的flags。
PFlag 与针对命令行选项的 POSIX 建议的 GNU 扩展兼容。有关更精确的说明，请参见下面的“Command-line flag syntax命令行标志语法”部分。
PFlag 与 Go 语言采用相同的 BSD 许可风格，可以在 LICENSE 文件中找到。


Installation
pflag可以使用标准的go get命令。
Install by running通过运行安装：
	go get github.com/spf13/pflag
Run tests by running运行测试:
	go test github.com/spf13/pflag


Usage用法：
pflag是Go的本地flag包的一个替代品。如果您以“flag”的名称导入pflag，那么所有代码都应该继续运行，没有任何更改。
	import flag "github.com/spf13/pflag"
有一个例外：如果您直接实例化Flag结构，则需要设置另一个字段“Shorthand”。大多数代码从未直接实例化此结构，而是使用String（）、BoolVar（）和Var（）等函数，因此不受影响。

定义flag使用flag.String(), Bool(), Int(), etc.
这声明了一个整数标志-flagname，存储在指针ip中，类型为*int。
	var ip *int = flag.Int("flagname", 1234, "help message for flagname")

如果愿意，可以使用Var（）函数将标志绑定到变量。
	var flagvar int
	func init() {
		flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	}

或者，您可以创建满足Value接口（带指针接收器）的自定义标志，并通过
	flag.Var(&flagVal, "name", "help message for flagname")
对于此类标志，默认值只是变量的初始值。

定义所有标志后，调用，将命令行解析为定义的标志。
	flag.Parse()

然后可以直接使用标志。如果你使用的是标志本身，它们都是指针；如果绑定到变量，它们就是值。
	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)

如果您有FlagSet，但发现很难得到代码中的所有指针，则可以使用帮助函数获取存储在Flag中的值。如果你有pflag.FlagSet带有一个名为“flagname”的int
类型的标志，您可以使用GetInt（）获取int值。但请注意，“flagname”必须存在，并且必须是int类型。GetString（“flagname”）将失败。
	i, err := flagset.GetInt("flagname")
解析后，flag后面的arg可用作[]flag.Args()或单独作为flag.Arg(i)。参数从0到flag.NArg() - 1进行索引。

pflag包还定义了一些不在flag中的新函数，这些函数为flag提供了一个字母的缩写。您可以通过在定义标志的任何函数的名称后面加上“P”来使用它们。
	var ip = flag.IntP("flagname", "f", 1234, "help message")
	var flagvar bool
	func init() {
		flag.BoolVarP(&flagvar, "boolname", "b", true, "help message")
	}
	flag.VarP(&flagVal, "varname", "v", "help message")
速记字母可以在命令行上与单破折号(-)一起使用。布尔速记标志可以与其他速记标志组合。
默认的命令行标志集由top-level函数控制。FlagSet类型允许定义独立的标志集，例如在命令行界面中实现子命令。FlagSet的方法类似于命令行标志集的top-level函数。




Setting no option default values for flags 未设置标志的选项默认值
创建一个标志之后，就可以为给定的标志设置 pFlag.NoOptDefVal。这样做会稍微改变标志的含义。如果标志具有NoOptDefVal，并且在命令行上设置了该标
志而没有选项，该标志将被设置为NoOptDefVal。例如:
	var ip = flag.IntP("flagname", "f", 1234, "help message")
	flag.Lookup("flagname").NoOptDefVal = "4321"
会导致如下结果
	Parsed Arguments	Resulting Value
	--flagname=1357		ip=1357
	--flagname			ip=4321
	[nothing]			ip=1234



Command line flag syntax	命令行标志语法
--flag    // boolean flags, or flags with no option default values	//布尔标志，或没有选项默认值的标志
--flag x  // only on flags without a default value					//仅在没有默认值的标志上
--flag=x
与flag包不同，选项前的单破折号(-)意味着与双破折号(--)不同的东西。单破折号表示标志的一系列速记字母。除最后一个速记字母外，所有字母都必须是布尔标志或具有默认值的标志
// boolean or flags where the 'no option default value' is set		//布尔值或标志，其中设置了“无选项默认值”
-f
-f=true
-abc		// but -b true is INVALID	// 但 -b true 是无效的


// non-boolean and flags without a 'no option default value'		//非布尔值和不带“no option default value”的标志
-n 1234
-n=1234
-n1234

// mixed
-abcs "hello"
-absd="hello"
-abcs1234
标志解析在终止符“--”之后停止。与flag包不同，标志可以在命令行中此终止符之前的任何位置穿插参数。
整数标志接受1234、0664、0x1234，并且可以是负数。布尔标志（in their long form）接受1, 0, t, f, true, false, TRUE, FALSE, True, False。
Duration标志接受对time.ParseDuration有效的任何输入。




Mutating or "Normalizing" Flag names	变异或“规范化”标志名称变异或“规范化”标志名称
可以设置自定义标志名“normalization function”它允许标记名在代码中创建时和在命令行上使用时都可以变异为某种“标准化”形式。“标准化的”形式用于比
较。下面是使用自定义规范化函数的两个示例。
Example #1: You want -, _, and . in flags to compare the same. aka --my-flag == --my_flag == --my.flag
示例#1：您希望-、_和.标志进行比较。aka --my-flag == --my_flag == --my.flag
func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}
myFlagSet.SetNormalizeFunc(wordSepNormalizeFunc)

Example #2: You want to alias two flags. aka --old-flag-name == --new-flag-name
示例#2：您想给两个标志取别名。aka--旧标志名==--新标志名
func aliasNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "old-flag-name":
		name = "new-flag-name"
		break
	}
	return pflag.NormalizedName(name)
}
myFlagSet.SetNormalizeFunc(aliasNormalizeFunc)




Deprecating a flag or its shorthand
可以弃用一个flag，或它的简写。弃用flag/shorthand将其隐藏在帮助文本中，并在使用弃用标志或速记时打印使用信息。
Example #1: You want to deprecate a flag named "badflag" as well as inform the users what flag they should use instead.
示例#1：您希望弃用名为“badflag”的标志，并告知用户应该使用什么标志。
// deprecate a flag by specifying its name and a usage message	//通过指定标志的名称和用法消息来否决标志
	flags.MarkDeprecated("badflag", "please use --good-flag instead")
这将从帮助文本中隐藏“badflag”，当使用"badflag"时打印"Flag --badflag has been deprecated, please use --good-flag instead" 。

Example #2: You want to keep a flag name "noshorthandflag" but deprecate its shortname "n".
示例#2：您希望保留一个标志名“noshorthandflag”，但不建议使用其简称“n”。
// deprecate a flag shorthand by specifying its flag name and a usage message	//通过指定标志名称和用法消息来反对标志的简写
flags.MarkShorthandDeprecated("noshorthandflag", "please use --noshorthandflag only")
这将从帮助文本中隐藏缩写“n”，当使用"n"时打印"Flag shorthand -n has been deprecated, please use --noshorthandflag only"
请注意，这里的用法信息非常重要，不应为空。




Hidden flags
可以将标志标记为隐藏，这意味着它仍将正常工作，但不会显示在用法/帮助文本中。
Example: You have a flag named "secretFlag" that you need for internal use only and don't want it showing up in help text, or for its usage text to be available.
示例：您有一个名为“secretFlag”的标志，您只需要在内部使用，不希望它出现在帮助文本中，也不希望它的用法文本可用。
// hide a flag by specifying its name	//通过指定标志的名称隐藏标志
flags.MarkHidden("secretFlag")




Disable sorting of flags
pflag允许您禁用help和usag消息的标志排序。
Example:

flags.BoolP("verbose", "v", false, "verbose output")
flags.String("coolflag", "yeaah", "it's really cool flag")
flags.Int("usefulflag", 777, "sometimes it's very useful")
flags.SortFlags = false
flags.PrintDefaults()
Output:

-v, --verbose           verbose output
--coolflag string   it's really cool flag (default "yeaah")
--usefulflag int    sometimes it's very useful (default 777)




Supporting Go flags when using pflag
为了支持使用Go的标志包定义的标志，必须将它们添加到pflag标志集中。这对于支持第三方依赖项（例如golang/glog）定义的标志通常是必要的。
Example: You want to add the Go flags to the CommandLine flagset 示例：您希望将Go标志添加到CommandLine标志集

import (
	goflag "flag"
	flag "github.com/spf13/pflag"
)
var ip *int = flag.Int("flagname", 1234, "help message for flagname")
func main() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()
}


More info
您可以在godoc上查看pflag包的完整参考文档godoc.org，或者通过go的标准文档系统运行godoc-http=:6060并浏览
http://localhost:6060/pkg/github.com/spf13/pflag安装后。