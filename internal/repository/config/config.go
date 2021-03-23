package config

import "time"

func BookingStartTime() time.Time {
	t, _ := time.Parse(StartTimeFormat, configuration().GoLive.StartIsoDatetime)
	return t
}

func ServerAddr() string {
	return ":" + configuration().Server.Port
}

func BookingCode() string {
	return configuration().GoLive.BookingCode
}

func GoLiveTime() time.Time {
	start, _ := time.Parse(StartTimeFormat, configuration().GoLive.StartIsoDatetime)
	return start
}

func IsCorsDisabled() bool {
	return configuration().Security.DisableCors
}

func JWTPublicKey() string {
	return `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnzyis1ZjfNB0bBgKFMSv
vkTtwlvBsaJq7S5wA+kzeVOVpVWwkWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHc
aT92whREFpLv9cj5lTeJSibyr/Mrm/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIy
tvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0
e+lf4s4OxQawWD79J9/5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWb
V6L11BWkpzGXSW4Hv43qa+GSYOD2QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9
MwIDAQAB
-----END PUBLIC KEY-----`
}
