package protorand

//go:generate go run github.com/bufbuild/buf/cmd/buf@v1.11.0 generate

import (
	"testing"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"

	testpb "github.com/sryoya/protorand/testdata"
)

func TestEmbedValues(t *testing.T) {
	p := New()
	p.Seed(0)

	input := &testpb.TestMessage{}
	res, err := p.Gen(input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// check if the input is not mutated.
	if !proto.Equal(input, &testpb.TestMessage{}) {
		t.Errorf("The input was unexpectedly mutated.")
	}

	got := res.(*testpb.TestMessage)

	// assert all the fields got some value
	if got.SomeInt32 == 0 {
		t.Errorf("Field SomeInt32 is not set")
	}
	if got.SomeSint32 == 0 {
		t.Errorf("Field SomeSint32 is not set")
	}
	if got.SomeUint32 == 0 {
		t.Errorf("Field SomeUint32 is not set")
	}
	if got.SomeFloat32 == 0 {
		t.Errorf("Field SomeFloat32 is not set")
	}
	if got.SomeFixed32 == 0 {
		t.Errorf("Field SomeFixed32 is not set")
	}
	if got.SomeSfixed32 == 0 {
		t.Errorf("Field SomeSfixed32 is not set")
	}
	if got.SomeInt64 == 0 {
		t.Errorf("Field SomeInt64 is not set")
	}
	if got.SomeSint64 == 0 {
		t.Errorf("Field SomeSint64 is not set")
	}
	if got.SomeUint64 == 0 {
		t.Errorf("Field SomeUint64 is not set")
	}
	if got.SomeFloat64 == 0 {
		t.Errorf("Field SomeFloat64 is not set")
	}
	if got.SomeFixed64 == 0 {
		t.Errorf("Field SomeFixed64 is not set")
	}
	if got.SomeSfixed64 == 0 {
		t.Errorf("Field SomeSfixed64 is not set")
	}
	if got.SomeStr == "" {
		t.Errorf("Field SomeStr is not set")
	}
	if got.SomeMsg == nil {
		t.Errorf("Field SomeMsg is not set")
	}
	if got.SomeMsg.SomeInt == 0 {
		t.Errorf("Field SomeMsg.SomeInt is not set")
	}
	if got.SomeMsg.SubChild == nil {
		t.Errorf("Field SomeMsg.SubChild is not set")
	}
	if got.SomeMsg.SubChild.SomeInt == 0 {
		t.Errorf("Field SomeMsg.SubChild.SomeInt is not set")
	}
	if len(got.SomeSlice) == 0 {
		t.Errorf("Field SomeSlice is not set")
	}
	if len(got.SomeMsgs) == 0 {
		t.Errorf("Field SomeMsgs is not set")
	}
	if len(got.SomeMap) == 0 {
		t.Errorf("Field SomeMap is not set")
	}
	if got.SomeEnum > 3 { // undeclared enum value
		t.Errorf("Field SomeEnum is not set")
	}
	if got.SomeEnum2 > 1 { // undeclared enum value
		t.Errorf("Field SomeEnum2 is not set")
	}
	if got.SomeOneOf == nil {
		t.Errorf("Field SomeOneOf is not set")
	}
	if got.Timestamp == nil {
		t.Errorf("Field Timestamp is not set")
	}
	if got.Duration == nil {
		t.Errorf("Field Duration is not set")
	}
}

func TestEnums(t *testing.T) {
	p := New()
	seenEnumValues := uint32(0)
	allValuesMask := uint32(1<<testpb.SomeEnum(0).Descriptor().Values().Len()) - 1
	for i := 0; i < 1000; i++ {
		res, err := p.Gen(&testpb.TestMessage{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		seenEnumValues |= (1 << res.(*testpb.TestMessage).SomeEnum)
		if seenEnumValues == allValuesMask {
			return
		}
	}
	t.Errorf("all enum values were not observed")
}

func TestOneofs(t *testing.T) {
	p := New()
	seenOneofValues := uint32(0)
	oneofDesc := (*testpb.TestMessage)(nil).ProtoReflect().Descriptor().Oneofs().ByName("some_one_of")
	allValuesMask := uint32(1<<oneofDesc.Fields().Len()) - 1
	indexesByNumber := map[protowire.Number]int{}
	for i := 0; i < oneofDesc.Fields().Len(); i++ {
		indexesByNumber[oneofDesc.Fields().Get(i).Number()] = i
	}
	for i := 0; i < 1000; i++ {
		res, err := p.Gen(&testpb.TestMessage{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		seenOneofValues |= (1 << indexesByNumber[res.(*testpb.TestMessage).ProtoReflect().WhichOneof(oneofDesc).Number()])
		if seenOneofValues == allValuesMask {
			return
		}
	}
	t.Errorf("all oneof values were not observed")
}
