package factory

import "testing"

func TestNewIphone(t *testing.T) {
	iphone := NewIphone(110)
	if iphone == nil {
		t.Error("创建新IPhone失败")
		return
	}

	iphone.LoginQQ()
}
