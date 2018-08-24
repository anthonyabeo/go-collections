package tests

import (
	"go-collections/counter"
	"sync"
	"testing"
)

var c = counter.GetInstance()

func BenchmarkIncrement1(b *testing.B) { benchmarkTest(1, b) }
func BenchmarkIncrement2(b *testing.B) { benchmarkTest(2, b) }
func BenchmarkIncrement3(b *testing.B) { benchmarkTest(3, b) }
func BenchmarkIncrement4(b *testing.B) { benchmarkTest(4, b) }
func BenchmarkIncrement5(b *testing.B) { benchmarkTest(5, b) }
func BenchmarkIncrement6(b *testing.B) { benchmarkTest(6, b) }
func BenchmarkIncrement7(b *testing.B) { benchmarkTest(7, b) }

func benchmarkTest(numGoroutines int, b *testing.B) {
	var inc sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		inc.Add(1)

		go func() {
			defer inc.Done()
			for n := 0; n < b.N; n++ {
				c.Increment()
			}
		}()
	}

	inc.Wait()
}
