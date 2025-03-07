package config

import (
	"flag"
	"fmt"
	"os"
	"scraper/internal/object"
	"scraper/internal/raw_handler"
	"strings"
)

func ParseLogType(s string) raw_handler.RawDataHandler {
	switch s {
	case "verbose":
		return raw_handler.NewRawDataVerbose()
	case "file":
		return raw_handler.NewRawDataFiled()
	default:
		return raw_handler.NewRawDataIgnore()
	}
}

type Config interface {
	ParseFromArgs()
	Close()
	GetUrlsToScrape() []object.Job
	GetRawDataHandler() raw_handler.RawDataHandler
}

type ConfigImpl struct {
	rawDataHandler raw_handler.RawDataHandler
	urlsToScrape   []object.Job
}

func (c *ConfigImpl) Close() {
	c.rawDataHandler.Close()
}

func DefaultConfig() *ConfigImpl {
	return &ConfigImpl{
		rawDataHandler: raw_handler.NewRawDataIgnore(),
	}
}

func (c *ConfigImpl) ParseFromArgs() {
	logType := flag.String("log", "ignore", "Log type")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n    %s -log <string> 'regex1|regex2|https://url.com' 'regex1|regex2|https://another-url.com' ...\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	c.rawDataHandler = ParseLogType(*logType)

	// Get remaining arguments after flags as URLs
	if flag.NArg() > 0 {
		for _, url := range flag.Args() {
			parts := strings.Split(url, "|")
			size := len(parts)

			regexes := parts[:size-1]
			pageUrl := parts[size-1]

			c.urlsToScrape = append(c.urlsToScrape, object.NewJob(pageUrl, regexes))
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}
}

func (c *ConfigImpl) GetUrlsToScrape() []object.Job {
	return c.urlsToScrape
}

func (c *ConfigImpl) GetRawDataHandler() raw_handler.RawDataHandler {
	return c.rawDataHandler
}
