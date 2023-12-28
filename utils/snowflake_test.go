package utils

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestId(t *testing.T) {
	var g sync.WaitGroup
	var testSum = 10000
	var ids = struct {
		data map[int64]*interface{}
		sync.RWMutex
	}{
		data: make(map[int64]*interface{}, testSum),
	}

	tt := NewTest(t)
	w, err := NewIDWorker(0)
	tt.EqualNil(err)

	g.Add(testSum)
	for i := 0; i < testSum; i++ {
		go func(t *testing.T) {
			id, err := w.ID()
			tt.EqualNil(err)
			ids.Lock()
			if _, ok := ids.data[id]; ok {
				t.Error("repeated")
				os.Exit(1)
			}
			ids.data[id] = new(interface{})
			ids.Unlock()
			g.Done()
		}(t)
	}
	g.Wait()

	w, err = NewIDWorker(2)
	tt.EqualNil(err)
	id, _ := w.ID()
	tim, ts, workerId, seq := ParseID(id)
	tt.EqualNil(err)
	t.Log(id, tim, ts, workerId, seq)
}

func TestIdWorker_timeReGen(t *testing.T) {
	iw, err := NewIDWorker(0)
	if err != nil {
		t.Error(err)
		return
	}
	ts, err := iw.ID()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(time.Now().UnixNano() / 1000 / 1000)
	fmt.Println(ts)
	id, i, workerId, seq := ParseID(ts)
	fmt.Println(id, i, workerId, seq)
	fmt.Printf("id:%s, i:%d,workerId:%d,seq:%d\n", id, i, workerId, seq)
	aa := "中文"
	// 转成unicode
	unicode := strings.ToValidUTF8(aa, "")
	fmt.Println(unicode)
	fmt.Println("\u4e2d\u6587")
}
