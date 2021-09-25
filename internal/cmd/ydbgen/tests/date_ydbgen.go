// Code generated by ydbgen; DO NOT EDIT.

package tests

import (
	"strconv"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/resultset"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

var (
	_ = strconv.Itoa
	_ = time.Now
	_ = table.NewQueryParameters
	_ = resultset.Result.Scan
	_ = types.StringValue
)

func (t *Times) Scan(res resultset.Result) (err error) {
	_ = res.ScanWithDefaults(&t.Date)
	return res.Err()
}

func (t *Times) QueryParameters() *table.QueryParameters {
	var v0 types.Value
	{
		var v1 types.Value
		var x0 =t.Date
		ok0 := !t.Date.IsZero()
		if ok0 {
			v1 = types.OptionalValue(types.DateValueFromTime(x0))
		} else {
			v1 = types.NullValue(types.TypeDate)
		}
		v0 = v1
	}
	return table.NewQueryParameters(
		table.ValueParam("$date", v0),
	)
}

func (t *Times) StructValue() types.Value {
	var v0 types.Value
	{
		var v1 types.Value
		{
			var v2 types.Value
			var x0 =t.Date
			ok0 := !t.Date.IsZero()
			if ok0 {
				v2 = types.OptionalValue(types.DateValueFromTime(x0))
			} else {
				v2 = types.NullValue(types.TypeDate)
			}
			v1 = v2
		}
		v0 = types.StructValue(
			types.StructFieldValue("date", v1),
		)
	}
	return v0
}

func (t *Times) StructType() types.Type {
	var t0 types.Type
	{
		fs0 := make([]types.StructOption, 1)
		var t1 types.Type
		{
			tp0 := types.TypeDate
			t1 = types.Optional(tp0)
		}
		fs0[0] = types.StructField("date", t1)
		t0 = types.Struct(fs0...)
	}
	return t0
}
