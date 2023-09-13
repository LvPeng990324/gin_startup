-- 示例数据库补丁文件，展示了如何使用数据库补丁

-- 本次补丁要对数据库进行的操作，示例就不进行实际操作了，select一下就好
select * from db_version;

-- 每次补丁最后都要插入一条版本号记录到db_version表里
insert into db_version (version) values ('000002');
