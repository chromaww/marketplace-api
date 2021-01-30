package constants

type OrderStatus struct {
	New      uint8
	Checked  uint8
	Paid     uint8
	Canceled uint8
}

var ORDER_STATUS = OrderStatus{
	New:      0,
	Checked:  1,
	Paid:     2,
	Canceled: 3,
}
