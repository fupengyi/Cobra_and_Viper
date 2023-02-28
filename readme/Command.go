package readme

type Command struct {
	// Use is the one-line usage message.		Use是单行usage信息
	// Recommended syntax is as follows:		Recommended语法如下：
	//   [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
	//   ... indicates that you can specify multiple values for the previous argument.
	//   |   indicates mutually exclusive information. You can use the argument to the left of the separator or the
	//       argument to the right of the separator. You cannot use both arguments in a single use of the command.
	//   { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
	//       optional, they are enclosed in brackets ([ ]).
	//。

	//	 []	 [可选参数]	例子：
	//	 ... 参数切片		例子：[-F file | -D dir]...
	//   |   互斥参数，不能在一次使用命令时同时使用这两个参数。	例子：[-F file | -D dir]
	//   { } 分隔一组互斥参数。如果参数是可选参数，则将其括在[]中。
	// Example: add [-F file | -D dir]... [-f format] profile
	Use string

	// Aliases是一个别名数组，可以用来代替Use中的第一个单词。
	// Aliases is an array of aliases that can be used instead of the first word in Use.
	Aliases []string

	// SuggestFor是一个command names数组for which this command will be suggested -
	// SuggestFor is an array of command names for which this command will be suggested -
	// 类似于aliases，但只suggests
	// similar to aliases but only suggests.
	SuggestFor []string

	// Short是“help”输出中显示的简短描述。
	// Short is the short description shown in the 'help' output.
	Short string

	// 此子命令在其父命令的“help”输出中分组的组id。
	// The group id under which this subcommand is grouped in the 'help' output of its parent.
	GroupID string

	// Long是“help＜这个command＞”输出中显示的长消息。
	// Long is the long message shown in the 'help <this-command>' output.
	Long string

	// Example是如何使用命令的示例。
	// Example is examples of how to use the command.
	Example string

	// ValidArgs是shell完成中接受的所有有效非标志参数的列表
	// ValidArgs is list of all valid non-flag arguments that are accepted in shell completions
	ValidArgs []string
	// ValidArgsFunction是一个可选函数，为shell完成提供有效的非标志参数
	// ValidArgsFunction is an optional function that provides valid non-flag arguments for shell completion.
	// 这是使用ValidArgs的动态版本
	// It is a dynamic version of using ValidArgs.
	// 命令只能使用ValidArgs和ValidArgFunction之一。
	// Only one of ValidArgs and ValidArgsFunction can be used for a command.
	ValidArgsFunction func(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)

	// 预期参数
	// Expected arguments
	Args PositionalArgs

	// ArgAliases是ValidArgs的别名列表。
	// ArgAliases is List of aliases for ValidArgs.
	// 不建议用户在shell completion，但可以接受手动输入
	// These are not suggested to the user in the shell completion, but accepted if entered manually.
	ArgAliases []string

	// BashCompletionFunction is custom bash functions used by the legacy bash autocompletion generator.
	// 为了与其他shell兼容，建议改用ValidArgsFunction
	// For portability with other shells, it is recommended to instead use ValidArgsFunction
	BashCompletionFunction string

	// 如果此命令已弃用，并且应在使用时打印此字符串，则定义已弃用。
	// Deprecated defines, if this command is deprecated and should print this string when used.
	Deprecated string

	// 注释是键/值对，应用程序可以使用它们来标识或分组命令。
	// Annotations are key/value pairs that can be used by applications to identify or group commands.
	Annotations map[string]string

	/*
		Version定义此命令的版本。如果此值为非空且此命令未定义 “version” flag，则将“version” boolean flag添加到命令中，如果指定，将打印“Version”变量的内容。
		如果命令未定义“v” flag，也将添加一个简写 “v” flag。
	*/
	// Version defines the version for this command. If this value is non-empty and the command does not
	// define a "version" flag, a "version" boolean flag will be added to the command and, if specified,
	// will print content of the "Version" variable. A shorthand "v" flag will also be added if the
	// command does not define one.
	Version string

	// Run函数按以下顺序执行：
	// The *Run functions are executed in the following order:
	//   * PersistentPreRun()
	//   * PreRun()
	//   * Run()
	//   * PostRun()
	//   * PersistentPostRun()
	// All functions get the same args, the arguments after the command name.
	//
	// PersistentPreRun: children of this command will inherit and execute. 	// PersistentPreRun:此命令的子级将继承并执行。
	PersistentPreRun func(cmd *Command, args []string)
	// PersistentPreRunE: PersistentPreRun but returns an error.				// PersistentPreRunE:此命令的子级将继承并执行,返回值error
	PersistentPreRunE func(cmd *Command, args []string) error
	// PreRun: children of this command will not inherit.						// PreRun:此命令的子级不会继承。
	PreRun func(cmd *Command, args []string)
	// PreRunE: PreRun but returns an error.									// PreRunE:此命令的子级不会继承, 返回值error
	PreRunE func(cmd *Command, args []string) error
	// Run: Typically the actual work function. Most commands will only implement this.  // Run:通常为实际工作函数。大多数命令将仅实现此功能。
	Run func(cmd *Command, args []string)
	// RunE: Run but returns an error.											// RunE:通常为实际工作函数。大多数命令将仅实现此功能，返回值error
	RunE func(cmd *Command, args []string) error
	// PostRun: run after the Run command.										// PostRun：在Run命令之后运行。
	PostRun func(cmd *Command, args []string)
	// PostRunE: PostRun but returns an error.									// PostRunE：在Run命令之后运行,返回值error。
	PostRunE func(cmd *Command, args []string) error
	// PersistentPostRun: children of this command will inherit and execute after PostRun.  // PersistentPostRun:此命令的子级将在PostRun之后继承并执行。
	PersistentPostRun func(cmd *Command, args []string)
	// PersistentPostRunE: PersistentPostRun but returns an error.				// PersistentPostRun:此命令的子级将在PostRun之后继承并执行,返回值error
	PersistentPostRunE func(cmd *Command, args []string) error

	// FParseErrWhitelist flag parse errors to be ignored						// FParseErrWhitelist: flag parse errors被忽略
	FParseErrWhitelist FParseErrWhitelist

	// CompletionOptions is a set of options to control the handling of shell completion  // CompletionOptions:是一组选项，用于控制shell完成的处理
	CompletionOptions CompletionOptions

	// TraverseChildren parses flags on all parents before executing child command. // TraverseChildren:在执行子命令之前先解析所有父级上的flags。
	TraverseChildren bool

	// Hidden defines, if this command is hidden and should NOT show up in the list of available commands. // Hidden:如果此命令是隐藏的，并且不应显示在可用命令列表中。
	Hidden bool

	// SilenceErrors is an option to quiet errors down stream.  				// SilenceErrors:quiet errors down stream
	SilenceErrors bool

	// SilenceUsage is an option to silence usage when an error occurs.  		// SilenceUsage:silence usage when an error occurs
	SilenceUsage bool

	// DisableFlagParsing disables the flag parsing.							// DisableFlagParsing:禁用标志解析。
	// If this is true all flags will be passed to the command as arguments.	// 如果为true，所有flags都将作为参数传递给命令。
	DisableFlagParsing bool

	// DisableAutoGenTag:如果gen tag（“由spf13/cobra自动生成…”）将通过生成此命令的文档来打印。
	// DisableAutoGenTag defines, if gen tag ("Auto generated by spf13/cobra...") will be printed by generating docs for this command.
	DisableAutoGenTag bool

	// DisableFlagsInUseLine:当打印help或生成docs时，将禁用向命令的用法行添加[flags]
	// DisableFlagsInUseLine will disable the addition of [flags] to the usage line of a command when printing help or generating docs
	DisableFlagsInUseLine bool

	// DisableSuggestions:禁用基于Levenshtein距离的建议，这些建议为“unknown command”消息。
	// DisableSuggestions disables the suggestions based on Levenshtein distance that go along with 'unknown command' messages.
	DisableSuggestions bool
	// SuggestionsMinimumDistance:定义显示建议的最小levenshtein距离。必须大于0。
	// SuggestionsMinimumDistance defines minimum levenshtein distance to display suggestions. Must be > 0.
	SuggestionsMinimumDistance int
	// 包含筛选或未报告的字段
	// contains filtered or unexported fields
}
