package object

type BaseData interface {
	GetPageURL() string
	GetRawData() string
	GetStatus() int
}

type PageData interface {
	BaseData
	GetDataMap() map[string]any
}

type PageDataImpl struct {
	pageURL string
	rawData string
	status  int
	dataMap map[string]any
}

func (d *PageDataImpl) GetPageURL() string {
	return d.pageURL
}

func (d *PageDataImpl) GetRawData() string {
	return d.rawData
}

func (d *PageDataImpl) GetStatus() int {
	return d.status
}

func (d *PageDataImpl) GetDataMap() map[string]any {
	return d.dataMap
}

func NewData(pageURL string, rawData string, status int, dataMap map[string]any) PageData {
	return &PageDataImpl{
		pageURL: pageURL,
		rawData: rawData,
		status:  status,
		dataMap: dataMap,
	}
}

type ErrorData interface {
	BaseData
	GetError() error
}

type ErrorDataImpl struct {
	err     error
	rawData string
	pageUrl string
	status  int
}

func (e *ErrorDataImpl) GetPageURL() string {
	return e.pageUrl
}

func (e *ErrorDataImpl) GetRawData() string {
	return e.rawData
}

func (e *ErrorDataImpl) GetStatus() int {
	return e.status
}

func (e *ErrorDataImpl) GetDataMap() map[string]any {
	return nil
}

func (e *ErrorDataImpl) GetError() error {
	return e.err
}

func NewErrorData(err error, rawData string, pageUrl string, status int) ErrorData {
	return &ErrorDataImpl{
		err:     err,
		rawData: rawData,
		pageUrl: pageUrl,
		status:  status,
	}
}
