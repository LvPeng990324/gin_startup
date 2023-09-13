-- 数据库版本表
-- 版本号格式：六位数字，例如：000001，对应的sql文件名为：000001.sql，版本从文件名就可见，便于管理
-- 在每次有数据库更新时候，都要保存对应的sql文件到 db/migration 目录下，操作完对应的sql命令后，都要新增一条版本记录到这个表里
-- 不同部署环境拉到更新后对比版本号，如果不是最新的就按版本号顺序执行对应的sql文件，更新到最新版本
create table db_version
(
    id   int auto_increment not null primary key,
    version varchar(32) not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- 插入当前最新版本，每次对本文件有修改时，都要更新这个版本号为最新版本号
insert into db_version (version) values ('000001');

-- 用户表
create table user
(
    id   int auto_increment not null primary key,
    name varchar(32) not null,
    phone varchar(16) not null,
    password varchar(64) not null,
    is_deleted boolean not null default false,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;