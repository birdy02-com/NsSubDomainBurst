# NsSubDomainBurst

# 1. 简介
一个实现DNS爆破的库，可过滤掉一部分泛解析的域名结果
```
>> NsSubDomain.exe -h
Usage of NsSubDomain.exe:
  -d string
        输入单个子域名用于解析
  -s string
        输入根域名，进行域名爆破
```
* 内置了多个DNS服务器
* 默认最大并发数：20
* 内置子域名：3001个
* 只支持A记录
