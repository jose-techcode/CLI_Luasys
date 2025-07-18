package sysinfo

import (
	"bytes"
	"fmt"

	"github.com/shirou/gopsutil/v3/net"
)

func GetNETInfo() string {
	var out bytes.Buffer

	    out.WriteString("-----INPUT/OUTPUT-----\n")

        // Net (IOCounters)

		net_info_IOCounters, err := net.IOCounters(true)
		if err != nil {
			out.WriteString("Error retrieving information about net: " + err.Error() + "\n")
	        return out.String()
		}

		for _, IO := range net_info_IOCounters {
			fmt.Printf("Input/Output: %s\n", IO.Name)
		}
		
		out.WriteString("\n-----TCP-----\n")

		// Net (connections TCP)

		net_info_connections_tcp, err := net.Connections("tcp")
		if err != nil {
			out.WriteString("Error retrieving information about net: " + err.Error() + "\n")
	        return out.String()
		}

		for _, tcp := range net_info_connections_tcp {
			out.WriteString(fmt.Sprintf(("TCP Fd: %d  TCP Family: %d  TCP Type: %d  TCP Status: %s  TCP Uids: %v  TCP Pid: %d\n"), tcp.Fd, tcp.Family, tcp.Type, tcp.Status, tcp.Uids, tcp.Pid))
		}
		
		out.WriteString("\n-----UDP-----\n")

		// Net (connections UDP)

		net_info_connections_udp, err := net.Connections("udp")
		if err != nil {
			out.WriteString("Error retrieving information about net: " + err.Error() + "\n")
	        return out.String()
		}

		for _, udp := range net_info_connections_udp {
			out.WriteString(fmt.Sprintf("UDP Fd: %d  UDP Family: %d  UDP Type: %d  UDP Uids: %v  UDP Pid: %d\n", udp.Fd, udp.Family, udp.Type, udp.Uids, udp.Pid))
		}
		
		return out.String()
	}