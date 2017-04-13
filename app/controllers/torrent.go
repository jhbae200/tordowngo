package controllers

import (
	"tordowngo/app/model"

	"github.com/revel/modules/jobs/app/jobs"
	"github.com/revel/revel"
	"tordowngo/app/task"
	"tordowngo/app/util/filesave"
)

type FileInfo struct {
	File   []model.File
	Status string
}

type TorrentController struct {
	GormController
}

func (c TorrentController) List() revel.Result {
	files := []model.File{}
	if err := c.Txn.Order("file_num desc").Find(&files).Error; err != nil {
		revel.ERROR.Println(err)
	}
	return c.Render(files)
}

func (c TorrentController) Create() revel.Result {
	dbFiles := make([]model.File, len(c.Params.Files["file"]))
	var msg string
	var err error
	for idx, file := range c.Params.Files["file"] {
		dbFiles[idx], err = filesave.Save(file)
		if err != nil {
			revel.ERROR.Println(err)
			msg = err.Error()
			break
		} else {
			if err := c.Txn.Create(&dbFiles[idx]).Error; err != nil {
				revel.ERROR.Println(err)
				msg = err.Error()
				break
			}
			jobs.Now(task.File{
				File: dbFiles[idx],
			})
		}
	}
	msg = "Successfully uploaded"
	return c.RenderJSON(FileInfo{
		File:   dbFiles,
		Status: msg,
	})
}
