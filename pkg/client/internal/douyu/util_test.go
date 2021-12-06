package douyu

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	fmt.Println(hex.Dump(encode(`type@=loginreq/roomid@=5473080/dfl@=sn@AA=105@ASss@AA=1/username@=/password@=/ltkid@=/biz@=/stk@=/devid@=10bb104656f9fe396fc4167800021601/ct@=0/pt@=2/cvr@=0/tvr@=7/apd@=/rt@=1637305464/vk@=bc579a41c0274bd3156906a5ddfefd5a/ver@=20190610/aver@=218101901/dmbt@=chrome/dmbv@=95/er@=1/`)))
}
func TestDecode(t *testing.T) {
	b := encode(`type@=loginreq/roomid@=5473080/dfl@=sn@AA=105@ASss@AA=1/username@=/password@=/ltkid@=/biz@=/stk@=/devid@=10bb104656f9fe396fc4167800021601/ct@=0/pt@=2/cvr@=0/tvr@=7/apd@=/rt@=1637305464/vk@=bc579a41c0274bd3156906a5ddfefd5a/ver@=20190610/aver@=218101901/dmbt@=chrome/dmbv@=95/er@=1/`)
	// fmt.Println(hex.Dump(b))
	fmt.Println(decode(b))
	// type@=joingroup/rid@={room_id}/gid@=-9999/
}
