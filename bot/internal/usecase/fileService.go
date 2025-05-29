package usecase

import (
	"github.com/Relkos12/bot-voice2txt/internal/entity"
	"github.com/Relkos12/bot-voice2txt/internal/port"
)

type FileService struct {
	fileSender port.FileSender //dependence on port.FileSender
}

// NewFileService is contructor for FileService with
func NewFileService(fs port.FileSender) *FileService {
	return &FileService{fileSender: fs}
}

// SendFile send a file to user
func (fs *FileService) SendFile(chatID int64, file entity.ConvertedText) error {
	//sending a file
	if err := fs.fileSender.SendFile(chatID, file.File); err != nil {
		return err
	}
	return nil
}
