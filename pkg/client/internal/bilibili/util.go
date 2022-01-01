package bilibili

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"github.com/andybalholm/brotli"
	"github.com/iyear/pure-live-core/pkg/conf"
	"io"
)

// ops
const (
	wsOpHeartbeat        = 2 // 心跳
	wsOpHeartbeatReply   = 3 // 心跳回应
	wsOpMessage          = 5 // 弹幕消息等
	wsOpEnterRoom        = 7 // 请求进入房间
	wsOpEnterRoomSuccess = 8 // 进房回应
)

// Header
const (
	wsPackHeaderTotalLen = 16 // 头部字节大小
	wsPackageLen         = 4
	wsHeaderLen          = 2
	wsVerLen             = 2
	wsOpLen              = 4
	wsSequenceLen        = 4
)

// ws header default
const (
	wsHeaderDefaultSequence = 1
)

// version protocol
const (
	wsVerPlain  = 0
	wsVerInt    = 1
	wsVerZlib   = 2
	wsVerBrotli = 3
)

var dp2bili = map[int]int{
	conf.DanmakuTypeTop:    5, // 顶部
	conf.DanmakuTypeRight:  1, // 滚动
	conf.DanmakuTypeBottom: 4, // 底部
}
var bili2dp = map[int]int{
	5: conf.DanmakuTypeTop,    // 顶部
	1: conf.DanmakuTypeRight,  // 滚动
	4: conf.DanmakuTypeBottom, // 底部
}

func dp2bilibili(tp int) int {
	return dp2bili[tp]
}
func bilibili2dp(tp int) int {
	return bili2dp[tp]
}

func write(l int, n int) []byte {
	b := make([]byte, l)
	switch l {
	case 2:
		binary.BigEndian.PutUint16(b, uint16(n))
	case 4:
		binary.BigEndian.PutUint32(b, uint32(n))
	case 8:
		binary.BigEndian.PutUint64(b, uint64(n))
	}
	return b
}
func encode(ver, op uint8, body []byte) []byte {
	header := make([]byte, 0, wsPackHeaderTotalLen)
	header = append(header, write(wsPackageLen, len(body)+wsPackHeaderTotalLen)...)
	header = append(header, write(wsHeaderLen, wsPackHeaderTotalLen)...)
	header = append(header, write(wsVerLen, int(ver))...)
	header = append(header, write(wsOpLen, int(op))...)
	header = append(header, write(wsSequenceLen, wsHeaderDefaultSequence)...)

	return append(header, body...)
}

// decode 必须确保len(b)>16
func decode(b []byte) (ver uint16, op uint32, body []byte) {
	return binary.BigEndian.Uint16(b[wsPackageLen+wsHeaderLen : wsPackageLen+wsHeaderLen+wsVerLen]),
		binary.BigEndian.Uint32(b[wsPackageLen+wsHeaderLen+wsVerLen : wsPackageLen+wsHeaderLen+wsVerLen+wsOpLen]),
		b[wsPackHeaderTotalLen:]
}
func zlibDe(src []byte) ([]byte, error) {
	var (
		r   io.ReadCloser
		o   bytes.Buffer
		err error
	)
	if r, err = zlib.NewReader(bytes.NewReader(src)); err != nil {
		return nil, err
	}
	if _, err = io.Copy(&o, r); err != nil {
		return nil, err
	}
	return o.Bytes(), nil
}
func brotliDe(src []byte) ([]byte, error) {
	o := new(bytes.Buffer)
	r := brotli.NewReader(bytes.NewReader(src))
	if _, err := io.Copy(o, r); err != nil {
		return nil, err
	}
	return o.Bytes(), nil
}
