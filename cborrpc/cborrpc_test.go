// Copyright 2015 Diffeo, Inc.
// This software is released under an MIT/X11 open source license.

package cborrpc

import (
	"github.com/satori/go.uuid"
	"github.com/ugorji/go/codec"
	"gopkg.in/check.v1"
	"reflect"
	"testing"
)

type Suite struct {
	cbor *codec.CborHandle
}

func (s *Suite) SetUpTest(c *check.C) {
	s.cbor = new(codec.CborHandle)
	err := SetExts(s.cbor)
	c.Assert(err, check.IsNil)
}

func (s *Suite) encoderTest(c *check.C, obj interface{}, expecteds ...[]byte) {
	var actual []byte
	encoder := codec.NewEncoderBytes(&actual, s.cbor)
	err := encoder.Encode(obj)
	c.Assert(err, check.IsNil)
	c.Check(actual, DeepEqualAny, expecteds)
}

func concat(slices ...[]byte) (result []byte) {
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return
}

func (s *Suite) TestRpcRequestToBytes(c *check.C) {
	req := Request{
		Method: "test",
		ID:     1,
		Params: []interface{}{},
	}
	// Since this is a map with 3 elements, there are 6 possible
	// orderings of iterating through it.  The map element pairs are
	method := []byte{
		// byte string "method"
		0x46, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
		// bytes "test"
		0x44, 0x74, 0x65, 0x73, 0x74,
	}
	id := []byte{
		// byte string "id"
		0x42, 0x69, 0x64,
		// positive integer 1
		0x01,
	}
	params := []byte{
		// byte string "params"
		0x46, 0x70, 0x61, 0x72, 0x61, 0x6D, 0x73,
		// array of length 0
		0x80,
	}
	// and the "header" is
	header := []byte{
		// tag 24
		0xD8, 0x18,
		// byte string of length 25
		0x58, 0x19,
		// map of 3 pairs
		0xA3,
	}
	// Also note that the string/bytes types above match what are
	// generated by the encoder, but are a little goofy.  Since
	// the Python 2 receiver for the most part doesn't care what
	// goes back, this doesn't matter, much, which suggests we
	// should make it consistent.
	expecteds := [][]byte{
		concat(header, method, id, params),
		concat(header, method, params, id),
		concat(header, id, method, params),
		concat(header, id, params, method),
		concat(header, params, method, id),
		concat(header, params, id, method),
	}
	s.encoderTest(c, req, expecteds...)
}

func (s *Suite) TestEmptyTupleToBytes(c *check.C) {
	tuple := PythonTuple{[]interface{}{}}
	expected := []byte{
		// tag 128
		0xD8, 0x80,
		// array of length 0
		0x80,
	}
	s.encoderTest(c, tuple, expected)
}

func (s *Suite) TestReallyEmptyTupleToBytes(c *check.C) {
	tuple := PythonTuple{}
	expected := []byte{
		// tag 128
		0xD8, 0x80,
		// array of length 0
		0x80,
	}
	s.encoderTest(c, tuple, expected)
}

func (s *Suite) TestListOfTupleToBytes(c *check.C) {
	tuple := PythonTuple{[]interface{}{}}
	list := []PythonTuple{tuple}
	expected := []byte{
		// array of length 1
		0x81,
		// tag 128
		0xD8, 0x80,
		// array of length 0
		0x80,
	}
	s.encoderTest(c, list, expected)
}

func (s *Suite) TestBytesToEmptyTuple(c *check.C) {
	bytes := []byte{
		// tag 128
		0xD8, 0x80,
		// array of length 0
		0x80,
	}
	expected := PythonTuple{[]interface{}{}}
	var actual PythonTuple
	encoder := codec.NewDecoderBytes(bytes, s.cbor)
	err := encoder.Decode(&actual)
	c.Assert(err, check.IsNil)
	c.Check(actual, check.DeepEquals, expected)
}

func (s *Suite) TestDecodeTupleReq(c *check.C) {
	bytes := []byte{
		// tag 24
		0xD8, 0x18,
		// byte string of length 31
		0x58, 0x1F,
		// map of 3 pairs
		0xA3,
		// byte string "method"
		0x46, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
		// bytes "test"
		0x44, 0x74, 0x65, 0x73, 0x74,
		// byte string "id"
		0x42, 0x69, 0x64,
		// positive integer 1
		0x01,
		// byte string "params"
		0x46, 0x70, 0x61, 0x72, 0x61, 0x6D, 0x73,
		// array of length 1
		0x81,
		// tag 128
		0xD8, 0x80,
		// array of length 2
		0x82,
		// string "k"
		0x61, 0x6B,
		// map of 0 pairs
		0xA0,
	}
	encoder := codec.NewDecoderBytes(bytes, s.cbor)
	var req Request
	err := encoder.Decode(&req)
	c.Assert(err, check.IsNil)
	c.Check(req.Method, check.Equals, "test")
	c.Check(req.ID, check.Equals, uint(1))
	c.Check(req.Params, check.HasLen, 1)
	if len(req.Params) > 0 {
		c.Check(req.Params[0], check.DeepEquals, PythonTuple{Items: []interface{}{"k", map[interface{}]interface{}{}}})
	}
}

func (s *Suite) TestEncodeUUID(c *check.C) {
	aUUID := uuid.NewV4()
	expected := []byte{
		// tag 37
		0xD8, 0x25,
		// byte string of length 16
		0x50,
	}
	expected = append(expected, aUUID.Bytes()...)
	s.encoderTest(c, aUUID, expected)
}

func (s *Suite) TestDecodeUUID(c *check.C) {
	bytes := []byte{
		// tag 37
		0xD8, 0x25,
		// byte string of length 16
		0x50,
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
	}
	expected := uuid.UUID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05,
		0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D,
		0x0E, 0x0F}
	var actual uuid.UUID
	encoder := codec.NewDecoderBytes(bytes, s.cbor)
	err := encoder.Decode(&actual)
	c.Assert(err, check.IsNil)
	c.Check(actual, check.DeepEquals, expected)
}

// deepEqualAny is a gocheck checker that passes if the provided value
// is reflect.DeepEqual any of a provided set of expected values.
type deepEqualAny struct {
	*check.CheckerInfo
}

func (checker *deepEqualAny) Info() *check.CheckerInfo {
	return checker.CheckerInfo
}

func (checker *deepEqualAny) Check(params []interface{}, names []string) (bool, string) {
	obtained := params[0]
	// We can't blindly cast params[1] to a []interface{}; we need
	// to do an intermediate reflect
	expecteds := reflect.ValueOf(params[1])
	if expecteds.Kind() != reflect.Slice {
		return false, "DeepEqualAny needs a slice of expecteds"
	}
	for i := 0; i < expecteds.Len(); i++ {
		expected := expecteds.Index(i).Interface()
		if reflect.DeepEqual(obtained, expected) {
			return true, ""
		}
	}
	return false, ""
}

var DeepEqualAny check.Checker = &deepEqualAny{
	&check.CheckerInfo{
		Name:   "DeepEqualAny",
		Params: []string{"obtained", "expecteds"},
	},
}

// gocheck boilerplate

func Test(t *testing.T) {
	check.TestingT(t)
}

func init() {
	check.Suite(&Suite{})
}
