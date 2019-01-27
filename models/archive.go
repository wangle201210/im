package models

import (
	"container/list"
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
	EVENT_OLD
	EVENT_IMG
	EVENT_NEWIMG
)

type Event struct {
	Type      EventType // JOIN, LEAVE, MESSAGE
	User      string
	Room      int64
	Timestamp int // Unix timestamp (secs)
	Content   string
}
//只记录20条
const archiveSize = 20

// Event archives.
var archive = list.New()
// 保存一条记录，如果长度超过设置项则删除最老的一项，并在尾部追加一项
// NewArchive saves new event to archive list.
func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}
// 顺序遍历获取全部记录
// GetEvents returns all events after lastReceived.
func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
