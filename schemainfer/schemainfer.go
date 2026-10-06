// Package schemainfer 从记录里推断字段模式。
package schemainfer

// Counter 记录比较次数。
type Counter struct {
	Scanned int
}

// Field 是一个字段的当前类型。
type Field struct {
	Name string
	Kind string
}

func rank(kind string) int {
	switch kind {
	case "null":
		return 0
	case "bool":
		return 1
	case "int":
		return 2
	case "float":
		return 3
	default:
		return 4
	}
}

func widen(a, b string) string {
	if rank(b) > rank(a) {
		return b
	}
	return a
}

// Infer 推断窗口内字段的顺序与类型。
func Infer(records [][]Field, window int, c *Counter) []Field {
	if window > 0 && len(records) > window {
		records = records[len(records)-window:]
	}
	order := []string{}
	best := map[string]string{}
	for _, rec := range records {
		for _, f := range rec {
			c.Scanned++
			if kind, ok := best[f.Name]; ok {
				best[f.Name] = widen(kind, f.Kind)
			} else {
				order = append(order, f.Name)
				best[f.Name] = f.Kind
			}
		}
	}
	out := make([]Field, 0, len(order))
	for _, name := range order {
		out = append(out, Field{Name: name, Kind: best[name]})
	}
	return out
}
