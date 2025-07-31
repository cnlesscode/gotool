package static

import (
	"path"
	"path/filepath"
)

// Local 实现 StaticCloud 接口
type Local struct {
	Config Config
}

func (l *Local) InitClient(config Config) {
	l.Config = config
}

func (l *Local) UploadFile(fileUrl string) error {
	return nil
}

func (l *Local) DownloadFile(fileUrl string) error {
	return nil
}

func (l *Local) RemoveFile(fileUrl string, removeLocalFile bool) error {
	return nil
}

func (l *Local) MoveFile(fileUrl string, targerDir string) (string, error) {
	fileUrl, _ = InitFileUrl(fileUrl, l.Config)
	targerUrl := path.Join(targerDir, filepath.Base(fileUrl))
	targerUrl, _ = InitFileUrl(targerUrl, l.Config)
	err := MoveFile(fileUrl, targerUrl)
	if err != nil {
		return "", err
	}
	return targerUrl, nil
}
