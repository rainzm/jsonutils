package jsonutils

import (
	"fmt"
	"reflect"
	"testing"
	"yunion.io/x/pkg/gotypes"
)

type Internal struct {
	Name string
	Age int
}

type External struct {
	Job string
	Inter Internal
	Params []int
}

func (e *External) String() string {
	return Marshal(e).String()
}

func (e *External) IsZero() bool {
	return len(e.Job) == 0
}

func init() {
	ty := reflect.TypeOf((*External)(nil))
	gotypes.RegisterSerializable(ty, func() gotypes.ISerializable {
		return &External{}
	})
}

func genJSONStr() string {
	e := External{
		Job:    "Viewer",
		Inter:  Internal{Name: "zhengcc", Age: 22},
		Params: []int{1, 2, 3},
	}
	return Marshal(e).String()
}

func TestHello(t *testing.T) {
	ty := reflect.TypeOf((*External)(nil))
	s, err := JSONDeserialize(ty, genJSONStr())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", s)
}
