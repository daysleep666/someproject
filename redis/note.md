redis持久化

### 快照
    将某一时刻的数据全部写入文件

**快照配置**

save 900 1 

    如果距离上次快照有900秒且有一次操作就进行快照操作

stop-writes-on-bgsave-error yes

    如果快照失败了，是否允许用户继续执行写操作。默认是yes。

rdbcompression yes

    是否在存储的时候压缩下。

rdbchecksum yes

    是否在存储和读取的时候启用校验。

dbfilename dump.rdb

    存储的文件的名字

dir /usr/local/var/db/redis/

    文件存储路径


BGSAVE:redis会调用fork来创建一个子进程，子进程来创建快照，父进程继续处理用户输入。

SAVE:在快照结束之前阻塞其他操作。

创建快照的方法

- 客户端向服务器发送BGSAVE命令
- 客户端向服务器发送SAVE命令
- 服务器关闭的时候，执行BGSAVE
- 触发自动快照，执行BGSAVE
- 从服务器向主服务器发送SYNC命令的时候，主服务器会执行BGSAVE。

----

### AOF持久化

    只追加文件，在执行写命令的时候，将写命令追加写入到磁盘上。

**AOF配置**

appendonly no

    是否启用AOF

appendfilename "appendonly.aof"

    存储的文件名

appendfsync everysec

    每秒中将缓冲区数据写入硬盘，最多丢失一秒的数据，保证了速度和安全。

appendfsync no

    等操作系统来执行。如果系统发生崩溃，会丢失一定数据。不推荐

appendfsync always

     每有一个新命令都会立刻写入硬盘，这样保证不会丢失数据，但是受限于io速度。

no-appendfsync-on-rewrite no

    当BGSAVE或者BGREWRITEAOF执行的时候不执行fsync()。默认为no，不会丢失数据，但是因为竞争磁盘io，所以会发生阻塞。

auto-aof-rewrite-percentage 100

auto-aof-rewrite-min-size 64mb

    当AOF文件大于64mb且大小是上一次重写时文件大小的一倍 ，执行重写操作。

aof-load-truncated yes

    是否允许redis从AOF的中间某块开始读数据。

aof-use-rdb-preamble no

    混合持久化。当AOF重写操作后创建一个包含RDB和AOF的文件，其中RDB数据在AOF文件的开头，储存了重写操作时的数据库状态，在重写操作后执行的命令以AOF格式追加到AOF文件的末尾。

BGREWRITEAOF 重写操作

为了避免AOF文件过大，占据全部硬盘空间的问题。根据策略定期执行重写操作，移除冗余的命令，来减小AOF体积。