package util

import "regexp"

var ipRe *regexp.Regexp

func GetIP(s string) string {
	if ipRe == nil {
		ipRe = regexp.MustCompile(`(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)
	}
	return ipRe.FindString(s)
}
