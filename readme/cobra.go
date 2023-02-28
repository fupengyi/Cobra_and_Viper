package readme
Cobra 是一个用于创建强大的现代 CLI 应用程序的库。
Cobra 用于许多 Go 项目，例如 Kubernetes、Hugo 和 GitHub CLI 等等。此列表包含更广泛的使用 Cobra 的项目列表。

Overview
Cobra 是一个库，提供了一个简单的界面来创建类似于 git & go 工具的强大的现代 CLI 界面。
Cobra provides:
	简单的基于子命令的 CLI：应用程序服务器、应用程序获取等。
	完全符合 POSIX 标准的标志（包括短版和长版）
	嵌套子命令
	全局、本地和级联标志
	智能建议（应用服务器...您是指应用服务器吗？）
	命令和标志的自动help生成
	子命令的分组help
	-h、--help 等的自动help标志识别。
	为您的应用程序自动生成的 shell 自动完成功能（bash、zsh、fish、powershell）
	为您的应用程序自动生成手册页
	命令别名，这样您就可以在不破坏它们的情况下进行更改
	定义您自己的帮助、用法等的灵活性。
	可选择与 viper 无缝集成以实现 12 要素应用程序

Concepts
Cobra 建立在命令、参数和标志的结构之上。
命令代表动作，参数是事物，标志是这些动作的修饰符。
最好的应用程序在使用时读起来像句子，因此，用户凭直觉知道如何与之交互。
要遵循的模式是 APPNAME VERB NOUN --ADJECTIVE 或 APPNAME COMMAND ARG --FLAG。
一些很好的现实世界的例子可能会更好地说明这一点。
在以下示例中，“server”是一个命令，“port”是一个标志：
	hugo server --port=1313
在这个命令中，我们告诉 Git 克隆 url bare。
	git clone URL --bare

Commands
命令是应用程序的中心点。应用程序支持的每个交互都将包含在命令中。一个命令可以有子命令并且可以选择运行一个动作。
在上面的示例中，server 是命令。
更多关于 cobra.Command

Flags
标志是一种修改命令行为的方法。 Cobra 支持完全符合 POSIX 的标志以及 Go 标志包。 Cobra 命令可以定义持续到子命令的标志和仅对该命令可用的标志。
在上面的示例中，port 是标志。
标志功能由 pflag 库提供，它是标志标准库的一个分支，它在添加 POSIX 合规性的同时保持相同的接口。

Installing
使用 Cobra 很容易。首先，使用 go get 安装最新版本的库。
	go get -u github.com/spf13/cobra@latest
接下来，在您的应用程序中包含 Cobra：
	import "github.com/spf13/cobra"

Usage
cobra-cli 是一个命令行程序，用于生成 cobra 应用程序和命令文件。它将引导您的应用程序脚手架以快速开发基于 Cobra 的应用程序。这是将 Cobra 合
并到您的应用程序中的最简单方法。

它可以通过运行来安装：
	go install github.com/spf13/cobra-cli@latest
有关使用 Cobra-CLI 生成器的完整详细信息，请阅读 The Cobra Generator README
有关使用 Cobra 库的完整详细信息，请阅读  The Cobra User Guide。

License
Cobra 在 Apache 2.0 许可下发布。请参阅 LICENSE.txt
