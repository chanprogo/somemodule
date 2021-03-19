package kafkaconsumer

import (
	"container/list"
	"time"
)

// MyCacheData ...
var MyCacheData chan MyInfo

// MyInfo ...
type MyInfo struct {
	MyID     int    `json:"myID"`
	MyName   string `json:"myName"`
	MyTime   string `json:"myTime"`
	MyEnable bool   `json:"myEnable"`
}

// CacheDataStruct ...
type CacheDataStruct struct {
	List    list.List
	MinTime string
	SendLen int
}

// GCacheData ...
var GCacheData *CacheDataStruct

func init() {
	node := new(CacheDataStruct)
	node.List = *list.New()
	node.List.Init()
	node.SendLen = 0
	GCacheData = node
}

// SendCacheData ...
func SendCacheData(to chan MyInfo) {

	GCacheData.SendLen = GCacheData.List.Len()
	if GCacheData.SendLen == 0 {
		return
	}

	for i := 0; i < GCacheData.SendLen; i++ {
		iter := GCacheData.List.Front()
		totalmsg := iter.Value.(*MyInfo)
		to <- *totalmsg
		GCacheData.List.Remove(iter)
	}

	GCacheData.SendLen = GCacheData.List.Len()
}

// GoCacheData ...
func GoCacheData(to chan MyInfo, sec int64) {
	t := time.NewTimer(time.Duration(sec) * time.Second)

	for {

		select {
		case totalmsg := <-MyCacheData:
			GCacheData.List.PushBack(&totalmsg)
		case <-t.C:
			SendCacheData(to)
			t.Reset(time.Duration(sec) * time.Second)
		}

	}

}
