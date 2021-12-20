![](https://socialify.git.ci/iyear/biligo/image?description=1&font=Raleway&forks=1&issues=1&logo=https://s4.ax1x.com/2021/12/06/orLSGF.png&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Dark)

## 🎉 简介

![](https://img.shields.io/github/go-mod/go-version/iyear/biligo?style=flat-square)
![](https://img.shields.io/badge/license-GPL-lightgrey.svg?style=flat-square)
![](https://img.shields.io/github/v/release/iyear/biligo?color=red&style=flat-square)
![](https://img.shields.io/github/last-commit/iyear/biligo?style=flat-square)

**该项目仅供学习，请勿用于商业用途。任何使用该项目造成的后果由使用者自行承担。**

> 一个想让直播回归纯粹的项目

没有礼物、粉丝团、弹窗，只有直播、弹幕

## ✨ 特性

- 🔎   直播间信息获取、直播流获取、发送弹幕
- ⌛    平台 `Websocket` 协议封装，支持转发弹幕消息、直播间热度消息
- 🗝️   解决跨域问题，支持直播流本地转发
- 📂   简易的收藏夹功能支持
- 🔨   支持设置 `Socks5` 代理 (未测试)
- 🧱   良好的项目结构设计，解耦直播平台和核心功能
- ⚙️ 同时它也是一个简单的命令行工具。
- ......

## 🛠️ 部署

### 后端

下载 [Release](https://baidu.com) 的最新打包文件，解压后重命名 `config.yaml.example` 为 `config.yaml` ，填写相关信息。

```sh
chmod +x ./pure-live
./pure-live run
```

打开对应的本地地址 `localhost:<port>` ，即可看到前端界面，开始使用 `pure-live` 吧！

### 前端
`Release` 都已经内置了默认的前端页面

前端自己快速看了一下 `Vue` 一把梭写出来的，仅仅是能用的水平，代码结构也很庞杂凌乱，期待更好的第三方前端页面出现。

前端仓库: https://github.com/iyear/pure-live-frontend

**其他前端页面：**

- ......

## ⚙️ 命令行

### get
#### 获取直播信息、直播流、弹幕流

`-p` :平台名。涉及的平台参数在 [API文档](./docs/API.md#直播平台)  中查询

`-r` : 房间号。长短号均可

`--stream` : 下载对应的直播流(暂时只支持 `flv`)，不传入则不下载，传入文件名

`--danmaku` : 抓取对应的弹幕流，以 `xlsx` 格式保存，不传入则不抓取，传入文件名

`--roll` : 抓取弹幕是否显示弹幕滚动信息

```sh
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

## 📝 文档

如何写一个自己的前端? [API文档](./docs/API.md)

如何添加新的平台支持? [Client文档](./docs/Client.md)

## 🔌 TODO

### 基本直播功能(直播流+弹幕接收)

- [x] 哔哩哔哩
- [x] 虎牙
- [x] 斗鱼
- [ ] 企鹅电竞
- [ ] Twitch (等待第三方库支持 `m3u8` 拉流)
- [ ] 咪咕体育

### 发送弹幕

- [x] 哔哩哔哩
- [ ] 虎牙
- [ ] 斗鱼

## 参考

https://github.com/wbt5/real-url

https://github.com/flxxyz/douyudm

https://github.com/BacooTang/huya-danmu

## 🔖 LICENSE

GPL-3.0 License