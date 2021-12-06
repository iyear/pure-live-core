package douyu

import (
	"encoding/binary"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

func encode(v interface{}) []byte {
	j, _ := json.Marshal(v)
	msg := strings.NewReplacer(`":"`, `@=`,
		`","`, `/`,
		`@`, `@A`,
		`/`, `@S`,
		`{"`, "",
		`"}`, "").Replace(string(j))
	total := make([]byte, 4)
	binary.LittleEndian.PutUint32(total, uint32(len(msg)+9))
	header := []byte{0xb1, 0x02, 0x00, 0x00}
	end := []byte{0x00}
	// 需要加两次数据包长度
	data := append(total, total...)
	data = append(data, header...)
	data = append(data, []byte(msg)...)
	data = append(data, end...)
	return data
}

// TODO 嵌套decode 未实现嵌套结构的解析，但是目前只用来解析弹幕够用了
func decode(msg []byte) []string {
	var r []string
	packs := regexp.MustCompile(`(type@=.*?)\x00`).FindAllString(string(msg), -1)
	for _, pack := range packs {
		p := pack[0 : len(pack)-1] // 去除尾 0x00
		p = strings.NewReplacer(`@=`, `":"`,
			`/`, `","`,
			`@A`, `@`,
			`@S`, `/`).Replace(p)
		p = `{"` + p[:len(p)-2] + `}` // 去除末尾多余的符号
		r = append(r, p)
	}
	return r
}

var col = map[int]int64{
	1: 16723502,
	2: 52479,
	3: 6749952,
	5: 13369599,
	6: 16139391,
	4: 16737792,
}

func colorConv(color string) int64 {
	if color == "" {
		return 16777215
	}
	c, err := strconv.Atoi(color)
	if err != nil {
		return 16777215
	}
	return col[c]
}
