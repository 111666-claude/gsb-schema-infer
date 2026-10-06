# schema-infer

从一批记录里推断模式：字段顺序、字段类型、以及只统计最近若干个记录。只用标准库。

```
go test ./...
go vet ./...
go run ./cmd/schemainfer --sample=widen
go run ./cmd/schemainfer --sample=nullfield
go run ./cmd/schemainfer --sample=order
go run ./cmd/schemainfer --sample=window
go run ./cmd/schemainfer --sample=work
```

## 口径

- **类型格**：类型按 `null < bool < int < float < string` 加宽；`null` 是中性值，不参与加宽。
- **字段顺序**：按窗口内首次出现的顺序排列，与出现次数、字母序无关。
- **窗口**：只统计最近 <窗口> 个记录；窗口为 0 表示不限。

## 不变量

- 字段顺序在记录顺序固定时唯一。
- 同一字段的类型等于窗口内出现过的所有类型按类型格取上界。
- `scanned` 不随字段数平方级放大：两千个字段的 `scanned` 不超过 6000。

## 输出契约

场景打印一行 JSON：`order` 是字段名按首见序的列表，`kinds` 是字段名到类型的映射。
`--sample=work` 打印 `{"fields": N, "scanned": N}`。
