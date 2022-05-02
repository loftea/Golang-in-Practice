package lib

type GoTickets interface {
	Take()
	Return()
	Active() bool
	Total() uint32
	Remainder() uint32
}

type myGoTickets struct {
	total    uint32
	ticketCh chan struct{}
	active   bool
}

func (gt *myGoTickets) init(total uint32) bool {
	if gt.active {
		return false
	}
	if total == 0 {
		return false
	}
	ch := make(chan struct{}, total)
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
	gt.ticketCh = ch
	gt.total = total
	gt.active = true
	return true
}
