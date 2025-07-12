package cmd

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/net"
	"github.com/spf13/cobra"
)

// Raiz do comando

var netCmd = &cobra.Command{
	Use: "net",
	Short: "Exibe informações da rede.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("-----Input/Output-----")

        // Net (IOCounters)

		net_info_IOCounters, err := net.IOCounters(true)
		if err != nil {
			fmt.Println("Error retrieving information about Net.", err)
			return
		}

		for _, IO := range net_info_IOCounters {
			fmt.Printf("Input/Output: %s\n", IO.Name)
		}
		fmt.Println("-----TCP-----")

		// Net (connections TCP)

		net_info_connections_tcp, err := net.Connections("tcp")
		if err != nil {
			fmt.Println("Error retrieving information about Net.", err)
			return
		}

		for _, tcp := range net_info_connections_tcp {
			fmt.Printf("TCP Fd: %d  TCP Family: %d  TCP Type: %d  TCP Status: %s  TCP Uids: %d  TCP Pid: %d\n", tcp.Fd, tcp.Family, tcp.Type, tcp.Status, tcp.Uids, tcp.Pid)
		}
		fmt.Println("-----UDP-----")

		// Net (connections UDP)

		net_info_connections_udp, err := net.Connections("udp")
		if err != nil {
			fmt.Println("Error retrieving information about Net.", err)
			return
		}

		for _, udp := range net_info_connections_udp {
			fmt.Printf("UDP Fd: %d  UDP Family: %d  UDP Type: %d  UDP Uids: %d  UDP Pid: %d\n", udp.Fd, udp.Family, udp.Type, udp.Uids, udp.Pid)
		}
	},
}

// Registro do subcomando

func init() {
	rootCmd.AddCommand(netCmd)
}