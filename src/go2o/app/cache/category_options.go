/**
 * Copyright 2014 @ ops Inc.
 * name :
 * author : newmin
 * date : 2014-02-05 21:53
 * description :
 * history :
 */
package cache

import (
	"bytes"
	"fmt"
	"go2o/core/domain/interface/sale"
	"go2o/core/service/dps"
)

type CategoryFormatFunc func(buf *bytes.Buffer, c *sale.ValueCategory, level int)

func readToCategoryDropList(partnerId int) []byte {
	categorys := dps.SaleService.GetCategories(partnerId)
	buf := bytes.NewBuffer([]byte{})
	var f CategoryFormatFunc = func(buf *bytes.Buffer, c *sale.ValueCategory, level int) {
		buf.WriteString(fmt.Sprintf(
			`<option class="opt%d" value="%d">%s</option>`,
			level,
			c.Id,
			c.Name,
		))
	}
	itrCategory(buf, categorys, &sale.ValueCategory{Id: 0}, f, 0)

	return buf.Bytes()
}

func itrCategory(buf *bytes.Buffer, categorys []*sale.ValueCategory, c *sale.ValueCategory, f CategoryFormatFunc, level int) {
	if c.Id != 0 {
		f(buf, c, level)
	}

	for _, k := range categorys {
		if k.ParentId == c.Id {
			itrCategory(buf, categorys, k, f, level+1)
		}
	}
}

// 获取分类下拉选项
func GetDropOptionsOfCategory(partnerId int) []byte {
	return readToCategoryDropList(partnerId)
}
