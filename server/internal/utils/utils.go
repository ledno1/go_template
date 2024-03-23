package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"sort"
	"strings"
)

var (
	Utils = cUtils{}
)

type cUtils struct{}

func (c *cUtils) Sign(ctx context.Context, params map[string]string, key string) (sign string, err error) {
	// 1. 将请求字段按Ascii码方式进行升序排序
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 2. 按照key1=val1&key2=val2&key3=val3...的方式生成加密字符串
	var kvPairs []string
	for _, k := range keys {
		// 过滤空值
		if params[k] == "" {
			continue
		}
		kvPairs = append(kvPairs, fmt.Sprintf("%s=%s", k, params[k]))
	}
	joinedString := strings.Join(kvPairs, "&") + "&key=" + key
	// 3. 将生成的字符串进行MD5加密，并转换成大写
	hasher := md5.New()
	hasher.Write([]byte(joinedString))
	signature := hex.EncodeToString(hasher.Sum(nil))
	signature = strings.ToUpper(signature)

	return signature, nil
}

func (c *cUtils) Verify(ctx context.Context, obj interface{}, key string) (result bool, err error) {
	//将obj 转换 map[string]string
	objMap := gconv.Map(obj)
	signReq := gconv.String(objMap["sign"])
	// 删除sign
	delete(objMap, "sign")
	objString := make(map[string]string)
	// 将objMap 转换成map[string]string 类型
	for k, v := range objMap {
		objString[k] = gconv.String(v)
	}
	if sign, err := c.Sign(ctx, objString, key); err != nil {
		return false, err
	} else {
		return sign == signReq, nil
	}
}
