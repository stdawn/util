/*
@Time : 2022/9/23 13:44
@Author : LiuKun
@File : if
@Software: GoLand
@Description:
*/

package util

// If If 三目表达，泛型
func If[T interface{}](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// IfString 三目运算,返回String
func IfString(condition bool, trueVal, falseVal string) string {
	return If(condition, trueVal, falseVal)
}

// IfInt 三目运算,返回Int
func IfInt(condition bool, trueVal, falseVal int) int {
	return If(condition, trueVal, falseVal)
}

// IfUInt64 三目运算,返回Uint64
func IfUInt64(condition bool, trueVal, falseVal uint64) uint64 {
	return If(condition, trueVal, falseVal)
}

// IfFloat 三目运算,返回Float64
func IfFloat(condition bool, trueVal, falseVal float64) float64 {
	return If(condition, trueVal, falseVal)
}
