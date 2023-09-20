/*
@Time : 2022/9/23 13:47
@Author : LiuKun
@File : math
@Software: GoLand
@Description:
*/

package util

import "math"

// FloatAvg Float的平均值
func FloatAvg(fs ...float64) float64 {
	if len(fs) == 0 {
		return 0
	}
	if len(fs) == 1 {
		return fs[0]
	}

	sum := 0.0
	for _, f := range fs {
		sum += f
	}
	return sum / float64(len(fs))
}

// FloatMax Float取最大值
func FloatMax(fs ...float64) float64 {
	if len(fs) == 0 {
		return 0
	}
	if len(fs) == 1 {
		return fs[0]
	}

	return math.Max(FloatMax(fs[1:]...), fs[0])
}

// FloatMin Float取最小值
func FloatMin(fs ...float64) float64 {
	if len(fs) == 0 {
		return 0
	}
	if len(fs) == 1 {
		return fs[0]
	}

	return math.Min(FloatMin(fs[1:]...), fs[0])
}

// IntMax Int取最大值 *
func IntMax(is ...int) int {
	fs := make([]float64, 0)
	for _, i := range is {
		fs = append(fs, float64(i))
	}
	return int(FloatMax(fs...))
}

// IntMin  Int取最小值
func IntMin(is ...int) int {
	fs := make([]float64, 0)
	for _, i := range is {
		fs = append(fs, float64(i))
	}
	return int(FloatMin(fs...))
}

// GetPercentFloat 数量百分比, Float
func GetPercentFloat(single, total float64) float64 {
	if math.Abs(total) < 1e-6 {
		return 0
	}

	return single / total
}

// GetPercent 数量百分比, Int
func GetPercent(single, total int) float64 {
	return GetPercentFloat(float64(single), float64(total))
}

// CompareInt  比较Int大小, i1<i2:-1, i1=i2:0, i1>i2:1,
func CompareInt(i1, i2 int) int {
	return CompareFloat(float64(i1), float64(i2))
}

// CompareFloat 比较Float64大小, f1<f2:-1, f1=f2:0, f1>f2:1,
func CompareFloat(f1, f2 float64) int {

	if f1 < f2 {
		return -1
	}
	if f1 > f2 {
		return 1
	}
	return 0
}

// SolveBinaryEquations 解二元一次方程组, 克莱姆法则
func SolveBinaryEquations(a1, b1, c1, a2, b2, c2 float64) (float64, float64, bool) {
	d := a1*b2 - a2*b1
	if math.Abs(d) < 1e-5 {
		return 0, 0, false
	}

	return (c1*b2 - c2*b1) / d, (a1*c2 - a2*c1) / d, true
}

// GetKxAddBValue y = kx + b
func GetKxAddBValue(k, b, x float64) float64 {
	return k*x + b
}

// GetPrice 获取增加涨跌幅后的价格(2位小数)
func GetPrice(price float64, addRange float64) float64 {
	return math.Round(price*(1+addRange)*100) / 100
}
