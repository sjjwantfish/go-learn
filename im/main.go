package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

type IMPkg struct {
	PkgLen          [4]byte
	HeaderLen       [2]byte
	ProtocolVersion [2]byte
	Operation       [4]byte
	SequenceId      [4]byte
	BodyLen         [4]byte
	Body            []byte
}

func UnpackGoim(data []byte) *IMPkg {
	pkg := new(IMPkg)
	var pkgLen [4]byte
	copy(pkgLen[:], data[:4])
	pkg.PkgLen = pkgLen

	var headerLen [2]byte
	copy(headerLen[:], data[4:6])
	pkg.HeaderLen = headerLen

	var pv [2]byte
	copy(pv[:], data[8:10])
	pkg.ProtocolVersion = pv

	var operation [4]byte
	copy(operation[:], data[10:14])
	pkg.Operation = operation

	var seqId [4]byte
	copy(seqId[:], data[14:18])
	pkg.SequenceId = seqId

	var bodyLen [4]byte
	copy(bodyLen[:], data[18:22])
	pkg.BodyLen = bodyLen

	var body []byte
	copy(body[:], data[22:])
	pkg.Body = body

	return pkg
}


func HandleConn(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte("unpack goim\n"))
	if err != nil {
		fmt.Println("error: ", err)
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, conn)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("Received: ", buf)
	pkg := UnpackGoim(buf.Bytes())
	// your logic
	fmt.Println(*pkg)
}

func main() {
	addr := "0.0.0.0:8080" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
		}
		go HandleConn(conn) //开启多个协程。
	}
}
