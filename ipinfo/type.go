package ipinfo

import "DcCheckerStatus/ipinfo/types"

type Context struct {
	ListKnowIP     map[string]*types.ServerInfo
	TrustedASN     []string
	TrustedIP      []string
	MostUsedASN    map[string]int64
	MergedDCStatus []types.DCStatus
}
