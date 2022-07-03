package simpleFactory

func TestNewIphone() {
	iphone := NewIphone(110)
	if iphone == nil {
		fmt.Println("创建新IPhone失败")
		return
	}

	iphone.LoginQQ()
}
