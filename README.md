douban-group-spider 
===================
[豆瓣小组](https://www.douban.com/group/) 抓取, 基于 Golang

## 使用

运行

```
go build
./douban-group-spider
```

帮助

```
./douban-group-spider -h
```

## 参数

```
  # 小组设置
  -groups string
        小组 URL 上的标识, 默认是租房小组
  -pages int
        一次抓取多少页, 一页25个帖子, 默认1
  
  # 存储设置      
  -es_addr string
        默认存储是 ElasticSearch, 默认 127.0.0.1:27017
  -es_index string
        ElasticSearch 索引, 如果不存在会自动建立 mapping, 默认 db_rent
  -mongo bool
        使用 MongoDB, 但其中文搜索支持比较麻烦, 所以只做正则查询    
  -mg_addr string
        MongoDB 地址, 默认 127.0.0.1:27017
  -mg_db string
        MongoDB 数据库, 默认 db_rent
  -mg_usr string
        MongoDB 用户名, 默认空
  -mg_pwd string
        MongoDB 密码, 默认空
  
  # 抓取设置      
  -freq int
        抓取频率, 单位秒, 默认60
  -t_con int
        小组抓取线程数, 默认1
  -g_con int
        帖子抓取线程数, 默认1
  -proxy string
        代理设置文件地址, 每一行格式是[地址 端口], 例子 test/proxy.txt
  
  # 帖子过滤设置              
  -author_filter string
        作者过滤文件, 支持名字和连接过滤, 例子 test/block_authors.txt
  -title_filter string
        标题过滤文件, 例子 test/block_titles.txt
  -content_filter string
        标题过滤文件, 例子 test/block_titles.txt
  -reply_filter int
        最大回复数过滤
  -pic_filter bool
        有图片过滤
  -last_utime_filter string
        最后更新时间过滤, 格式 2006-01-02 15:04:05

```

## 自定义

自定义 task (`task/task.go#Task`)

* 存储接口 `storage/storage.go#StorageSave`

* 代理均衡接口 `proxy/balancer.go#Balancer`
