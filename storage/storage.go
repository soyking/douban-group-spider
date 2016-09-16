package storage

import "github.com/soyking/douban-group-spider/group"

type StorageSave interface {
	Save([]*group.Topic) error
}
