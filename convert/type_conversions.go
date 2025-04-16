package convert

import (
	"fmt"
	"strconv"
)

func Convert_to_float_64(input any) float64 {
	switch v := input.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case int32:
		return float64(v)
	case int16:
		return float64(v)
	case int8:
		return float64(v)
	case uint:
		return float64(v)
	case uint64:
		return float64(v)
	case uint32:
		return float64(v)
	case uint16:
		return float64(v)
	case uint8:
		return float64(v)
	case string:
		if val, err := strconv.ParseFloat(v, 64); err == nil {
			return val
		}
	default:
		// Fallback: try to convert to string and parse
		if s := fmt.Sprintf("%v", input); s != "" {
			if val, err := strconv.ParseFloat(s, 64); err == nil {
				return val
			}
		}
	}
	return 0.0
}

func Convert_to_int_64(value any) int64 {
	return value.(int64)
}
