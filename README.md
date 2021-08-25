# ginparam
一个实验性质的gin数据自动映射结构体的包
通过一个结构体来对请求参数进行获取和校验，并在未来考虑生成swagger的文档（如何优雅的将返回值也弄在里面目前还没想好）
## struct
结构体大致结构如下：
```go
package main
type Address struct {
	city string
	country string
}
type User struct{
	ID int64 `gp:"name=id;form;"`
	name string
}
```
tag可选项如下：

| 序号|名称|可选项|解释|示例|
---|---|---|---|---
|1|字段类型|json、form、header、param、cookie|指从哪些地方获取字段参数(考虑到以后swagger需求，这里只能选一个)|注意：默认情况下值为json，但是一旦出现某个参数为form，则默认参数都会变为form。（主要考虑到请求body应该不允同时存在form和json）|
|2|是否可为空|null|如果有null则代表可以为空值|
|3|是否必填|required|
|4|数据参数名称|name|

结构体字段可选:
int家族，[]byte,string,新的结构体，一旦出现[]byte则所有字段类型必须为form
如果参数为结构体，则默认为一个json，如果想把结构体字段和父结构体字段放一起，需要字段：super(取自superclass)

考虑实现依赖