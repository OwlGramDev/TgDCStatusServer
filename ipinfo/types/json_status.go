package types

type JSONStatus struct {
	Status      []DCStatus `json:"status"`
	LastRefresh int64      `json:"last_refresh"`
}
