package compare

import (
	"fmt"
	"strings"

	"github.com/xgfone/go-tools/extremum"
)

// CompareSlice is the same as Compare, but only compare the slice.
func CompareSlice(v1, v2 interface{}) int {
	switch _v1 := v1.(type) {
	case []int:
		if _v2, ok := v2.([]int); !ok {
			panic("Type is not the same")
		} else {
			return CompareIntSlice(_v1, _v2)
		}
	case []uint:
		if _v2, ok := v2.([]uint); !ok {
			panic("Type is not the same")
		} else {
			return CompareUintSlice(_v1, _v2)
		}
	case []int8:
		if _v2, ok := v2.([]int8); !ok {
			panic("Type is not the same")
		} else {
			return CompareInt8Slice(_v1, _v2)
		}
	case []uint8:
		if _v2, ok := v2.([]uint8); !ok {
			panic("Type is not the same")
		} else {
			return CompareUint8Slice(_v1, _v2)
		}
	case []int16:
		if _v2, ok := v2.([]int16); !ok {
			panic("Type is not the same")
		} else {
			return CompareInt16Slice(_v1, _v2)
		}
	case []uint16:
		if _v2, ok := v2.([]uint16); !ok {
			panic("Type is not the same")
		} else {
			return CompareUint16Slice(_v1, _v2)
		}
	case []int32:
		if _v2, ok := v2.([]int32); !ok {
			panic("Type is not the same")
		} else {
			return CompareInt32Slice(_v1, _v2)
		}
	case []uint32:
		if _v2, ok := v2.([]uint32); !ok {
			panic("Type is not the same")
		} else {
			return CompareUint32Slice(_v1, _v2)
		}
	case []int64:
		if _v2, ok := v2.([]int64); !ok {
			panic("Type is not the same")
		} else {
			return CompareInt64Slice(_v1, _v2)
		}
	case []uint64:
		if _v2, ok := v2.([]uint64); !ok {
			panic("Type is not the same")
		} else {
			return CompareUint64Slice(_v1, _v2)
		}
	case []string:
		if _v2, ok := v2.([]string); !ok {
			panic("Type is not the same")
		} else {
			return CompareStringSlice(_v1, _v2)
		}
	case []float32:
		if _v2, ok := v2.([]float32); !ok {
			panic("Type is not the same")
		} else {
			return CompareFloat32Slice(_v1, _v2)
		}
	case []float64:
		if _v2, ok := v2.([]float64); !ok {
			panic("Type is not the same")
		} else {
			return CompareFloat64Slice(_v1, _v2)
		}
	default:
		panic(fmt.Sprintf("Type is not supported: %v\n", _v1))
	}
}

// CompareIntSlice is the same as CompareSlice, but only compare the int slice.
func CompareIntSlice(v1, v2 []int) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareUintSlice is the same as CompareSlice, but only compare the uint slice.
func CompareUintSlice(v1, v2 []uint) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareInt8Slice is the same as CompareSlice, but only compare the int8 slice.
func CompareInt8Slice(v1, v2 []int8) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareUint8Slice is the same as CompareSlice, but only compare the uint8 slice.
func CompareUint8Slice(v1, v2 []uint8) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareInt16Slice is the same as CompareSlice, but only compare the int16 slice.
func CompareInt16Slice(v1, v2 []int16) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareUint16Slice is the same as CompareSlice, but only compare the uint16 slice.
func CompareUint16Slice(v1, v2 []uint16) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareInt32Slice is the same as CompareSlice, but only compare the int32 slice.
func CompareInt32Slice(v1, v2 []int32) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareUint32Slice is the same as CompareSlice, but only compare the uint32 slice.
func CompareUint32Slice(v1, v2 []uint32) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareInt16Slice is the same as CompareSlice, but only compare the int64 slice.
func CompareInt64Slice(v1, v2 []int64) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareUint64Slice is the same as CompareSlice, but only compare the uint64 slice.
func CompareUint64Slice(v1, v2 []uint64) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareFloat32Slice is the same as CompareSlice, but only compare the float32 slice.
func CompareFloat32Slice(v1, v2 []float32) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareFloat64Slice is the same as CompareSlice, but only compare the float64 slice.
func CompareFloat64Slice(v1, v2 []float64) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}

// CompareStringSlice is the same as CompareSlice, but only compare the string slice.
func CompareStringSlice(v1, v2 []string) int {
	len1, len2 := len(v1), len(v2)
	_len := extremum.MinInt(len1, len2)
	for i := 0; i < _len; i++ {
		diff := strings.Compare(v1[i], v2[i])
		if diff != 0 {
			return diff
		}
	}

	if len1 == len2 {
		return 0
	} else if len1 < len2 {
		return -1
	} else {
		return 1
	}
}
