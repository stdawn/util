/**
 * @Time: 2023/9/20 15:40
 * @Author: LiuKun
 * @File: util_test.go
 * @Description:
 */

package util

import "testing"

func TestAll(t *testing.T) {

	s := "2255454455555556666"
	i := GetInt64FromV(s)
	t.Log(i)

	i1 := GetIntFromV(s)
	t.Log(i1)

}
