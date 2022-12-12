package observer

import "testing"

func TestObserver(*testing.T) {
	ca := &CustomerA{}
	cb := &CustomerB{}

	no := NewsOffice{}

	no.AddCustomer(ca)
	no.AddCustomer(cb)
	no.newspaperCome()
}
