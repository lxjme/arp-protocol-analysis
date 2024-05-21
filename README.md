<a name="readme-top"></a>


<div align="center">

[![GitHub Stars](https://img.shields.io/github/stars/lxjme/arp-protocol-analysis.svg)](https://github.com/lxjme/arp-protocol-analysis/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/lxjme/arp-protocol-analysis.svg)](https://github.com/lxjme/arp-protocol-analysis/network/members)
[![GitHub Contributors](https://img.shields.io/github/license/lxjme/arp-protocol-analysis?style=social)](https://github.com/lxjme/arp-protocol-analysis/contributors)

</div>

<!-- PROJECT LOGO -->
<br />

<div align="center">
  <a href="https://github.com/lxjme/arp-protocol-analysis">
    <img  src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">网络协议解析</h3>

</div>



## 关于项目
<a name="about-the-project"></a>

```

此仓库准备做成使用golang解析各种网络协议工具库（未完成）
将来准备使用vue模仿wireshark构建界面
```




## 环境
<a name="run-env"></a>
- win10
- go1.21.4

## 使用
<a name="usage"></a>
### 查看Mac地址和Ip地址
```
> ipconfig /all

以太网适配器 Ethernet0:

   连接特定的 DNS 后缀 . . . . . . . :
   描述. . . . . . . . . . . . . . . : ******************************
   物理地址. . . . . . . . . . . . . : 00-0C-29-01-51-22   本机Mac地址
   DHCP 已启用 . . . . . . . . . . . : 是
   自动配置已启用. . . . . . . . . . : 是
   本地链接 IPv6 地址. . . . . . . . : ********************************
   IPv4 地址 . . . . . . . . . . . . : 192.168.3.131(首选) 本机IP
   子网掩码  . . . . . . . . . . . . : 255.255.255.0
   获得租约的时间  . . . . . . . . . : ******************************
   租约过期的时间  . . . . . . . . . : ******************************
   默认网关. . . . . . . . . . . . . : ******************************
   DHCP 服务器 . . . . . . . . . . . : ******************************
   DHCPv6 IAID . . . . . . . . . . . : ******************************
   DHCPv6 客户端 DUID  . . . . . . . : ******************************
   DNS 服务器  . . . . . . . . . . . : ******************************
   TCPIP 上的 NetBIOS  . . . . . . . : ******************************
```