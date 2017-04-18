package filesave

import (
	"io"
	"mime/multipart"
	"os"
	"strings"

	"tordowngo/app/model"
)

type Sizer interface {
	Size() int64
}

func Save(header *multipart.FileHeader) (model.File, error) {
	file, err := header.Open()
	if err != nil {
		return model.File{}, err
	}
	defer file.Close()
	dst, err := os.Create(header.Filename)
	if err != nil {
		return model.File{}, err

	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		return model.File{}, err
	}
	os.Chmod(header.Filename, (os.FileMode)(0644))
	//mime-type torrent update.
	mimeType := header.Header.Get("Content-Type")
	dot := strings.LastIndex(header.Filename, ".")
	if !(dot == -1 || dot+1 >= len(header.Filename)) {
		extension := header.Filename[dot+1:]
		if extension == "torrent" {
			mimeType = "application/x-bittorrent"
		}
	}
	return model.File{
		FileName: header.Filename,
		FileType: mimeType,
		FileSize: file.(Sizer).Size(),
	}, nil
}
