CREATE TABLE trade(
    trade_id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    amount INT NOT NULL,
    createtime INT NOT NULL,
    PRIMARY KEY (trade_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO trade (
    name,
    amount,
    createtime
)
VALUES
(
    'aaa',
    '1',
    '5'
),
(
    'bbb',
    '1',
    '5'
),
(
    'ccc',
    '1',
    '5'
),
(
    'ddd',
    '2',
    '6'
);

// 排序 默认升序， 降序关键字DESC
//在数量一致的情况下，时间小的在前面。
SELECT * FROM trade ORDER BY amount,createtime; 

SELECT * FROM trade ORDER BY createtime,amount; 

SELECT * FROM trade ORDER BY amount ASC, createtime DESC;

// 分组求和 SUM
SELECT name,SUM(amount) FROM trade GROUP BY name ORDER BY amount DESC;

// 分组求平均值 AVG
SELECT name,AVG(amount) FROM trade GROUP BY name;

// 求最大值
SELECT name, MAX(amount) FROM trade GROUP BY name;

// 求最小值
SELECT name, MIN(amount) FROM trade GROUP BY name;

// 个数 
SELECT name, COUNT(amount) FROM trade GROUP BY name;

// 去重复
SELECT DISTINCT name FROM trade;

// 重命名
ALTER TABLE toy CHANGE toy_id id int(11);

// 限制
SELECT name,SUM(amount) FROM trade GROUP BY name ORDER BY amount DESC LIMIT 1,2;

外键是表中的一列，它引用另一个表的主键

外键约束 用于确认引用的是主表的哪一列

// 外键
CREATE TABLE location(
    location_id INT NOT NULL AUTO_INCREMENT,
    location_name VARCHAR(50) NOT NULL,
    trade_id INT NOT NULL,
    CONSTRAINT trade_trade_id_fk    // 可以用这个名称解除约束
    FOREIGN KEY (trade_id)          // 外键
    REFERENCES trade (trade_id),    // 哪个表的哪个主键
    PRIMARY KEY (location_id)
);

INSERT INTO location (
    location_name,
    trade_id
)
VALUES
(
    'china',
    1
);

第一范式 1NF
规则1:数据列只包含具有原子性的值
规则2:没有重复的数据组

组合键:多个数据列构成的主键，组合各列后形成具有唯一性的键。

函数依赖:当某列的数据必须随着另一列的数据的改变而改变时,表示第一列函数依赖于第二列。

T.x -> t.y
在关系表T中，y列函数依赖于x列

部分函数依赖:非主键的列依赖于组合主键的某个部分(但不是完全依赖于组合主键)

传递依赖:任何非键列与另一个非键列有关联。(改变任何非键列可能造成其他列的改变)

第二范式 2NF
规则1:先符合1NF
规则2:没有部分函数依赖

第三范式
规则1:符合第二范式
规则2:没有传递函数依赖

CREATE TABLE user3 (
    id INT AUTO_INCREMENT PRIMARY KEY
) AS
    SELECT name AS name, amount AS amount FROM trade GROUP BY name ORDER BY name;

INSERT INTO user (name)
    SELECT name FROM trade GROUP BY name ORDER BY name;

INSERT INTO user
(name)
VALUES
(
'eee'
);

交叉联接
SELECT t.name, u.name 
FROM trade AS t CROSS JOIN user AS u GROUP BY t.name;

内联结
SELECT user.name,toy.name
FROM ut
INNER JOIN
user,toy
WHERE ut.user_id=user.id AND ut.toy_id=toy.id;

不等联接
SELECT user.name,toy.name
FROM ut
INNER JOIN
user,toy
WHERE ut.user_id<>user.id AND ut.toy_id<>toy.id;

自然联接 列名相同
SELECT user.name,toy.name
FROM ut
NATURAL JOIN
user,toy;

子查询
SELECT toy_id FROM ut WHERE user_id IN (SELECT id FROM user ORDER BY id DESC)  LIMIT 2;

别名
SELECT u.name AS username,t.name AS toyname 
FROM ut 
INNER JOIN user AS u, toy AS t
WHERE ut.toy_id =t.id AND ut.user_id=u.id;

SELECT id,name,(SELECT id FROM ut WHERE user_id=u.id) AS tid FROM user u;

非关联子查询：如果子查询可以独立运行且不会引用外层查询的任何结果，即成为非关联子查询