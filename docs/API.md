## 简介

`core` 为前端提供了所有与直播有关的API以及简单的收藏夹CRUD功能，你可以使用API定制自己的前端甚至将其使用在移动平台。

## 直播平台

所有涉及到传递、获取平台参数的接口均使用以下枚举值：

|  参数名  |  平台名  |
| :------: | :------: |
| bilibili | 哔哩哔哩 |
|   huya   |   虎牙   |
|  douyu   |   斗鱼   |
|  egame   |   企鹅电竞   |
|  inke   |   映客   |

## 直播间号

|  平台  |  直播号 |
| :------: | :------: |
| bilibili | https://live.bilibili.com/6 (`6`为直播间号) |
|   huya   | https://www.huya.com/25975826 (`25975826`为直播间号) https://www.huya.com/kaerlol (`kaerlol`为直播间号)   |
|  douyu   | https://www.douyu.com/312212 (`312212`为直播间号)   |
|  egame   | https://egame.qq.com/383204988 (`383204988`为直播间号)   |
|  inke   | https://www.inke.cn/liveroom/index.html?uid=297594356&id=1642924692900921 (uid`297594356`为直播间号)   |


## 参数说明

所有参数后端没有默认值，即前端均需传入一个确定的值

## 版本检查

### GetVersion 获取core版本信息

> GET /api/version

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "runtime": "go1.17.3 windows/amd64",
    "ver": "v0.1.0"
  }
}
```

## HTTP请求转发

直播推荐和分区等功能应当在前端中实现，为了避免跨域限制， `core` 提供了 `http` 请求转发功能，用于请求除 `core` `localhost` `局域网` 外的其他接口。

即使客户端不会有跨域限制，依旧推荐使用该方法请求网页，因为这样可以统一管理所有HTTP请求的具体参数和代理。

> Any /api/v1/proxy

你需要在 `Header` 的 `PL-URL` 属性设置你需要访问的实际 `url` ，注意: 需要显式设置 `scheme` `host` `query`.

`Request Method` `Body` 与 其他 `Header` 都是你实际需要请求的参数。

例如:

```shell
curl --request POST -sL \
     --url 'http://localhost:8081/api/v1/proxy'\
     --header 'PL-URL: https://api.uomg.com/api/long2dwz'\
     --header 'Content-Type: application/x-www-form-urlencoded'\
     --data 'dwzapi=urlcn&url=https://www.baidu.com'
```

## 直播信息类

### GetRoomInfo 获取直播间信息

> GET /api/v1/live/room_info

**Query:**

| 参数名 |              内容              |   示例   |
| :----: | :----------------------------: | :------: |
|  plat  |             平台名             | bilibili |
|  room  | 直播间(短号、长号、完整号均可) |   462    |

请求示例：`/api/v1/live/room_info?plat=bilibili&room=462`

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "status": 1,
    "room": "763679",
    "upper": "老骚豆腐",
    "link": "https://live.bilibili.com/763679",
    "title": "【豆腐】杀鸡+第五+躲猫猫！新游！"
  }
}
```

status 开播情况 0:未开播 1:开播

room 真实房间号

upper 主播名

link 直播间链接

title 直播间标题

### GetRoomInfos 批量获取直播间信息

> **POST** /api/v1/live/room_infos

**Body:**

以下结构的数组形式

| 参数名 |              内容              |  示例   |
| :----: | :----------------------------: |:-----:|
|  id  |      此次请求数组唯一ID，字符串形式，用于获取响应    | "2" |
|  plat  |             平台名             | "douyu" |
|  room  | 直播间(短号、长号、完整号均可) | "8302"  |

请求示例：`/api/v1/live/room_infos`

```json
[
  {
    "id": "1",
    "plat": "douyu",
    "room": "8302"
  },
  {
    "id": "2",
    "plat": "bilibili",
    "room": "732602"
  },
  {
    "id": "3",
    "plat": "douyu",
    "room": "96291"
  }
]
```

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "1": {
      "status": 0,
      "room": "8302",
      "upper": "祈风1v9",
      "link": "https://www.douyu.com/8302",
      "title": "国服第一VN！一个打八百多个！"
    },
    "2": {
      "status": 1,
      "room": "732602",
      "upper": "大祥哥来了",
      "link": "https://live.bilibili.com/732602",
      "title": "申鹤第一生产线"
    },
    "3": {
      "status": 1,
      "room": "96291",
      "upper": "东北大鹌鹑",
      "link": "https://www.douyu.com/96291",
      "title": "东北大鹌鹑 抄作业都抄不明白"
    }
  }
}
```

响应为 `map` 格式，`id` 对应请求时的 `id`

status 开播情况 0:未开播 1:开播

room 真实房间号

upper 主播名

link 直播间链接

title 直播间标题

### GetPlayURL 获取直播流信息

> GET /api/v1/live/play_url

获取直播流前应当先调用 `GetRoomInfo` 获取真实房间号，不同平台下房间号规则不同，这由 `core` 内部统一。

**Query:**

| 参数名 |        内容        |   示例   |
| :----: | :----------------: | :------: |
|  plat  |       平台名       | bilibili |
|  room  | 直播间(真实房间号) |  763679  |

请求示例：`/api/v1/live/play_url?plat=bilibili&room=763679`

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "qn": 0,
    "desc": "原画",
    "origin": "https://d1--cn-gotcha03.bilivideo.com/live-bvc/723585/live_4578433_9339544.flv?cdn=cn-gotcha03&expires=16386...",
    "cors": false,
    "type": "flv"
  }
}
```

qn 清晰度，暂时无切换清晰度功能

desc 清晰度描述，暂时无切换清晰度功能

origin 直播流地址

cors 是否有跨域限制 true:有 false:无 若为true必须通过本地流量转发才能播放，若为false直接播放或本地转发均可

type 直播流编码格式

### Play 直播流本地转发

一些直播流开启了跨域限制，获取的直播流无法直接在前端加载，所以 `core` 提供了本地的流量转发功能。

> GET /api/v1/live/play

**Query:**

| 参数名 |                    内容                     |                             示例                             |
| :----: | :-----------------------------------------: | :----------------------------------------------------------: |
|  url   |          平台直播流地址(`url`编码)          | https://d1--cn-gotcha03.bilivideo.com/live-bvc/723585/live_4578433_9339544.flv?cdn=cn-gotcha03&expires=16386... |
|  type  | 编码格式(目前只支持`flv` ,`m3u8`等待支持)， |                             flv                              |

请求示例：`/api/v1/live/play?type=flv&url=https://d1--cn-gotcha03.bilivideo.com/live-bvc/723585/live_4578433_9339544.flv?cdn=cn-gotcha03&expires=16386...`

**Response:**

直接在该次请求下开始返回直播流，即直接把例如请求示例的`url`传入播放器即可

## 直播监听类

### Serve 监听直播实时信息

> Websocket /api/v1/live/serve

**Query:**

| 参数名 |        内容        |   示例   |
| :----: | :----------------: | :------: |
|  plat  |       平台名       | bilibili |
|  room  | 直播间(真实房间号) |  763679  |

**Response:**

成功连接则开始  `websocket` 连接，`core`  会不断向客户端传输直播实时信息，`Websocket Message` 类型为 `Text Message`，由 `JSON` 编码，基本格式为

```json
{
  "type": "",
  "data": {}
}
```

同时， `Upgrade Websocket` 的 `Response Header` 会在 `Set-Cookie` 头中设置 `uuid` 字段，这个值是后续直播操作类的唯一标识，`core`
会保存其对应的客户端和 `websocket` 连接。

前端应当在成功建立连接后立刻保存 `uuid` 值，而不是依靠 `cookie` 保存，否则在多个标签页中会造成 `uuid` 的覆盖。

`uuid`示例: `8luunypcms`

消息目前包括：弹幕、心跳检测(直播平台心跳包由`core`维护)、直播间热度(仅支持部分平台)

- 弹幕消息(`danmaku`)

  示例：

  ```json
  {
      "type": "danmaku",
      "data": {
          "content": "妙啊",
          "type": 0,
          "color": 5566168
      }
  }
  ```

  content 弹幕内容

type 0:右侧飞行弹幕 1:顶部弹幕 2:底部弹幕

color 弹幕十进制颜色

- 热度消息(`hot`)

  示例:

  ```json
  {
      "type": "hot",
      "data": {
          "hot": 501989
      }
  }
  ```

  hot 热度值

- 心跳检测(`check`)

  对于客户端没有用处，只是用于 `core` 及时释放无用的 `websocket` 连接。

  示例:

  ```json
  {
      "type": "check"
  }
  ```

## 直播操作类

### SendDanmaku 发送弹幕

> POST /api/v1/danmaku/send

**Body:** (JSON编码)

| 参数名  |              内容              |                 示例                 |
| :-----: | :----------------------------: | :----------------------------------: |
|   id    | Serve响应 `Cookie` 中的 `uuid` | 8luunypcms |
| content |            弹幕内容            |            哔哩哔哩干杯~             |
|  type   | 弹幕位置(1:顶部 0:滚动 2:底部) |                  0                   |
|  color  |        弹幕十进制颜色值        |            16777215(白色)            |

请求示例：`/api/v1/live/danmaku/send`

```json
{
  "id": "8luunypcms",
  "content": "哔哩哔哩干杯~",
  "type": 0,
  "color": 16777215
}
```

**Response:**

```json
{
  "code": 0,
  "msg": "ok"
}
```

## 收藏夹类

`core` 提供了简单的收藏夹功能，可以让使用者存下自己喜欢的直播间

- 数据库在每次启动时都会检查默认收藏夹的存在，即必须至少存在一个默认收藏夹(order=1)

- 包含 `order` 字段的项前端应当提供优先级的设置，`order` 越大，优先级越高，显示位置越靠前。

### AddFavList 添加收藏夹

> POST /api/v1/fav/list/add

**Body:** (JSON编码)

| 参数名 |           内容           |    示例    |
| :----: | :----------------------: | :--------: |
| title  | 收藏夹标题(2-60字节长度) | 测试收藏夹 |
| order  |     排序大小(0-100)      |     30     |

请求示例：`/api/v1/fav/list/add`

```json
{
  "title": "测试收藏夹",
  "order": 30
}
```

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": 7,
    "title": "测试ecccqqq",
    "order": 40,
    "created_at": 1638615636,
    "updated_at": 1638615636
  }
}
```

### GetAllFavLists 获取所有收藏夹的信息

> GET /api/v1/fav/list/get_all

请求示例：`/api/v1/fav/list/get_all`

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": [
    {
      "id": 1,
      "title": "默认收藏夹",
      "order": 1,
      "created_at": 1636995643,
      "updated_at": 1636995643
    },
    {
      "id": 6,
      "title": "测试cccqqq",
      "order": 40,
      "created_at": 1636970104,
      "updated_at": 1636970104
    },
    {
      "id": 7,
      "title": "测试ecccqqq",
      "order": 40,
      "created_at": 1638615636,
      "updated_at": 1638615636
    }
  ]
}
```

### GetFavList 获取收藏夹详细信息

> GET /api/v1/fav/list/get

**Query:**

| 参数名 | 内容 | 示例 | | :----: | :------: | :--: | | id | 收藏夹ID | 1 |

请求示例：`/api/v1/fav/list/get?id=1`

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": 1,
    "title": "默认收藏夹",
    "order": 1,
    "created_at": 1636995643,
    "updated_at": 1636995643,
    "favorites": [
      {
        "id": 7,
        "fid": 1,
        "order": 10,
        "plat": "bilibili",
        "room": "469",
        "upper": "test1",
        "created_at": 1637995998,
        "updated_at": 1637995998
      },
      {
        "id": 9,
        "fid": 1,
        "order": 10,
        "plat": "bilibili",
        "room": "469",
        "upper": "test2",
        "created_at": 1637995999,
        "updated_at": 1637995999
      },
      {
        "id": 12,
        "fid": 1,
        "order": 10,
        "plat": "bilibili",
        "room": "469",
        "upper": "test3",
        "created_at": 1637996001,
        "updated_at": 1637996001
      }
    ]
  }
}
```

### DelFavList 删除收藏夹

> POST /api/v1/fav/list/del

**Body:** (JSON编码)

| 参数名 | 内容 | 示例 | | :----: | :-------------------------------------------: | :--: | | id | 收藏夹ID(为1时会报不允许删除默认收藏夹的错误) | 3
|

请求示例：`/api/v1/fav/list/del`

```json
{
  "id": 6
}
```

**Response:**

```json
{
  "code": 0,
  "msg": "ok"
}
```

### EditFavList 编辑收藏夹

> POST /api/v1/fav/list/edit

**Body:** (JSON编码)

如果某些字段不变依旧需要传入旧值

| 参数名 |      内容      |  示例   |
| :----: | :------------: | :-----: |
|   id   |    收藏夹ID    |    3    |
| title  | 新的收藏夹标题 | 测试new |
| order  | 新的收藏夹排序 |   70    |

请求示例：`/api/v1/fav/list/edit`

```json
{
  "id": 6,
  "title": "测试new",
  "order": 70
}
```

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": 6,
    "title": "测试new",
    "order": 70,
    "created_at": 1636970104,
    "updated_at": 1638616320
  }
}
```

## 收藏项类

### AddFav 添加收藏项

> POST /api/v1/fav/list/add

**Body:** (JSON编码)

| 参数名 |        内容        |   示例   |
| :----: | :----------------: | :------: |
|  fid   |      收藏夹ID      |    6     |
| order  |  排序大小(0-100)   |    30    |
|  plat  |       平台名       | bilibili |
|  room  | 房间号(长短号均可) |   469    |
| upper  |       主播名       | 老骚豆腐 |

请求示例：`/api/v1/fav/add`

```json
{
  "fid": 6,
  "order": 30,
  "plat": "bilibili",
  "room": "469",
  "upper": "老骚豆腐"
}
```

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": 21,
    "fid": 6,
    "order": 30,
    "plat": "bilibili",
    "room": "469",
    "upper": "老骚豆腐",
    "created_at": 1638619332,
    "updated_at": 1638619332
  }
}
```

### GetFav 获取收藏项详细信息

> GET /api/v1/fav/get

**Query:**

| 参数名 | 内容 | 示例 | | :----: | :------: | :--: | | id | 收藏项ID | 21 |

请求示例：`/api/v1/fav/get?id=21`

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": 21,
    "fid": 6,
    "order": 30,
    "plat": "bilibili",
    "room": "469",
    "upper": "老骚豆腐",
    "created_at": 1638619332,
    "updated_at": 1638619332
  }
}
```

### DelFav 删除收藏项

> POST /api/v1/fav/del

**Body:** (JSON编码)

| 参数名 | 内容 | 示例 | | :----: | :------: | :--: | | id | 收藏项ID | 21 |

请求示例：`/api/v1/fav/del`

```json
{
  "id": 21
}
```

**Response:**

```json
{
  "code": 0,
  "msg": "ok"
}
```

### EditFav 编辑收藏项

> POST /api/v1/fav/edit

**Body:** (JSON编码)

如果某些字段不变依旧需要传入旧值

| 参数名 |         内容          | 示例  |
| :----: | :-------------------: | :---: |
|   id   |       收藏项ID        |  20   |
| order  | 新的收藏项排序(0-100) |  70   |
|  plat  |    新的收藏项平台     | douyu |
|  room  |   新的收藏项房间号    |  101  |
| upper  |    新的收藏项主播     |  pdd  |

请求示例：`/api/v1/fav/list/edit`

```json
{
  "id": 20,
  "order": 70,
  "plat": "douyu",
  "room": "101",
  "upper": "pdd"
}
```

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": 20,
    "fid": 3,
    "order": 70,
    "plat": "douyu",
    "room": "101",
    "upper": "pdd",
    "created_at": 1638093115,
    "updated_at": 1638619841
  }
}
```

## 资源监控类

### GetOSInfo 获取系统信息

> GET /api/v1/os/info

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "uptime": 75898,
    "os": "windows",
    "platform": "Microsoft Windows 10 Home China",
    "platform_version": "10.0.19042 Build 19042",
    "kernel_version": "10.0.19042 Build 19042",
    "kernel_arch": "x86_64"
  }
}
```

### GetSysMem 获取系统内存占用

> GET /api/v1/os/mem/sys

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "total": 16487870464,
    "total_str": "15.36 GB",
    "avl": 3288862720,
    "avl_str": "3.06 GB"
  }
}
```

### GetSelfMem 获取自身内存占用

> GET /api/v1/os/mem/self

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "mem": 22511616,
    "mem_str": "21.47 MB"
  }
}
```

### GetSysCPU 获取系统CPU占用

> GET /api/v1/os/cpu/sys

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "percent": 5.88235294117647
  }
}
```

### GetSelfCPU 获取自身CPU占用

> GET /api/v1/os/cpu/self

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "percent": 0.040586327084460326
  }
}
```

### GetOSAll 获取全部信息

> GET /api/v1/os/all

**Response:**

```json
{
  "code": 0,
  "msg": "ok",
  "data": {
    "info": {
      "uptime": 76138,
      "os": "windows",
      "platform": "Microsoft Windows 10 Home China",
      "platform_version": "10.0.19042 Build 19042",
      "kernel_version": "10.0.19042 Build 19042",
      "kernel_arch": "x86_64"
    },
    "self_cpu": {
      "percent": 0.038924943854536424
    },
    "self_mem": {
      "mem": 22695936,
      "mem_str": "21.64 MB"
    },
    "sys_cpu": {
      "percent": 3.125
    },
    "sys_mem": {
      "total": 16487870464,
      "total_str": "15.36 GB",
      "avl": 3263381504,
      "avl_str": "3.04 GB"
    }
  }
}
```
