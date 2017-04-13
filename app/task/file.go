package task

import (
	"github.com/revel/revel"
	"os"
	"tordowngo/app/model"
	"tordowngo/app/util/ftp"
)

type File struct {
	File model.File
}

func (j File) Run() {
	revel.INFO.Println("hi! I'm anyc job!")
	if j.File.FileType == "application/x-bittorrent" {
		revel.INFO.Println("It is torrent file")
	} else {
		if err := j.move(); err != nil {
			revel.ERROR.Panic(err)
		}
	}
	revel.INFO.Println("end my job!")
}

func (j File) move() error {
	file, err := os.Open(j.File.FileName)
	if err != nil {
		return err
	}
	if err = ftp.Save(j.File.FileName, file); err != nil {
		return err
	}
	return nil
}
