你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# 介绍 

fileHttpServer 是一个使用golang编写的简单文件存储系统，包括get, put, post等方式来上传文件 

# 背景 

场景1：在工作之后，特别是大厂，都要使用跳板机链接到开发机上进行编程，经常一些东西要从本地传到开发机
一些东西要从开发机传到本地，这个文件传输就是一个问题，使用fileHttpServer可以很好的解决这个问题， 
在开发机上部署上fileHttpServer，然后在本地可以通过命令一键将文件传输到开发机上，同时还生成下载链接，
方便和同事们共享文件 

场景2：我们平时开发自己的代码，使用阿里云或者百度云的服务器，一样可以用fileHttpServer来传输文件 

场景3：平时我们和其他人分享一些文件的时候无非是通过微信或者QQ直接把文件发过去，可是上G大小的文件是受限制的
我们可以通过自己在云服务器上部署fileHttpServer，然后生成文件的下载链接，直接把下载链接给要分享的人，非常方便。

# 部署 

* `git clone https://github.com/youngyangyang04/fileHttpServer.git`
* 安装最新版本的go语言开发环境（1.11.5以上） 

# 运行 
* 执行 `sh start.sh`  
此时fileHttpServer就运行起来了

# 使用方式 

`curl -T $fileName http://$ip:8085/upload/$fileName`   

$fileName 是要传输的文件   
$ip是fileHttpServer部署所在机器的IP   
8085是fileHttpServer的默认服务端口号  

