package task

import (
	"os"

	"tordowngo/app/model"
	"tordowngo/app/util/ftp"
	utilTorrent "tordowngo/app/util/torrent"

	"github.com/anacrolix/torrent"
	"github.com/revel/revel"
	"golang.org/x/time/rate"
	"tordowngo/app/database"
)

type File struct {
	File    model.File
	UserNum string
}

func (j File) Run() {
	revel.INFO.Println(j.File.FileName + "staring job..")
	db := database.GetDb()
	if j.File.FileType == "application/x-bittorrent" {
		j.File.FileStatus = 9
		db.Save(&j.File)
		config := torrent.Config{}
		config.DownloadRateLimiter = rate.NewLimiter(rate.Limit(3*1024*1024), 1<<20)
		fileNames, err := utilTorrent.Down(j.File.FileName, config)
		if err != nil {
			revel.ERROR.Panic(err)
			j.File.FileStatus = 8
			db.Save(&j.File)
		}
		j.File.FileStatus = 7
		db.Save(&j.File)
		for _, fileName := range fileNames {
			if err := j.move(fileName); err != nil {
				revel.ERROR.Panic(err)
			}
		}
	} else {
		j.File.FileStatus = 7
		db.Save(&j.File)
		if err := j.move(j.File.FileName); err != nil {
			revel.ERROR.Panic(err)
		}
	}
	revel.INFO.Println(j.File.FileName + "end job..")
}

func (j File) move(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	if err = ftp.Save(fileName, file); err != nil {
		return err
	}
	return nil
}
