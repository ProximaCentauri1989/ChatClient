package main

import (
	"net"
	"os"
	"syscall"
)

//Checks whether port is already in use
func IsAlreadyBinded(err error) bool {
	opErr, ok := err.(*net.OpError)
	if !ok {
		return false
	}
	sysCallErr, ok := opErr.Err.(*os.SyscallError)
	if !ok {
		return false
	}
	return sysCallErr.Err == syscall.EADDRINUSE
}
