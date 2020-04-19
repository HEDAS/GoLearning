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
