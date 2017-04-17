package torrent

import (
	"github.com/anacrolix/torrent"

	"github.com/revel/revel"
	"time"
)

func Down(fileName string, config torrent.Config) ([]string, error) {
	client, _ := torrent.NewClient(&config)
	defer client.Close()
	tor, err := client.AddTorrentFromFile(fileName)
	if err != nil {
		return nil, err
	}
	<-tor.GotInfo()
	tor.DownloadAll()
	for {
		if tor.BytesCompleted() == tor.Info().TotalLength() {
			break
		} else {
			revel.INFO.Println(tor.BytesCompleted())
			time.Sleep(1 * time.Second)
		}
	}
	downFileName := []string{}
	for _, file := range tor.Files() {
		downFileName = append(downFileName, file.DisplayPath())
	}
	return downFileName, nil
}
