/*
@Time : 2022/9/23 13:45
@Author : LiuKun
@File : value
@Software: GoLand
@Description:
*/

package util

import (
	"encoding/binary"
	"strconv"
	"strings"
)

func GetInt64FromMapByDimKey(data map[string]interface{}, dimKey string) int64 {
	for k, v := range data {
		if strings.Contains(k, dimKey) {
			return GetInt64FromV(v)
		}
	}
	return 0
}

func GetIntFromMapByDimKey(data map[string]interface{}, dimKey string) int {
	for k, v := range data {
		if strings.Contains(k, dimKey) {
			return GetIntFromV(v)
		}
	}
	return 0
}

func GetFloatFromMapByDimKey(data map[string]interface{}, dimKey string) float64 {
	for k, v := range data {
		if strings.Contains(k, dimKey) {
			return GetFloatFromV(v)
		}
	}
	return 0
}

func GetStringFromMapByDimKey(data map[string]interface{}, dimKey string) string {
	for k, v := range data {
		if strings.Contains(k, dimKey) {
			return GetStringFromV(v)
		}
	}
	return ""
}

func GetIntFromMap(data map[string]interface{}, key string, otherKeys ...string) int {

	keys := []string{key}
	if len(otherKeys) > 0 {
		keys = append(keys, otherKeys...)
	}

	for _, k := range keys {
		v, ok := data[k]
		if ok {
			return GetIntFromV(v)
		}
	}
	return 0
}

func GetInt64FromMap(data map[string]interface{}, key string, otherKeys ...string) int64 {

	keys := []string{key}
	if len(otherKeys) > 0 {
		keys = append(keys, otherKeys...)
	}

	for _, k := range keys {
		v, ok := data[k]
		if ok {
			return GetInt64FromV(v)
		}
	}

	return 0
}

func GetFloatFromMap(data map[string]interface{}, key string, otherKeys ...string) float64 {

	keys := []string{key}
	if len(otherKeys) > 0 {
		keys = append(keys, otherKeys...)
	}
	for _, k := range keys {
		v, ok := data[k]
		if ok {
			return GetFloatFromV(v)
		}
	}
	return 0.0
}

func GetStringFromMap(data map[string]interface{}, key string, otherKeys ...string) string {

	keys := []string{key}
	if len(otherKeys) > 0 {
		keys = append(keys, otherKeys...)
	}
	for _, k := range keys {
		v, ok := data[k]
		if ok {
			return GetStringFromV(v)
		}
	}
	return ""
}

func GetIMapFromMap(data map[string]interface{}, key string, otherKeys ...string) map[string]interface{} {

	keys := []string{key}
	if len(otherKeys) > 0 {
		keys = append(keys, otherKeys...)
	}
	for _, k := range keys {
		v, ok := data[k]
		if ok {
			return GetIMapFromV(v)
		}
	}
	return make(map[string]interface{})
}

func GetISliceFromMap(data map[string]interface{}, key string, otherKeys ...string) []interface{} {

	keys := []string{key}
	if len(otherKeys) > 0 {
		keys = append(keys, otherKeys...)
	}

	for _, k := range keys {
		v, ok := data[k]
		if ok {
			return GetISliceFromV(v)
		}
	}
	return make([]interface{}, 0)
}

func GetIntFromSlice(data []interface{}, index int) int {
	if index >= len(data) {
		return 0
	}
	return GetIntFromV(data[index])
}

func GetInt64FromSlice(data []interface{}, index int) int64 {
	if index >= len(data) {
		return 0
	}
	return GetInt64FromV(data[index])
}

func GetFloatFromSlice(data []interface{}, index int) float64 {
	if index >= len(data) {
		return 0.0
	}
	return GetFloatFromV(data[index])
}

func GetStringFromSlice(data []interface{}, index int) string {
	if index >= len(data) {
		return ""
	}
	return GetStringFromV(data[index])
}

func GetIntFromV(v interface{}) int {

	switch value := v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return intValue(value)
	case float32:
		return int(value)
	case float64:
		return int(value)
	case string:
		intV, e := strconv.Atoi(value)
		if e != nil {
			return 0
		}
		return intV
	case []byte:
		return int(binary.LittleEndian.Uint64(value))
	case bool:
		if value {
			return 1
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

func GetInt64FromV(v interface{}) int64 {

	switch value := v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return int64Value(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case string:
		intV, e := strconv.ParseInt(value, 10, 64)
		if e != nil {
			return 0
		}
		return intV
	case []byte:
		return int64(binary.LittleEndian.Uint64(value))
	case bool:
		if value {
			return 1
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

func GetFloatFromV(v interface{}) float64 {
	switch value := v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return float64(intValue(value))
	case float32:
		return float64(value)
	case float64:
		return value
	case string:

		strV := strings.ReplaceAll(value, "+", "")

		if strings.HasSuffix(value, "%") {
			strV = strings.ReplaceAll(value, "%", "")
			floatV, e := strconv.ParseFloat(strV, 64)
			if e != nil {
				return 0.0
			}
			return floatV / 100
		}
		floatV, e := strconv.ParseFloat(strV, 64)
		if e != nil {
			return 0.0
		}
		return floatV
	case []byte:
		return float64(binary.LittleEndian.Uint64(value))
	case bool:
		if value {
			return 1.0
		}
		return 0.0
	case nil:
		return 0.0
	default:
		return 0.0
	}
}

func GetStringFromV(v interface{}) string {
	switch value := v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return strconv.Itoa(intValue(value))
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case string:
		return strings.TrimSpace(value)
	case []byte:
		return string(value)
	case bool:
		return strconv.FormatBool(value)
	case nil:
		return ""
	default:
		return ""
	}
}

func GetIMapFromV(v interface{}) map[string]interface{} {
	m, ok := v.(map[string]interface{})
	if ok {
		return m
	}
	return make(map[string]interface{})
}

func GetISliceFromV(v interface{}) []interface{} {
	s, ok := v.([]interface{})
	if ok {
		return s
	}
	return make([]interface{}, 0)
}

func intValue(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	}
	return 0
}

func int64Value(value interface{}) int64 {
	switch v := value.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	}
	return 0
}
