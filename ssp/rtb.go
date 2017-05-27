package ssp

// minimal OpenRTB 2.3

type RTBBidRequest struct {
	ID          string          `json:"id"`
	Impressions []RTBImpression `json:"imp,omitempty"`
}

type RTBImpression struct {
	ID     string     `json:"id,omitempty"`
	Banner *RTBBanner `json:"banner,omitempty"`
}

type RTBBanner struct {
	Width  int `json:"w,omitempty"`
	Height int `json:"h,omitempty"`
}

type RTBBidResponse struct {
	ID       string       `json:"id,omitempty"`
	Seatbids []RTBSeatbid `json:"seatbid,omitempty"`
}

type RTBSeatbid struct {
	Bids []RTBBid `json:"bid,omitempty"`
}

type RTBBid struct {
	ImpressionID    string  `json:"impid,omitempty"`
	Price           float64 `json:"price,omitempty"`
	AdMarkup        string  `json:"adm,omitempty"`
	NotificationURL string  `json:"nurl,omitempty"`
}