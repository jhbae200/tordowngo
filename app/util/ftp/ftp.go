package ftp

import (
	"os"

	"github.com/revel/revel"
	"gopkg.in/dutchcoders/goftp.v1"
)

var con *goftp.FTP

func Save(path string, file *os.File) error {
	var err error
	if err = connect(); err != nil {
		return err
	}
	defer disConnect()
	revel.INFO.Println("start ftp save")
	if err = con.Stor("/01094587374/tordowngo/"+path, file); err != nil {
		return err
	}
	return nil
}

func connect() error {
	var err error
	source, _ := revel.Config.String("ftp.source")
	user, _ := revel.Config.String("ftp.user")
	password, _ := revel.Config.String("ftp.password")
	if con, err = goftp.Connect(source); err != nil {
		return err
	}
	if err = con.Login(user, password); err != nil {
		return err
	}
	return nil
}

func disConnect() error {
	var err error
	if con != nil {
		if err = con.Quit(); err != nil {
			return err
		}
	}
	return nil
}
