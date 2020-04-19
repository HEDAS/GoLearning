package introduction

import (
	"fmt"
	"math/rand"
	"time"
)

func Basic() {
	introduction()
	concurrency()
	//producerAndCustomerDemo()
	project()
	bad()
	lib()
	regular()
	//runServer()
	compiler()
	install()
	structure()
}

func introduction() {
	/*
		简介
		Go语言不但能让你访问底层操作系统，还提供了强大的网络编程和并发编程支持。
		Go语言的用途众多，可以进行网络编程、系统编程、并发编程、分布式编程。

		具有“部署简单、并发性好、语言设计良好、执行性能好”等优势

		开源项目
		Docker、Go-Ethereum、Thrraform 和 Kubernetes

		创始人
		1) Ken Thompson
		贝尔实验室 Unix 团队成员，C语言、Unix 和 Plan 9 的创始人之一，在 20 世纪 70 年代，设计并实现了最初的 UNIX 操作系统，
		仅从这一点说，他对计算机科学的贡献怎么强调都不过分。他还与 Rob Pike 合作设计了 UTF-8 编码方案。
		2) Rob Pike
		Go语言项目总负责人，贝尔实验室 Unix 团队成员，除帮助设计 UTF-8 外，
		还帮助开发了分布式多用户操作系统 Plan 9、Inferno 操作系统和 Limbo 编程语言，
		并与人合著了《The Unix Programming Environment》，对 UNIX 的设计理念做了正统的阐述。
		3) Robert Griesemer
		就职于 Google，参与开发 Java HotSpot 虚拟机，对语言设计有深入的认识，
		并负责 Chrome 浏览器和 Node.js 使用的 Google V8 JavaScript 引擎的代码生成部分。

		优点
		快速编译，高效执行，易于开发

		Go语言支持交叉编译，比如说你可以在运行 Linux 系统的计算机上开发可以在 Windows 上运行的应用程序。

		这是第一门完全支持 UTF-8 的编程语言，这不仅体现在它可以处理使用 UTF-8 编码的字符串，
		就连它的源码文件格式都是使用的 UTF-8 编码。
	*/

	/*
		语言特性
		将“++”、“--”从运算符降级为语句，
		保留指针，但默认阻止指针运算
		将切片和字典作为内置类型
		从运行时的层面进行优化

		它用类协程的方式来处理并发单元，却又在运行时层面做了更深度的优化处理。
		无须处理回调，无须关注线程切换，仅一个关键字，简单而自然

		搭配 channel，实现 CSP 模型。将并发单元间的数据耦合拆解开来，各司其职
		若说有所不足，那就是应该有个更大的计划，将通信从进程内拓展到进程外，实现真正意义上的分布式

		高并发下的内存分配和管理
		Go 选择了 tcmalloc，它本就是为并发而设计的高性能内存分配组件。

		可以说，内存分配器是运行时三大组件里变化最少的部分。
		刨去因配合垃圾回收器而修改的内容，内存分配器完整保留了 tcmalloc 的原始架构。
		使用 cache 为当前执行线程提供无锁分配，多个 central 在不同线程间平衡内存单元复用。
		在更高层次里，heap 则管理着大块内存，用以切分成不同等级的复用内存块。
		快速分配和二级内存平衡机制，让内存分配器能优秀地完成高压力下的内存管理任务。

		在最近几个版本中，编译器优化卓有成效。它会竭力将对象分配在栈上，以降低垃圾回收压力，减少管理消耗，提升执行性能。
		可以说，除偶尔因性能问题而被迫采用对象池和自主内存管理外，我们基本无须参与内存管理操作。

		垃圾回收
		垃圾回收一直是个难题。早年间，Java 就因垃圾回收低效被嘲笑了许久，后来 Sun 连续收纳了好多人和技术才发展到今天。
		可即便如此，在 Hadoop 等大内存应用场景下，垃圾回收依旧捉襟见肘、步履维艰。

		相比 Java，Go 面临的困难要更多。因指针的存在，所以回收内存不能做收缩处理。幸好，指针运算被阻止，否则要做到精确回收都难。

		每次升级，垃圾回收器必然是核心组件里修改最多的部分。
		从并发清理，到降低 STW 时间，直到 Go 的 1.5 版本实现并发标记，逐步引入三色标记和写屏障等等，
		都是为了能让垃圾回收在不影响用户逻辑的情况下更好地工作。
		尽管有了努力，当前版本的垃圾回收算法也只能说堪用，离好用尚有不少距离。
	*/

	/*
		静态链接
		Go 刚发布时，静态链接被当作优点宣传。只须编译后的一个可执行文件，无须附加任何东西就能部署。
		这似乎很不错，只是后来风气变了。连着几个版本，编译器都在完善动态库 buildmode 功能，场面一时变得有些尴尬。

		暂不说未完工的 buildmode 模式，静态编译的好处显而易见。
		将运行时、依赖库直接打包到可执行文件内部，简化了部署和发布操作，无须事先安装运行环境和下载诸多第三方库。
		这种简单方式对于编写系统软件有着极大好处，因为库依赖一直都是个麻烦。

		标准库
		功能完善、质量可靠的标准库为编程语言提供了充足动力。
		在不借助第三方扩展的情况下，就可完成大部分基础功能开发，这大大降低了学习和使用成本。
		最关键的是，标准库有升级和修复保障，还能从运行时获得深层次优化的便利，这是第三方库所不具备的。

		Go 标准库虽称不得完全覆盖，但也算极为丰富。
		其中值得称道的是 net/http，仅须简单几条语句就能实现一个高性能 Web Server，这从来都是宣传的亮点。
		更何况大批基于此的优秀第三方 Framework 更是将 Go 推到 Web/Microservice 开发标准之一的位置。

		工具链
		完整的工具链对于日常开发极为重要。
		Go 在此做得相当不错，无论是编译、格式化、错误检查、帮助文档，还是第三方包下载、更新都有对应的工具。
		其功能未必完善，但起码算得上简单易用。

		内置完整测试框架，
		其中包括单元测试、性能测试、代码覆盖率、数据竞争，以及用来调优的 pprof，
		这些都是保障代码能正确而稳定运行的必备利器。

		除此之外，还可通过环境变量输出运行时监控信息，尤其是垃圾回收和并发调度跟踪，可进一步帮助我们改进算法，获得更佳的运行期表现。
	*/
}

func concurrency() {
	/*
		Go语言的并发是基于 goroutine 的，goroutine 类似于线程，但并非线程。可以将 goroutine 理解为一种虚拟线程。
		Go语言运行时会参与调度 goroutine，并将 goroutine 合理地分配到每个 CPU 中，最大限度地使用 CPU 性能。

		多个 goroutine 中，Go语言使用通道（channel）进行通信，
		通道是一种内置的数据结构，可以让用户在不同的 goroutine 之间同步发送具有类型的消息。
	*/
}

func project() {
	/*
		早期的Go语言开源项目只是通过Go语言与传统项目进行C语言库绑定实现，例如 Qt、Sqlite 等；

		https://github.com/docker/docker
		https://github.com/golang/go
		https://github.com/kubernetes/kubernetes
		https://github.com/coreos/etcd
		https://github.com/astaxie/beego
		https://github.com/go-martini/martini
		https://github.com/CodisLabs/codis
		https://github.com/derekparker/delve

		fackbook的开源go项目
		https://github.com/facebookgo
		最具代表性的就是著名平滑重启工具 grace

		小米
		http://open-falcon.org/

		360
		https://github.com/Qihoo360/poseidon

		在服务器编程方面，Go语言适合处理日志、数据打包、虚拟机处理、文件系统、分布式系统、数据库代理等；
		网络编程方面，Go语言广泛应用于 Web 应用、API 应用、下载应用等；
		此外，Go语言还可用于内存数据库和云平台领域，目前国外很多云平台都是采用 Go 开发。

		除了上面介绍到的，Go语言还可以用来开发底层，例如以太坊、超级账本等都是基于Go语言开发的。
		很多基于区块链的 DApps（去中心化应用）和工具都是用的Go语言来实现的。

		云计算基础设施领域，代表项目：docker、kubernetes、etcd、consul、cloudflare CDN、七牛云存储等。
		基础软件，代表项目：tidb、influxdb、cockroachdb 等。
		微服务，代表项目：go-kit、micro、monzo bank 的 typhon、bilibili 等。
		互联网基础设施，代表项目：以太坊、hyperledger 等。
	*/
}

func bad() {
	/*
		下面，我们来客观地看一下目前Go语言需要加强或改进的地方（虽然有些 Gopher 并不这么认为）。

		1) 从分布式计算的角度来看，Go语言的成熟度不及 Erlang（现在已经出现了一些这方面的Go语言代码包，我们已经可以看到光明的未来了）。

		2) 从程序运行速度的角度来看，Go语言虽然已与 Java 不相上下，但还不及 C（差距正在不断地缩小）。

		3) 从第三方库的角度来看，Go语言的库数量还远远不及其他几门主流语言（比如 Java、Python、Ruby 等）。
		不过与Go语言的年纪相比，用它实现的第三方库已经相当多了，并且它们的数量在持续地飞速增长中。

		另外，在更深的层面，Go语言标准库中也有些不尽如人意的的地方，具体如下。

		1) 从语言语法角度来看，Go语言语法里的语法糖并不多，这让许多 Python、Ruby 爱好者们对它不屑一顾。
		另外，变量赋值方式多得有点儿累赘了。
		最让人遗憾的也是我比较在意的一个地方是，Go语言不支持自定义的泛型类型。

		2) 从并发编程角度来看，Go语言提供的并发模型很强大，但也有一些编写规则需要了解。否则，很容易踩进“坑”里。
		其实不提倡把这叫作“坑”。因为这些所谓的“坑”，大都是我们由于对原理不熟悉而自己挖出来的。

		3) 从垃圾回收角度看，Go语言的垃圾回收采用的是并发的标记清除算法（Concurrent Mark and Sweep，CMS）。
		虽然是并发的操作，时间比串行操作短很多，但是还是会在垃圾回收期间停止所有用户程序的操作。
		这一点多少会影响到对实时性要求比较高的应用。不过，在Go语言 1.3 之后的版本中，这方面的问题已经得到了极大的改善。
	*/
}

func lib() {
	/*
		在 Windows 下，标准库的位置在Go语言根目录下的子目录 pkg\windows_amd64 中；
		在 Linux 下，标准库在Go语言根目录下的子目录 pkg\linux_amd64 中（如果是安装的是 32 位，则在 linux_386 目录中）。
		一般情况下，标准包会存放在 $GOROOT/pkg/$GOOS_$GOARCH/ 目录下。

		Go语言标准库常用的包及功能
		Go语言标准库包名	功  能
		bufio	带缓冲的 I/O 操作
		bytes	实现字节操作
		container	封装堆、列表和环形列表等容器
		crypto	加密算法
		database	数据库驱动和接口
		debug	各种调试文件格式访问及调试功能
		encoding	常见算法如 JSON、XML、Base64 等
		flag	命令行解析
		fmt	格式化操作
		go	Go语言的词法、语法树、类型等。可通过这个包进行代码信息提取和修改
		html	HTML 转义及模板系统
		image	常见图形格式的访问及生成
		io	实现 I/O 原始访问接口及访问封装
		math	数学库
		net	网络库，支持 Socket、HTTP、邮件、RPC、SMTP 等
		os	操作系统平台不依赖平台操作封装
		path	兼容各操作系统的路径操作实用函数
		plugin	Go 1.7 加入的插件系统。支持将代码编译为插件，按需加载
		reflect	语言反射支持。可以动态获得代码中的类型信息，获取和修改变量的值
		regexp	正则表达式封装
		runtime	运行时接口
		sort	排序接口
		strings	字符串转换、解析及实用函数
		time	时间接口
		text	文本模板及 Token 词法器
	*/
}

func regular() {
	/*
		Go语言无须解决方案、工程文件和 Make File，
		只要将工程文件按照 GOPATH 的规则进行填充，即可使用 go build/go install 进行编译，
		编译完成的二进制可执行文件统一放在 bin 文件夹下。

		Go语言可以利用自己的特性实现并发编译，并发编译的最小元素是包。
		从 Go 1.9 版本开始，最小并发编译元素缩小到函数，整体编译速度提高了 20%。
	*/

	/*
		去掉循环冗余括号，int 声明被简化为:=
		for a := 0;a<10;a++{
		    // 循环代码
		}

		去掉表达式冗余括号
		if 表达式{
		    // 表达式成立
		}

		Go语言中，左括号必须紧接着语句不换行。其他样式的括号将被视为代码编译错误。

		在Go语言中，自增操作符不再是一个操作符，而是一个语句。
		i++
	*/
}

func runServer() {
	go Server()
	for {
		var q string
		fmt.Print("输入q退出：")
		fmt.Scanf("%s", &q)
		if q == "q" {
			break
		}
	}
}

func compiler() {
	/*
		1) 抽象语法树
		在计算机科学中，抽象语法树（Abstract Syntax Tree，AST），或简称语法树（Syntax tree），是源代码语法结构的一种抽象表示。
		它以树状的形式表现编程语言的语法结构，树上的每个节点都表示源代码中的一种结构。

		之所以说语法是“抽象”的，是因为这里的语法并不会表示出真实语法中出现的每个细节。
		比如，嵌套括号被隐含在树的结构中，并没有以节点的形式呈现。
		而类似于 if else 这样的条件判断语句，可以使用带有两个分支的节点来表示。

		抽象语法树可以应用在很多领域，比如浏览器，智能编辑器，编译器。

		1+3*(4-1)+2
		参见：./AST.gif

		2) 静态单赋值
		静态单赋值形式（static single assignment form，通常简写为 SSA form 或是 SSA）
		是中介码（IR，intermediate representation）的属性，它要求每个变量只分配一次，并且变量需要在使用之前定义。
		x := 1
		x := 2
		y := x

		从上面的描述所知，第一行赋值行为是不需要的，因为 x 在第二行被二度赋值并在第三行被使用，在 SSA 下，将会变成下列的形式：
		x1 := 1
		x2 := 2
		y1 := x2
		从使用 SSA 的中间代码我们就可以非常清晰地看出变量 y1 的值和 x1 是完全没有任何关系的，
		所以在机器码生成时其实就可以省略第一步，这样就能减少需要执行的指令来优化这一段代码。

		根据 Wikipedia（维基百科）对 SSA 的介绍来看，在中间代码中使用 SSA 的特性能够为整个程序实现以下的优化：
		常数传播（constant propagation）
		值域传播（value range propagation）
		稀疏有条件的常数传播（sparse conditional constant propagation）
		消除无用的程式码（dead code elimination）
		全域数值编号（global value numbering）
		消除部分的冗余（partial redundancy elimination）
		强度折减（strength reduction）
		寄存器分配（register allocation）

		因为 SSA 的主要作用就是代码的优化，所以是编译器后端（主要负责目标代码的优化和生成）的一部分。
		当然，除了 SSA 之外代码编译领域还有非常多的中间代码优化方法，
		优化编译器生成的代码是一个非常古老并且复杂的领域，这里就不展开介绍了。

		3) 指令集架构
		最后要介绍的一个预备知识就是指令集架构了，指令集架构（Instruction Set Architecture，简称 ISA），
		又称指令集或指令集体系，是计算机体系结构中与程序设计有关的部分，
		包含了基本数据类型，指令集，寄存器，寻址模式，存储体系，中断，异常处理以及外部 I/O。
		指令集架构包含一系列的 opcode 即操作码（机器语言），以及由特定处理器执行的基本命令。

		指令集架构常见种类如下：
		复杂指令集运算（Complex Instruction Set Computing，简称 CISC）；
		精简指令集运算（Reduced Instruction Set Computing，简称 RISC）；
		显式并行指令集运算（Explicitly Parallel Instruction Computing，简称 EPIC）；
		超长指令字指令集运算（VLIW）。

		不同的处理器（CPU）使用了大不相同的机器语言，所以我们的程序想要在不同的机器上运行，就需要将源代码根据架构编译成不同的机器语言。

		编译原理
		Go语言编译器的源代码在 cmd/compile 目录中，目录下的文件共同构成了Go语言的编译器，
		学过编译原理的人可能听说过编译器的前端和后端，
		编译器的前端一般承担着词法分析、语法分析、类型检查和中间代码生成几部分工作，
		而编译器后端主要负责目标代码的生成和优化，也就是将中间代码翻译成目标机器能够运行的机器码。
		compile.gif

		Go的编译器在逻辑上可以被分成四个阶段：
		词法与语法分析、类型检查和 AST 转换、通用 SSA 生成和最后的机器代码生成，
		下面我们来分别介绍一下这四个阶段做的工作。

		1) 词法与语法分析
		所有的编译过程其实都是从解析代码的源文件开始的，词法分析的作用就是解析源代码文件，
		它将文件中的字符串序列转换成 Token 序列，方便后面的处理和解析，我们一般会把执行词法分析的程序称为词法解析器（lexer）。

		而语法分析的输入就是词法分析器输出的 Token 序列，这些序列会按照顺序被语法分析器进行解析，
		语法的解析过程就是将词法分析生成的 Token 按照语言定义好的文法（Grammar）自下而上或者自上而下的进行规约，
		每一个 Go 的源代码文件最终会被归纳成一个 SourceFile 结构：
		SourceFile = PackageClause ";" { ImportDecl ";" } { TopLevelDecl ";" }

		标准的 Golang 语法解析器使用的就是 LALR(1) 的文法，
		语法解析的结果其实就是上面介绍过的抽象语法树（AST），
		每一个 AST 都对应着一个单独的Go语言文件，这个抽象语法树中包括当前文件属于的包名、定义的常量、结构体和函数等。

		如果在语法解析的过程中发生了任何语法错误，都会被语法解析器发现并将消息打印到标准输出上，整个编译过程也会随着错误的出现而被中止。

		2) 类型检查
		当拿到一组文件的抽象语法树 AST 之后，Go语言的编译器会对语法树中定义和使用的类型进行检查，
		类型检查分别会按照顺序对不同类型的节点进行验证，按照以下的顺序进行处理：

		常量、类型和函数名及类型；
		变量的赋值和初始化；
		函数和闭包的主体；
		哈希键值对的类型；
		导入函数体；
		外部的声明；

		通过对每一棵抽象节点树的遍历，我们在每一个节点上都会对当前子树的类型进行验证保证当前节点上不会出现类型错误的问题，
		所有的类型错误和不匹配都会在这一个阶段被发现和暴露出来。

		类型检查的阶段不止会对树状结构的节点进行验证，同时也会对一些内建的函数进行展开和改写，
		例如 make 关键字在这个阶段会根据子树的结构被替换成 makeslice 或者 makechan 等函数。

		其实类型检查不止对类型进行了验证工作，还对 AST 进行了改写以及处理Go语言内置的关键字，
		所以，这一过程在整个编译流程中是非常重要的，没有这个步骤很多关键字其实就没有办法工作。

		3) 中间代码生成
		当我们将源文件转换成了抽象语法树，对整个语法树的语法进行解析并进行类型检查之后，
		就可以认为当前文件中的代码基本上不存在无法编译或者语法错误的问题了，
		Go语言的编译器就会将输入的 AST 转换成中间代码。

		Go语言编译器的中间代码使用了 SSA(Static Single Assignment Form) 的特性，
		如果我们在中间代码生成的过程中使用这种特性，就能够比较容易的分析出代码中的无用变量和片段并对代码进行优化。

		在类型检查之后，就会通过一个名为 compileFunctions 的函数开始对整个Go语言项目中的全部函数进行编译，
		这些函数会在一个编译队列中等待几个后端工作协程的消费，这些 Goroutine 会将所有函数对应的 AST 转换成使用 SSA 特性的中间代码。

		4) 机器码生成
		Go语言源代码的 cmd/compile/internal 目录中包含了非常多机器码生成相关的包，
		不同类型的 CPU 分别使用了不同的包进行生成 amd64、arm、arm64、mips、mips64、ppc64、s390x、x86 和 wasm，
		也就是说Go语言能够在几乎全部常见的 CPU 指令集类型上运行。

		编译器入口
		Go语言的编译器入口是 src/cmd/compile/internal/gc 包中的 main.go 文件，
		这个 600 多行的 Main 函数就是Go语言编译器的主程序，这个函数会先获取命令行传入的参数并更新编译的选项和配置，
		随后就会开始运行 parseFiles 函数对输入的所有文件进行词法与语法分析得到文件对应的抽象语法树：

		func Main(archInit func(*Arch)) {
		    // ...

		    lines := parseFiles(flag.Args())

		接下来就会分九个阶段对抽象语法树进行更新和编译，就像我们在上面介绍的，整个过程会经历类型检查、SSA 中间代码生成以及机器码生成三个部分：
		检查常量、类型和函数的类型；
		处理变量的赋值；
		对函数的主体进行类型检查；
		决定如何捕获变量；
		检查内联函数的类型；
		进行逃逸分析；
		将闭包的主体转换成引用的捕获变量；
		编译顶层函数；
		检查外部依赖的声明；

		了解了剩下的编译过程之后，我们重新回到词法和语法分析后的具体流程，
		在这里编译器会对生成语法树中的节点执行类型检查，
		除了常量、类型和函数这些顶层声明之外，它还会对变量的赋值语句、函数主体等结构进行检查：
		for i := 0; i < len(xtop); i++ {
		    n := xtop[i]
		    if op := n.Op; op != ODCL && op != OAS && op != OAS2 && (op != ODCLTYPE || !n.Left.Name.Param.Alias) {
		        xtop[i] = typecheck(n, ctxStmt)
		    }
		}
		for i := 0; i < len(xtop); i++ {
		    n := xtop[i]
		    if op := n.Op; op == ODCL || op == OAS || op == OAS2 || op == ODCLTYPE && n.Left.Name.Param.Alias {
		        xtop[i] = typecheck(n, ctxStmt)
		    }
		}
		for i := 0; i < len(xtop); i++ {
		    n := xtop[i]
		    if op := n.Op; op == ODCLFUNC || op == OCLOSURE {
		        typecheckslice(Curfn.Nbody.Slice(), ctxStmt)
		    }
		}
		checkMapKeys()
		for _, n := range xtop {
		    if n.Op == ODCLFUNC && n.Func.Closure != nil {
		        capturevars(n)
		    }
		}
		escapes(xtop)
		for _, n := range xtop {
		    if n.Op == ODCLFUNC && n.Func.Closure != nil {
		        transformclosure(n)
		    }
		}

		类型检查会对传入节点的子节点进行遍历，这个过程会对 make 等关键字进行展开和重写，
		类型检查结束之后并没有输出新的数据结构，只是改变了语法树中的一些节点，
		同时这个过程的结束也意味着源代码中已经不存在语法错误和类型错误，中间代码和机器码也都可以正常的生成了。
		   initssaconfig()
		    peekitabs()
		    for i := 0; i < len(xtop); i++ {
		        n := xtop[i]
		        if n.Op == ODCLFUNC {
		            funccompile(n)
		        }
		    }
		    compileFunctions()
		    for i, n := range externdcl {
		        if n.Op == ONAME {
		            externdcl[i] = typecheck(externdcl[i], ctxExpr)
		        }
		    }
		    checkMapKeys()
		}

		在主程序运行的最后，会将顶层的函数编译成中间代码并根据目标的 CPU 架构生成机器码，
		不过这里其实也可能会再次对外部依赖进行类型检查以验证正确性。

		总结
		Go语言的编译过程其实是非常有趣并且值得学习的，
		通过对Go语言四个编译阶段的分析和对编译器主函数的梳理，我们能够对 Golang 的实现有一些基本的理解，
		掌握编译的过程之后，Go语言对于我们来讲也不再那么神秘，所以学习其编译原理的过程还是非常有必要的。
	*/
}

func install() {
	/*
		D:/Go 目录
		这个目录的结构遵守 GOPATH 规则，后面的章节会提到这个概念。目录中各个文件夹的含义如下表所示。

		Go 开发包的安装目录的功能及说明
		目录名	说明
		api	每个版本的 api 变更差异
		bin	go 源码包编译出的编译器（go）、文档工具（godoc）、格式化工具（gofmt）
		doc	英文版的 Go 文档
		lib	引用的一些库文件
		misc	杂项用途的文件，例如 Android 平台的编译、git 的提交钩子等
		pkg	Windows 平台编译好的中间文件
		src	标准库的源码
		test	测试用例
	*/
}

func structure() {
	/*
		我们前面讲搭建Go语言开发环境时提到的环境变量 GOPATH，项目的构建主要是靠它来实现的。
		这么说吧，如果想要构建一个项目，就需要将这个项目的目录添加到 GOPATH 中，多个项目之间可以使用;分隔。

		如果不配置 GOPATH，即使处于同一目录，代码之间也无法通过绝对路径相互调用。
		目录结构
		一个Go语言项目的目录一般包含以下三个子目录：
		src 目录：放置项目和库的源文件；
		pkg 目录：放置编译后生成的包/库的归档文件；
		bin 目录：放置编译后生成的可执行文件。

		三个目录中我们需要重点关注的是 src 目录，其他两个目录了解即可，下面来分别介绍一下这三个目录。
		src 目录
		用于以包（package）的形式组织并存放 Go 源文件，这里的包与 src 下的每个子目录是一一对应。
		例如，若一个源文件被声明属于 log 包，那么它就应当保存在 src/log 目录中。

		并不是说 src 目录下不能存放 Go 源文件，一般在测试或演示的时候也可以把 Go 源文件直接放在 src 目录下，
		但是这么做的话就只能声明该源文件属于 main 包了。正常开发中还是建议大家把 Go 源文件放入特定的目录中。

		包是Go语言管理代码的重要机制，其作用类似于Java中的 package 和 C/C++ 的头文件。
		Go 源文件中第一段有效代码必须是package <包名> 的形式，如 package hello。

		另外需要注意的是，Go语言会把通过go get 命令获取到的库源文件下载到 src 目录下对应的文件夹当中。

		pkg 目录
		用于存放通过go install 命令安装某个包后的归档文件。归档文件是指那些名称以“.a”结尾的文件。

		该目录与 GOROOT 目录（也就是Go语言的安装目录）下的 pkg 目录功能类似，区别在于这里的 pkg 目录专门用来存放项目代码的归档文件。

		编译和安装项目代码的过程一般会以代码包为单位进行，
		比如 log 包被编译安装后，将生成一个名为 log.a 的归档文件，并存放在当前项目的 pkg 目录下。

		bin 目录
		与 pkg 目录类似，在通过go install 命令完成安装后，保存由 Go 命令源文件生成的可执行文件。
		在类 Unix 操作系统下，这个可执行文件的名称与命令源文件的文件名相同。
		而在 Windows 操作系统下，这个可执行文件的名称则是命令源文件的文件名加 .exe 后缀。

		上面我们提到了命令源文件和库源文件，它们到底是什么呢？
		命令源文件：如果一个 Go 源文件被声明属于 main 包，并且该文件中包含 main 函数，则它就是命令源码文件。
		命令源文件属于程序的入口，可以通过Go语言的go run 命令运行或者通过go build 命令生成可执行文件。
		库源文件：库源文件则是指存在于某个包中的普通源文件，并且库源文件中不包含 main 函数。

		不管是命令源文件还是库源文件，在同一个目录下的所有源文件，其所属包的名称必须一致的。
	*/
}

func dependence() {
	/*
	godep 是一个Go语言官方提供的通过 vender 模式来管理第三方依赖的工具，类似的还有由社区维护的准官方包管理工具 dep。

	安装godep工具
	我们可以通过go get 命令来获取 godep 工具。
	go get github.com/tools/godep

	命令执行成功后会将 godep 工具的源码下载到 D:\GOPATH\pkg\mod\github.com\tools 的 src 目录下对应的文件夹中，
	同时还会在 GOPATH 的 bin 目录下生成一个名为 godep.exe 的可执行文件，如下图所示。

	为了方便使用 godep 工具，我们需要将存放 godep.exe 文件的目录添加到环境变量 PATH 中。

	godep工具的基本命令
	完成上面的操作后，我们就可以在命令行窗口（CMD）中使用 godep 工具了，godep 支持的命令如下表所示：

	命令	作用
	godep save	将依赖包的信息保存到 Godeps.json 文件中
	godep go	使用保存的依赖项运行 go 工具
	godep get	下载并安装指定的包
	godep path	打印依赖的 GOPATH 路径
	godep restore	在 GOPATH 中拉取依赖的版本
	godep update	更新选定的包或 go 版本
	godep diff	显示当前和以前保存的依赖项集之间的差异
	godep version	查看版本信息

	使用godep help [命令名称]可以查看命令的帮助信息
	 */
}

// 生产者与消费者
func producerAndCustomerDemo() {
	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建producer()函数的并发goroutine
	go producer("cat", channel)
	go producer("dog", channel)
	// 数据消费函数
	customer(channel)
}

// producer 数据生产者
func producer(header string, channel chan<- string) {
	// 无限循环，不停地生产数据
	for {
		// 将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		// 等待1秒
		time.Sleep(time.Second)
	}
}

// 数据消费者
func customer(channel <-chan string) {
	// 不停地获取数据
	for {
		// 从通道中取出数据, 此处会阻塞直到信道中返回数据
		message := <-channel
		// 打印数据
		fmt.Println(message)
	}
}
