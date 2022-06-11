#问题记录


# 1.go get失败
## 1.1 问题原因
导入gin包时由于github无法访问的原因timeout
```go
mac@mac-2 GoBlog % go get -u github.com/gin-gonic/gin
go: module github.com/gin-gonic/gin: Get "https://proxy.golang.org/github.com/gin-gonic/gin/@v/list": dial tcp 142.251.43.17:443: i/o timeout
```
## 1.2 解决方法
> 使用[七牛云](https://goproxy.cn)提供的Go模块代理,我的操作系统是mac，所以使用export命令，其他系统请参考官网。
```bash
 $ export GO111MODULE=on
 $ export GOPROXY=https://goproxy.cn
```
![alt 成功解决](/Users/mac/myth/Blog/GoBlog/static/logpic/getSuccess.png)
# 2. runtime error

## 2.1 问题原因
> 初步估计是使用指针的问题，某个指针变量声明之后, 还没有经过赋值时默认指向nil, 直接调用指针就会报错”
```azure
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x139bd3c]
```
## 2.2 问题解决
> 不使用defer

# 3. 数据库查询字段问题
## 3.1 问题原因
> 定义的struct和数据库字段不一致
```azure
select data fail scannable dest type struct with >1 columns (2) in result
```
## 问题解决
修改字段

