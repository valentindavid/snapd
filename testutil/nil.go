package testutil

import (
	"fmt"
	"reflect"

	"gopkg.in/check.v1"
)

type nilChecker struct {
	*check.CheckerInfo
}
type notNilChecker struct {
	*check.CheckerInfo
}

var IsIntrfcNil check.Checker = &nilChecker{
	CheckerInfo: &check.CheckerInfo{Name: "IsIntrfcNil", Params: []string{"value"}},
}
var IsIntrfcNotNil check.Checker = &notNilChecker{
	CheckerInfo: &check.CheckerInfo{Name: "IsIntrfcNotNil", Params: []string{"value"}},
}

func checkIntrfcNil(param interface{}) (bool, string) {
	value := reflect.ValueOf(param)
	// (interface{})(nil) should be invalid. But if it is also nil, it is not an inteface nil.
	if value.IsValid() && value.IsNil() {
		ty := reflect.TypeOf(param)
		return false, fmt.Sprintf("Got a typed nil of type: %v", ty)
	}
	return true, ""
}

func (c *nilChecker) Check(params []interface{}, names []string) (result bool, error string) {
	if r, message := checkIntrfcNil(params[0]); !r {
		return false, message
	}
        if params[0] != nil {
		return false, fmt.Sprintf("Value was not nil: %v", params[0])
	}
	return true, ""
}

func (c *notNilChecker) Check(params []interface{}, names []string) (result bool, error string) {
	if r, message := checkIntrfcNil(params[0]); !r {
		return false, message
	}
        if params[0] == nil {
		return false, fmt.Sprintf("Value was nil")
	}
	return true, ""
}
