/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * util/intutil_test.go                                   *
 *                                                        *
 * intutil test for Go.                                   *
 *                                                        *
 * LastModified: Aug 17, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package util

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

func BenchmarkGetIntBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetIntBytes(int64(i))
		GetIntBytes(int64(-i))
		GetIntBytes(math.MaxInt32 - int64(i))
		GetIntBytes(math.MinInt32 + int64(i))
		GetIntBytes(math.MaxInt64 - int64(i))
		GetIntBytes(math.MinInt64 + int64(i))
	}
}

func BenchmarkGetUintBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetUintBytes(uint64(i))
		GetUintBytes(uint64(-i))
		GetUintBytes(math.MaxUint32 - uint64(i))
		GetUintBytes(math.MaxUint32 + uint64(i))
		GetUintBytes(math.MaxUint64 - uint64(i))
		GetUintBytes(math.MaxUint64 + uint64(i))
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(int64(i), 10)
		strconv.FormatInt(int64(-i), 10)
		strconv.FormatInt(math.MaxInt32-int64(i), 10)
		strconv.FormatInt(math.MinInt32+int64(i), 10)
		strconv.FormatInt(math.MaxInt64-int64(i), 10)
		strconv.FormatInt(math.MinInt64+int64(i), 10)
	}
}

func BenchmarkFormatUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatUint(uint64(i), 10)
		strconv.FormatUint(uint64(-i), 10)
		strconv.FormatUint(math.MaxUint32-uint64(i), 10)
		strconv.FormatUint(math.MaxUint32+uint64(i), 10)
		strconv.FormatUint(math.MaxUint64-uint64(i), 10)
		strconv.FormatUint(math.MaxUint64+uint64(i), 10)
	}
}

func BenchmarkGetIntBytesParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var i int64
		for pb.Next() {
			GetIntBytes(i)
			GetIntBytes(-i)
			GetIntBytes(math.MaxInt32 - i)
			GetIntBytes(math.MinInt32 + i)
			GetIntBytes(math.MaxInt64 - i)
			GetIntBytes(math.MinInt64 + i)
			i++
		}
	})
}

func BenchmarkGetUintBytesParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var i uint64
		for pb.Next() {
			GetUintBytes(i)
			GetUintBytes(-i)
			GetUintBytes(math.MaxUint32 - i)
			GetUintBytes(math.MaxUint32 + i)
			GetUintBytes(math.MaxUint64 - i)
			GetUintBytes(math.MaxUint64 + i)
			i++
		}
	})
}

func BenchmarkFormatIntParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var i int64
		for pb.Next() {
			strconv.FormatInt(i, 10)
			strconv.FormatInt(-i, 10)
			strconv.FormatInt(math.MaxInt32-i, 10)
			strconv.FormatInt(math.MinInt32+i, 10)
			strconv.FormatInt(math.MaxInt64-i, 10)
			strconv.FormatInt(math.MinInt64+i, 10)
			i++
		}
	})
}

func BenchmarkFormatUintParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var i uint64
		for pb.Next() {
			strconv.FormatUint(i, 10)
			strconv.FormatUint(-i, 10)
			strconv.FormatUint(math.MaxUint32-i, 10)
			strconv.FormatUint(math.MaxUint32+i, 10)
			strconv.FormatUint(math.MaxUint64-i, 10)
			strconv.FormatUint(math.MaxUint64+i, 10)
			i++
		}
	})
}

func TestGetIntBytes(t *testing.T) {
	data := []int64{
		0, 9, 10, 99, 100, 999, 1000, -1000, 10000, -10000,
		123456789, -123456789, math.MaxInt32, math.MinInt32,
		math.MaxInt64, math.MinInt64}
	for _, i := range data {
		b := GetIntBytes(i)
		if !reflect.DeepEqual(b, []byte(strconv.FormatInt(i, 10))) {
			t.Error("b must be []byte(\"" + strconv.FormatInt(i, 10) + "\")")
		}
	}
}

func TestGetUintBytes(t *testing.T) {
	data := []uint64{
		0, 9, 10, 99, 100, 999, 1000, 10000, 123456789,
		math.MaxInt32, math.MaxUint32, math.MaxInt64, math.MaxUint64}
	for _, i := range data {
		b := GetUintBytes(i)
		if !reflect.DeepEqual(b, []byte(strconv.FormatUint(i, 10))) {
			t.Error("b must be []byte(\"" + strconv.FormatUint(i, 10) + "\")")
		}
	}
}
