虽然欢迎您提供自己的组织，但通常基于Cobra的应用程序将遵循以下组织结构：
▾ appName/
  ▾ cmd/
	root.go
	add.go
	your.go
	commands.go
	here.go
  main.go


在Cobra应用程序中，通常是主应用程序。go文件非常简单。它有一个目的：初始化Cobra。
package main
import (
	"{pathToYourApp}/cmd"
)
func main() {
	cmd.Execute()
}


Using the Cobra Generator 使用Cobra生成器
Cobra CLI 是它自己的程序，它将创建您的应用程序并添加您想要的任何命令。这是将 Cobra 合并到应用程序中最简单的方法。
有关使用 Cobra 生成器的完整详细信息，请参阅 The Cobra-CLI Generator README


Using the Cobra Library 使用Cobra库
要手动实现Cobra，您需要创建一个bare main.go文件和rootCmd文件。您可以根据需要提供其他命令。
1.创建rootCmd
Cobra不需要任何特殊的构造函数。只需创建命令即可。
理想情况下，您可以将其放在app/cmd/root.go中：
var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here //在此处执行操作
	},
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
您还将在init()函数中定义flags和handle配置。例如：cmd/root.go
package cmd
import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
var (
	// Used for flags.用于flags
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,
		}
	)
// Execute executes the root command.执行根命令。
func Execute() error {
	return rootCmd.Execute()
}
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)
}
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.	//使用标志中的配置文件。
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.				//查找主目录。
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).//在主目录中搜索名为“.cobra”的配置(不带扩展名)。
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

2.创建 main.go
使用root命令，您需要让main函数执行它。为了清楚起见，应该在根目录上运行Execute，尽管它可以在任何命令上调用。
在cobra应用程序中，通常是主应用程序。go文件非常简单。它有一个目的：初始化Cobra。
package main
import (
	"{pathToYourApp}/cmd"
)
func main() {
	cmd.Execute()
}

3.Create additional commands
可以定义其他命令，通常每个命令在cmd/目录中都有自己的文件。
如果要创建version命令，可以创建cmd/version。使用以下内容填充它：
package cmd
import (
	"fmt"
	"github.com/spf13/cobra"
)
func init() {
	rootCmd.AddCommand(versionCmd)
}
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

4.返回和处理错误,如果希望向命令的调用方返回错误，可以把Run换成RunE,然后可以在执行函数调用时捕获错误。
	RunE: func(cmd *cobra.Command, args []string) error {
			if err := someFunc(); err != nil {
				return err
			}
			return nil
	},



Working with Flags 使用Flags
Flags提供修饰符来控制命令的操作方式。
1.为command指定flags
由于flags是在不同的位置定义和使用的，因此我们需要在外部定义一个具有正确范围的变量来分配要使用的flags。
	var Verbose bool
	var Source string
分配flag有两种不同的方法。

2.Persistent Flags
永久flag可以是“持久”的，这意味着该标志将可用于分配给它的命令以及该命令下的每个命令。对于全局标志，在根上指定一个标志作为持久标志。
rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

3.Local Flags
也可以在本地分配flag，这只适用于该特定命令。
localCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")

4.父命令上的本地标志Local Flag on Parent Commands
默认情况下，Cobra只解析目标命令上的本地标志，父命令上的任何本地标志都将被忽略。通过启用Command.TraverseChildren，Cobra将在执行目标命令之前解析每个命令的本地标志。
command := cobra.Command{
	Use: "print [OPTIONS] [COMMANDS]",
	TraverseChildren: true,
}

5.使用配置绑定标志
你也可以用viper绑定你的flags：
var author string
func init() {
	rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
}
在上面本例中，持久flag author与viper绑定。注意：当用户提供--author标志时，变量author不会设置为config中的值。更多详情请看《viper documentation》

6.Required flags
默认情况下，flag是可选的。如果您希望您的命令在未设置flag时报告错误，请标记为required：
rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
rootCmd.MarkFlagRequired("region")
或者，对于永久标志：
rootCmd.PersistentFlags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
rootCmd.MarkPersistentFlagRequired("region")

7.Flag Groups 如果您必须同时提供不同的标志（例如，如果他们提供了--username标志，则必须同时提供--password标志），那么Cobra可以强制执行该要求：
rootCmd.Flags().StringVarP(&u, "username", "u", "", "Username (required if password is set)")
rootCmd.Flags().StringVarP(&pw, "password", "p", "", "Password (required if username is set)")
rootCmd.MarkFlagsRequiredTogether("username", "password")
如果不同的标志表示互斥选项，例如将输出格式指定为--json或--yaml，但不能同时指定为两者，则也可以防止它们一起提供：
rootCmd.Flags().BoolVar(&u, "json", false, "Output in JSON")
rootCmd.Flags().BoolVar(&pw, "yaml", false, "Output in YAML")
rootCmd.MarkFlagsMutuallyExclusive("json", "yaml")
在这两种情况：
	本地标志和持久标志都在
		注意：该组仅在定义了每个标志的命令上强制执行
	一个标志可能出现在多个组中
	一个组可以包含任意数量的标志




Positional and Custom Arguments预期参数和自定义参数
可以使用Command的Args字段指定预期参数的验证。内置验证器如下：
	1.Number of arguments:
		NoArgs 					如果存在任何参数，则报告error。
		ArbitraryArgs 			接受任意数量的参数。
		MinimumNArgs（int） 		如果提供的位置参数少于N个，则报告错误。
		MaximumNArgs（int） 		如果提供了N个以上的位置参数，则报告错误。
		ExactArgs（int）		 	如果没有正好N个位置参数，则报告错误。
		RangeArgs（min，max）	如果参数数不在min和max之间，则报告错误。
	2.Content of the arguments:
		OnlyValidArgs			如果命令的ValidArg字段中未指定任何预期参数，则报告错误，可以选择将其设置为预期参数的有效值列表。
如果Args为undefined或nil，则默认为ArbitraryArgs。
此外，MatchAll（pargs…PositionalArgs）允许将现有检查与任意其他检查相结合。例如，如果您想报告一个错误，如果没有正好N个预期参数，或者如果有
任何预期参数不在Command的ValidArgs字段中，您可以在ExactArgs和OnlyValidArg上调用MatchAll，如下所示：
var cmd = &cobra.Command{
	Short: "hello",
	Args: cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}
可以设置满足func（cmd *cobra.Command，args []string）error的任何自定义验证器。例如：
var cmd = &cobra.Command{
	Short: "hello",
	Args: func(cmd *cobra.Command, args []string) error {
			// Optionally run one of the validators provided by cobra 可选运行cobra提供的一个验证器
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			// Run the custom validation logic	运行自定义验证逻辑
			if myapp.IsValidColor(args[0]) {
				return nil
			}
			return fmt.Errorf("invalid color specified: %s", args[0])
		},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}



Example
在下面的示例中，我们定义了三个命令。两个位于顶层，一个（cmdTimes）是顶层命令之一的子命令。在这种情况下，root是不可执行的，这意味着子命令
是required。这是通过不为“rootCmd”提供“Run”来实现的。
我们只为单个命令定义了一个标志。
有关flags的更多文档，请访问https://github.com/spf13/pflag
package main
import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)
func main() {
	var echoTimes int
	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
				For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}
	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
				Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}
	var cmdTimes = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "Echo anything to the screen more times",
		Long: `echo things multiple times back to the user by providing
				a count and a string.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
				for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
				}
		},
	}
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint, cmdEcho)
	cmdEcho.AddCommand(cmdTimes)
	rootCmd.Execute()
}
//	有关更完整的大型应用程序示例，请查看Hugo。



Help Command
当您有子命令时，Cobra会自动向应用程序添加help命令。当用户运行“app help”时将调用此函数。此外，help还将支持所有其他命令作为输入。例如，您有
一个名为“create”的命令，没有任何额外的配置；当调用“app help create”时，Cobra将工作。每个命令都会自动添加“--help” flag。
1.Example
以下输出由Cobra自动生成。除了commend和flag定义之外，不需要任何东西。
$ cobra-cli help

Cobra is a CLI library for Go that empowers applications. 		Cobra是一个用于Go的CLI库，它支持应用程序。
This application is a tool to generate the needed files to quickly create a Cobra application.	这个应用程序是一个生成所需文件以快速创建Cobra应用程序的工具。

Usage:
cobra-cli [command]

Available Commands:
add         Add a command to a Cobra Application
completion  Generate the autocompletion script for the specified shell
help        Help about any command
init        Initialize a Cobra Application

Flags:
-a, --author string    author name for copyright attribution (default "YOUR NAME")
--config string    config file (default is $HOME/.cobra.yaml)
-h, --help             help for cobra-cli
-l, --license string   name of license for the project
--viper            use Viper for configuration

Use "cobra-cli [command] --help" for more information about a command.
“help”就像其他命令一样。它周围没有特殊的逻辑或行为。事实上，如果你愿意，你可以自己提供。

2.help中的命令分组
Cobra支持可用命令的分组。组必须由AddGroup显式定义，并由子命令的GroupId元素设置。这些组将按定义的顺序显示。如果使用生成的help或completion
命令，则可以分别通过SetHelpCommandGroupId和SetCompletionCommandGroupId设置组ID。

3.自定义help  您可以为默认命令提供自己的help命令或模板，以用于以下函数：
	cmd.SetHelpCommand(cmd *Command)
	cmd.SetHelpFunc(f func(*Command, []string))
	cmd.SetHelpTemplate(s string)
后两个命令也适用于任何子命令。



Usage Message
当用户提供无效标志或无效命令时，Cobra会向用户显示 "usage"。
你可以从上面的help中认识到这一点。这是因为默认help将用法作为其输出的一部分嵌入。
1.自定义用法  您可以提供自己的使用函数或模板供Cobra使用。与help一样，函数和模板可以通过公共方法重写：
	cmd.SetUsageFunc(f func(*Command) error)
	cmd.SetUsageTemplate(s string)


Version Flag
如果根命令上设置了version字段，则Cobra会添加顶级的“--version”标志。运行带有“--version”标志的应用程序将使用版本模板将版本打印到stdout。
可以使用cmd自定义模板。SetVersionTemplate（s string）函数。


PreRun and PostRun Hooks	运行前和运行后hook函数
可以在命令的主运行函数之前或之后运行函数。PersistentPreRun和PreRun函数将在运行之前执行。PersistentPostRun和PostRuns将在运行后执行。如果
Persistent*Run函数未声明自己的函数，则它们将由子函数继承。这些功能按以下顺序运行：
		PersistentPreRun PreRun Run PostRun PersistentPostRun
下面是使用所有这些功能的两个命令的示例。执行子命令时，它将运行根命令的PersistentPreRun，而不是根命令的PersistentPostRun：
package main

import (
	"fmt"
	"github.com/spf13/cobra"
)
func main() {
	var rootCmd = &cobra.Command{
		Use:   "root [sub]",
		Short: "My root command",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd Run with args: %v\n", args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
		},
	}
	var subCmd = &cobra.Command{
		Use:   "sub [no options!]",
		Short: "My subcommand",
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside subCmd PreRun with args: %v\n", args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside subCmd Run with args: %v\n", args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside subCmd PersistentPostRun with args: %v\n", args)
		},
	}
	rootCmd.AddCommand(subCmd)
	rootCmd.SetArgs([]string{""})
	rootCmd.Execute()
	fmt.Println()
	rootCmd.SetArgs([]string{"sub", "arg1", "arg2"})
	rootCmd.Execute()
}
输出：
Inside rootCmd PersistentPreRun with args: []
Inside rootCmd PreRun with args: []
Inside rootCmd Run with args: []
Inside rootCmd PostRun with args: []
Inside rootCmd PersistentPostRun with args: []

Inside rootCmd PersistentPreRun with args: [arg1 arg2]
Inside subCmd PreRun with args: [arg1 arg2]
Inside subCmd Run with args: [arg1 arg2]
Inside subCmd PostRun with args: [arg1 arg2]
Inside subCmd PersistentPostRun with args: [arg1 arg2]



Suggestions when "unknown command" happens	发生“未知命令”时的建议
当“未知命令”错误发生时，Cobra将打印自动建议。这使Cobra可以在出现错误时类似于git命令。例如：
$ hugo srever
Error: unknown command "srever" for "hugo"

Did you mean this?

server
Run 'hugo --help' for usage.
Suggestions基于现有子命令自动生成，并使用Levenshtein距离的实现。匹配最小距离2（忽略大小写）的每个注册命令都将显示为建议。
如果需要在命令中禁用建议或调整字符串距离，请使用：
command.DisableSuggestions = true
command.SuggestionsMinimumDistance = 1
您还可以使用SuggestFor属性显式设置建议使用给定命令的名称。这允许对字符串距离不近的字符串提供建议，但在命令集中有意义，但不需要别名。例子：
$ kubectl remove
Error: unknown command "remove" for "kubectl"

Did you mean this?
delete

Run 'kubectl help' for usage.




Cobra可以基于子命令、标志等生成文档。请参阅文档生成文档了解更多信息。
Cobra可以为以下shell生成shell完成文件：bash、zsh、fish、PowerShell。如果您向命令中添加更多信息，这些补全将非常强大和灵活。在Shell Completions中了解更多信息。
Cobra使用shell完成系统来定义一个框架，允许您向用户提供活动帮助。活动帮助是在使用程序时打印的消息（提示、警告等）。请在“活动帮助”中阅读更多信息。
