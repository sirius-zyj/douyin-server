package mysql2redis

import (
	"github.com/duanhf2012/origin/util/sync"
)

type RawLogEntity struct {
	TableName string
	Action    string
	Rows      [][]interface{}
	Header    []string
	HeaderMap map[string]int
	//ValueMap  map[string]interface{}

	ref bool
}

func (slf *RawLogEntity) Reset() {

}

func (slf *RawLogEntity) IsRef() bool {
	return slf.ref
}

func (slf *RawLogEntity) Ref() {
	slf.ref = true
}

func (slf *RawLogEntity) UnRef() {
	slf.ref = false
}

type WriteAdapter interface {
	Write([]*RawLogEntity) error
	Close() error
}

var RawLogEntityPoll = sync.NewPoolEx(make(chan sync.IPoolData, 1000), func() sync.IPoolData {
	return &RawLogEntity{}
})
