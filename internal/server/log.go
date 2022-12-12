package server

import (
	"fmt"
	"sync"
)

var ErrOffsetNotFound = fmt.Errorf("offset not found")

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.SetOffset(&record, c.GetCurrentOffset())

	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.IsOffsetOutOfBound(offset) {
		return Record{}, ErrOffsetNotFound
	}

	return c.GetRecord(offset), nil
}

func (c *Log) GetRecord(offset uint64) Record {
	return c.records[offset]
}

func (c *Log) SetOffset(record *Record, newOffset int) {
	fmt.Println("SetOffset", newOffset)
	record.Offset = uint64(newOffset)
}

func (c *Log) GetCurrentOffset() int {
	return len(c.records)
}

func (c *Log) IsOffsetOutOfBound(offset uint64) bool {
	return offset >= uint64(c.GetCurrentOffset())
}
