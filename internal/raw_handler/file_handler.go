package raw_handler

import (
	"fmt"
	"log"
	"os"
	"scraper/internal/object"
	"time"
)

type RawDataFiled struct {
	logger log.Logger
	file   *os.File
}

func (r *RawDataFiled) HandleData(data object.BaseData) {
	r.logger.Println(data.GetRawData())
}

func (r *RawDataFiled) Close() {
	err := r.file.Close()
	if err != nil {
		log.Printf("\nError occurred while closing file : %v", err)
	}
}

func NewRawDataFiled() *RawDataFiled {
	var prefix = "scraper"
	var file_type = "log"

	var file_name = fmt.Sprintf("%s.%d.%s", prefix, time.Now().UnixMilli(), file_type)
	base_path, err := os.Getwd()

	if err != nil {
		log.Fatalf("\nUnable to find CWD %v", err)
	}

	var file_uri = base_path + "/" + file_name
	file, err := os.OpenFile(file_uri, os.O_WRONLY, os.ModeAppend)

	if err != nil {
		log.Fatalf("\nUnable to open file %s : %v", file_uri, err)
	}

	return &RawDataFiled{
		logger: *log.New(file, "", log.Flags()),
		file:   file,
	}
}
