# go gin脚手架项目

## 配置
- viper
> 初始需要把`config`目录下的`application_tmp.yaml`复制出一份`application.yaml`并更改内容

## 数据库
- sqlx
> 抄了Django框架的数据库迁移做法，`init.sql`文件有`db_version`表，标记了当前项目的数据库版本，`migration`目录里存放了`版本号.sql`的迁移文件，记录了每次修改对应的sql脚本，应用后同时更新`db_version`表。

## 缓存
- redis


## 说明
> 自己根据自己习惯搭的一个项目，属于自用吧，如果看得顺眼也欢迎使用，会不定期更新新的东西，但是应该还是会保持轻量简洁。