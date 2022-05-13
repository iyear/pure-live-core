## 🎉 简介

![](https://img.shields.io/github/go-mod/go-version/iyear/pure-live-core?style=flat-square)
![](https://img.shields.io/badge/license-GPL-lightgrey.svg?style=flat-square)
![](https://img.shields.io/github/v/release/iyear/pure-live-core?color=red&style=flat-square)
![](https://img.shields.io/github/last-commit/iyear/pure-live-core?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/iyear/pure-live-core)](https://goreportcard.com/report/github.com/iyear/pure-live-core)

![](https://img.shields.io/github/workflow/status/iyear/pure-live-core/Docker%20Build?label=docker%20build&style=flat-square)
![](https://img.shields.io/docker/v/iyear/pure-live?label=docker%20tag&style=flat-square)
![](https://img.shields.io/docker/image-size/iyear/pure-live?style=flat-square&label=docker%20image%20size)

**该项目仅供学习，请勿用于商业用途。任何使用该项目造成的后果由使用者自行承担。**

> 一个想让直播回归纯粹的项目

没有礼物、粉丝团、弹窗，只有直播、弹幕

## ✨ 特性

- 🔎   直播间信息、直播流、弹幕流、发送弹幕
- ⌛    平台 `Websocket` 协议封装，支持转发弹幕消息、直播间热度消息
- 🗝️   解决跨域问题，支持直播流本地转发
- 📂   简易的收藏夹功能支持
- 🎯   资源占用低，5开百万热度直播间、蓝光直播流转发、弹幕全开占用 `40M` 内存
- 🧬   跨平台支持，甚至可以运行在路由器上
- 🔨   支持设置 `Socks5` 代理 (未测试)
- 🧱   良好的项目结构设计，解耦直播平台和核心功能
- ⚙️ 同时它也是一个简单的命令行工具。

......

## 🛠️ 部署

### Docker

支持 `amd64` `386` `arm64` `arm/v6` `arm/v7` 架构

```shell
#启动
docker run --name pure-live -p <HOST_PORT>:8800 -d --restart=always iyear/pure-live:latest
#或添加-v参数
docker run --name pure-live -p <HOST_PORT>:8800 -v /HOST/PATH/DATA:/data -v /HOST/PATH/LOG:/log -d --restart=always iyear/pure-live:latest

#查看log
docker logs -f pure-live

#设置账户配置文件
docker cp PATH/TO/account.yaml pure-live:/config/account.yaml
docker restart pure-live

#设置服务器配置文件
docker cp PATH/TO/server.yaml pure-live:/config/server.yaml
docker restart pure-live

#备份数据库
docker cp pure-live:/data/data.db .

#备份配置文件
docker cp pure-live:/config .

#复制log到宿主机
docker cp pure-live:/log .
```

### 二进制部署

下载 [Release](https://github.com/iyear/pure-live-core/releases) 的最新打包文件

解压后重命名 `config` 目录下的 `server.yaml.example` 为 `server.yaml` , `config/account.yaml.example` 为 `account.yaml` ,填写相关信息。

```shell
chmod +x ./pure-live
./pure-live run
```

打开对应的本地地址 `localhost:<port>` ，即可看到前端界面，开始使用 `pure-live` 吧！

`pure-live` 的初衷是本地或局域网的直播流推送，对于 `websocket` 推送没有做压缩或优化处理。

将 `pure-live` 运行在局域网内的 `NAS` 或其他小型服务器上，就可以使整个局域网内享受到 `pure-live` 的支持。
### 前端
`Release` 都已经内置了默认的前端页面

如果前端有小BUG修复，请前往前端仓库下载最新版本替换 `static` 目录下的所有文件

前端自己快速看了一下 `Vue` 一把梭写出来的，仅仅是能用的水平，代码结构也很庞杂凌乱，期待更好的第三方前端页面出现。

前端仓库: https://github.com/iyear/pure-live-frontend

**其他前端页面：**

- ......

## ⚙️ 命令行(仅支持二进制文件)

查看版本:
```shell
./pure-live -v
```

```
v0.1.0.211224-beta
go1.17.3 windows/amd64
```

查看帮助:
```shell
./pure-live -h
./pure-live run -h
./pure-live get -h
./pure-live export -h
```

### run
#### 启动本地服务器

`-s` : 服务器配置文件路径，默认为 `config/server.yaml`

`-a` : 账号配置文件路径，默认为 `config/account.yaml`

```shell
./pure-live run
./pure-live run -s myserver.yml
./pure-live run -s my/myserver.yml -a my/myaccount.yml
```

### get
#### 获取直播信息、直播流、弹幕流

`-p` :平台名。涉及的平台参数在 [API文档](./docs/API.md#直播平台)  中查询

`-r` : 房间号。长短号均可

`--stream` : 下载对应的直播流(暂时只支持 `flv`)，不传入则不下载，传入文件名。此方式下载的 `flv` 文件较大，如需要更精细的控制请使用 `ffmpeg`

`--danmaku` : 抓取对应的弹幕流，以 `xlsx` 格式保存，不传入则不抓取，传入文件名

`--roll` : 抓取弹幕是否显示弹幕滚动信息

```shell
./pure-live get -p bilibili -r 6
./pure-live get -p bilibili -r 6 --stream b.flv
./pure-live get -p bilibili -r 6 --stream b.flv --danmaku dm.xlsx
./pure-live get -p bilibili -r 6 --danmaku dm.xlsx --roll
./pure-live get -p bilibili -r 6 --stream b.flv --danmaku dm.xlsx --roll
```

成功获得相关信息

```
Room: 7734200
Upper: 哔哩哔哩英雄联盟赛事
Title: 直播：全明星周末选人仪式
Link: https://live.bilibili.com/7734200
Stream: https://d1--cn-gotcha03.bilivideo.com/live-bvc/842331/live_50329118_9516950.flv?cdn=cn-gotch......
```

### export
#### 导出收藏及收藏夹信息

`-d` : 数据库路径。默认 `data/data.db`

`-p` : 导出路径。默认 `export.xlsx`

```shell
./pure-live export
./pure-live export -d mydata/data.db
./pure-live export -d mydata/data.db -p mydata.xlsx
```

## 🌲 生态

目前 `pure-live` 的生态并不完善，最终的目标是做到开源社区驱动的维护模式。

在发展到一定规模后， `pure-live` 将会以 `organization` 的形式维护 `core` 与不同平台的客户端。

## 📝 文档

如何写一个自己的前端? [API文档](./docs/API.md)

如何添加新的平台支持? [Client文档](./docs/Client.md)

移动平台 `gomobile` 支持? [TODO](./docs)

## 📷 预览

[WEB前端预览](img/frontend)

## 🔩 贡献

### ISSUE
请使用 `issue` 发起任何问题，非重要事情请勿私聊。

- 提出新的特性帮助 `pure-live` 成长。特性的支持效率取决于其重要程度。
- 提出 `BUG` 解决使用中的问题。 `BUG` 的修复将优先考虑。
- ......

### PR

在 `dev` 分支签出一个自己的分支，请勿向 `master` 发起 `PR`

## 🔌 TODO

### 基本直播功能(直播流+弹幕接收)

- [x] 哔哩哔哩
- [x] 虎牙
- [x] 斗鱼
- [x] 企鹅电竞
- [x] 映客
- [ ] 网易CC
- [ ] Twitch (等待第三方库支持 `m3u8` 拉流)
- [ ] 咪咕体育

### 发送弹幕

- [x] 哔哩哔哩
- [ ] 虎牙
- [ ] 斗鱼

### get

- [ ] 弹幕JSON保存
- [ ] 弹幕ASS保存

## 📈 趋势

![stars](https://starchart.cc/iyear/pure-live-core.svg)

## 🧑 贡献者

<a href="https://github.com/iyear/pure-live-core/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=iyear/pure-live-core"  alt="contrib"/>
</a>

## 🗒️ 参考

https://github.com/wbt5/real-url

https://github.com/flxxyz/douyudm

https://github.com/BacooTang/huya-danmu

## 🔖 LICENSE

AGPL-3.0 License