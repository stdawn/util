/**
 * @Time: 2023/5/15 14:13
 * @Author: LiuKun
 * @File: parse.go
 * @Description:
 */

package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ParseCharlesJson 解析CharlesJson, 根据tagUrl、tagMethod返回response
func ParseCharlesJson(jsonStr, tagUrl string, tagMethod string) ([]string, error) {

	var arrayResult []interface{}
	err := json.Unmarshal([]byte(jsonStr), &arrayResult)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("JsonToArray err: %s", err.Error()))
	}

	responseArray := make([]string, 0)

	for _, v := range arrayResult {
		mp := GetIMapFromV(v)
		method := GetStringFromMap(mp, "method")
		if method != tagMethod {
			continue
		}
		scheme := GetStringFromMap(mp, "scheme")
		host := GetStringFromMap(mp, "host")
		path := GetStringFromMap(mp, "path")
		urlStr := fmt.Sprintf("%s://%s%s", scheme, host, path)
		if urlStr != tagUrl {
			continue
		}

		if GetIntFromMap(GetIMapFromMap(mp, "response"), "status") == 200 {
			responseStr := GetStringFromMap(GetIMapFromMap(GetIMapFromMap(mp, "response"), "body"), "text")
			responseArray = append(responseArray, responseStr)
		}
	}
	return responseArray, nil
}

// ParseBrowserHarJson 解析从浏览器下载的网络请求JSON，提取第一个符合条件的请求的params、Header, response
func ParseBrowserHarJson(harStr, tagUrl string, tagMethod string) (string, map[string]string, string, error) {

	ps, hs, rs, err := ParseBrowserHarJsonAll(harStr, tagUrl, tagMethod, true)
	if err != nil {
		return "", nil, "", err
	}
	return ps[0], hs[0], rs[0], nil

}

// ParseBrowserHarJsonAll ParseBrowserHarJson 解析从浏览器下载的网络请求JSON，提取所有符合条件的请求的params、Header, response， first=true只返回第一个
func ParseBrowserHarJsonAll(harStr, tagUrl string, tagMethod string, first bool) ([]string, []map[string]string, []string, error) {

	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(harStr), &mapResult)
	if err != nil {
		return nil, nil, nil, errors.New(fmt.Sprintf("JsonToMap err: %s", err.Error()))
	}

	entries := GetISliceFromMap(GetIMapFromMap(mapResult, "log"), "entries")

	paramSlice := make([]string, 0)
	headerSlice := make([]map[string]string, 0)
	responseSlice := make([]string, 0)

	for _, entry := range entries {
		m := GetIMapFromV(entry)
		request := GetIMapFromMap(m, "request")
		method := GetStringFromMap(request, "method")
		if method != tagMethod {
			continue
		}
		urlStr := GetStringFromMap(request, "url")
		if urlStr != tagUrl && !strings.HasPrefix(urlStr, tagUrl) {
			continue
		}

		response := GetIMapFromMap(m, "response")
		if GetIntFromMap(response, "status") != 200 {
			continue
		}
		content := GetIMapFromMap(response, "content")
		responseStr := GetStringFromMap(content, "text")

		hs := GetISliceFromMap(request, "headers")
		hMap := make(map[string]string)
		for _, h := range hs {
			hm := GetIMapFromV(h)
			hMap[GetStringFromMap(hm, "name")] = GetStringFromMap(hm, "value")
		}

		params := GetStringFromMap(GetIMapFromMap(request, "postData"), "text")

		paramSlice = append(paramSlice, params)
		headerSlice = append(headerSlice, hMap)
		responseSlice = append(responseSlice, responseStr)

		if first {
			//第一条直接返回
			return paramSlice, headerSlice, responseSlice, nil
		}

	}

	if len(paramSlice) < 1 {
		return nil, nil, nil, errors.New(fmt.Sprintf("没找到正确的请求[%s:%s]", tagUrl, tagMethod))
	}

	return paramSlice, headerSlice, responseSlice, nil

}

