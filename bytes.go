/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/10/26
   Description :
-------------------------------------------------
*/

package zutils

var Bytes = new(bytesUtil)

type bytesUtil struct{}

// 将data写入到buff的start位置, 超出的数据会被忽略, 返回实际写入的长度
func (*bytesUtil) Write(buff []byte, start int, data []byte) int {
	if start >= len(buff) {
		return 0
	}
	return copy(buff[start:], data)
}

// 将data的dataOffset开始长length的数据写入到buff的start位置, 超出的数据会被忽略, 返回实际写入的长度
func (*bytesUtil) WriteAt(buff []byte, start int, data []byte, dataOffset int, length int) int {
	if start >= len(buff) || length == 0 || dataOffset >= len(data) {
		return 0
	}

	data = data[dataOffset:]
	if length < len(data) {
		data = data[:length]
	}
	return copy(buff[start:], data)
}
