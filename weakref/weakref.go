package weakref

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

type WeakRef struct {
	t uintptr // type
	d uintptr // data
}

func NewWeakRef(v interface{}) *WeakRef {
	i := (*[2]uintptr)(unsafe.Pointer(&v))
	w := &WeakRef{^i[0], ^i[1]}
	runtime.SetFinalizer((*uintptr)(unsafe.Pointer(i[1])), func(_ *uintptr) {
		atomic.StoreUintptr(&w.d, 0)
		w.t = 0
	})
	return w
}

func (w *WeakRef) Get() (v interface{}) {
	t := w.t
	d := atomic.LoadUintptr(&w.d)
	if d != 0 {
		i := (*[2]uintptr)(unsafe.Pointer(&v))
		i[0] = ^t
		i[1] = ^d
	}
	return
}
