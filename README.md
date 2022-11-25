# 基于物联网技术的图书馆管理系统


---

用户API

- [x] 用户注册
- [x] 用户登录
- [x] 用户详情

---

书籍API

- [x] 创建书籍信息
- [x] 删除书籍信息
- [x] 查询所有书籍
- [x] 模糊查询书籍
- [x] 借出书籍
- [x] 归还书籍
- [x] 添加馆存书籍
- [x] 减少馆存书籍

---

数据API
- [x] 查询单个传感器的所有数据
- [x] 查询单个传感器n位最新数据


---

```
|- bin 二进制文件
|- common 通用文件
|- config 配置文件
|- controller 控制层
|- dao 数据库访问对象
|- dto 数据传输对象
|- middleware 过滤器
|- po 持久层对象
|- response 返回体
|- router 路由文件
|- script 脚本文件
...
```


# used frame
gin
gorm
swagger
sqlalchemy


## API
http://localhost:8080/swagger/index.html



# 智慧图书馆部署文档

## 前置部署
---
sql文件夹下运行gradation.sql脚本，创建数据库

sql文件夹下另外的文件夹为以表为单位的脚本


## Http服务器搭建
---

1. 安装go语言SDK
2. 在根目录config文件夹下的yaml文件中修改自己的数据库信息
3. 通过根目录python文件夹下的GetDataSaveToDbs.py文件获取云平台数据
	```
	python GetDataSaveToDbs.py
	```
4. 运行根目录sh文件夹下的run.bat文件或直接运行bin文件夹下的二进制文件