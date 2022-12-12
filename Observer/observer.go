package observer

import (
	"fmt"
	"sync"
)

// 报纸订阅，报社为被观察者，订阅的人为观察者

type Customer interface {
	OnChange()
}

type CustomerA struct {
}

func (la *CustomerA) OnChange() {
	fmt.Println("我是客户A, 我收到报纸了")
}

type CustomerB struct {
}

func (la *CustomerB) OnChange() {
	fmt.Println("我是客户B, 我收到报纸了")
}

type NewsOffice struct {
	Customers []Customer
	lock      sync.Mutex
}

func (o *NewsOffice) AddCustomer(l Customer) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.Customers = append(o.Customers, l)
}

func (o *NewsOffice) notifyAllCustomers() {
	wg := sync.WaitGroup{}
	for _, customer := range o.Customers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			customer.OnChange()
		}()
	}
	wg.Wait()
}

func (o *NewsOffice) newspaperCome() {

	o.notifyAllCustomers()

}
