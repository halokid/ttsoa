
TTsoa
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


提供者容器安装/运行
======
    cd ./ttsoa/provider;
	go build;
	./provider

目前在端口 7777 运行容器服务


Quick Start 快速开始
======

###### 目录说明 ####


├─consumer　　　　　消费者程序
│  └─php　　　　　　　PHP范例
├─monitor　　　　　　服务信息展示中心
├─provider　　　　　　提供者容器程序
│  └─services　　　　　提供者服务程序（多语言支持），包含PHP和java范例
└─test　　　　　　　 开发过程中的一些测试代码


###### 提供者程序编写/部署 #######

提供者容器运行之后， 会在服务器的 7777 端口监听服务，假如我们要写一个PHP的提供者程序

在 provider/services/php 目录之下， java的程序也按照这种代码部署形式， 比如我们要执行User类下的ulist方式，编写类文件放在PHP目录下

    
    namespace TTsoa;
    
    class User
    {
      public function ulist() {
        echo 'User list';
      }
    }
    
    

保证提供者程序在部署的服务器 可以用 `php xx.php` 的形式运行， 提供者容器最终是请求服务器的CLI的，就等于CGI程序的原理一样， PHP要加入系统的env里面，能直接运行

**java的范例具体看代码**



###### 消费者程序编写/部署 ######

TTSOA把负载均衡， 高可用这些逻辑处理封装在消费者端， 所以消费者端分两部分
1.  消费者端的功能封装逻辑（ 负责负载均衡， 权限访问，跟monitor中心通信等逻辑），每一种语言的消费者端封装都包括这些功能逻辑的实现

2.  消费者请求提供者的服务逻辑（包括信息的加密，加速等）

具体部署步骤为:
- 定义服务接口文件， json格式说明


      {
    "user_php":{                     //程序用的自定义服务名称
      "name":"用户服务",              //服务名称
      "engine":"php",                //服务语言引擎
      "hosts":["127.0.0.1:7777","192.168.0.203:7777"],     //服务接口的可用服务器列表
      "host_items":{                        //每个服务器提供的详细服务，可分布式
        "127.0.0.1:7777":["ulist","detail"],
        "192.168.0.203:7777":["ulist","detail"]
      },
      "all_items":["ulist","detail","reg","login","logout"]    //该服务的所有接口列表
    },



假如是PHP的消费者端，代码部署在    ./consumer/php/ 下面
    
    <?php
    
    require_once 'inc/ttConsumer.php';    //load packing file
    
    $jStr = file_get_contents('./consumer.json');    //get config json
    $serviceList = json_decode($jStr, true);
    $userSvs = $serviceList['user_java'];      //load user services config, java sample
    
    $cs = new TTsoa\Consumer($userSvs['hosts']);
    
    //java provider sample
    $res = $cs->setEngine($userSvs['engine'])->balanceServ('random')->getService('/TTsoa/UserImpl');
    
    echo $res;
    
    

整个部署过程这样就已经快速完成了，目前关于代码位置的设置，后期可考虑再灵活一下


###### 性能测试 ######

待完善.....


即将编写的功能
=========

- 三种负载均衡的算法，随机，最小访问，权重值
- 各种语言的消费者端封装程序的完成（目前只做了一下PHP的，可用而已）
- monitor程序编写（打算用CI3.X + bootstrap）
- 性能优化， 提供者容器， 消费者访问接口等关键功能点


项目是会实际生产环境持续应用，持续优化。。。。,有兴趣的同学可以指导指导,多多参与，多谢。



更多联系邮箱82049406@qq.com


















