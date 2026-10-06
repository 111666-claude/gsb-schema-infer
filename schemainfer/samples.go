package schemainfer

import "fmt"

// Sample 是一个固定场景，用来复现题面里的锚点。
type Sample struct {
	Name string
	Run  func() map[string]any
}

func f(name, kind string) Field { return Field{Name: name, Kind: kind} }

func render(fields []Field) map[string]any {
	order := []string{}
	kinds := map[string]string{}
	for _, fd := range fields {
		order = append(order, fd.Name)
		kinds[fd.Name] = fd.Kind
	}
	return map[string]any{"order": order, "kinds": kinds}
}

// Samples 返回全部场景。
func Samples() []Sample {
	return []Sample{
		{Name: "widen", Run: func() map[string]any {
			recs := [][]Field{{f("a", "int")}, {f("a", "float")}}
			return render(Infer(recs, 0, &Counter{}))
		}},
		{Name: "nullfield", Run: func() map[string]any {
			recs := [][]Field{{f("a", "null")}, {f("a", "int")}}
			return render(Infer(recs, 0, &Counter{}))
		}},
		{Name: "order", Run: func() map[string]any {
			recs := [][]Field{{f("b", "int"), f("a", "int")}, {f("a", "int")}}
			return render(Infer(recs, 0, &Counter{}))
		}},
		{Name: "window", Run: func() map[string]any {
			recs := [][]Field{{f("a", "float")}, {f("a", "int")}, {f("a", "int")}}
			return render(Infer(recs, 2, &Counter{}))
		}},
	}
}

// Work 跑规模线场景。
func Work(n int) map[string]any {
	recs := make([][]Field, 0, n)
	for i := 0; i < n; i++ {
		recs = append(recs, []Field{f(fmt.Sprintf("f%04d", i), "int")})
	}
	c := &Counter{}
	Infer(recs, 0, c)
	return map[string]any{"fields": n, "scanned": c.Scanned}
}
