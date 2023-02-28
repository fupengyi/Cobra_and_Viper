Viper v2 feedback
Viper正在向v2进发，我们很想听听你想在其中看到什么。在此处分享您的想法：https://forms.gle/R6faU74qPRPAzchZ9
Go configuration with fangs!  用fangs配置！
许多Go项目使用Viper构建，包括：Hugo, EMC RexRay, Imgur’s Incus, Nanobox/Nanopack, Docker Notary, BloomApi, doctl, Clairctl, Mercure
Install:
go get github.com/spf13/viper	注意：Viper使用Go模块来管理依赖关系。


What is Viper?
Viper是Go应用程序的完整配置解决方案，包括12个Factor应用程序。它被设计为在应用程序中工作，并且可以处理所有类型的配置需求和格式。它支持：
1.setting defaults设置默认值,
2.reading from JSON, TOML, YAML, HCL, envfile and Java properties config files读取JSON、TOML、YAML、HCL、envfile和Java属性配置文件
3.live watching and re-reading of config files (optional)实时观看和重新读取配置文件（可选）
4.reading from environment variables读取环境变量
5.reading from remote config systems (etcd or Consul), and watching changes从远程配置系统（etcd或Consul）读取并查看更改
6.reading from command line flags读取命令行标志
7.reading from buffer从缓冲区读取
8.setting explicit values设置显式值
Viper可以被认为是您所有应用程序配置需求的注册中心。


Why Viper?
在构建现代应用程序时，您不必担心配置文件格式；你想专注于构建很棒的软件。Viper是来帮忙的。
Viper为您提供以下服务：
1.查找、加载和解组JSON、TOML、YAML、HCL、INI、envfile或Java属性格式的配置文件。
2.提供设置不同配置选项的默认值的机制。
3.提供一种机制，为通过命令行标志指定的选项设置重写值。
4.提供别名系统，以便在不破坏现有代码的情况下轻松重命名参数。
5.很容易区分用户何时提供了与默认值相同的命令行或配置文件。

Viper使用以下优先顺序。每个项目优先于其下面的项目：
1.explicit call to Set显式调用Set		2.flag标志	3.env环境	4.config配置		5.key/value store存储	6.default默认值
重要提示：Viper配置密钥不区分大小写。目前正在讨论是否将其设为可选。


Putting Values into Viper	向 Viper 注入 Values
1.建立默认值
良好的配置系统将支持默认值。密钥不需要默认值，但如果未通过配置文件、环境变量、远程配置或标志设置密钥，则该值很有用。例如：
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

2.Reading Config Files正在读取配置文件
Viper需要最少的配置，因此它知道在哪里查找配置文件。Viper支持JSON、TOML、YAML、HCL、INI、envfile和Java属性文件。Viper可以搜索多个路径，但目
前单个Viper实例仅支持单个配置文件。Viper不默认任何配置搜索路径，将默认决定留给应用程序。
下面是一个如何使用Viper搜索和读取配置文件的示例。不需要任何特定路径，但应在需要配置文件的地方至少提供一个路径。
viper.SetConfigName("config") 			// 配置文件的名称（无扩展名）
viper.SetConfigType("yaml") 			// 如果配置文件的名称中没有扩展名，则需要
viper.AddConfigPath("/etc/appname/")   	// 查找配置文件的路径
viper.AddConfigPath("$HOME/.appname")  	// 多次调用以添加多个搜索路径
viper.AddConfigPath(".")                // 可以在工作目录中查找配置
err := viper.ReadInConfig() 			// 查找并读取配置文件
if err != nil { 						// 处理读取配置文件的错误
	panic(fmt.Errorf("fatal error config file: %w", err))
}
您可以这样处理未找到配置文件的特定情况：
if err := viper.ReadInConfig(); err != nil {
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config file not found; ignore error if desired		//未找到配置文件；如果需要，忽略错误
	} else {
		// Config file was found but another error was produced	//找到了配置文件，但产生了另一个错误
	}
}
// Config file found and successfully parsed					//找到配置文件并成功分析
注[自1.6]：您也可以有一个没有扩展名的文件，并以编程方式指定格式。对于那些位于用户家中且没有任何扩展名（如.bashrc）的配置文件

3.Writing Config Files正在写入配置文件
读取配置文件很有用，但有时您希望存储运行时所做的所有修改。为此，可以使用一系列命令，每个命令都有自己的目的：
	1.WriteConfig-将当前viper配置写入预定义路径（如果存在）。没有预定义路径时出错。将覆盖当前配置文件（如果存在）。
	2.SafeWriteConfig-将当前viper配置写入预定义路径。没有预定义路径时出错。不会覆盖当前配置文件（如果存在）。
	3.WriteConfigAs-将当前viper配置写入给定的文件路径。将覆盖给定文件（如果存在）。
	4.SafeWriteConfigAs-将当前viper配置写入给定的文件路径。不会覆盖给定文件（如果存在）。
根据经验，所有标记为safe的文件都不会覆盖任何文件，而是在不存在的情况下创建，而默认行为是创建或截断。一个小示例部分：
viper.WriteConfig() // 将当前配置写入由“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") // 将出错，因为它已被写入
viper.SafeWriteConfigAs("/path/to/my/.other_config")

4.Watching and re-reading config files监视和重新读取配置文件
Viper 支持让应用程序在运行时实时读取配置文件的能力。
需要重新启动服务器才能使配置生效的日子已经一去不复返了，使用 viper 的应用程序可以在运行时读取配置文件的更新，而且不会错过任何一次更新。
只需告诉 viper 实例监视 Config。您还可以为 Viper 提供一个函数，以便在每次发生更改时运行该函数。
确保在调用 WatchConfig ()之前添加了所有的 configPath
viper.OnConfigChange(func(e fsnotify.Event) {
	fmt.Println("Config file changed:", e.Name)
})
viper.WatchConfig()

5.Reading Config from io.Reader从 io.Reader 读取配置
Viper 预定义了许多配置源，例如文件、环境变量、标志和远程 K/V 存储，但是您并不绑定到它们。您还可以实现自己所需的配置源，并将其提供给 viper。
viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
// any approach to require this configuration into your program.要求在程序中使用此配置的任何方法。
var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)
viper.ReadConfig(bytes.NewBuffer(yamlExample))
viper.Get("name") // this would be "steve"		name应该是“ steve”

6.Setting Overrides设置覆盖
它们可以来自命令行标志，也可以来自您自己的应用程序逻辑。
viper.Set("Verbose", true)
viper.Set("LogFile", LogFile)

7.Registering and Using Aliases注册和使用别名
别名允许多个键引用单个值
viper.RegisterAlias("loud", "Verbose")
viper.Set("verbose", true) // same result as next line结果与下一行相同
viper.Set("loud", true)   // same result as prior line结果与上一行相同
viper.GetBool("loud") // true
viper.GetBool("verbose") // true

8.Working with Environment Variables使用环境变量
Viper 完全支持环境变量。这使12因素应用程序开箱即用。有五种方法可以帮助 ENV 的工作:
	1.AutomaticEnv()										//自动检查是否有名称与之匹配的环境变量
	2.BindEnv(string...) : error							//第一个参数是键名，其余的是要绑定到这个键的环境变量的名称
	3.SetEnvPrefix(string)									//设置前缀并且自动大写,Viper 提供了一种机制来尝试确保 ENV 变量是唯一的
	4.SetEnvKeyReplacer(string...) *strings.Replacer		//重写Env键
	5.AllowEmptyEnv(bool)									//默认情况下，空环境变量被认为是未设置的，并将回退到下一个配置源。若要将空环境变量视为设置，请使用 AllowEmptyEnv 方法。

当使用 ENV 变量时，重要的是要认识到 Viper 将 ENV 变量视为区分大小写的。
Viper 提供了一种机制来尝试确保 ENV 变量是唯一的。通过使用 SetEnvPrefix(string)，可以告诉 Viper 在读取环境变量时使用前缀。BindEnv 和 AutomaticEnv
都将使用此前缀。
	BindEnv 接受一个或多个参数。第一个参数是键名，其余的是要绑定到这个键的环境变量的名称。如果提供了多个，它们将按照指定的顺序优先。环境变量的名字
是区分大小写的。如果没有提供 ENV 变量名，那么 Viper 将自动假定 ENV 变量匹配以下格式: 前缀 + “_”+ ALL CAPS 中的键名。当显式提供 ENV变量名(第
二个参数)时，它不会自动添加前缀。例如，如果第二个参数是“ ID”，Viper 将查找 ENV 变量“ ID”。
使用 ENV 变量时需要注意的一点是，每次访问该值时都会读取它。当调用 BindEnv 时，Viper 不会修复该值。
	AutomaticEnv 是一个强大的助手，特别是当与 SetEnvPrefix 组合时。当召唤时，viper.Get会随时检查是否有环境变量。它将适用以下规则。如果
设置了大写和前缀为 EnvPrefix 的键，它将检查是否有名称与之匹配的环境变量。
	SetEnvKeyReplace 允许您使用strings.Replacer对象重写一定范围内的 Env 键。如果您希望在 Get ()调用中使用-或其他内容，但希望环境变量使
用_分隔符，那么这很有用。在 viper_test.go中可以找到使用它的示例。
	或者，您可以使用带有 NewWithOptions 工厂函数的 EnvKeyReplace。与 SetEnvKeyReplace 不同，它接受 StringReplace 接口，允许您编写自定
义字符串替换逻辑。
	默认情况下，空环境变量被认为是未设置的，并将回退到下一个配置源。若要将空环境变量视为设置，请使用 AllowEmptyEnv 方法。
Env example:
SetEnvPrefix("spf") // will be uppercased automatically 将自动大写
BindEnv("id")
os.Setenv("SPF_ID", "13") // typically done outside of the app 通常在应用程序之外完成
id := Get("id") // 13

9.Working with Flags
Viper 具有绑定到标志的能力。具体来说，Viper 支持在 Cobra 库中使用的标志。
与 BindEnv 一样，该值不是在调用绑定方法时设置的，而是在访问该方法时设置的。这意味着您可以尽早进行绑定，即使在 init ()函数中也是如此。
对于单个标志，BindPFlag ()方法提供了此功能。例如：
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
您还可以绑定一组现有的 pFlag (pFlag. FlagSet) 例如：
pflag.Int("flagname", 1234, "help message for flagname")
pflag.Parse()
viper.BindPFlags(pflag.CommandLine)
i := viper.GetInt("flagname") // retrieve values from viper instead of pflag	从 viper 而不是 pFlag 中检索值
在 Viper 中使用 pFlag 并不排除使用标准库中使用标志包的其他包。PFlag 包可以通过导入这些标志来处理为标志包定义的标志。这是通过调用一个名为
AddGoFlagSet ()的 pFlag 包提供的便利函数来实现的。例如：
package main
import (
	"flag"
	"github.com/spf13/pflag"
)
func main() {
	// using standard library "flag" package	使用标准库“flag”包
	flag.Int("flagname", 1234, "help message for flagname")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	i := viper.GetInt("flagname") // retrieve value from viper	从viper获取value
	// ...
}
Flag interfaces
Viper 提供了两个 Go 接口来绑定其他标志系统，如果你不使用Pflags的话。
FlagValue 表示一个标志，这是一个关于如何实现这个接口的非常简单的例子:
type myFlag struct {}
func (f myFlag) HasChanged() bool { return false }
func (f myFlag) Name() string { return "my-flag-name" }
func (f myFlag) ValueString() string { return "my-flag-value" }
func (f myFlag) ValueType() string { return "string" }
一旦你的标志实现了这个接口，你可以简单地告诉 Viper 绑定它:
viper.BindFlagValue("my-flag-name", myFlag{})
FlagValueSet表示一组标志。这是一个关于如何实现此接口的非常简单的示例：
type myFlagSet struct {
	flags []myFlag
}
func (f myFlagSet) VisitAll(fn func(FlagValue)) {
	for _, flag := range flags {
		fn(flag)
	}
}
一旦您的标志集实现了这个接口，您可以简单地告诉 Viper 绑定它:
fSet := myFlagSet{
	flags: []myFlag{myFlag{}, myFlag{}},
}
viper.BindFlagValues("my-flags", fSet)

10.Remote Key/Value Store Support 远程键/值存储支持
要在 Viper 中启用远程支持，请对 Viper/remote 包执行空白导入:	import _ "github.com/spf13/viper/remote"
Viper 将读取从 Key/Value 存储(如 etcd 或 Consule)中的路径检索到的配置字符串(如 JSON、 TOML、 YAML、 HCL 或 envfile)。这些值优先于
默认值，但是被从磁盘、标志或环境变量检索到的配置值覆盖。
Viper 使用 crypt 从 K/V 存储中检索配置，这意味着您可以存储加密的配置值，如果您有正确的 gpg 密钥环，则可以自动解密它们。加密是可选的。
您可以将远程配置与本地配置结合使用，也可以独立于本地配置使用。
Crypt 有一个命令行助手，您可以使用它在 K/V 存储中放置配置。在 http://127.0.0.1:4001上，crypt 默认为 etcd。
$ go get github.com/bketelsen/crypt/bin/crypt
$ crypt set -plaintext /config/hugo.json /Users/hugo/settings/config.json
确认你的value已设定:$ crypt get -plaintext /config/hugo.json
有关如何设置加密值或如何使用 Consull 的示例，请参见 crypt 文档。

11.Remote Key/Value Store Example - Unencrypted 远程密钥/值存储示例-未加密
viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，所以支持的扩展名是“ json”、“ toml”、“ yaml”、“ yml”、“ properties”、“ props”、“ prom”、“ env”、“ dotenv”
err := viper.ReadRemoteConfig()
viper.AddRemoteProvider("etcd3", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，所以支持的扩展名是“ json”、“ toml”、“ yaml”、“ yml”、“ properties”、“ props”、“ prom”、“ env”、“ dotenv”
err := viper.ReadRemoteConfig()
Consul您需要使用包含所需配置的 JSON 值将一个键设置为 Consul key/value 存储。例如，创建一个 Consul key/value store键 MY _ CONSUL _ KEY，其值为:
{
	"port": 8080,
	"hostname": "myhostname.com"
}
viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
viper.SetConfigType("json") // Need to explicitly set this to json	需要将其显式设置为 json
err := viper.ReadRemoteConfig()
fmt.Println(viper.Get("port")) // 8080
fmt.Println(viper.Get("hostname")) // myhostname.com
Firestore云存储:
viper.AddRemoteProvider("firestore", "google-cloud-project-id", "collection/document")
viper.SetConfigType("json") // Config's format: "json", "toml", "yaml", "yml"	配置的格式: “ json”、“ toml”、“ yaml”、“ yml”
err := viper.ReadRemoteConfig()
当然，您也可以使用 SecureRemoteProvider

12.Remote Key/Value Store Example - Encrypted	远程密钥/值存储示例-加密
viper.AddSecureRemoteProvider("etcd","http://127.0.0.1:4001","/config/hugo.json","/etc/secrets/mykeyring.gpg")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，所以支持的扩展名是“ json”、“ toml”、“ yaml”、“ yml”、“ properties”、“ props”、“ prom”、“ env”、“ dotenv”
err := viper.ReadRemoteConfig()

13.Watching Changes in etcd - Unencrypted	观察 etcd 中的更改-未加密
var runtime_viper = viper.New()			// 或者，您可以创建一个新的 viper 实例。
runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
runtime_viper.SetConfigType("yaml") 	// 因为在字节流中没有文件扩展名，所以支持的扩展名是“ json”、“ toml”、“ yaml”、“ yml”、“ properties”、“ props”、“ prom”、“ env”、“ dotenv”
err := runtime_viper.ReadRemoteConfig()	// 第一次从远程配置读取。
runtime_viper.Unmarshal(&runtime_conf)	// unmarshal config
go func(){								// 打开一个 goroutine 来永远观看远程更改
	for {
		time.Sleep(time.Second * 5) 		// 在每次请求后延迟
		err := runtime_viper.WatchRemoteConfig()// 目前，仅使用 etcd 支持进行测试
		if err != nil {
			log.Errorf("unable to read remote config: %v", err)
			continue
		}
		//将新配置解组到我们的运行时配置结构中。您也可以使用通道来实现信号，以通知系统的更改
		runtime_viper.Unmarshal(&runtime_conf)
	}
}()


Getting Values From Viper	从 Viper 获取 Values
在 Viper 中，有几种根据值的类型获取值的方法。现有以下函数和方法:
Get(key string) : interface{}
GetBool(key string) : bool
GetFloat64(key string) : float64
GetInt(key string) : int
GetIntSlice(key string) : []int
GetString(key string) : string
GetStringMap(key string) : map[string]interface{}
GetStringMapString(key string) : map[string]string
GetStringSlice(key string) : []string
GetTime(key string) : time.Time
GetDuration(key string) : time.Duration
IsSet(key string) : bool
AllSettings() : map[string]interface{}
需要注意的一点是，如果没有找到，每个 Get 函数将返回一个零值。为了检查给定的密钥是否存在，提供了 IsSet ()方法。例如：
viper.GetString("logfile") 			// 不区分大小写的设置与获取
if viper.GetBool("verbose") {
	fmt.Println("verbose enabled")
}
1.Accessing nested keys		访问嵌套密钥
访问器方法还接受深度嵌套键的格式化路径:
{
	"host": {
		"address": "localhost",
		"port": 5799
	},
	"datastore": {
		"metric": {
			"host": "127.0.0.1",
			"port": 3099
		},
		"warehouse": {
			"host": "198.0.0.1",
			"port": 2112
		}
	}
}
Viper 可以通过传递. 分隔的键路径来访问嵌套字段:	GetString("datastore.metric.host") // (returns "127.0.0.1")
这符合上面建立的优先级规则; 对路径的搜索将通过剩余的配置注册中心级联，直到找到。
例如，给定这个配置文件，datastore.metric.host 和 datastore.metric.port 都已经定义(并且可能被覆盖)。如果在默认情况下另外定义
了datastore.metric.protocol，Viper也会找到它。
但是，如果 datastore.metrics 被覆盖(通过一个标志、一个环境变量、 Set ()方法，...)并且有一个立即的值，那么 datastore.metrics 的所有子键
都将变得未定义，它们将被更高优先级的配置级别“隐藏”。
Viper 可以通过路径中的数字访问数组索引。例如:
{
	"host": {
		"address": "localhost",
		"ports": [
			5799,
			6029
		]
	},
	"datastore": {
		"metric": {
			"host": "127.0.0.1",
			"port": 3099
		},
		"warehouse": {
			"host": "198.0.0.1",
			"port": 2112
		}
	}
}
GetInt("host.ports.1") // returns 6029
最后，如果存在一个与分隔的键路径匹配的键，那么它的值将被返回。
{
	"datastore.metric.host": "0.0.0.0",
	"host": {
		"address": "localhost",
		"port": 5799
	},
	"datastore": {
		"metric": {
			"host": "127.0.0.1",
			"port": 3099
		},
		"warehouse": {
			"host": "198.0.0.1",
			"port": 2112
		}
	}
}
GetString("datastore.metric.host") // returns "0.0.0.0"

2.Extracting a sub-tree		提取sub-tree
在开发可重用模块时，提取配置的子集并将其传递给模块通常很有用。通过这种方式，可以使用不同的配置多次实例化模块。
例如，应用程序可能为不同目的使用多个不同的缓存存储区:
cache:
	cache1:
		max-items: 100
		item-size: 64
	cache2:
		max-items: 200
		item-size: 80
我们可以将缓存名传递给一个模块(例如NewCache(“cache1”)，但是访问配置键需要奇怪的连接，并且与全局配置的分离程度较低。
因此，与其这样做，不如将 Viper 实例传递给代表配置子集的构造函数:
cache1Config := viper.Sub("cache.cache1")
if cache1Config == nil { // Sub returns nil if the key cannot be found	如果找不到密钥，Sub 将返回 null
	panic("cache configuration not found")
}
cache1 := NewCache(cache1Config)
注意: 始终检查 Sub 的返回值。如果找不到键，它会返回 nil。
在内部，NewCache 函数可以直接处理 max-item 和 item-size 键:
func NewCache(v *Viper) *Cache {
	return &Cache{
		MaxItems: v.GetInt("max-items"),
		ItemSize: v.GetInt("item-size"),
	}
}
产生的代码很容易测试，因为它与主配置结构解耦，并且更容易重用(出于同样的原因)。

3.Unmarshaling
您还可以选择将所有值或特定值解组到结构、映射等。
有两种方法可以做到这一点:
Unmarshal(rawVal interface{}) : error
UnmarshalKey(key string, rawVal interface{}) : error
例如：
type config struct {
	Port int
	Name string
	PathMap string `mapstructure:"path_map"`
}
var C config
err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("unable to decode into struct, %v", err)
}
如果要解组键本身包含点(默认的键分隔符)的配置，必须更改分隔符:
v := viper.NewWithOptions(viper.KeyDelimiter("::"))
v.SetDefault("chart::values", map[string]interface{}{
	"ingress": map[string]interface{}{
		"annotations": map[string]interface{}{
			"traefik.frontend.rule.type":                 "PathPrefix",
			"traefik.ingress.kubernetes.io/ssl-redirect": "true",
		},
	},
})
type config struct {
	Chart struct{
		Values map[string]interface{}
	}
}
var C config
v.Unmarshal(&C)
Viper 还支持将数据解组为嵌入式结构:
/*
   Example config实力配置:

   module模块:
       enabled: true
       token: 89h3f98hbwf987h3f98wenf89ehf
*/
type config struct {
	Module struct {
		Enabled bool
		moduleConfig `mapstructure:",squash"`
	}
}
// moduleConfig could be in a module specific package	// moduleConfig 可能位于模块特定的包中
type moduleConfig struct {
	Token string
}
var C config
err := viper.Unmarshal(&C)
if err != nil {
	t.Fatalf("unable to decode into struct, %v", err)
}
Viper 使用底层 github.com/mitchellh/mapstructure 来解组值，默认情况下使用 mapstructtag。

4.Decoding custom formats	解码自定义格式
Viper 经常需要的一个特性是添加更多的值格式和解码器。例如，解析字符(点、逗号、分号等)将字符串分隔成片。
这在使用映射结构解码钩子的 Viper 中已经可用。
阅读更多关于这篇blog的细节。

5.Marshalling to string 	编组到字符串
您可能需要将 viper 中保存的所有设置编组成一个字符串，而不是将它们写入文件。您可以对 AllSettings ()返回的配置使用您喜欢的格式的编组器。
import (
	yaml "gopkg.in/yaml.v2"
	// ...
)
func yamlStringSettings() string {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	return string(bs)
}


Viper or Vipers?
Viper已经准备好开箱使用了。开始使用 Viper 不需要配置或初始化。由于大多数应用程序都希望使用单个中央存储库进行配置，因此 viper 包提供了这一功能。它类似于单例模式。
在上面的所有示例中，它们都演示了如何在单例样式方法中使用 viper。
1.Working with multiple vipers	和多条Viper一起工作
您还可以创建许多不同的vipers在您的应用程序中使用。每一个都有自己独特的一组配置和值。每个都可以从不同的配置文件、键值存储等读取。Viper 软件包支持的所有功能都反映为 viper 上的方法。
例如：
x := viper.New()
y := viper.New()
x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")
//...
在处理多条vipers时，由用户来跟踪不同的vipers。


Q & A
Why is it called “Viper”?
答: 毒蛇Viper是被设计成眼镜蛇Cobra的伴侣。虽然两者都可以完全独立地进行操作，但是它们结合在一起构成了一个强大的对，可以处理应用程序的大部分基础需求。

Why is it called “Cobra”?
答：还有比这更适合指挥官的名字吗？

Viper 支持区分大小写的密钥吗？
答：没有
Viper 从各种源合并配置，其中许多源要么不区分大小写，要么使用与其他源不同的大小写(例如 env vars)。为了在使用多个源时提供最佳体验，决定使所有键不区分大小写。
已经有好几个实现大小写敏感性的尝试，但不幸的是，它并不是那么微不足道。我们可能会尝试在 Viper v2中实现它，但是尽管最初的噪音，它似乎并没有被要求那么多。
你可以通过填写以下反馈表格为大小写敏感性投票:  https://forms.gle/r6fau74qprpazchz9

同时对viper进行读写安全吗？
不安全，您需要自己同步对viper的访问(例如使用同步包)。并发读写可能会导致恐慌。