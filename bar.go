package bar

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
)

var BetterAnyResolver *ar

type ar struct{}

func (*ar) Resolve(typeUrl string) (proto.Message, error) {
	mname := typeUrl
	if slash := strings.LastIndex(mname, "/"); slash >= 0 {
		mname = mname[slash+1:]
	}
	mt := proto.MessageType(mname)
	if mt == nil {
		return new(valMsg), nil
	}
	return reflect.New(mt.Elem()).Interface().(proto.Message), nil
}

type valMsg struct {
	v []byte
}

func (*valMsg) ProtoMessage()             {}
func (*valMsg) XXX_WellKnownType() string { return "BytesValue" }
func (m *valMsg) Reset()                  { *m = valMsg{} }
func (m *valMsg) String() string {
	return fmt.Sprintf("%x", m.v) // not compatible w/ pb oct
}
func (m *valMsg) Unmarshal(b []byte) error {
	m.v = append(m.v, b...)
	return nil
}
