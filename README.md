# auto

## 工具介绍

本工具采用 Go开发，集成子域名收集、端口扫描、服务提取、URL 验活去重、指纹识别、WAF 识别、漏洞扫描等功能，只需要输入根域名即可完成初步的信息收集，并检测漏洞。本项目借用多种优秀工具，实现赏金自动化和减轻日常使用工具手工处理的繁琐，懒人必备。

## 调用工具

![Auto](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/Auto.png)



## 执行流程

![赏金自动化](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/%E8%B5%8F%E9%87%91%E8%87%AA%E5%8A%A8%E5%8C%96.png)

## 使用效果

> 本工具仅支持 Linux 使用，测试环境 Debian 10

![ESNvTsXuUC](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/ESNvTsXuUC.png)

## 工具下载

123盘：123盘：https://www.123pan.com/s/lNI5Vv-AK3tv.html 提取码:j11X

压缩包内存放有信息收集的相关工具，给权限，安装依赖过后执行 `./auto` 即可。

## 工具使用

信息收集：这里推荐把 `OneForAll、Subfinder` 的 api 填写完全，`fofa` 这种要会员的的 key 可以在 `github` 找泄露或者咸鱼上买，免费的都尽量给配上。`ksubdomain` 可以修改 `config.yaml` 替换成 `oneforall` 提供的大字典，然后再爆破多级域名。`alterx` 这种根据已知子域生成子域名列表的工具也可以加上，说不了发现一些正常情况发现不了的资产。

漏洞扫描：这里是使用了 3 款漏扫工具和一个款爬虫工具，`nuclei、xscan` 的结果都是解析为 `xray` 的 html 格式，方便观看，`nuclei` 由于扫出的东西比较多，但没多大用，这里是生成 `nuclei-漏洞等级` 的 html，方便看重点。`xray + crawlergo`  这里是自己写的，没有加那个更换请求头的功能，但是可以在 `xray` 扫描结束后 `kill` 掉 `xray` 的监听进程，比较方便。

因为通常情况下执行时间较长，配置一下 server 酱，

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

## 依赖工具

`tools` 目录结构：

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

config.yaml 里面就是工具参数和漏扫工具的启用。

注意工具需要和 config.yaml 匹配，自行修改。



### 信息收集

信息收集工具这里会在压缩包里面直接提供，需要安装的只有 `masscan、nmap` ，下面是压缩包中的工具来源。解压后需要手动把依赖安装好，然后自行测试工具是否能够正常运行。

#### 子域名收集

##### oneforall

项目地址：https://github.com/shmilylty/OneForAll

依赖安装：https://github.com/shmilylty/OneForAll/blob/master/docs/installation_dependency.md

> Debian

```shell
sudo apt update
sudo apt install git python3-pip -y
cd OneForAll/
sudo apt install python3-dev python3-pip python3-testresources -y
sudo python3 -m pip install -U pip setuptools wheel -i https://mirrors.aliyun.com/pypi/simple/
sudo pip3 install --ignore-installed -r requirements.txt -i https://mirrors.aliyun.com/pypi/simple/
python3 oneforall.py --help
```

##### subfinder

项目地址：https://github.com/projectdiscovery/subfinder

##### ksubdomain

项目地址：https://github.com/knownsec/ksubdomain

#### 端口扫描

这里 debian 系列的可以直接使用 `apt` 安装：

```
apt install masscan -y
apt insstall nmap -y
```

#### 其他工具

##### httpx

项目地址：https://github.com/projectdiscovery/httpx

```
go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest
```

这里使用了 `-fr` 参数，遵循 `http` 重定向，所有 30x 跳转的网站会直接跳转。

##### TideFinger

go 版本的比较好用，可以通过 `Tide安全团队` 公众号获取。

##### wafw00f

项目地址：https://github.com/EnableSecurity/wafw00f

```
python setup.py install
```

好像是 python2 环境 ？

也可以直接 `apt` 安装：

```
apt install wafw00f -y
```

这样就需要修改一些配置文件。

### 漏洞扫描

漏洞扫描这里的工具不提供，需要自己找，然后放到相同的目录结构即可。

```
`--  vuln
    ├── crawlergo
    ├── nuclei
    │   └── nuclei
    ├── xray
    │   ├── xray
    └── xscan
        └── xscan
```

#### crawlergo

项目地址：https://github.com/Qianlitp/crawlergo

`crawlergo` 直接下载可执行文件即可，然后就是安装 `chromium` 内核。

这里下载 `chrome` 的 deb 包安装即可：

```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo dpkg -i google-chrome-stable_current_amd64.deb
```

然后找一下路径：

```
which google-chrome
```

![image-20230730102152014](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/image-20230730102152014.png)

修改 `config.yaml`  中 `-c` 参数后面的路径为上面 `which` 的路径即可。

#### xray

项目地址：https://github.com/chaitin/xray

#### nuclei

项目地址：https://github.com/projectdiscovery/nuclei

#### xscan

扫 xss 的，很强。

https://mp.weixin.qq.com/s/A6Kjej2pfcCjuY7qey5irw

可加入 "Hacking自动化就是好玩" 知识星球获取

### server酱

https://sct.ftqq.com/sendkey

绑定微信，获取 key，关注 "方糖" 公众号即可。

## 结果说明

运行结束后会生成 `results` 目录，结构如下：

```shell
.
|-- httpx
|   |-- urlInfo.json			# httpx 运行结果
|   `-- urls.txt				# 验活去重后的 url
|-- services
|   |-- EtherNetIP-1.txt		# nmap 指纹识别的结果,文件名就是对应服务
|   |-- http.txt
|   `-- urls.txt				# 拼接获取的 url, 这里的 url 是过度阶段, 不要使用
|-- subdomains
|   `-- subdomains.txt			# 收集到的子域名
|-- vuln						# 漏洞扫描结果
|   |-- nuclei-info.html		# nuclei info 级别的漏洞信息, html 样式为 xray 的
|   |-- nuclei-low.html
|   |-- nuclei-medium.html
|   |-- tools					# nuclei, xscan 漏洞扫描的 json 结果
|   |   |-- nuclei.json
|   |   `-- xscan.json
|   |-- xray.html				# xray 漏扫结果
|   `-- xscan.html				# xscan 漏扫结果
`-- webinfo						# 信息收集结果
    |-- webinfo.csv				
    `-- webinfo.json
```

`results` 这里有很多过度阶段的结果，需要关注的就只有这些：

```shell
|-- services
|   |-- EtherNetIP-1.txt		# nmap 指纹识别的结果,文件名就是对应服务
|   |-- http.txt
|-- vuln						# 漏洞扫描结果
|   |-- nuclei-*.html		
|   |-- xray.html				# xray 漏扫结果
|   `-- xscan.html				# xscan 漏扫结果
`-- webinfo						
    |-- webinfo.csv				
    `-- webinfo.json
```

`webinfo.csv` 如下：

![image-20230730103850930](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/image-20230730103850930.png)

`webinfo.json` 是方便使用脚本解析结果生成的。

除此之外还有一个 `tmp` 目录，都是是工具运行时生成的结果文件。

