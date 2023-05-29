package static

import "os"

type Local struct {
	BaseUrl string
}

// Remove AliOSS File
func (m Local) RemoveFile(fileUrl string) error {
	return os.Remove(fileUrl)
}

func (m Local) UploadFile(fileUrl string) error {
	return nil
}
