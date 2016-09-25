// +build !goid

package gls

import (
	"runtime"
	"sync"
	"testing"
)

func BenchmarkSpawnShim(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Go(func() { wg.Done() })
	}
	wg.Wait()
}

func BenchmarkPointerResolution(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(1)
	f := func() {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			getPtr()
		}
		b.StopTimer()
		wg.Done()
	}
	Go(f)
	wg.Wait()
}

func BenchmarkPointerResolutionContention(b *testing.B) {
	var wg sync.WaitGroup

	f := func() {
		for i := 0; i < b.N; i++ {
			getPtr()
		}
		wg.Done()
	}

	b.ResetTimer()
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		Go(f)
	}
	wg.Wait()
}

func BenchmarkPutShim(b *testing.B) {
	nroutines := runtime.NumCPU() * 4
	var wg sync.WaitGroup
	wg.Add(nroutines)
	b.ResetTimer()
	for i := 0; i < nroutines; i++ {
		Go(func() {
			for i := 0; i < b.N; i++ {
				Put(1, 1)
			}
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkGetShim(b *testing.B) {
	nroutines := runtime.NumCPU() * 4
	var wg sync.WaitGroup
	wg.Add(nroutines)
	b.ResetTimer()
	for i := 0; i < nroutines; i++ {
		Go(func() {
			for i := 0; i < b.N; i++ {
				Get(1)
			}
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkDeleteShim(b *testing.B) {
	nroutines := runtime.NumCPU() * 4
	var wg sync.WaitGroup
	wg.Add(nroutines)
	b.ResetTimer()
	for i := 0; i < nroutines; i++ {
		Go(func() {
			for i := 0; i < b.N; i++ {
				Delete(1)
			}
			wg.Done()
		})
	}
	wg.Wait()
}
