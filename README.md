# scanMaxThread

主要用于网络安全扫描器测试最高的线程。

### 使用说明
- -t: 用于测试的目标 ip:port
- -n: 每次测试都会以该参数递增，默认为 200
- -s: 响应时间超过该时长即为最大支持线程 ，单位为秒，默认为 3
- 命令示例 ./scanMaxThread -t exptest.com


