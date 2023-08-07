# Project Road Map

## Phase One 

create websocket client and server

> change project struct and for who , need to know , write expalin under the folder 

complete phase one but still remain some problems 

## Phase two

in this stage ,we need to rebuild some struct of the project and add code to cooperate with SDN controller


### add cobra for cmd
依据之前反馈的问题，调节bug 以及功能测试过于困难，加上运行的参数和相关指令绑定在了一起

对于这个问题，借用现有的 cobra 库进行优化

### 日志文件通过 klog 输出
输出格式如下：
``` go
I0807 20:44:17.121542 2137864 server.go:56] Version: v1.13.0-beta.0.59+61df0b6a1e68c6-dirty
I0807 20:44:17.121579 2137864 server.go:92] [1] Prepare agent to run
```
具体参考的文章为：
[讲得非常详细](https://zhengtianbao.com/kubernetes/2021/10/17/kubernetes-cookbook%E4%B9%8B%E6%97%A5%E5%BF%97%E7%AF%87.html)