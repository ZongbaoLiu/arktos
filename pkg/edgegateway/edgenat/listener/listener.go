package listener

import (
	"bufio"
	"fmt"
	"io"
	"k8s.io/kubernetes/pkg/edgegateway/edgenat/config"
	"net"
	"syscall"
	"unsafe"
)

type sockAddr struct {
	family uint16
	data   [14]byte
}

const (
	SoOriginalDst = 80
)

//func Start() {
//	conn, err := config.Config.Listener.Accept()
//	if err != nil {
//
//	}
//	ip, port, err := realServerAddress(&conn)
//	println(ip, port)
//	bufio.NewReader(conn)
//	go io.Copy(conn, conn)
//
//}

// realServerAddress returns an intercepted connection's original destination.
func RealServerAddress(conn *net.Conn) (string, int, error) {
	tcpConn, ok := (*conn).(*net.TCPConn)
	if !ok {
		return "", -1, fmt.Errorf("not a TCPConn")
	}

	file, err := tcpConn.File()
	if err != nil {
		return "", -1, err
	}

	// To avoid potential problems from making the socket non-blocking.
	tcpConn.Close()
	*conn, err = net.FileConn(file)
	if err != nil {
		return "", -1, err
	}

	defer file.Close()
	fd := file.Fd()

	var addr sockAddr
	size := uint32(unsafe.Sizeof(addr))
	err = getSockOpt(int(fd), syscall.SOL_IP, SoOriginalDst, uintptr(unsafe.Pointer(&addr)), &size)
	if err != nil {
		return "", -1, err
	}

	var ip net.IP
	switch addr.family {
	case syscall.AF_INET:
		ip = addr.data[2:6]
	default:
		return "", -1, fmt.Errorf("unrecognized address family")
	}

	port := int(addr.data[0])<<8 + int(addr.data[1])
	syscall.SetNonblock(int(fd), true)

	return ip.String(), port, nil
}

func getSockOpt(s int, level int, name int, val uintptr, vallen *uint32) (err error) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_GETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}
