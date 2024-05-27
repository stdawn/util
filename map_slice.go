/**
 * @Time: 2024/5/27 9:32
 * @Author: LiuKun
 * @File: map_slice.go
 * @Description:
 */

package util

// MapKeysSlice 获取map的keys
func MapKeysSlice[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}

// MapValuesSlice 获取map的values
func MapValuesSlice[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// SliceToMap 将slice转换为map
func SliceToMap[K comparable, V any](s []K, f func(k K) V) map[K]V {
	m := make(map[K]V)
	for _, k := range s {
		m[k] = f(k)
	}
	return m
}
