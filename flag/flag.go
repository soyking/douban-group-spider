package flag

import (
	"flag"
)

const (
	NO_FILE = "NO_FILE"

	// ==== DOUBAN Setting ====
	FLAG_GROUPS_NAME    = "groups"
	FLAG_GROUPS_DEFAULT = "beijingzufang"
	FLAG_GROUPS_USAGE   = "group name, split by ,"

	FLAG_PAGES_NAME    = "pages"
	FLAG_PAGES_DEFAULT = 1
	FLAG_PAGES_USAGE   = "fetch pages of a group"

	// ==== ElasticSearch Setting ====
	FLAG_ES_ADDR_NAME    = "es_addr"
	FLAG_ES_ADDR_DEFAULT = "127.0.0.1:9200"
	FLAG_ES_ADDR_USAGE   = "es address"

	FLAG_ES_INDEX_NAME    = "es_index"
	FLAG_ES_INDEX_DEFAULT = "db_rent"
	FLAG_ES_INDEX_USAGE   = "es index"

	// ==== MongoDB Setting ====
	FLAG_USE_MONGO_NAME    = "mongo"
	FLAG_USE_MONGO_DEFAULT = false
	FLAG_USE_MONGO_USAGE   = "use mongo storage"

	FLAG_MONGO_ADDR_NAME    = "mg_addr"
	FLAG_MONGO_ADDR_DEFAULT = "127.0.0.1:27017"
	FLAG_MONGO_ADDR_USAGE   = "MongoDB address, split by ,"

	FLAG_MONGO_USERNAME_NAME    = "mg_usr"
	FLAG_MONGO_USERNAME_DEFAULT = ""
	FLAG_MONGO_USERNAME_USAGE   = "MongoDB username"

	FLAG_MONGO_PASSWORD_NAME    = "mg_pwd"
	FLAG_MONGO_PASSWORD_DEFAULT = ""
	FLAG_MONGO_PASSWORD_USAGE   = "MongoDB password"

	FLAG_MONGO_DATABASE_NAME    = "mg_db"
	FLAG_MONGO_DATABASE_DEFAULT = "db_rent"
	FLAG_MONGO_DATABASE_USAGE   = "MongoDB database"

	// ==== Crawling Setting ====
	FLAG_FREQUENCY_NAME    = "freq"
	FLAG_FREQUENCY_DEFAULT = 60
	FLAG_FREQUENCY_USAGE   = "spider frequency(in second)"

	FLAG_GROUPS_CONCURRENCY_NAME    = "g_con"
	FLAG_GROUPS_CONCURRENCY_DEFAULT = 1
	FLAG_GROUPS_CONCURRENCY_USAGE   = "concurrency for groups crawling"

	FLAG_TOPICS_CONCURRENCY_NAME    = "t_con"
	FLAG_TOPICS_CONCURRENCY_DEFAULT = 1
	FLAG_TOPICS_CONCURRENCY_USAGE   = "concurrency for topics crawling"

	FLAG_PROXY_NAME    = "proxy"
	FLAG_PROXY_DEFAULT = NO_FILE
	FLAG_PROXY_USAGE   = "proxy file, [addr port] split by new line"

	// ==== Filter Setting ====
	FLAG_AUTHOR_FILTER_NAME    = "author_filter"
	FLAG_AUTHOR_FILTER_DEFAULT = NO_FILE
	FLAG_AUTHOR_FILTER_USAGE   = "author filter file path, split by new line"

	FLAG_TITLE_FILTER_NAME    = "title_filter"
	FLAG_TITLE_FILTER_DEFAULT = NO_FILE
	FLAG_TITLE_FILTER_USAGE   = "title filter file path, split by new line"

	FLAG_CONTENT_FILTER_NAME    = "content_filter"
	FLAG_CONTENT_FILTER_DEFAULT = NO_FILE
	FLAG_CONTENT_FILTER_USAGE   = "content filter file path, split by new line"

	FLAG_REPLY_FILTER_NAME    = "reply_filter"
	FLAG_REPLY_FILTER_DEFAULT = 0
	FLAG_REPLY_FILTER_USAGE   = "max reply of a topic"

	FLAG_PIC_FILTER_NAME    = "pic_filter"
	FLAG_PIC_FILTER_DEFAULT = false
	FLAG_PIC_FILTER_USAGE   = "topic with picture"

	FLAG_LAST_UPDATE_TIME_FILTER_NAME    = "last_utime_filter"
	FLAG_LAST_UPDATE_TIME_FILTER_DEFAULT = ""
	FLAG_LAST_UPDATE_TIME_FILTER_USAGE   = "last update time filter, format: 2006-01-02 15:04:05"
)

type Flag struct {
	GroupsName string
	Pages      int

	EsAddr  string
	EsIndex string

	MongoDBOn       bool
	MongoDBAddr     string
	MongoDBUsername string
	MongoDBPassword string
	MongoDBDatabase string

	Frequency         int64
	GroupsConcurrency int
	TopicsConcurrency int
	ProxyFile         string

	AuthorFilterFile     string
	TitleFilterFile      string
	ContentFilterFile    string
	ReplyFilter          int
	PicFilter            bool
	LastUpdateTimeFilter string
}

func ParseFlag() *Flag {
	f := new(Flag)

	flag.StringVar(&f.GroupsName, FLAG_GROUPS_NAME, FLAG_GROUPS_DEFAULT, FLAG_GROUPS_USAGE)
	flag.IntVar(&f.Pages, FLAG_PAGES_NAME, FLAG_PAGES_DEFAULT, FLAG_PAGES_USAGE)

	flag.StringVar(&f.EsAddr, FLAG_ES_ADDR_NAME, FLAG_ES_ADDR_DEFAULT, FLAG_ES_ADDR_USAGE)
	flag.StringVar(&f.EsIndex, FLAG_ES_INDEX_NAME, FLAG_ES_INDEX_DEFAULT, FLAG_ES_INDEX_USAGE)

	flag.BoolVar(&f.MongoDBOn, FLAG_USE_MONGO_NAME, FLAG_USE_MONGO_DEFAULT, FLAG_USE_MONGO_USAGE)
	flag.StringVar(&f.MongoDBAddr, FLAG_MONGO_ADDR_NAME, FLAG_MONGO_ADDR_DEFAULT, FLAG_MONGO_ADDR_USAGE)
	flag.StringVar(&f.MongoDBUsername, FLAG_MONGO_USERNAME_NAME, FLAG_MONGO_USERNAME_DEFAULT, FLAG_MONGO_USERNAME_USAGE)
	flag.StringVar(&f.MongoDBPassword, FLAG_MONGO_PASSWORD_NAME, FLAG_MONGO_PASSWORD_DEFAULT, FLAG_MONGO_PASSWORD_USAGE)
	flag.StringVar(&f.MongoDBDatabase, FLAG_MONGO_DATABASE_NAME, FLAG_MONGO_DATABASE_DEFAULT, FLAG_MONGO_DATABASE_USAGE)

	flag.Int64Var(&f.Frequency, FLAG_FREQUENCY_NAME, FLAG_FREQUENCY_DEFAULT, FLAG_FREQUENCY_USAGE)
	flag.IntVar(&f.GroupsConcurrency, FLAG_GROUPS_CONCURRENCY_NAME, FLAG_GROUPS_CONCURRENCY_DEFAULT, FLAG_GROUPS_CONCURRENCY_USAGE)
	flag.IntVar(&f.TopicsConcurrency, FLAG_TOPICS_CONCURRENCY_NAME, FLAG_TOPICS_CONCURRENCY_DEFAULT, FLAG_TOPICS_CONCURRENCY_USAGE)
	flag.StringVar(&f.ProxyFile, FLAG_PROXY_NAME, FLAG_PROXY_DEFAULT, FLAG_PROXY_USAGE)

	flag.StringVar(&f.AuthorFilterFile, FLAG_AUTHOR_FILTER_NAME, FLAG_AUTHOR_FILTER_DEFAULT, FLAG_AUTHOR_FILTER_USAGE)
	flag.StringVar(&f.TitleFilterFile, FLAG_TITLE_FILTER_NAME, FLAG_TITLE_FILTER_DEFAULT, FLAG_TITLE_FILTER_USAGE)
	flag.StringVar(&f.ContentFilterFile, FLAG_CONTENT_FILTER_NAME, FLAG_CONTENT_FILTER_DEFAULT, FLAG_CONTENT_FILTER_USAGE)
	flag.IntVar(&f.ReplyFilter, FLAG_REPLY_FILTER_NAME, FLAG_REPLY_FILTER_DEFAULT, FLAG_REPLY_FILTER_USAGE)
	flag.BoolVar(&f.PicFilter, FLAG_PIC_FILTER_NAME, FLAG_PIC_FILTER_DEFAULT, FLAG_PIC_FILTER_USAGE)
	flag.StringVar(&f.LastUpdateTimeFilter, FLAG_LAST_UPDATE_TIME_FILTER_NAME, FLAG_LAST_UPDATE_TIME_FILTER_DEFAULT, FLAG_LAST_UPDATE_TIME_FILTER_USAGE)

	flag.Parse()

	if f.Frequency <= 0 {
		f.Frequency = FLAG_FREQUENCY_DEFAULT
	}
	if f.GroupsConcurrency <= 0 {
		f.GroupsConcurrency = FLAG_GROUPS_CONCURRENCY_DEFAULT
	}
	if f.TopicsConcurrency <= 0 {
		f.TopicsConcurrency = FLAG_TOPICS_CONCURRENCY_DEFAULT
	}

	return f
}
