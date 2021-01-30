package constants

type ShipmentStatus struct {
	Manifested                uint8 // Barang baru didaftarkan
	OnProcess                 uint8 // Barang sedang dalam pengiriman
	OnTransit                 uint8 // Barang sedang transit di kota tertentu
	Received                  uint8 // Barang telah sampai di kota tujuan
	Delivered                 uint8 // Barang telah diterima di alamat yg dituju
	ClosedOnceDeliveryAttempt uint8 // Rumah penerima barang sedang kosong atau tutup saat pengiriman barang
}

var SHIPMENT_STATUS = ShipmentStatus{
	Manifested:                0,
	OnProcess:                 1,
	OnTransit:                 2,
	Received:                  3,
	Delivered:                 4,
	ClosedOnceDeliveryAttempt: 5,
}
