package ipinfo

func translateStatus(status int64) int64 {
	switch status {
	case 1:
		return 2
	case 2:
		return 1
	default:
		return 0
	}
}
