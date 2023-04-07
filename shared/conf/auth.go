package conf

import "github.com/i-Things/things/shared/utils"

type AuthConf struct {
	IpRange []string `json:",optional"` //白名单ip 及ip段
	Users   []AuthUserInfo
}

type AuthUserInfo struct {
	UserName string // 内部为服务名
	Password string // 密码
}

// Auth 在名单内返回true
func (a *AuthConf) Auth(userName, password, ipaddr string) bool {
	var userCompare bool
	for _, user := range a.Users {
		if userName == user.UserName {
			userCompare = false
			if password == user.Password {
				userCompare = true
			}
			break
		}
	}
	if !userCompare {
		return false
	}
	if len(a.IpRange) == 0 {
		//如果没有,表示不开启ip白名单模式
		return true
	}
	for _, whiteIp := range a.IpRange {
		if utils.MatchIP(ipaddr, whiteIp) {
			return true
		}
	}
	return false
}
