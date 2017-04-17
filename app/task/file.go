package task

import (
	"github.com/revel/revel"
	"os"
	"tordowngo/app/model"
	"tordowngo/app/util/ftp"
	utilTorrent "tordowngo/app/util/torrent"

	"github.com/anacrolix/torrent"
	"golang.org/x/time/rate"
)

type File struct {
	File    model.File
	UserNum string
}

func (j File) Run() {
	revel.INFO.Println("hi! I'm anyc job!")
	if j.File.FileType == "application/x-bittorrent" {
		revel.INFO.Println(j.UserNum)
		config := torrent.Config{}
		config.DownloadRateLimiter = rate.NewLimiter(rate.Limit(3*1024*1024), 1<<20)
		fileNames, err := utilTorrent.Down(j.File.FileName, config)
		if err != nil {
			revel.ERROR.Panic(err)
		}
		for _, fileName := range fileNames {
			if err := j.move(fileName); err != nil {
				revel.ERROR.Panic(err)
			}
		}
	} else {
		if err := j.move(j.File.FileName); err != nil {
			revel.ERROR.Panic(err)
		}
	}
	revel.INFO.Println("end my job!")
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
