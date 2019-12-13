package jsonutils

import (
	"encoding/json"
	"testing"
)

type JSON struct {
	Foo int
	Bar string
	Baz float64
}

var j JSON

func init() {
	j = JSON{
		Foo: 123,
		Bar: "benchmark",
		Baz: 123.456,
	}
}

func BenchmarkJsonMarshall(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&j)
	}
}

func BenchmarkJsonUnmarshall(b *testing.B) {
	bys, _  := json.Marshal(j)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := JSON{}
		_ = json.Unmarshal(bys, &j)
	}
}

func BenchmarkJsonutilsMarshall(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Marshal(j)
	}
}

func TestJsonutilsMarshall(t *testing.T) {
	Marshal(j)
}

func BenchmarkJsonutilsUnmarshall(b *testing.B) {
	o := Marshal(j)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := JSON{}
		o.Unmarshal(&j)
	}
}