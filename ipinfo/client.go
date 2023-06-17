package ipinfo

import "DcCheckerStatus/ipinfo/types"

func Client() *Context {
	var tgDCStatus []types.DCStatus
	for i := 0; i < 5; i++ {
		tgDCStatus = append(tgDCStatus, types.DCStatus{
			Id: int8(i + 1),
		})
	}
	r := &Context{
		ListKnowIP:     make(map[string]*types.ServerInfo),
		MostUsedASN:    make(map[string]int64),
		MergedDCStatus: tgDCStatus,
	}
	r.readBackup(false)
	r.readBackup(true)
	return r
}
