// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

import (
	"bytes"
)

type float64Ge struct {
	V1 Float64
	V2 Float64
}

func (this *float64Ge) Eval(buf []byte) bool {
	return this.V1.Eval(buf) >= this.V2.Eval(buf)
}

func init() {
	Register("ge", new(float64Ge))
}

type float32Ge struct {
	V1 Float32
	V2 Float32
}

func (this *float32Ge) Eval(buf []byte) bool {
	return this.V1.Eval(buf) >= this.V2.Eval(buf)
}

func init() {
	Register("ge", new(float32Ge))
}

type int64Ge struct {
	V1 Int64
	V2 Int64
}

func (this *int64Ge) Eval(buf []byte) bool {
	return this.V1.Eval(buf) >= this.V2.Eval(buf)
}

func init() {
	Register("ge", new(int64Ge))
}

type uint64Ge struct {
	V1 Uint64
	V2 Uint64
}

func (this *uint64Ge) Eval(buf []byte) bool {
	return this.V1.Eval(buf) >= this.V2.Eval(buf)
}

func init() {
	Register("ge", new(uint64Ge))
}

type int32Ge struct {
	V1 Int32
	V2 Int32
}

func (this *int32Ge) Eval(buf []byte) bool {
	return this.V1.Eval(buf) >= this.V2.Eval(buf)
}

func init() {
	Register("ge", new(int32Ge))
}

type uint32Ge struct {
	V1 Uint32
	V2 Uint32
}

func (this *uint32Ge) Eval(buf []byte) bool {
	return this.V1.Eval(buf) >= this.V2.Eval(buf)
}

func init() {
	Register("ge", new(uint32Ge))
}

type bytesGe struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesGe) Eval(buf []byte) bool {
	return bytes.Compare(this.V1.Eval(buf), this.V2.Eval(buf)) >= 0
}

func init() {
	Register("ge", new(bytesGe))
}

type float64Gt struct {
	V1 Float64
	V2 Float64
}

func (this *float64Gt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) > this.V2.Eval(buf)
}

func init() {
	Register("gt", new(float64Gt))
}

type float32Gt struct {
	V1 Float32
	V2 Float32
}

func (this *float32Gt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) > this.V2.Eval(buf)
}

func init() {
	Register("gt", new(float32Gt))
}

type int64Gt struct {
	V1 Int64
	V2 Int64
}

func (this *int64Gt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) > this.V2.Eval(buf)
}

func init() {
	Register("gt", new(int64Gt))
}

type uint64Gt struct {
	V1 Uint64
	V2 Uint64
}

func (this *uint64Gt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) > this.V2.Eval(buf)
}

func init() {
	Register("gt", new(uint64Gt))
}

type int32Gt struct {
	V1 Int32
	V2 Int32
}

func (this *int32Gt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) > this.V2.Eval(buf)
}

func init() {
	Register("gt", new(int32Gt))
}

type uint32Gt struct {
	V1 Uint32
	V2 Uint32
}

func (this *uint32Gt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) > this.V2.Eval(buf)
}

func init() {
	Register("gt", new(uint32Gt))
}

type bytesGt struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesGt) Eval(buf []byte) bool {
	return bytes.Compare(this.V1.Eval(buf), this.V2.Eval(buf)) > 0
}

func init() {
	Register("gt", new(bytesGt))
}

type float64Le struct {
	V1 Float64
	V2 Float64
}

func (this *float64Le) Eval(buf []byte) bool {
	return this.V1.Eval(buf) <= this.V2.Eval(buf)
}

func init() {
	Register("le", new(float64Le))
}

type float32Le struct {
	V1 Float32
	V2 Float32
}

func (this *float32Le) Eval(buf []byte) bool {
	return this.V1.Eval(buf) <= this.V2.Eval(buf)
}

func init() {
	Register("le", new(float32Le))
}

type int64Le struct {
	V1 Int64
	V2 Int64
}

func (this *int64Le) Eval(buf []byte) bool {
	return this.V1.Eval(buf) <= this.V2.Eval(buf)
}

func init() {
	Register("le", new(int64Le))
}

type uint64Le struct {
	V1 Uint64
	V2 Uint64
}

func (this *uint64Le) Eval(buf []byte) bool {
	return this.V1.Eval(buf) <= this.V2.Eval(buf)
}

func init() {
	Register("le", new(uint64Le))
}

type int32Le struct {
	V1 Int32
	V2 Int32
}

func (this *int32Le) Eval(buf []byte) bool {
	return this.V1.Eval(buf) <= this.V2.Eval(buf)
}

func init() {
	Register("le", new(int32Le))
}

type uint32Le struct {
	V1 Uint32
	V2 Uint32
}

func (this *uint32Le) Eval(buf []byte) bool {
	return this.V1.Eval(buf) <= this.V2.Eval(buf)
}

func init() {
	Register("le", new(uint32Le))
}

type bytesLe struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesLe) Eval(buf []byte) bool {
	return bytes.Compare(this.V1.Eval(buf), this.V2.Eval(buf)) <= 0
}

func init() {
	Register("le", new(bytesLe))
}

type float64Lt struct {
	V1 Float64
	V2 Float64
}

func (this *float64Lt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) < this.V2.Eval(buf)
}

func init() {
	Register("lt", new(float64Lt))
}

type float32Lt struct {
	V1 Float32
	V2 Float32
}

func (this *float32Lt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) < this.V2.Eval(buf)
}

func init() {
	Register("lt", new(float32Lt))
}

type int64Lt struct {
	V1 Int64
	V2 Int64
}

func (this *int64Lt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) < this.V2.Eval(buf)
}

func init() {
	Register("lt", new(int64Lt))
}

type uint64Lt struct {
	V1 Uint64
	V2 Uint64
}

func (this *uint64Lt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) < this.V2.Eval(buf)
}

func init() {
	Register("lt", new(uint64Lt))
}

type int32Lt struct {
	V1 Int32
	V2 Int32
}

func (this *int32Lt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) < this.V2.Eval(buf)
}

func init() {
	Register("lt", new(int32Lt))
}

type uint32Lt struct {
	V1 Uint32
	V2 Uint32
}

func (this *uint32Lt) Eval(buf []byte) bool {
	return this.V1.Eval(buf) < this.V2.Eval(buf)
}

func init() {
	Register("lt", new(uint32Lt))
}

type bytesLt struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesLt) Eval(buf []byte) bool {
	return bytes.Compare(this.V1.Eval(buf), this.V2.Eval(buf)) < 0
}

func init() {
	Register("lt", new(bytesLt))
}

type float64Eq struct {
	V1 Float64
	V2 Float64
}

func (this *float64Eq) Eval(buf []byte) bool {
	return this.V1.Eval(buf) == this.V2.Eval(buf)
}

func init() {
	Register("eq", new(float64Eq))
}

type float32Eq struct {
	V1 Float32
	V2 Float32
}

func (this *float32Eq) Eval(buf []byte) bool {
	return this.V1.Eval(buf) == this.V2.Eval(buf)
}

func init() {
	Register("eq", new(float32Eq))
}

type int64Eq struct {
	V1 Int64
	V2 Int64
}

func (this *int64Eq) Eval(buf []byte) bool {
	return this.V1.Eval(buf) == this.V2.Eval(buf)
}

func init() {
	Register("eq", new(int64Eq))
}

type uint64Eq struct {
	V1 Uint64
	V2 Uint64
}

func (this *uint64Eq) Eval(buf []byte) bool {
	return this.V1.Eval(buf) == this.V2.Eval(buf)
}

func init() {
	Register("eq", new(uint64Eq))
}

type int32Eq struct {
	V1 Int32
	V2 Int32
}

func (this *int32Eq) Eval(buf []byte) bool {
	return this.V1.Eval(buf) == this.V2.Eval(buf)
}

func init() {
	Register("eq", new(int32Eq))
}

type uint32Eq struct {
	V1 Uint32
	V2 Uint32
}

func (this *uint32Eq) Eval(buf []byte) bool {
	return this.V1.Eval(buf) == this.V2.Eval(buf)
}

func init() {
	Register("eq", new(uint32Eq))
}

type boolEq struct {
	V1 Bool
	V2 Bool
}

func (this *boolEq) Eval(buf []byte) bool {
	return this.V1.Eval(buf) == this.V2.Eval(buf)
}

func init() {
	Register("eq", new(boolEq))
}

type stringEq struct {
	V1 String
	V2 String
}

func (this *stringEq) Eval(buf []byte) bool {
	return this.V1.Eval(buf) == this.V2.Eval(buf)
}

func init() {
	Register("eq", new(stringEq))
}

type bytesEq struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesEq) Eval(buf []byte) bool {
	return bytes.Equal(this.V1.Eval(buf), this.V2.Eval(buf))
}

func init() {
	Register("eq", new(bytesEq))
}