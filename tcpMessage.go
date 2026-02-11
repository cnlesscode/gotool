package gotool

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
)

const MaxMessageSize = 10 << 20 // 10MB

// 响应TCP连接二进制消息
func WriteTCPResponse(conn net.Conn, msg []byte) error {
	msgLength := uint32(len(msg))
	if msgLength > MaxMessageSize {
		return errors.New("message too long")
	}
	buf := make([]byte, 4+msgLength)
	binary.LittleEndian.PutUint32(buf[:4], msgLength)
	copy(buf[4:], msg)
	_, err := conn.Write(buf)
	return err
}

// 读取TCP连接二进制消息
func ReadTCPResponse(conn net.Conn) ([]byte, error) {
	// 1. 读取4字节长度头
	header := make([]byte, 4)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, err
	}

	// 2. 解析长度（小端序）
	contentLength := binary.LittleEndian.Uint32(header)
	if contentLength == 0 {
		return []byte{}, nil
	}

	if contentLength > MaxMessageSize {
		return nil, errors.New("message too long")
	}

	// 3. 读取完整消息体
	contentBuf := make([]byte, contentLength)
	if _, err := io.ReadFull(conn, contentBuf); err != nil {
		return nil, err
	}

	return contentBuf, nil
}
