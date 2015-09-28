// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// MapOfBoolToString

type MapOfBoolToString struct {
	m types.Map
}

func NewMapOfBoolToString() MapOfBoolToString {
	return MapOfBoolToString{types.NewMap()}
}

type MapOfBoolToStringDef map[bool]string

func (def MapOfBoolToStringDef) New() MapOfBoolToString {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.Bool(k), types.NewString(v))
	}
	return MapOfBoolToString{types.NewMap(kv...)}
}

func (m MapOfBoolToString) Def() MapOfBoolToStringDef {
	def := make(map[bool]string)
	m.m.Iter(func(k, v types.Value) bool {
		def[bool(k.(types.Bool))] = v.(types.String).String()
		return false
	})
	return def
}

func MapOfBoolToStringFromVal(p types.Value) MapOfBoolToString {
	// TODO: Validate here
	return MapOfBoolToString{p.(types.Map)}
}

func (m MapOfBoolToString) NomsValue() types.Value {
	return m.m
}

func (m MapOfBoolToString) Equals(p MapOfBoolToString) bool {
	return m.m.Equals(p.m)
}

func (m MapOfBoolToString) Ref() ref.Ref {
	return m.m.Ref()
}

// A Noms Value that describes MapOfBoolToString.
var __typeRefForMapOfBoolToString = types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.BoolKind), types.MakePrimitiveTypeRef(types.StringKind))

func (m MapOfBoolToString) TypeRef() types.TypeRef {
	return __typeRefForMapOfBoolToString
}

func init() {
	types.RegisterFromValFunction(__typeRefForMapOfBoolToString, func(v types.Value) types.NomsValue {
		return MapOfBoolToStringFromVal(v)
	})
}

func (m MapOfBoolToString) Empty() bool {
	return m.m.Empty()
}

func (m MapOfBoolToString) Len() uint64 {
	return m.m.Len()
}

func (m MapOfBoolToString) Has(p bool) bool {
	return m.m.Has(types.Bool(p))
}

func (m MapOfBoolToString) Get(p bool) string {
	return m.m.Get(types.Bool(p)).(types.String).String()
}

func (m MapOfBoolToString) Set(k bool, v string) MapOfBoolToString {
	return MapOfBoolToString{m.m.Set(types.Bool(k), types.NewString(v))}
}

// TODO: Implement SetM?

func (m MapOfBoolToString) Remove(p bool) MapOfBoolToString {
	return MapOfBoolToString{m.m.Remove(types.Bool(p))}
}

type MapOfBoolToStringIterCallback func(k bool, v string) (stop bool)

func (m MapOfBoolToString) Iter(cb MapOfBoolToStringIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(bool(k.(types.Bool)), v.(types.String).String())
	})
}

type MapOfBoolToStringIterAllCallback func(k bool, v string)

func (m MapOfBoolToString) IterAll(cb MapOfBoolToStringIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(bool(k.(types.Bool)), v.(types.String).String())
	})
}

type MapOfBoolToStringFilterCallback func(k bool, v string) (keep bool)

func (m MapOfBoolToString) Filter(cb MapOfBoolToStringFilterCallback) MapOfBoolToString {
	nm := NewMapOfBoolToString()
	m.IterAll(func(k bool, v string) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// MapOfStringToValue

type MapOfStringToValue struct {
	m types.Map
}

func NewMapOfStringToValue() MapOfStringToValue {
	return MapOfStringToValue{types.NewMap()}
}

type MapOfStringToValueDef map[string]types.Value

func (def MapOfStringToValueDef) New() MapOfStringToValue {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v)
	}
	return MapOfStringToValue{types.NewMap(kv...)}
}

func (m MapOfStringToValue) Def() MapOfStringToValueDef {
	def := make(map[string]types.Value)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v
		return false
	})
	return def
}

func MapOfStringToValueFromVal(p types.Value) MapOfStringToValue {
	// TODO: Validate here
	return MapOfStringToValue{p.(types.Map)}
}

func (m MapOfStringToValue) NomsValue() types.Value {
	return m.m
}

func (m MapOfStringToValue) Equals(p MapOfStringToValue) bool {
	return m.m.Equals(p.m)
}

func (m MapOfStringToValue) Ref() ref.Ref {
	return m.m.Ref()
}

// A Noms Value that describes MapOfStringToValue.
var __typeRefForMapOfStringToValue = types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakePrimitiveTypeRef(types.ValueKind))

func (m MapOfStringToValue) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToValue
}

func init() {
	types.RegisterFromValFunction(__typeRefForMapOfStringToValue, func(v types.Value) types.NomsValue {
		return MapOfStringToValueFromVal(v)
	})
}

func (m MapOfStringToValue) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToValue) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToValue) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToValue) Get(p string) types.Value {
	return m.m.Get(types.NewString(p))
}

func (m MapOfStringToValue) Set(k string, v types.Value) MapOfStringToValue {
	return MapOfStringToValue{m.m.Set(types.NewString(k), v)}
}

// TODO: Implement SetM?

func (m MapOfStringToValue) Remove(p string) MapOfStringToValue {
	return MapOfStringToValue{m.m.Remove(types.NewString(p))}
}

type MapOfStringToValueIterCallback func(k string, v types.Value) (stop bool)

func (m MapOfStringToValue) Iter(cb MapOfStringToValueIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v)
	})
}

type MapOfStringToValueIterAllCallback func(k string, v types.Value)

func (m MapOfStringToValue) IterAll(cb MapOfStringToValueIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v)
	})
}

type MapOfStringToValueFilterCallback func(k string, v types.Value) (keep bool)

func (m MapOfStringToValue) Filter(cb MapOfStringToValueFilterCallback) MapOfStringToValue {
	nm := NewMapOfStringToValue()
	m.IterAll(func(k string, v types.Value) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}
