Go进阶笔记-工程化



## 工程化项目目录

https://github.com/golang-standards/project-layout 规范写的，这个规范虽然不是官方强制规范的，但是确实很多开源项目都在采用的，所以我们在生产中正式的项目都应该尽可能遵循这个目录规范标准进行代码的编写。关于这个目录的规范使用，自己会在后续实际使用中逐渐完善。


**cmd**

每个应用程序的目录名，都应该与你想要的可执行文件的名称想匹配，如，我们的项目中可能存在多个app，那么在cmd目录结构中如下：
```
./cmd
├── myapp1
│   ├── main.go
│   └── myapp1
└── myapp2
    ├── main.go
    └── myaap2
```



注意： 不要在这个目录中放置太多代码。如果你认为代码可以导入并在其他项目中使用，那么它应该位于 /pkg 目录中。如果代码不是可重用的，或者你不希望其他人重用它，请将该代码放到 /internal 目录中。

**internal**

私有应用程序和库代码，这个是你不希望其他人在其应用程序或库中导入代码。

如果存在多个app，那么在internal目录下其实也是与之相对应的如下图所示：

```
./internal
├── myapp1
│   ├── biz
│   ├── data
│   └── service
└── myapp2
    ├── biz
    ├── data
    └── service
```

当然如果是单一的服务，可以直接在internal目录下添加biz, data, service， 而不用再简历app目录


**pkg**

外部应用程序可以使用的库代码(例如 /pkg/mypubliclib)。其他项目会导入这些库
`/internla/pkg` 一般用于项目内的 跨多个应用的公共共享代码，但其作用域仅在单个项目工程内

### Kit Project Layout

基础库 kit 为独立项目，公司级建议只有一个，按照功能目录来拆分会带来不少的管理工作，因此建议合并整合。
Kit 项目必须具备的特点：
- 统一
- 标准库方式布局
- 高度抽象
- 支持插件

[![rugI74.png](https://s3.ax1x.com/2020/12/14/rugI74.png)](https://imgchr.com/i/rugI74)



### Service Application Project Layout

[![ru2081.png](https://s3.ax1x.com/2020/12/14/ru2081.png)](https://imgchr.com/i/ru2081)

**api**
API 协议定义目录，xxapi.proto protobuf 文件，以及生成的 go 文件。我们通常把 api 文档直接在 proto 文件中描述

**configs**
配置文件模板或默认配置

**test**
额外的外部测试应用程序和测试数据。你可以随时根据需求构造 /test 目录。对于较大的项目，有一个数据子目录是有意义的。例如，你可以使用 /test/data 或 /test/testdata (如果你需要忽略目录中的内容)。请注意，Go 还会忽略以“.”或“_”开头的目录或文件，因此在如何命名测试数据目录方面有更大的灵活性。


一个 gitlab 的 project 里可以放置多个微服务的app(类似 monorepo)。也可以按照 gitlab 的 group 里建立多个 project，每个 project 对应一个 app。

一个 gitlab 的 project 里可以放置多个微服务的app(类似 monorepo)。也可以按照 gitlab 的 group 里建立多个 project，每个 project 对应一个 app。

和 app 平级的目录 pkg 存放业务有关的公共库（非基础框架库）。如果应用不希望导出这些目录，可以放置到 myapp/internal/pkg 中。

微服务中的 app 服务类型分为4类：interface、service、job、admin。
- interface: 对外的 BFF 服务，接受来自用户的请求，比如暴露了 HTTP/gRPC 接口。
- service: 对内的微服务，仅接受来自内部其他服务或者网关的请求，比如暴露了gRPC 接口只对内服务。
- admin：区别于 service，更多是面向运营测的服务，通常数据权限更高，隔离带来更好的代码级别安全。
- job: 流式任务处理的服务，上游一般依赖 message broker。
- task: 定时任务，类似 cronjob，部署到 task 托管平台中。

cmd 应用目录负责程序的: 启动、关闭、配置初始化等.


**常见的三层架构**

model: 放对应“存储层”的结构体，是对存储的一一映射
dao: 数据读写层，数据库和缓存全部在这层统一处理，包括 cache miss 处理。
service: 组合各种数据访问来构建业务逻辑。
server: 依赖 proto 定义的服务作为入参，提供快捷的启动服务全局方法。
api: 定义了 API proto 文件，和生成的 stub 代码，它生成的 interface，其实现者在 service 中。

service 的方法签名因为实现了 API 的 接口定义，DTO 直接在业务逻辑层直接使用了，更有 dao 直接使用，最简化代码

`DTO(Data Transfer Object)`：数据传输对象，泛指用于展示层API 层与服务层(业务逻辑层)之间的数据传输对象。
其实这个概念指的就是我们要专门针对API层定义数据结构，不能让API层依赖于Model层的数据结构定义。

`DO(Domain Object)`: 领域对象，就是从现实世界中抽象出来的有形或无形的业务实体。缺乏 `DTO` -> `DO` 的对象转换

这里引入了DO 即领域对象，即我们其实可以稍微调整一下三层架构的目录结构：

app 目录下有 api、cmd、configs、internal 目录，目录里一般还会放置 README、CHANGELOG、OWNERS。

- internal: 是为了避免有同业务下有人跨目录引用了内部的 biz、data、service 等内部 struct。
- biz: 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，repo 接口在这里定义，使用依赖倒置的原则。 关于DO对象应该在biz 这一层进行定义，持久化的interface也是应该这一层定义。 注意：业务逻辑层依赖持久化层，但是这里不应该依赖具体的实现，而应该依赖于一个抽象，这就是依赖倒置
- data: 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。注意：关于数据持久化的接口实在biz层定义，但是具体的实现实在data层做。
- service: 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑。注意：service层关注的是grpc/http的实现类，这里主要做的就是DTO到DO的转换，这部分代码不应该有非常复杂的业务逻辑，应该只有简单的编排逻辑。
- PO(Persistent Object): 持久化对象，它跟持久层（通常是关系型数据库）的数据结构形成一一对应的映射关系，如果持久层是关系型数据库，那么数据表中的每个字段（或若干个）就对应 PO 的一个（或若干个）属性。

[![ruokmq.png](https://s3.ax1x.com/2020/12/14/ruokmq.png)](https://imgchr.com/i/ruokmq)


### Lifecycle 生命周期

Lifecycle 需要考虑服务应用的对象初始化以及生命周期的管理，所有 HTTP/gRPC 依赖的前置资源初始化，包括 data、biz、service，之后再启动监听服务。可以使用 https://github.com/google/wire ，来管理所有资源的依赖注入

核心是为了：
- 方便测试
- 单次初始化和复用


## API 设计

之前说过通过gRPC做接口的好处：
- 服务而非对象、消息而非引用：促进微服务的系统间粗粒度消息交互设计理念。
- 负载无关的：不同的服务需要使用不同的消息类型和编码，例如 protocol buffers、JSON、XML和Thrift。
- 流: Streaming API。
- 阻塞式和非阻塞式：支持异步和同步处理在客户端和服务端间交互的消息序列。
- 元数据交换：常见的横切关注点，如认证或跟踪，依赖数据交换。
- 标准化状态码：客户端通常以有限的方式响应 API 调用返回的错误。

而之前的目录规范也是推荐将proto 文件放在api目录下


项目目录下的proto会通过gitlab 的ci/cd 添加hook 将项目的proto自动化推送到这个统一的apis目录下。这样也不会暴露源码，同时各个新项目的api都可以同步到这个目录，这样如果需要调用彼此的接口只需要取这个目录下找就可以了。
同时在这个统一的apis项目里可通过ci/cd 自动构建各个语言的客户端和服务端代码，这样对于使用也是非常方便。

为了统一检索和规范 API，内部可以建立了一个统一的 apis 仓库，整合所有对内对外 API。
- API 仓库，方便跨部门协作。
- 版本管理，基于 git 控制。
- 规范化检查，API lint。
- API design review，变更 diff。
- 权限管理，目录 OWNERS。

### APIs Porject Layout

同一个apis目录规范如下：

[![rMz8ED.png](https://s3.ax1x.com/2020/12/15/rMz8ED.png)](https://imgchr.com/i/rMz8ED)


### API 命名
包名为应用的标识(APP_ID)，用于生成 gRPC 请求路径，或者 proto 之间进行引用 Message。文件中声明的包名称应该与产品和服务名称保持一致。带有版本的 API 的软件包名称必须以此版本结尾。

my.package.v1，为 API 目录，定义service相关接口，用于提供业务使用。

package <package_name>.<version>;

[![rQpK0K.png](https://s3.ax1x.com/2020/12/15/rQpK0K.png)](https://imgchr.com/i/rQpK0K)

[![rQp3fH.png](https://s3.ax1x.com/2020/12/15/rQp3fH.png)](https://imgchr.com/i/rQp3fH)


gRPC 默认使用 Protobuf v3 格式，因为去除了 required 和 optional 关键字，默认全部都是 optional 字段。如果没有赋值的字段，默认会基础类型字段的默认值，比如 0 或者 “”。

Protobuf v3 中，建议使用：https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/wrappers.proto
Warpper 类型的字段，即包装一个 message，使用时变为指针。


### API Errors

如果您的 API 服务依赖于其他服务，则不应盲目地将这些服务的错误传播到您的客户端。在翻译错误时，我们建议执行以下操作：
- 隐藏实现详细信息和机密信息。
- 调整负责该错误的一方。例如，从另一个服务接收 INVALID_ARGUMENT 错误的服务器应该将 INTERNAL 传播给它自己的调用者。

全局错误码，是松散、易被破坏契约的，基于我们上述讨论的，在每个服务传播错误的时候，做一次翻译，这样保证**每个服务 + 错误枚举**，应该是唯一的，而且在 proto 定义中是可以写出来文档的。


## 配置管理


**环境变量(配置)**

Region、Zone、Cluster、Environment、Color、Discovery、AppID、Host，等之类的环境信息，都是通过在线运行时平台打入到容器或者物理机，供 kit 库读取使用。
    
**静态配置**
资源需要初始化的配置信息，比如 http/gRPC server、redis、mysql 等，这类资源在线变更配置的风险非常大，我通常不鼓励 on-the-fly 变更，很可能会导致业务出现不可预期的事故，变更静态配置和发布 bianry app 没有区别，应该走一次迭代发布的流程。
    
**动态配置**
应用程序可能需要一些在线的开关，来控制业务的一些简单策略，会频繁的调整和使用，我们把这类是基础类型(int, bool)等配置，用于可以动态变更业务流的收归一起，同时可以考虑结合类似 https://pkg.go.dev/expvar 来结合使用。
    
**全局配置**
通常，我们依赖的各类组件、中间件都有大量的默认配置或者指定配置，在各个项目里大量拷贝复制，容易出现意外，所以我们使用全局配置模板来定制化常用的组件，然后再特化的应用里进行局部替换。


这里有一个思路，是通过pb文件定义配置文件