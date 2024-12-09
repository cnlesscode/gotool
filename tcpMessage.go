package gotool

import (
	"bytes"
	"encoding/binary"
	"net"
)

// 响应TCP连接二进制消息
func WriteTCPResponse(conn net.Conn, msg []byte) error {
	// int32 占4字节
	var messageLength int32 = int32(len(msg))
	var response = new(bytes.Buffer)
	err := binary.Write(response, binary.LittleEndian, messageLength)
	if err != nil {
		return err
	}
	// 写入内容
	err = binary.Write(response, binary.LittleEndian, msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(response.Bytes())
	return err
}

// 读取TCP连接二进制消息
func ReadTCPResponse(conn net.Conn) ([]byte, error) {
	// 读取客户端发送的消息
	buf := make([]byte, 4)
	n, err := conn.Read(buf)
	if err != nil || n != 4 {
		return nil, err
	}
	var contentLength int32
	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &contentLength)
	if err != nil || n != 4 {
		return nil, err
	}
	// 02. 读取消息
	contentBuf := make([]byte, contentLength)
	_, err = conn.Read(contentBuf)
	if err != nil || n != 4 {
		return nil, err
	}
	return contentBuf, nil
}
