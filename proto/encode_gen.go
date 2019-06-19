// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Code generated by generate-types. DO NOT EDIT.

package proto

import (
	"math"
	"unicode/utf8"

	"google.golang.org/protobuf/internal/encoding/wire"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var wireTypes = map[protoreflect.Kind]wire.Type{
	protoreflect.BoolKind:     wire.VarintType,
	protoreflect.EnumKind:     wire.VarintType,
	protoreflect.Int32Kind:    wire.VarintType,
	protoreflect.Sint32Kind:   wire.VarintType,
	protoreflect.Uint32Kind:   wire.VarintType,
	protoreflect.Int64Kind:    wire.VarintType,
	protoreflect.Sint64Kind:   wire.VarintType,
	protoreflect.Uint64Kind:   wire.VarintType,
	protoreflect.Sfixed32Kind: wire.Fixed32Type,
	protoreflect.Fixed32Kind:  wire.Fixed32Type,
	protoreflect.FloatKind:    wire.Fixed32Type,
	protoreflect.Sfixed64Kind: wire.Fixed64Type,
	protoreflect.Fixed64Kind:  wire.Fixed64Type,
	protoreflect.DoubleKind:   wire.Fixed64Type,
	protoreflect.StringKind:   wire.BytesType,
	protoreflect.BytesKind:    wire.BytesType,
	protoreflect.MessageKind:  wire.BytesType,
	protoreflect.GroupKind:    wire.StartGroupType,
}

func (o MarshalOptions) marshalSingular(b []byte, fd protoreflect.FieldDescriptor, v protoreflect.Value) ([]byte, error) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		b = wire.AppendVarint(b, wire.EncodeBool(v.Bool()))
	case protoreflect.EnumKind:
		b = wire.AppendVarint(b, uint64(v.Enum()))
	case protoreflect.Int32Kind:
		b = wire.AppendVarint(b, uint64(int32(v.Int())))
	case protoreflect.Sint32Kind:
		b = wire.AppendVarint(b, wire.EncodeZigZag(int64(int32(v.Int()))))
	case protoreflect.Uint32Kind:
		b = wire.AppendVarint(b, uint64(uint32(v.Uint())))
	case protoreflect.Int64Kind:
		b = wire.AppendVarint(b, uint64(v.Int()))
	case protoreflect.Sint64Kind:
		b = wire.AppendVarint(b, wire.EncodeZigZag(v.Int()))
	case protoreflect.Uint64Kind:
		b = wire.AppendVarint(b, v.Uint())
	case protoreflect.Sfixed32Kind:
		b = wire.AppendFixed32(b, uint32(v.Int()))
	case protoreflect.Fixed32Kind:
		b = wire.AppendFixed32(b, uint32(v.Uint()))
	case protoreflect.FloatKind:
		b = wire.AppendFixed32(b, math.Float32bits(float32(v.Float())))
	case protoreflect.Sfixed64Kind:
		b = wire.AppendFixed64(b, uint64(v.Int()))
	case protoreflect.Fixed64Kind:
		b = wire.AppendFixed64(b, v.Uint())
	case protoreflect.DoubleKind:
		b = wire.AppendFixed64(b, math.Float64bits(v.Float()))
	case protoreflect.StringKind:
		if fd.Syntax() == protoreflect.Proto3 && !utf8.ValidString(v.String()) {
			return b, errors.InvalidUTF8(string(fd.FullName()))
		}
		b = wire.AppendBytes(b, []byte(v.String()))
	case protoreflect.BytesKind:
		b = wire.AppendBytes(b, v.Bytes())
	case protoreflect.MessageKind:
		var pos int
		var err error
		b, pos = appendSpeculativeLength(b)
		b, err = o.marshalMessage(b, v.Message())
		if err != nil {
			return b, err
		}
		b = finishSpeculativeLength(b, pos)
	case protoreflect.GroupKind:
		var err error
		b, err = o.marshalMessage(b, v.Message())
		if err != nil {
			return b, err
		}
		b = wire.AppendVarint(b, wire.EncodeTag(fd.Number(), wire.EndGroupType))
	default:
		return b, errors.New("invalid kind %v", fd.Kind())
	}
	return b, nil
}
