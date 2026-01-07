package kurohelperproxy

import "errors"

var (
	// ErrCreateSOCKS5DialerFailed 建立 SOCKS5 Dialer 失敗
	ErrCreateSOCKS5DialerFailed = errors.New("kurohelperproxy: failed to create SOCKS5 dialer")
)
