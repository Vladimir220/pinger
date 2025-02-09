package ping

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func Ping(ip string) (time.Duration, error) {
	/*ipAddr, err := net.ResolveIPAddr("ip4", ip)
	if err != nil {
		fmt.Println("Ошибка при разрешении IP:", err)
		return 0, err
	}

	// Открываем IP соединение для отправки и получения ICMP
	connection, err := net.DialIP("ip4:icmp", nil, ipAddr)
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	message := &icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  0,
			Data: []byte("HELLO-ITS_PING"),
		},
	}

	data, err := message.Marshal(nil)
	if err != nil {
		return 0, err
	}

	start := time.Now()

	_, err = connection.Write(data) // Используем Write вместо WriteTo
	if err != nil {
		return 0, err
	}

	connection.SetReadDeadline(time.Now().Add(2 * time.Second)) // Важно: таймаут!
	buf := make([]byte, 2000)
	n, err := connection.Read(buf) // Используем Read вместо ReadFrom
	if err != nil {
		return 0, err
	}

	duration := time.Since(start)

	reply, err := icmp.ParseMessage(1, buf[:n])
	if err != nil {
		return 0, err
	}

	if reply.Type == ipv4.ICMPTypeEchoReply {
		return duration, nil
	} else {
		return 0, fmt.Errorf("получен неожиданный ответ: %+v", reply)
	}*/

	ipAddr, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		fmt.Println("Ошибка при разрешении IP:", err)
		return 0, err
	}

	connection, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	message := &icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  0,
			Data: []byte("HELLO-ITS_PING"),
		},
	}

	data, err := message.Marshal(nil)
	if err != nil {
		return 0, err
	}

	start := time.Now()

	_, err = connection.WriteTo(data, ipAddr)
	if err != nil {
		return 0, err
	}

	buf := make([]byte, 1500)
	n, peer, err := connection.ReadFrom(buf)
	if err != nil {
		return 0, err
	}

	message, err = icmp.ParseMessage(1, buf[:n])
	if err != nil {
		return 0, err
	}

	duration := time.Since(start)

	if message.Type == ipv4.ICMPTypeEchoReply {
		return duration, nil
	} else {
		return 0, fmt.Errorf("получен неожиданный ответ от %s: %+v", peer, message)
	}
}
