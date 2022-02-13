package atom_test

import (
	. "github.com/huyx/gou/atom"
	"testing"
)

const (
	magic32 = 0xdedbeef
	magic64 = 0xdeddeadbeefbeef
)

// Int32 test

func TestAddInt32(t *testing.T) {
	var x struct {
		before int32
		i      int32
		after  int32
	}
	x.before = magic32
	x.after = magic32
	var j int32
	for delta := int32(1); delta+delta > delta; delta += delta {
		k := (*Int32)(&x.i).Add(delta)
		j += delta
		if x.i != j || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

func TestLoadInt32(t *testing.T) {
	var x struct {
		before int32
		i      int32
		after  int32
	}
	x.before = magic32
	x.after = magic32
	for delta := int32(1); delta+delta > delta; delta += delta {
		k := (*Int32)(&x.i).Load()
		if k != x.i {
			t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
		}
		x.i += delta
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

func TestStoreInt32(t *testing.T) {
	var x struct {
		before int32
		i      int32
		after  int32
	}
	x.before = magic32
	x.after = magic32
	v := int32(0)
	for delta := int32(1); delta+delta > delta; delta += delta {
		(*Int32)(&x.i).Store(v)
		if x.i != v {
			t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
		}
		v += delta
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

func TestSwapInt32(t *testing.T) {
	var x struct {
		before int32
		i      int32
		after  int32
	}
	x.before = magic32
	x.after = magic32
	var j int32
	for delta := int32(1); delta+delta > delta; delta += delta {
		k := (*Int32)(&x.i).Swap(delta)
		if x.i != delta || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
		j = delta
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

func TestCompareAndSwapInt32(t *testing.T) {
	var x struct {
		before int32
		i      int32
		after  int32
	}
	x.before = magic32
	x.after = magic32
	for val := int32(1); val+val > val; val += val {
		x.i = val
		if !(*Int32)(&x.i).CompareAndSwap(val, val+1) {
			t.Fatalf("should have swapped %#x %#x", val, val+1)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
		}
		x.i = val + 1
		if (*Int32)(&x.i).CompareAndSwap(val, val+2) {
			t.Fatalf("should not have swapped %#x %#x", val, val+2)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
		}
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

// Int64 test

func TestAddInt64(t *testing.T) {
	var x struct {
		before int64
		i      int64
		after  int64
	}
	x.before = magic64
	x.after = magic64
	var j int64
	for delta := int64(1); delta+delta > delta; delta += delta {
		k := (*Int64)(&x.i).Add(delta)
		j += delta
		if x.i != j || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func TestLoadInt64(t *testing.T) {
	var x struct {
		before int64
		i      int64
		after  int64
	}
	x.before = magic64
	x.after = magic64
	for delta := int64(1); delta+delta > delta; delta += delta {
		k := (*Int64)(&x.i).Load()
		if k != x.i {
			t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
		}
		x.i += delta
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func TestStoreInt64(t *testing.T) {
	var x struct {
		before int64
		i      int64
		after  int64
	}
	x.before = magic64
	x.after = magic64
	v := int64(0)
	for delta := int64(1); delta+delta > delta; delta += delta {
		(*Int64)(&x.i).Store(v)
		if x.i != v {
			t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
		}
		v += delta
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func TestSwapInt64(t *testing.T) {
	var x struct {
		before int64
		i      int64
		after  int64
	}
	x.before = magic64
	x.after = magic64
	var j int64
	for delta := int64(1); delta+delta > delta; delta += delta {
		k := (*Int64)(&x.i).Swap(delta)
		if x.i != delta || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
		j = delta
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func TestCompareAndSwapInt64(t *testing.T) {
	var x struct {
		before int64
		i      int64
		after  int64
	}
	x.before = magic64
	x.after = magic64
	for val := int64(1); val+val > val; val += val {
		x.i = val
		if !(*Int64)(&x.i).CompareAndSwap(val, val+1) {
			t.Fatalf("should have swapped %#x %#x", val, val+1)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
		}
		x.i = val + 1
		if (*Int64)(&x.i).CompareAndSwap(val, val+2) {
			t.Fatalf("should not have swapped %#x %#x", val, val+2)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
		}
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

// Uint32 test

func TestAddUint32(t *testing.T) {
	var x struct {
		before uint32
		i      uint32
		after  uint32
	}
	x.before = magic32
	x.after = magic32
	var j uint32
	for delta := uint32(1); delta+delta > delta; delta += delta {
		k := (*Uint32)(&x.i).Add(delta)
		j += delta
		if x.i != j || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

func TestSubUint32(t *testing.T) {
	var x struct {
		before uint32
		i      uint32
		after  uint32
	}
	x.before = magic32
	x.after = magic32
	x.i = 0xffffffff
	j := x.i
	for delta := uint32(1); delta+delta > delta; delta += delta {
		k := (*Uint32)(&x.i).Sub(delta)
		j -= delta
		if x.i != j || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

func TestLoadUint32(t *testing.T) {
	var x struct {
		before uint32
		i      uint32
		after  uint32
	}
	x.before = magic32
	x.after = magic32
	for delta := uint32(1); delta+delta > delta; delta += delta {
		k := (*Uint32)(&x.i).Load()
		if k != x.i {
			t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
		}
		x.i += delta
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

func TestStoreUint32(t *testing.T) {
	var x struct {
		before uint32
		i      uint32
		after  uint32
	}
	x.before = magic32
	x.after = magic32
	v := uint32(0)
	for delta := uint32(1); delta+delta > delta; delta += delta {
		(*Uint32)(&x.i).Store(v)
		if x.i != v {
			t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
		}
		v += delta
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

func TestSwapUint32(t *testing.T) {
	var x struct {
		before uint32
		i      uint32
		after  uint32
	}
	x.before = magic32
	x.after = magic32
	var j uint32
	for delta := uint32(1); delta+delta > delta; delta += delta {
		k := (*Uint32)(&x.i).Swap(delta)
		if x.i != delta || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
		j = delta
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

func TestCompareAndSwapUint32(t *testing.T) {
	var x struct {
		before uint32
		i      uint32
		after  uint32
	}
	x.before = magic32
	x.after = magic32
	for val := uint32(1); val+val > val; val += val {
		x.i = val
		if !(*Uint32)(&x.i).CompareAndSwap(val, val+1) {
			t.Fatalf("should have swapped %#x %#x", val, val+1)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
		}
		x.i = val + 1
		if (*Uint32)(&x.i).CompareAndSwap(val, val+2) {
			t.Fatalf("should not have swapped %#x %#x", val, val+2)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
		}
	}
	if x.before != magic32 || x.after != magic32 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic32, magic32)
	}
}

// Uint64 test

func TestAddUint64(t *testing.T) {
	var x struct {
		before uint64
		i      uint64
		after  uint64
	}
	x.before = magic64
	x.after = magic64
	var j uint64
	for delta := uint64(1); delta+delta > delta; delta += delta {
		k := (*Uint64)(&x.i).Add(delta)
		j += delta
		if x.i != j || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func TestSubUint64(t *testing.T) {
	var x struct {
		before uint64
		i      uint64
		after  uint64
	}
	x.before = magic64
	x.after = magic64
	x.i = 0xffffffffffffffff
	j := x.i
	for delta := uint64(1); delta+delta > delta; delta += delta {
		k := (*Uint64)(&x.i).Sub(delta)
		j -= delta
		if x.i != j || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func TestLoadUint64(t *testing.T) {
	var x struct {
		before uint64
		i      uint64
		after  uint64
	}
	x.before = magic64
	x.after = magic64
	for delta := uint64(1); delta+delta > delta; delta += delta {
		k := (*Uint64)(&x.i).Load()
		if k != x.i {
			t.Fatalf("delta=%d i=%d k=%d", delta, x.i, k)
		}
		x.i += delta
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func TestStoreUint64(t *testing.T) {
	var x struct {
		before uint64
		i      uint64
		after  uint64
	}
	x.before = magic64
	x.after = magic64
	v := uint64(0)
	for delta := uint64(1); delta+delta > delta; delta += delta {
		(*Uint64)(&x.i).Store(v)
		if x.i != v {
			t.Fatalf("delta=%d i=%d v=%d", delta, x.i, v)
		}
		v += delta
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func TestSwapUint64(t *testing.T) {
	var x struct {
		before uint64
		i      uint64
		after  uint64
	}
	x.before = magic64
	x.after = magic64
	var j uint64
	for delta := uint64(1); delta+delta > delta; delta += delta {
		k := (*Uint64)(&x.i).Swap(delta)
		if x.i != delta || k != j {
			t.Fatalf("delta=%d i=%d j=%d k=%d", delta, x.i, j, k)
		}
		j = delta
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}

func TestCompareAndSwapUint64(t *testing.T) {
	var x struct {
		before uint64
		i      uint64
		after  uint64
	}
	x.before = magic64
	x.after = magic64
	for val := uint64(1); val+val > val; val += val {
		x.i = val
		if !(*Uint64)(&x.i).CompareAndSwap(val, val+1) {
			t.Fatalf("should have swapped %#x %#x", val, val+1)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
		}
		x.i = val + 1
		if (*Uint64)(&x.i).CompareAndSwap(val, val+2) {
			t.Fatalf("should not have swapped %#x %#x", val, val+2)
		}
		if x.i != val+1 {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, val+1)
		}
	}
	if x.before != magic64 || x.after != magic64 {
		t.Fatalf("wrong magic: %#x _ %#x != %#x _ %#x", x.before, x.after, magic64, magic64)
	}
}
