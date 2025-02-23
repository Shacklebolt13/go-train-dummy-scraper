package raw_handler

import (
	"log"
	"scraper/internal/object"
)

type RawDataVerbose struct {
}

func (r *RawDataVerbose) HandleData(data object.BaseData) {
	log.Printf("Raw data: %+v\n", data)
}

func (r *RawDataVerbose) Close() {
	//ignore.
}

func NewRawDataVerbose() *RawDataVerbose {
	return &RawDataVerbose{}
}
