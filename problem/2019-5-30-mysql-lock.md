---
layout: post
title:  "Innodb-lock"
date:   2019-5-30 22:00:00
categories: 数据库
---

行锁是对mysql的索引加锁

如果对一个没有索引的字段加行锁,那实际上是加的是表锁

锁的分类:

1. 表锁

    锁住一张表

2. 行锁

    仅锁住一行或一些行的数据

3. 页锁

    锁住一个数据页

行锁的种类:

1. record锁

    仅锁住一行

2. gap锁

    锁住一些行

3. next-key锁

    record锁+gap锁

锁又被分为

1. s锁

    共享锁，读锁，同时可以有多个

2. x锁

    排他锁，写锁，同时只能有一个

3. 意向锁

    在加s/x锁之前必须加is/ix锁

----------

abc | 无索引 | 唯一索引 | 普通索引
-|-|-|-
对一个存在的字段加锁|表锁|record锁|next-key锁
对一个不存在字段加锁|表锁|gap锁|gap锁

唯一索引

    测试一

        数据:1,2,3,4,5

        执行:select * from t where id = 6 for update;

        加锁:(6,+∞)

    测试二

        数据:1,2,3,4,8

        执行:select * from t where id = 6 for update;

        加锁:(4, 8)

    测试三

        数据:1,2,3,4,5,6,8

        执行:select * from t where 7 <= id AND id <= 9 for update;

        加锁:(7,+无穷)

----------

普通索引

    测试一:
        数据:1,3,5

        执行:select * from t where num = 3 for update;

        加锁:[1,3) + 3 + [3,5)


    测试二:

        数据:1,5

        执行:select * from t where num = 3 for update;

        加锁:[1,5)

---------

gap锁和gap锁之间是兼容的

insert的时候加的是特殊的gap锁(意向插入锁)