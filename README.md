# rank

## 思考过程

1. 量大
2. 日活10万，写入压力大
3. 结束后建立排行榜，没有读写排序同时进行
4. 排序规则第一顺序分数，第二顺序时间
5. 需要查询排行榜自己名次前后10位

## 问题点

1. 数据总条目多，怎么读取（但总量不大，就读取一次，就这样吧）
2. 排序算法（测试后不需要排序算法百万条数排序时间不到1S）
3. 写入压力大（整个消息队列应该就没问题了）
4. 查排名可能需要关联其它表带出玩家昵称等（根据排名分表，查询至多三次）

## TODO

- [ ] 写入没有接消息队列
- [ ] 没有查询玩家昵称或者工会啥的
- [ ] 计算这一会理论上不能查询，没有加锁
- [ ] 没有实现真实数据库版本，目前是内存型的
- [ ] 等等等

## 测试

### 排序

没用 benchmark 因为没想好怎么生成 b.N 个待排序数组，因为 sort 是原地排序，同一个数组不能重复测试

``` bash
=== RUN   TestSort

  TestSort ✔


1 total assertion


  Test 1000000 record spend time     sort_test.go:42: spare time: 355



1 total assertion

--- PASS: TestSort (0.74s)
PASS
ok      github.com/zhangweiii/rank/sort 1.105s
```