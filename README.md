# auto

## 工具介绍

本工具采用 Go开发，集成子域名收集、端口扫描、服务提取、URL 验活去重、指纹识别、WAF 识别、漏洞扫描等功能，只需要输入根域名即可全方位收集相关资产，并检测漏洞。本项目借用多种优秀工具，实现赏金自动化和减轻日常使用工具手工处理的繁琐，懒人必备。

## 调用工具

![Auto](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/Auto.png)



## 执行流程

![赏金自动化](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/%E8%B5%8F%E9%87%91%E8%87%AA%E5%8A%A8%E5%8C%96.png)

## 工具说明

> 本工具仅支持 Linux 使用，测试环境 Debian 10

本工具执行的命令通过 `config.yaml` 获取，需要修改的地方是 `crawlergo` 那里的 `chrome` 路径，还有协程之类的。Server 酱不想使用的话可以不填，主要用其通知用户信息收集/漏扫结束/工具执行报错，还是很有必要的。

所使用的工具都应该在可编译后的可执行文件同目录下的 `tools` 目录下，`tools` 目录结构如下，结合着 `config.yaml` 就很容易理解，这里不提供 `xray` 和 `xscan` ，`xray` 可以自己找，`xscan` 可以加入 "Hacking自动化就是好玩" 知识星球获取。如果不使用 `xscan` 的话修改 `config.yaml` 即可。

其他工具的删减同理。

> `tools` 目录结构

```
.
├── other
│   ├── httpx
│   ├── TideFinger
│   └── wafw00f-2.2.0
│       ├── setup.py
│       ├── wafw00f
│       │   ├── main.py
├── subdomain
│   ├── alterx
│   ├── ksubdomain
│   ├── LICENSE
│   ├── OneForAll
│   │   ├── oneforall.py
│   └── subfinder
├── template
│   └── template.html
└── vuln
    ├── crawlergo
    ├── nuclei
    │   └── nuclei
    ├── xray
    │   ├── xray
    └── xscan
        └── xscan
```

工具命令：

```
Commands:
  collector   信息收集
    --target <target>  单个目标
    --targets <file>   目标文件
    --alterx <boolean> 是否使用 alterx 生成子域名列表 ( 默认关闭 )
  scanner     漏洞扫描
    --target <target>  单个目标
    --targets <file>   目标文件
  security    一条龙 ( 信息收集 + 漏洞扫描 )
    --target <target>  单个目标
    --targets <file>   目标文件
    --alterx <boolean> 是否使用 alterx 生成子域名列表 ( 默认关闭 )
```

比如下面的：

```
./auto security --targets targets.txt --alterx=true
```

当使用  `alterx` 工具时，会造成生成大量的子域名，所以会导致后面的执行速度过慢，不过不缺时间的话还是加上。

使用效果如下：

![image-20230729210650455](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/image-20230729210650455.png)

![image-20230729210716837](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/image-20230729210716837.png)

端口扫描部分向上面的 `ssh` 这种都是直接根据 `nmap` 结果动态生成的，其结果就保存在对应的文件中，如下：

![image-20230729210913072](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/image-20230729210913072.png)

![image-20230729211017137](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/image-20230729211017137.png)

内容是 `IP:PORT` 的形式，可以直接使用 `hydra` 进行爆破。

## 工具使用

信息收集：这里推荐把 `OneForAll、Subfinder` 的 api 填写完全，`fofa` 这种网络空间引擎的 key 可以在 `github` 找泄露或者咸鱼上买，免费的都尽量给配上。`ksubdomain` 可以修改 `config.yaml` 替换成 `oneforall` 提供的大字典，然后再爆破多级域名。`alterx` 这种根据已知子域生成子域名列表的工具也可以加上，有可能发现一些正常情况发现不了的资产。

漏洞扫描：这里是使用了 3 款漏扫工具和一个款爬虫工具，`nuclei、xscan` 的结果都是解析为 `xray` 的 html 格式，方便观看，`nuclei` 由于扫出的东西比较多，但没多大用，这里是生成 `nuclei-漏洞等级` 的 html，方便看重点。`xray + crawlergo`  这里是自己写的，没有加那个更换请求头的功能，但是可以在 `xray` 扫描结束后 `kill` 掉 `xray` 的监听进程，比较方便。漏扫工具

## 工具安装



