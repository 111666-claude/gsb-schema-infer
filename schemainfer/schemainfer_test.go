package schemainfer

import "testing"

func TestEmptyRecords(t *testing.T) {
	if got := Infer(nil, 3, &Counter{}); len(got) != 0 {
		t.Fatalf("空输入应无字段：%+v", got)
	}
}

func TestSingleFieldKind(t *testing.T) {
	got := Infer([][]Field{{{Name: "a", Kind: "int"}}}, 3, &Counter{})
	if len(got) != 1 || got[0].Kind != "int" {
		t.Fatalf("a 应是 int：%+v", got)
	}
}

func TestFieldCount(t *testing.T) {
	got := Infer([][]Field{{{Name: "a", Kind: "int"}, {Name: "b", Kind: "bool"}}}, 3, &Counter{})
	if len(got) != 2 {
		t.Fatalf("应有两个字段：%+v", got)
	}
}

func TestAllNullStaysNull(t *testing.T) {
	got := Infer([][]Field{{{Name: "a", Kind: "null"}}}, 3, &Counter{})
	if got[0].Kind != "null" {
		t.Fatalf("全 null 字段应是 null：%+v", got)
	}
}

func TestCounterCounts(t *testing.T) {
	c := &Counter{}
	Infer([][]Field{{{Name: "a", Kind: "int"}}}, 3, c)
	if c.Scanned == 0 {
		t.Fatal("应统计比较次数")
	}
}

func TestWindowZeroMeansAll(t *testing.T) {
	got := Infer([][]Field{{{Name: "a", Kind: "int"}}}, 0, &Counter{})
	if len(got) != 1 {
		t.Fatalf("窗口为 0 应统计全部：%+v", got)
	}
}
