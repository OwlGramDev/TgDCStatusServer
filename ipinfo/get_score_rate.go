package ipinfo

import "DcCheckerStatus/ipinfo/types"

func (context *Context) GetScoreRate(info types.ServerInfo) int64 {
	isTrustedProvider := false
	for i := 0; i < len(context.TrustedASN); i++ {
		if context.TrustedASN[i] == info.ASN {
			isTrustedProvider = true
			break
		}
	}
	scoreUsage := (context.MostUsedASN[info.ASN] * 100) / int64(len(context.ListKnowIP))
	isTrustedIp := false
	for i := 0; i < len(context.TrustedIP); i++ {
		if context.TrustedIP[i] == info.Ip {
			isTrustedIp = true
			break
		}
	}
	score := scoreUsage
	if isTrustedProvider {
		score += 100
	}
	if isTrustedIp {
		score += 100
	}
	return (score * 100) / 300
}
