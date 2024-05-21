<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/net.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">网络协议解析</h3>


</div>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">关于项目</a>
    </li>
    <li>
      <a href="#run-env">环境</a>
    </li>
    <li><a href="#usage">使用</a></li>
   
  </ol>
</details>


<!-- ABOUT THE PROJECT -->
## 关于项目
<a name="about-the-project"></a>

此仓库准备做成使用golang解析各种网络协议工具库（未完成）
将来准备使用vue模仿wireshark构建界面




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