package bytesx

import "encoding/hex"

func ContactSlice(datas ...[]byte) []byte {
	totalLen := 0
	for _, item := range datas {
		totalLen += len(item)
	}
	result := make([]byte, totalLen)
	count := 0
	for _, item := range datas {
		count += copy(result[count:], item)
	}
	return result
}


func MustHexDecode(data string) []byte {
	result, err := hex.DecodeString(data)
	if err != nil {
		return []byte{}
	}
	return result
}
