package ecode

var msg = map[int]string{
	Success:            "ok",
	UnknownError:       "unknown error",
	InvalidParams:      "invalid params",
	RateLimit:          "rate limit",
	ErrorSendDanmaku:   "failed to send danmaku",
	ErrorGetRoomInfo:   "failed to get room info",
	ErrorGetPlayURL:    "failed to get play url",
	ErrorAddFavList:    "failed to add fav list",
	ErrorGetAllFavList: "failed to get all fav lists",
	ErrorDelFavList:    "failed to del fav list",
	ErrorEditFavList:   "failed to edit fav list",
	ErrorGetFavList:    "failed to get fav list",
	ErrorAddFav:        "failed to add fav",
	ErrorDelFav:        "failed to del fav",
	ErrorEditFav:       "failed to edit fav",
	ErrorGetFav:        "failed to get fav",
	ErrorGetSysMem:     "failed to get sys mem",
	ErrorGetSelfMem:    "failed to get self mem",
	ErrorGetSysCPU:     "failed to get sys cpu",
	ErrorGetSelfCPU:    "failed to get self cpu",
	ErrorGetOSInfo:     "failed to get os info",
	ErrorGetOsAll:      "failed to get all os info",
}

func GetMsg(code int) string {
	return msg[code]
}
