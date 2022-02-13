package atom

import (
	"sync/atomic"
)

type (
	Bool    int32
	Int32   int32
	Int64   int64
	Uint32  uint32
	Uint64  uint64
	Uintptr uintptr
	Value   atomic.Value
)

// Bool

func (b *Bool) IsSet() bool {
	return atomic.LoadInt32((*int32)(b)) != 0
}

func (b *Bool) SetTrue() {
	atomic.StoreInt32((*int32)(b), 1)
}

func (b *Bool) SetFalse() {
	atomic.StoreInt32((*int32)(b), 0)
}

// Int32

func (addr *Int32) Add(delta int32) int32 {
	return atomic.AddInt32((*int32)(addr), delta)
}

func (addr *Int32) Load() int32 {
	return atomic.LoadInt32((*int32)(addr))
}

func (addr *Int32) Store(val int32) {
	atomic.StoreInt32((*int32)(addr), val)
}

func (addr *Int32) Swap(new int32) (old int32) {
	return atomic.SwapInt32((*int32)(addr), new)
}

func (addr *Int32) CompareAndSwap(old int32, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32((*int32)(addr), old, new)
}

// Int64

func (addr *Int64) Add(delta int64) int64 {
	return atomic.AddInt64((*int64)(addr), delta)
}

func (addr *Int64) Load() int64 {
	return atomic.LoadInt64((*int64)(addr))
}

func (addr *Int64) Store(val int64) {
	atomic.StoreInt64((*int64)(addr), val)
}

func (addr *Int64) Swap(new int64) (old int64) {
	return atomic.SwapInt64((*int64)(addr), new)
}

func (addr *Int64) CompareAndSwap(old int64, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64((*int64)(addr), old, new)
}

// Uint32

func (addr *Uint32) Add(delta uint32) uint32 {
	return atomic.AddUint32((*uint32)(addr), delta)
}

func (addr *Uint32) Sub(delta uint32) uint32 {
	return atomic.AddUint32((*uint32)(addr), ^(delta - 1))
}

func (addr *Uint32) Load() uint32 {
	return atomic.LoadUint32((*uint32)(addr))
}

func (addr *Uint32) Store(val uint32) {
	atomic.StoreUint32((*uint32)(addr), val)
}

func (addr *Uint32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32((*uint32)(addr), new)
}

func (addr *Uint32) CompareAndSwap(old uint32, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32((*uint32)(addr), old, new)
}

// Uint64

func (addr *Uint64) Add(delta uint64) uint64 {
	return atomic.AddUint64((*uint64)(addr), delta)
}

func (addr *Uint64) Sub(delta uint64) uint64 {
	return atomic.AddUint64((*uint64)(addr), ^(delta - 1))
}

func (addr *Uint64) Load() uint64 {
	return atomic.LoadUint64((*uint64)(addr))
}

func (addr *Uint64) Store(val uint64) {
	atomic.StoreUint64((*uint64)(addr), val)
}

func (addr *Uint64) Swap(new uint64) (old uint64) {
	return atomic.SwapUint64((*uint64)(addr), new)
}

func (addr *Uint64) CompareAndSwap(old uint64, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64((*uint64)(addr), old, new)
}

// Uintptr

func (addr *Uintptr) Add(delta uintptr) uintptr {
	return atomic.AddUintptr((*uintptr)(addr), delta)
}

func (addr *Uintptr) Load() uintptr {
	return atomic.LoadUintptr((*uintptr)(addr))
}

func (addr *Uintptr) Store(val uintptr) {
	atomic.StoreUintptr((*uintptr)(addr), val)
}

func (addr *Uintptr) Swap(new uintptr) (old uintptr) {
	return atomic.SwapUintptr((*uintptr)(addr), new)
}

func (addr *Uintptr) CompareAndSwap(old uintptr, new uintptr) (swapped bool) {
	return atomic.CompareAndSwapUintptr((*uintptr)(addr), old, new)
}
