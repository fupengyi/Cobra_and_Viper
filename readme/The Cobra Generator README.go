Cobra 生成器  Cobra Generator
Cobra提供了自己的程序，可以创建您的应用程序并添加您想要的任何命令。这是将Cobra合并到应用程序中的最简单方法。
使用命令 "go install github.com/spf13/cobra-cli@latest"安装cobra生成器.Go将自动将其安装在$GOPATH/bin目录中，该目录应该位于$PATH中。
安装后，您应该可以使用 "cobra-cli"命令，在命令行输入 "cobra-cli"进行确认。
Cobra生成器当前仅支持两种操作：


cobra-cli 初始化	cobra-cli init
"cobra-cli init[app]"命令将为您创建初始应用程序代码。这是一个非常强大的应用程序，它将以正确的结构填充您的程序，这样您就可以立即享受Cobra的所有好处。
它还可以将您指定的许可证应用于您的应用程序。
随着Go模块的引入，Cobra生成器已被简化以利用模块。Cobra生成器在Go模块中工作。
1.初始化模块 Initalizing a module
如果您已经有了模块，请跳过此步骤。
如果要初始化新的Go模块：
	1.Create a new directory		创建新目录
	2.cd into that directory		cd到该目录中
	3.run "go mod init <MODNAME>"	运行go mod init＜MODNAME＞
	e.g.  	cd $HOME/code
			mkdir myapp
			cd myapp
			go mod init github.com/spf13/myapp

2.初始化Cobra CLI应用程序 Initalizing a Cobra CLI application
在Go模块中运行"cobra-cli init"。这将创建一个新的主干项目供您编辑。
您应该能够立即运行新的应用程序。"go run main.go"
您需要打开并编辑`cmd/root.go`，提供你自己的描述和逻辑。
e.g.	cd $HOME/code/myapp
		cobra-cli init
		go run main.go
cobra-cli init也可以从子目录运行，例如cobra生成器本身的组织方式。如果您希望将应用程序代码与库代码分开，这将非常有用。

3.Optional flags:
您可以使用--author flag为其提供作者姓名 e.g.  cobra-cli init --author "Steve Francia spf@spf13.com"
您可以提供使用许可证--license  e.g.	  cobra-cli init --license apache
使用--viper flag自动设置viper
Viper是Cobra的伴侣，旨在提供对环境变量和配置文件的轻松处理，并将它们无缝连接到应用程序标志。




Add commands to a project向项目添加命令
初始化cobra应用程序后，您可以继续使用cobra生成器向应用程序添加其他命令。执行此操作的命令是"cobra-cli add"
假设您创建了一个应用程序，并且需要以下命令：1.app serve 2.app config 3.app config create
在项目目录（main.go文件所在的目录）中，您将运行以下命令：
	cobra-cli add serve
	cobra-cli add config
	cobra-cli add create -p 'configCmd'
cobra-cli add支持与cobra-cli init相同的所有Optional flags（如上所述）。
您会注意到最后一个命令有一个-p flag。这用于将父命令指定给新添加的命令。在这种情况下，我们希望将“create”命令分配给“config”命令。如果未指定，所有命令都具有rootCmd的默认父级。
默认情况下，cobra-cli将在提供的名称后面追加Cmd，举例config --> configCmd,并将此名称用作内部变量名称。指定父级时，请确保与代码中使用的变量名匹配。
注意：命令名使用camelCase（而不是snake_case/kebab大小写）。否则，您将遇到错误。例如，cobra-cli add-add-user不正确，但cobra-cli add addUser有效。
运行这三个命令后，应用程序结构如下所示：
▾ app/
▾ cmd/
	  config.go
      create.go
      serve.go
      root.go
  main.go
此时，您可以运行go run main.go，它将运行您的应用程序,config,create,serve,help serve都会运行
现在您已经启动并运行了一个基本的Cobra-based应用程序。下一步是在cmd中编辑文件并为应用程序自定义它们。
有关使用cobra库的完整详细信息，请阅读 The Cobra User Guide。




配置cobra生成器
如果您提供一个简单的配置文件，它将帮助您避免在flags中反复提供大量重复信息，那么Cobra生成器将更易于使用。
An example ~/.cobra.yaml file:
	author: Steve Francia <spf@spf13.com>
	license: MIT
	useViper: true
您还可以使用内置许可证。例如，GPLv2, GPLv3, LGPL, AGPL, MIT, 2-Clause BSD or 3-Clause BSD.
您可以通过将license设置为none来指定无许可证，也可以指定自定义许可证：
author: Steve Francia <spf@spf13.com>
year: 2020
license:
	header: This file is part of CLI application foo.
	text: |
	  {{ .copyright }}
	This is my license. There are many like it, but this one is mine.
	My license is my best friend. It is my life. I must master it as I must master my life.

在上述自定义许可证配置中，许可证文本中的版权(copyright)行由作者和年份属性生成。许可证(LICENSE)文件的内容为
Copyright © 2020 Steve Francia <spf@spf13.com>
This is my license. There are many like it, but this one is mine.
My license is my best friend. It is my life. I must master it as I must master my life.

header属性用作许可证头文件。不进行插值。这是go文件头的示例。
/*
Copyright © 2020 Steve Francia <spf@spf13.com>
This file is part of CLI application foo.
*/
