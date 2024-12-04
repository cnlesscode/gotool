package gotool

import (
	"encoding/binary"
	"io"
	"os"
	"strconv"
)

var binaryFileMap map[string]*BinaryFile

type BinaryDataFileIndexStruct struct {
	Position uint64 // 8字节
	Size     uint64 // 8字节
}

type BinaryFile struct {
	DataDir               string
	DataFilePath          string
	DataFile              *os.File
	IndexFilePath         string
	IndexFile             *os.File
	SuperDataNumberFile   *os.File
	CurrentPosition       int64
	SupperSaveIndexNumber int64
	MaxDataNumberPerFile  uint64
	WriteChannel          chan int64
}

// 初始化
func NewBinaryFile(dataDir string, maxDataNumberPerFile uint64) *BinaryFile {
	if _, ok := binaryFileMap[dataDir]; ok {
		return binaryFileMap[dataDir]
	}
	bf := &BinaryFile{
		DataDir:               dataDir,
		DataFile:              nil,
		IndexFile:             nil,
		CurrentPosition:       0,
		MaxDataNumberPerFile:  maxDataNumberPerFile,
		SupperSaveIndexNumber: -1,
	}

	bf.WriteChannel = make(chan int64, 1)

	return bf
}

// 关闭
func (bf *BinaryFile) Close() {
	if bf.SuperDataNumberFile != nil {
		bf.SuperDataNumberFile.Close()
	}
	if bf.DataFile != nil {
		bf.DataFile.Close()
		bf.DataFilePath = ""
	}
	if bf.IndexFile != nil {
		bf.IndexFile.Close()
		bf.IndexFilePath = ""
	}
	bf.SupperSaveIndexNumber = -1
}

// 写入内容
func (bf *BinaryFile) Write(data []byte) error {

	// 整理文件
	bf.PrepareForWrite()

	// 先写索引
	binaryDataFileIndex := BinaryDataFileIndexStruct{
		Position: uint64(bf.CurrentPosition),
		Size:     uint64(len(data)),
	}
	err := binary.Write(bf.IndexFile, binary.LittleEndian, binaryDataFileIndex.Position)
	if err != nil {
		return err
	}
	err = binary.Write(bf.IndexFile, binary.LittleEndian, binaryDataFileIndex.Size)
	if err != nil {
		return err
	}

	// 再写入数据
	_, err = bf.DataFile.Write(data)
	if err != nil {
		return err
	}
	bf.CurrentPosition += int64(binaryDataFileIndex.Size)

	// 更新最大存储值
	bf.SupperSaveIndexNumber++
	bf.SuperDataNumberFile.Seek(0, 8)
	binary.Write(bf.SuperDataNumberFile, binary.LittleEndian, bf.SupperSaveIndexNumber)
	return nil
}

// 通过索引读取数据
func (bf *BinaryFile) Read(startIndex uint64, length int) ([][]byte, error) {

	err := bf.InitFiles(startIndex, "read")
	if err != nil {
		return nil, err
	}

	var dataAll [][]byte = make([][]byte, 0)
	bfIndex := BinaryDataFileIndexStruct{
		Position: 0,
		Size:     0,
	}

	currentReadIndex := startIndex % bf.MaxDataNumberPerFile
	for i := 0; i < length; i++ {
		if currentReadIndex >= bf.MaxDataNumberPerFile {
			res, err := bf.Read(startIndex, length-i)
			if err != nil {
				break
			}
			dataAll = append(dataAll, res...)
		}
		offsetStart := currentReadIndex * 16
		_, err = bf.IndexFile.Seek(int64(offsetStart), 0)
		if err != nil {
			break
		}
		err := binary.Read(bf.IndexFile, binary.LittleEndian, &bfIndex.Position)
		if err != nil {
			break
		}
		err = binary.Read(bf.IndexFile, binary.LittleEndian, &bfIndex.Size)
		if err != nil {
			break
		}
		dataIn := make([]byte, bfIndex.Size)
		bf.DataFile.Seek(int64(bfIndex.Position), 0)
		_, err = bf.DataFile.Read(dataIn)
		if err != nil {
			break
		}
		dataAll = append(dataAll, dataIn)
		currentReadIndex++
		startIndex++
	}
	return dataAll, nil
}

func (bf *BinaryFile) PrepareForWrite() error {
	var err error
	// 首次写入从记录中读取最大存储值
	if bf.SupperSaveIndexNumber == -1 {
		// 根据当前索引计算索引文件及数据文件路径
		bf.SuperDataNumberFile, err = os.OpenFile(bf.DataDir+SystemSeparator+"supperNumber.bin", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			return err
		}
		// 读取最大存储值
		err = binary.Read(bf.SuperDataNumberFile, binary.LittleEndian, &bf.SupperSaveIndexNumber)
		if err != nil {
			if err == io.EOF {
				bf.SupperSaveIndexNumber = 0
				err = binary.Write(bf.SuperDataNumberFile, binary.LittleEndian, bf.SupperSaveIndexNumber)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	err = bf.InitFiles(uint64(bf.SupperSaveIndexNumber), "write")
	if err != nil {
		return err
	}

	if bf.SupperSaveIndexNumber%int64(bf.MaxDataNumberPerFile) == 0 {
		bf.CurrentPosition = 0
	}

	return nil
}

func (bf *BinaryFile) InitFiles(startIndex uint64, action string) error {
	var err error
	// 根据索引最大值创建索引文件
	fileNumber := startIndex / bf.MaxDataNumberPerFile
	// 数据文件
	oldDataFilePath := bf.DataFilePath
	bf.DataFilePath = bf.DataDir + SystemSeparator + "data_" + strconv.Itoa(int(fileNumber)) + ".bin"
	if oldDataFilePath != bf.DataFilePath {
		if bf.DataFile != nil {
			bf.DataFile.Close()
		}
		if action == "read" {
			bf.DataFile, err = os.OpenFile(bf.DataFilePath, os.O_RDWR, 0777)
		} else {
			bf.DataFile, err = os.OpenFile(bf.DataFilePath, os.O_CREATE|os.O_RDWR, 0777)
		}
		if err != nil {
			return err
		}
	}
	// 索引文件
	oldIndexFilePath := bf.IndexFilePath
	bf.IndexFilePath = bf.DataDir + SystemSeparator + "index_" + strconv.Itoa(int(fileNumber)) + ".bin"
	if oldIndexFilePath != bf.IndexFilePath {
		if bf.IndexFile != nil {
			bf.IndexFile.Close()
		}
		if action == "read" {
			bf.IndexFile, err = os.OpenFile(bf.IndexFilePath, os.O_RDWR, 0777)
		} else {
			bf.IndexFile, err = os.OpenFile(bf.IndexFilePath, os.O_CREATE|os.O_RDWR, 0777)
		}
		if err != nil {
			return err
		}
	}
	// 开始的 position
	return nil
}
