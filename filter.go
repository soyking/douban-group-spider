package main

import (
	"errors"
	"github.com/soyking/douban-group-spider/filter"
	"github.com/soyking/douban-group-spider/flag"
	"time"
)

func newFilter(f *flag.Flag) (filter.Filter, error) {
	filterFuncs := []filter.FilterFunc{}
	if f.AuthorFilterFile != flag.FLAG_AUTHOR_FILTER_DEFAULT {
		authors, err := readLines(f.AuthorFilterFile)
		if err != nil {
			return nil, err
		}
		filterFuncs = append(filterFuncs, filter.AuthorFilter(authors))
	}
	if f.TitleFilterFile != flag.FLAG_AUTHOR_FILTER_DEFAULT {
		titles, err := readLines(f.TitleFilterFile)
		if err != nil {
			return nil, err
		}
		filterFuncs = append(filterFuncs, filter.TitleFilter(titles))
	}
	if f.ContentFilterFile != flag.FLAG_AUTHOR_FILTER_DEFAULT {
		contents, err := readLines(f.ContentFilterFile)
		if err != nil {
			return nil, err
		}
		filterFuncs = append(filterFuncs, filter.ContentFilter(contents))
	}
	if f.ReplyFilter > flag.FLAG_REPLY_FILTER_DEFAULT {
		filterFuncs = append(filterFuncs, filter.ReplyLimitFilter(f.ReplyFilter))
	}
	if f.PicFilter {
		// 只对有图片要求的过滤
		filterFuncs = append(filterFuncs, filter.PicFilter(true))
	}
	if f.LastUpdateTimeFilter != flag.FLAG_LAST_UPDATE_TIME_FILTER_DEFAULT {
		t, err := time.Parse("2006-01-02 15:04:05", f.LastUpdateTimeFilter)
		if err != nil {
			return nil, errors.New("check your time format: " + err.Error())
		}
		filterFuncs = append(filterFuncs, filter.LastUpdateTimeFilter(t))
	}

	return filter.NewFilter(filterFuncs...), nil
}
