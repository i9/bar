package bar

import (
	"encoding/hex"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
)

const (
	anyHex = "0a2a747970652e676f6f676c65617069732e636f6d2f63656c6572782e7265776172642e577972654e6f746512060a0461626364"
)

func TestAny(t *testing.T) {
	b, _ := hex.DecodeString(anyHex)
	aMsg := new(any.Any)
	err := proto.Unmarshal(b, aMsg)
	if err != nil {
		t.Error("unmarshal err: ", err)
	}
	jm := jsonpb.Marshaler{
		AnyResolver: BetterAnyResolver,
	}
	jstr, err := jm.MarshalToString(aMsg)
	if err != nil {
		t.Error("json marshal err: ", err)
	}
	t.Log(jstr)
}
