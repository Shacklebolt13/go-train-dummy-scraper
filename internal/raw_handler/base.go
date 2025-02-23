package raw_handler

import "scraper/internal/object"

type RawDataHandler interface {
	HandleData(data object.BaseData)
	Close()
}

/*
Default implementation that can be used as a no-op handler when raw data processing needs to be bypassed.
*/
type RawDataIgnore struct{}

func (r *RawDataIgnore) HandleData(data object.BaseData) {
	//Ignores any raw Data
}

func (r *RawDataIgnore) Close() {
	//Ignore
}

func NewRawDataIgnore() *RawDataIgnore {
	return &RawDataIgnore{}
}
