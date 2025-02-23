package object

import "log"

const (
	EMAIL_REGEX = "email"
	PHONE_REGEX = "phone"
)

var regexMap = map[string]string{
	EMAIL_REGEX: `([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`,
	PHONE_REGEX: `(\d{3}-\d{3}-\d{4})`,
}

type Job interface {
	GetUrl() string
	GetRegexes() []string
}

type JobImpl struct {
	url     string
	regexes []string
}

func (j *JobImpl) GetUrl() string {
	return j.url
}

func (j *JobImpl) GetRegexes() []string {
	var regexes []string
	for _, r := range j.regexes {
		value, ok := regexMap[r]
		if !ok {
			log.Printf("\nRegex %s not found in regexMap\n", r)
			continue
		}
		regexes = append(regexes, value)
	}
	return regexes
}

func NewJob(url string, regexes []string) Job {
	return &JobImpl{
		url:     url,
		regexes: regexes,
	}
}
