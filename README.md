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

 
# 2. 运行
```
>>NsSubDomain.exe -s baidu.com
【SubDomain Burst】 泛解析测试
tmk7mfymgufci6t3y260hewsgakryh2r.baidu.com
【SubDomain Burst】 Count SubDomains Dict 3001 个
bd123002.baidu.com【SubDomain Burst】 baidu.com 去除泛解析
【SubDomain Burst】 baidu.com 数据输出
[+] 2024-11-18 11:25:44 yule.baidu.com 14.215.183.15 119.29.29.29
[+] 2024-11-18 11:26:26 flight.baidu.com 183.240.98.116 101.226.4.6
[+] 2024-11-18 11:26:26 rec.baidu.com 223.109.81.100 117.50.10.10
[+] 2024-11-18 11:27:13 showcase.baidu.com 220.181.33.106 182.254.116.116
[+] 2024-11-18 11:27:15 wj.baidu.com 220.181.33.51 210.2.4.8
[+] 2024-11-18 11:23:55 mobile.baidu.com 112.34.111.126 8.8.4.4
[+] 2024-11-18 11:24:07 IN.baidu.com 112.34.111.126 52.80.60.30
[+] 2024-11-18 11:24:11 in.baidu.com 112.34.111.126 52.80.52.52
[+] 2024-11-18 11:24:02 events.baidu.com 36.155.169.170 223.5.5.5
[+] 2024-11-18 11:24:11 open.baidu.com 36.155.169.170 114.114.115.115
[+] 2024-11-18 11:25:24 find.baidu.com 36.155.169.170 119.29.29.29
[+] 2024-11-18 11:26:56 jia.baidu.com 36.155.169.170 114.114.114.114
[+] 2024-11-18 11:27:09 eol.baidu.com 36.155.169.170 8.8.4.4
[+] 2024-11-18 11:25:02 www9.baidu.com 110.242.68.3 182.254.118.118
[+] 2024-11-18 11:25:26 www10.baidu.com 110.242.68.3 182.254.118.118
[+] 2024-11-18 11:24:18 q.baidu.com 112.34.112.191 114.114.114.110
[+] 2024-11-18 11:27:36 100.baidu.com 223.109.81.78 182.254.118.118
[+] 2024-11-18 11:24:13 brand.baidu.com 39.156.66.125 223.5.5.5
[+] 2024-11-18 11:24:43 insight.baidu.com 39.156.69.136 223.6.6.6
[+] 2024-11-18 11:25:08 ep.baidu.com 10.65.211.57 114.114.115.119
[+] 2024-11-18 11:25:16 mr.baidu.com 39.156.68.115 52.80.52.52
[+] 2024-11-18 11:24:11 f.baidu.com 36.155.169.143 117.50.11.11
[+] 2024-11-18 11:27:08 gx.baidu.com 36.155.169.143 114.114.115.115
[+] 2024-11-18 11:24:11 gh.baidu.com 39.156.68.20 218.30.118.6
[+] 2024-11-18 11:24:11 sd.baidu.com 103.235.46.69 8.8.4.4
[+] 2024-11-18 11:24:58 antivirus.baidu.com 103.235.46.69 223.5.5.5
[+] 2024-11-18 11:27:08 openapi.baidu.com 39.156.66.111 117.50.11.11
[+] 2024-11-18 11:23:54 www.baidu.com 36.155.132.3 210.2.4.8
[+] 2024-11-18 11:24:05 im.baidu.com 36.155.169.19 123.125.81.6
[+] 2024-11-18 11:24:13 0.baidu.com 111.13.100.240 114.114.115.119
[+] 2024-11-18 11:24:07 feeds.baidu.com 223.109.81.235 52.80.66.66
[+] 2024-11-18 11:24:13 abc.baidu.com 36.155.132.76 182.254.118.118
[+] 2024-11-18 11:27:08 pat.baidu.com 10.99.16.57 182.254.118.118
[+] 2024-11-18 11:27:13 qh.baidu.com 10.46.139.100 112.124.47.27
[+] 2024-11-18 11:27:24 custom.baidu.com 61.135.185.121 223.5.5.5
[+] 2024-11-18 11:24:07 ns5.baidu.com 220.181.33.63 182.254.116.116
[+] 2024-11-18 11:24:13 site.baidu.com 223.109.81.220 180.76.76.76
[+] 2024-11-18 11:24:17 developer.baidu.com 223.109.81.179 119.28.28.28
[+] 2024-11-18 11:27:40 ut.baidu.com 10.23.248.87 52.80.60.30
[+] 2024-11-18 11:23:58 download.baidu.com 223.109.81.217 114.114.114.119
[+] 2024-11-18 11:25:26 sv.baidu.com 223.109.81.217 223.5.5.5
[+] 2024-11-18 11:25:14 sec.baidu.com 110.242.68.183 114.114.114.119
[+] 2024-11-18 11:26:26 hospital.baidu.com 223.109.81.66 182.254.118.118
[+] 2024-11-18 11:27:20 zhaopin.baidu.com 223.109.82.173 218.30.118.6
[+] 2024-11-18 11:24:18 ks.baidu.com 111.13.108.113 114.114.115.119
[+] 2024-11-18 11:25:16 pics.baidu.com 36.155.169.194 112.124.47.27
[+] 2024-11-18 11:27:09 light.baidu.com 112.16.225.38 223.5.5.5
[+] 2024-11-18 11:27:16 um.baidu.com 39.156.66.101 223.6.6.6
[+] 2024-11-18 11:27:20 dcc.baidu.com 120.48.7.36 182.254.116.116
[+] 2024-11-18 11:24:07 ask.baidu.com 10.58.222.200 182.254.118.118
[+] 2024-11-18 11:24:11 global.baidu.com 103.235.46.83 8.8.4.4
[+] 2024-11-18 11:24:28 pay.baidu.com 220.181.33.88 112.124.47.27
[+] 2024-11-18 11:24:07 u.baidu.com 112.34.111.153 114.114.115.115
[+] 2024-11-18 11:24:13 cc.baidu.com 112.34.111.153 117.50.60.30
[+] 2024-11-18 11:27:13 token.baidu.com 112.34.113.75 114.114.115.115
[+] 2024-11-18 11:27:13 eye.baidu.com 36.155.169.107 117.50.60.30
[+] 2024-11-18 11:24:13 p.baidu.com 223.109.81.49 117.50.60.30
[+] 2024-11-18 11:25:38 pi.baidu.com 223.109.81.49 8.8.4.4
[+] 2024-11-18 11:26:25 zy.baidu.com 223.109.81.161 210.2.4.8
[+] 2024-11-18 11:23:54 ns2.baidu.com 220.181.33.31 223.6.6.6
[+] 2024-11-18 11:24:11 ca.baidu.com 111.202.115.77 180.76.76.76
[+] 2024-11-18 11:27:13 h5.baidu.com 111.45.3.142 180.76.76.76
[+] 2024-11-18 11:24:11 house.baidu.com 153.37.235.12 218.30.118.6
[+] 2024-11-18 11:26:00 match.baidu.com 36.155.169.4 52.80.66.66
[+] 2024-11-18 11:24:11 x.baidu.com 39.156.68.165 114.114.115.115
[+] 2024-11-18 11:24:22 client.baidu.com 10.242.112.16 114.114.114.110
[+] 2024-11-18 11:24:07 ss.baidu.com 14.215.183.11 114.215.126.16
[+] 2024-11-18 11:24:11 ntp.baidu.com 10.48.49.44 210.2.4.8
[+] 2024-11-18 11:24:11 design.baidu.com 39.156.66.116 182.254.116.116
[+] 2024-11-18 11:24:05 3g.baidu.com 36.152.44.46 218.30.118.6
[+] 2024-11-18 11:26:26 m3.baidu.com 36.152.44.46 223.6.6.6
[+] 2024-11-18 11:27:13 m4.baidu.com 36.152.44.46 117.50.11.11
[+] 2024-11-18 11:27:16 m5.baidu.com 36.152.44.46 117.50.10.10
[+] 2024-11-18 11:27:36 m6.baidu.com 36.152.44.46 223.5.5.5
[+] 2024-11-18 11:24:07 ms.baidu.com 36.155.132.104 182.254.116.116
[+] 2024-11-18 11:27:09 psa.baidu.com 112.34.113.160 117.50.10.10
[+] 2024-11-18 11:27:11 lisa.baidu.com 183.194.218.14 123.125.81.6
[+] 2024-11-18 11:23:54 baidu.com 110.242.68.66 182.254.116.116
[+] 2024-11-18 11:24:05 w.baidu.com 110.242.68.66 182.254.116.116
[+] 2024-11-18 11:24:11 wwww.baidu.com 110.242.68.66 52.80.66.66
[+] 2024-11-18 11:24:11 th.baidu.com 110.242.68.66 117.50.60.30
[+] 2024-11-18 11:26:04 brazil.baidu.com 110.242.68.66 52.80.52.52
[+] 2024-11-18 11:24:07 qa.baidu.com 220.181.107.197 114.114.115.119
[+] 2024-11-18 11:25:19 moa.baidu.com 180.101.49.156 117.50.10.10
[+] 2024-11-18 11:24:07 update.baidu.com 36.155.169.204 119.29.29.29
[+] 2024-11-18 11:25:38 lj.baidu.com 124.237.208.107 114.114.115.115
[+] 2024-11-18 11:27:04 i1.baidu.com 39.175.100.36 119.28.28.28
[+] 2024-11-18 11:27:13 i4.baidu.com 39.175.100.36 117.50.60.30
[+] 2024-11-18 11:27:15 t4.baidu.com 39.175.100.36 112.124.47.27
[+] 2024-11-18 11:27:17 q1.baidu.com 39.175.100.36 119.28.28.28
[+] 2024-11-18 11:27:09 activity.baidu.com 36.155.169.7 8.8.8.8
[+] 2024-11-18 11:24:30 pl.baidu.com 180.101.212.35 1.2.4.8
[+] 2024-11-18 11:24:51 voice.baidu.com 36.155.169.206 223.6.6.6
[+] 2024-11-18 11:25:24 hm.baidu.com 112.34.111.235 52.80.60.30
[+] 2024-11-18 11:24:13 partner.baidu.com 112.34.113.34 112.124.47.27
[+] 2024-11-18 11:25:32 robot.baidu.com 112.34.112.89 114.114.114.114
[+] 2024-11-18 11:27:08 sitemap.baidu.com 182.61.201.91 123.125.81.6
[+] 2024-11-18 11:24:07 ad.baidu.com 182.61.62.50 8.8.4.4
[+] 2024-11-18 11:24:07 file.baidu.com 182.61.62.50 114.114.115.110
[+] 2024-11-18 11:24:07 vc.baidu.com 182.61.62.50 218.30.118.6
[+] 2024-11-18 11:24:32 md.baidu.com 182.61.62.50 117.50.60.30
[+] 2024-11-18 11:24:07 ent.baidu.com 112.34.111.125 8.8.4.4
[+] 2024-11-18 11:27:09 mil.baidu.com 112.34.111.125 114.215.126.16
[+] 2024-11-18 11:24:11 map.baidu.com 223.109.81.130 180.76.76.76
[+] 2024-11-18 11:24:07 metrics.baidu.com 10.42.7.60 210.2.4.8
[+] 2024-11-18 11:24:13 cm.baidu.com 112.34.116.2 123.125.81.6
[+] 2024-11-18 11:24:13 link.baidu.com 111.45.3.85 52.80.60.30
[+] 2024-11-18 11:27:08 pmp.baidu.com 36.155.169.138 101.226.4.6
[+] 2024-11-18 11:27:15 fff.baidu.com 163.177.151.23 114.114.114.110
[+] 2024-11-18 11:27:09 bsc.baidu.com 10.42.7.122 8.8.8.8
[+] 2024-11-18 11:27:15 xiu.baidu.com 39.156.68.147 223.6.6.6
[+] 2024-11-18 11:23:59 home.baidu.com 111.206.209.69 52.80.52.52
[+] 2024-11-18 11:24:05 cloud.baidu.com 112.34.111.165 119.28.28.28
[+] 2024-11-18 11:24:13 user.baidu.com 10.242.113.33 114.114.114.119
[+] 2024-11-18 11:25:58 magic.baidu.com 223.109.81.173 180.76.76.76
[+] 2024-11-18 11:26:11 mcp.baidu.com 117.185.17.22 117.50.11.11
[+] 2024-11-18 11:27:11 zr.baidu.com 111.206.208.227 123.125.81.6
[+] 2024-11-18 11:27:16 child.baidu.com 36.155.169.38 210.2.4.8
[+] 2024-11-18 11:24:04 image.baidu.com 36.155.169.123 223.5.5.5
[+] 2024-11-18 11:27:18 picture.baidu.com 36.155.169.123 8.8.4.4
[+] 2024-11-18 11:24:13 tuan.baidu.com 111.13.100.93 182.254.116.116
[+] 2024-11-18 11:24:52 its.baidu.com 36.155.169.24 182.254.118.118
[+] 2024-11-18 11:27:13 t5.baidu.com 112.34.111.171 112.124.47.27
[+] 2024-11-18 11:24:13 sg.baidu.com 180.97.107.241 8.8.8.8
[+] 2024-11-18 11:24:13 sales.baidu.com 39.156.41.131 123.125.81.6
[+] 2024-11-18 11:25:12 xian.baidu.com 183.232.231.89 223.6.6.6
[+] 2024-11-18 11:27:15 bq.baidu.com 10.11.80.248 123.125.81.6
[+] 2024-11-18 11:23:59 t.baidu.com 183.240.98.41 8.8.4.4
[+] 2024-11-18 11:24:05 live.baidu.com 223.109.81.35 218.30.118.6
[+] 2024-11-18 11:24:13 money.baidu.com 39.156.68.124 101.226.4.6
[+] 2024-11-18 11:24:07 mx2.baidu.com 61.135.163.62 210.2.4.8
[+] 2024-11-18 11:27:09 tts.baidu.com 36.155.169.140 112.124.47.27
[+] 2024-11-18 11:27:15 qm.baidu.com 223.109.81.85 114.114.114.110
[+] 2024-11-18 11:27:15 scc.baidu.com 220.181.107.192 223.5.5.5
[+] 2024-11-18 11:26:10 np.baidu.com 10.23.239.95 182.254.118.118
[+] 2024-11-18 11:27:13 g7.baidu.com 39.156.66.95 182.254.116.116
[+] 2024-11-18 11:24:07 auth.baidu.com 112.34.111.209 114.114.115.110
[+] 2024-11-18 11:25:24 alert.baidu.com 220.181.33.230 182.254.116.116
[+] 2024-11-18 11:25:32 af.baidu.com 110.242.68.110 117.50.60.30
[+] 2024-11-18 11:27:08 du.baidu.com 180.101.51.199 1.2.4.8
[+] 2024-11-18 11:27:11 or.baidu.com 112.34.111.131 114.114.114.110
[+] 2024-11-18 11:24:11 share.baidu.com 39.156.68.163 52.80.66.66
[+] 2024-11-18 11:24:13 auto.baidu.com 14.215.183.213 182.254.118.118
[+] 2024-11-18 11:27:09 msc.baidu.com 39.156.66.154 101.226.4.6
[+] 2024-11-18 11:24:13 local.baidu.com 223.109.81.80 117.50.10.10
[+] 2024-11-18 11:25:04 eshop.baidu.com 223.109.81.80 218.30.118.6
[+] 2024-11-18 11:24:56 router.baidu.com 10.65.211.124 114.114.114.114
[+] 2024-11-18 11:24:05 go.baidu.com 36.155.132.42 52.80.66.66
[+] 2024-11-18 11:24:13 ly.baidu.com 36.155.132.42 52.80.66.66
[+] 2024-11-18 11:27:13 xing.baidu.com 36.155.132.42 114.114.115.115
[+] 2024-11-18 11:24:07 job.baidu.com 36.155.169.63 112.124.47.27
[+] 2024-11-18 11:24:11 education.baidu.com 36.155.169.63 114.114.115.115
[+] 2024-11-18 11:24:13 vote.baidu.com 36.155.169.63 119.29.29.29
[+] 2024-11-18 11:27:13 gaokao.baidu.com 36.155.169.63 223.5.5.5
[+] 2024-11-18 11:24:13 fz.baidu.com 110.242.68.125 52.80.66.66
```
检测时间、域名、解析IP、DNS服务器IP
，没做好输出展示，看起来有点乱，用的话自己整一下吧
