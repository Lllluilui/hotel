package delete

// Request 请求
type Request struct {
	// ID 酒店标识
	ID int64 `uri:"hotelId"`
}

// Response 响应
type Response struct{}
