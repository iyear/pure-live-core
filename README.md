## ð ç®ä»

![](https://img.shields.io/github/go-mod/go-version/iyear/pure-live-core?style=flat-square)
![](https://img.shields.io/badge/license-GPL-lightgrey.svg?style=flat-square)
![](https://img.shields.io/github/v/release/iyear/pure-live-core?color=red&style=flat-square)
![](https://img.shields.io/github/last-commit/iyear/pure-live-core?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/iyear/pure-live-core)](https://goreportcard.com/report/github.com/iyear/pure-live-core)

![](https://img.shields.io/github/workflow/status/iyear/pure-live-core/Docker%20Build?label=docker%20build&style=flat-square)
![](https://img.shields.io/docker/v/iyear/pure-live?label=docker%20tag&style=flat-square)
![](https://img.shields.io/docker/image-size/iyear/pure-live?style=flat-square&label=docker%20image%20size)

**è¯¥é¡¹ç®ä»ä¾å­¦ä¹ ï¼è¯·å¿ç¨äºåä¸ç¨éãä»»ä½ä½¿ç¨è¯¥é¡¹ç®é æçåæç±ä½¿ç¨èèªè¡æ¿æã**

> ä¸ä¸ªæ³è®©ç´æ­åå½çº¯ç²¹çé¡¹ç®

æ²¡æç¤¼ç©ãç²ä¸å¢ãå¼¹çªï¼åªæç´æ­ãå¼¹å¹

## â¨ ç¹æ§

- ð   ç´æ­é´ä¿¡æ¯ãç´æ­æµãå¼¹å¹æµãåéå¼¹å¹
- â    å¹³å° `Websocket` åè®®å°è£ï¼æ¯æè½¬åå¼¹å¹æ¶æ¯ãç´æ­é´ç­åº¦æ¶æ¯
- ðï¸   è§£å³è·¨åé®é¢ï¼æ¯æç´æ­æµæ¬å°è½¬å
- ð   ç®æçæ¶èå¤¹åè½æ¯æ
- ð¯   èµæºå ç¨ä½ï¼5å¼ç¾ä¸ç­åº¦ç´æ­é´ãèåç´æ­æµè½¬åãå¼¹å¹å¨å¼å ç¨ `40M` åå­
- ð§¬   è·¨å¹³å°æ¯æï¼çè³å¯ä»¥è¿è¡å¨è·¯ç±å¨ä¸
- ð¨   æ¯æè®¾ç½® `Socks5` ä»£ç (æªæµè¯)
- ð§±   è¯å¥½çé¡¹ç®ç»æè®¾è®¡ï¼è§£è¦ç´æ­å¹³å°åæ ¸å¿åè½
- âï¸ åæ¶å®ä¹æ¯ä¸ä¸ªç®åçå½ä»¤è¡å·¥å·ã

......

## ð ï¸ é¨ç½²

### Docker

æ¯æ `amd64` `386` `arm64` `arm/v6` `arm/v7` æ¶æ

```shell
#å¯å¨
docker run --name pure-live -p <HOST_PORT>:8800 -d --restart=always iyear/pure-live:latest
#ææ·»å -våæ°
docker run --name pure-live -p <HOST_PORT>:8800 -v /HOST/PATH/DATA:/data -v /HOST/PATH/LOG:/log -d --restart=always iyear/pure-live:latest

#æ¥çlog
docker logs -f pure-live

#è®¾ç½®è´¦æ·éç½®æä»¶
docker cp PATH/TO/account.yaml pure-live:/config/account.yaml
docker restart pure-live

#è®¾ç½®æå¡å¨éç½®æä»¶
docker cp PATH/TO/server.yaml pure-live:/config/server.yaml
docker restart pure-live

#å¤ä»½æ°æ®åº
docker cp pure-live:/data/data.db .

#å¤ä»½éç½®æä»¶
docker cp pure-live:/config .

#å¤å¶logå°å®¿ä¸»æº
docker cp pure-live:/log .
```

### äºè¿å¶é¨ç½²

ä¸è½½ [Release](https://github.com/iyear/pure-live-core/releases) çææ°æåæä»¶

è§£ååéå½å `config` ç®å½ä¸ç `server.yaml.example` ä¸º `server.yaml` , `config/account.yaml.example` ä¸º `account.yaml` ,å¡«åç¸å³ä¿¡æ¯ã

```shell
chmod +x ./pure-live
./pure-live run
```

æå¼å¯¹åºçæ¬å°å°å `localhost:<port>` ï¼å³å¯çå°åç«¯çé¢ï¼å¼å§ä½¿ç¨ `pure-live` å§ï¼

`pure-live` çåè¡·æ¯æ¬å°æå±åç½çç´æ­æµæ¨éï¼å¯¹äº `websocket` æ¨éæ²¡æååç¼©æä¼åå¤çã

å° `pure-live` è¿è¡å¨å±åç½åç `NAS` æå¶ä»å°åæå¡å¨ä¸ï¼å°±å¯ä»¥ä½¿æ´ä¸ªå±åç½åäº«åå° `pure-live` çæ¯æã
### åç«¯
`Release` é½å·²ç»åç½®äºé»è®¤çåç«¯é¡µé¢

å¦æåç«¯æå°BUGä¿®å¤ï¼è¯·åå¾åç«¯ä»åºä¸è½½ææ°çæ¬æ¿æ¢ `static` ç®å½ä¸çæææä»¶

åç«¯èªå·±å¿«éçäºä¸ä¸ `Vue` ä¸ææ¢­ååºæ¥çï¼ä»ä»æ¯è½ç¨çæ°´å¹³ï¼ä»£ç ç»æä¹å¾åºæåä¹±ï¼æå¾æ´å¥½çç¬¬ä¸æ¹åç«¯é¡µé¢åºç°ã

åç«¯ä»åº: https://github.com/iyear/pure-live-frontend

**å¶ä»åç«¯é¡µé¢ï¼**

- ......

## âï¸ å½ä»¤è¡(ä»æ¯æäºè¿å¶æä»¶)

æ¥ççæ¬:
```shell
./pure-live -v
```

```
v0.1.0.211224-beta
go1.17.3 windows/amd64
```

æ¥çå¸®å©:
```shell
./pure-live -h
./pure-live run -h
./pure-live get -h
./pure-live export -h
```

### run
#### å¯å¨æ¬å°æå¡å¨

`-s` : æå¡å¨éç½®æä»¶è·¯å¾ï¼é»è®¤ä¸º `config/server.yaml`

`-a` : è´¦å·éç½®æä»¶è·¯å¾ï¼é»è®¤ä¸º `config/account.yaml`

```shell
./pure-live run
./pure-live run -s myserver.yml
./pure-live run -s my/myserver.yml -a my/myaccount.yml
```

### get
#### è·åç´æ­ä¿¡æ¯ãç´æ­æµãå¼¹å¹æµ

`-p` :å¹³å°åãæ¶åçå¹³å°åæ°å¨ [APIææ¡£](./docs/API.md#ç´æ­å¹³å°)  ä¸­æ¥è¯¢

`-r` : æ¿é´å·ãé¿ç­å·åå¯

`--stream` : ä¸è½½å¯¹åºçç´æ­æµ(ææ¶åªæ¯æ `flv`)ï¼ä¸ä¼ å¥åä¸ä¸è½½ï¼ä¼ å¥æä»¶åãæ­¤æ¹å¼ä¸è½½ç `flv` æä»¶è¾å¤§ï¼å¦éè¦æ´ç²¾ç»çæ§å¶è¯·ä½¿ç¨ `ffmpeg`

`--danmaku` : æåå¯¹åºçå¼¹å¹æµï¼ä»¥ `xlsx` æ ¼å¼ä¿å­ï¼ä¸ä¼ å¥åä¸æåï¼ä¼ å¥æä»¶å

`--roll` : æåå¼¹å¹æ¯å¦æ¾ç¤ºå¼¹å¹æ»å¨ä¿¡æ¯

```shell
./pure-live get -p bilibili -r 6
./pure-live get -p bilibili -r 6 --stream b.flv
./pure-live get -p bilibili -r 6 --stream b.flv --danmaku dm.xlsx
./pure-live get -p bilibili -r 6 --danmaku dm.xlsx --roll
./pure-live get -p bilibili -r 6 --stream b.flv --danmaku dm.xlsx --roll
```

æåè·å¾ç¸å³ä¿¡æ¯

```
Room: 7734200
Upper: åå©åå©è±éèçèµäº
Title: ç´æ­ï¼å¨ææå¨æ«éäººä»ªå¼
Link: https://live.bilibili.com/7734200
Stream: https://d1--cn-gotcha03.bilivideo.com/live-bvc/842331/live_50329118_9516950.flv?cdn=cn-gotch......
```

### export
#### å¯¼åºæ¶èåæ¶èå¤¹ä¿¡æ¯

`-d` : æ°æ®åºè·¯å¾ãé»è®¤ `data/data.db`

`-p` : å¯¼åºè·¯å¾ãé»è®¤ `export.xlsx`

```shell
./pure-live export
./pure-live export -d mydata/data.db
./pure-live export -d mydata/data.db -p mydata.xlsx
```

## ð² çæ

ç®å `pure-live` ççæå¹¶ä¸å®åï¼æç»çç®æ æ¯åå°å¼æºç¤¾åºé©±å¨çç»´æ¤æ¨¡å¼ã

å¨åå±å°ä¸å®è§æ¨¡åï¼ `pure-live` å°ä¼ä»¥ `organization` çå½¢å¼ç»´æ¤ `core` ä¸ä¸åå¹³å°çå®¢æ·ç«¯ã

## ð ææ¡£

å¦ä½åä¸ä¸ªèªå·±çåç«¯? [APIææ¡£](./docs/API.md)

å¦ä½æ·»å æ°çå¹³å°æ¯æ? [Clientææ¡£](./docs/Client.md)

ç§»å¨å¹³å° `gomobile` æ¯æ? [TODO](./docs)

## ð· é¢è§

[WEBåç«¯é¢è§](img/frontend)

## ð© è´¡ç®

### ISSUE
è¯·ä½¿ç¨ `issue` åèµ·ä»»ä½é®é¢ï¼ééè¦äºæè¯·å¿ç§èã

- æåºæ°çç¹æ§å¸®å© `pure-live` æé¿ãç¹æ§çæ¯ææçåå³äºå¶éè¦ç¨åº¦ã
- æåº `BUG` è§£å³ä½¿ç¨ä¸­çé®é¢ã `BUG` çä¿®å¤å°ä¼åèèã
- ......

### PR

å¨ `dev` åæ¯ç­¾åºä¸ä¸ªèªå·±çåæ¯ï¼è¯·å¿å `master` åèµ· `PR`

## ð TODO

### åºæ¬ç´æ­åè½(ç´æ­æµ+å¼¹å¹æ¥æ¶)

- [x] åå©åå©
- [x] èç
- [x] æé±¼
- [x] ä¼é¹çµç«
- [x] æ å®¢
- [ ] ç½æCC
- [ ] Twitch (ç­å¾ç¬¬ä¸æ¹åºæ¯æ `m3u8` ææµ)
- [ ] åªåä½è²

### åéå¼¹å¹

- [x] åå©åå©
- [ ] èç
- [ ] æé±¼

### get

- [ ] å¼¹å¹JSONä¿å­
- [ ] å¼¹å¹ASSä¿å­

## ð è¶å¿

![stars](https://starchart.cc/iyear/pure-live-core.svg)

## ð§ è´¡ç®è

<a href="https://github.com/iyear/pure-live-core/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=iyear/pure-live-core"  alt="contrib"/>
</a>

## ðï¸ åè

https://github.com/wbt5/real-url

https://github.com/flxxyz/douyudm

https://github.com/BacooTang/huya-danmu

## ð LICENSE

AGPL-3.0 License