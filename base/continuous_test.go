package base

import (
	"testing"
)

func TestContinuousFeature_GetSysValFromString(t *testing.T) {
	vs := []string{"5.1", "3.5", "1.4", "0.2"}
	c := NewContinuousFeature("test")
	c.Precision = 1
	for _, str := range vs {
		err, val := c.GetSysValFromString(str)
		if err != nil {
			t.Errorf("receive %s error %v", str, err)
		} else {
			err, f := c.GetStringFromSysVal(val)
			if err != nil {
				t.Errorf("send %s, get error, %v", str, err)
			} else {
				if str != f {
					t.Errorf("send %s, get %s\n", str, f)
				}
			}
		}
	}
}
