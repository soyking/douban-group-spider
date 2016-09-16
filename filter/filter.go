package filter

import "github.com/soyking/douban-group-spider/group"

type Filter func([]*group.Topic) []*group.Topic

// 与逻辑过滤器，需要满足所有过滤要求
func NewFilter(filterFuncs ...FilterFunc) Filter {
	return func(topics []*group.Topic) []*group.Topic {
		if len(filterFuncs) == 0 {
			return topics
		}

		filteredTopics := []*group.Topic{}
		for _, topic := range topics {
			valid := true
			for _, filter := range filterFuncs {
				if !filter(topic) {
					valid = false
					break
				}
			}
			if valid {
				filteredTopics = append(filteredTopics, topic)
			}
		}

		return filteredTopics
	}
}
