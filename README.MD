
Ttsoa
=======
distributed API services all in one, include provider, consumer, monitor.
provider use golang, moniter use PHP, the services info will store in mysql database.
include some consumer test code base on PHP.

一个用golang和PHP开发的SOA基型架构， 包括供应者， 消费者，注册中心


特性
======
* 供应者容器用golang的协程实现
* 供应者服务支持多种语言（服务端）， 消费者也支持多种语言（客户端）
* 提供服务接口的信息展示，分析等
* 支持本地容器接口和远程 HTTP/SOCKET 接口
* 支持 负载均衡， 接口缓存等功能
* 服务API可分布在不同的服务器上，实现高可用


容器安装/运行
======
    cd ./ttsoa/provider;
	go build;
	./provider

目前在端口 7777 运行容器服务



未完待续......


[更多联系邮箱](82049406@qq.com)


















