package ipinfo

func (context *Context) getIndexById(dcID int8) int {
	for i := 0; i < len(context.MergedDCStatus); i++ {
		if context.MergedDCStatus[i].Id == dcID {
			return i
		}
	}
	return -1
}
