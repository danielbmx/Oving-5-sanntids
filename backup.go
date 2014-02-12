// Backup:

package main

import (
	"net"
	"fmt"
	"time"
	"os"
	"bufio"
	"strconv"


)

func main(){

	// go udp_reciever()

	backup("counting.txt")

	
}



func backup(path string){
	fmt.Println("jeg er her")
	

	Check_loop:	
		for {
			
			if ( !udp_reciever() ){
				break Check_loop
			}
			
		}
	
	current_num,_ := readLines("counting.txt")
	for {
		fmt.Println(current_num[0])
		current_num[0] = current_num[0] + 1
		time.Sleep(1000 * time.Millisecond)
		writeLines(current_num, "counting.txt")
		go udp_sender()
	}

	


}


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
	temp := scanner.Text()
	tempInt,_ := strconv.Atoi(temp)
    	lines = append(lines, tempInt)
  }
  return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []int, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}


func udp_sender() {
    serverAddr_udp, err := net.ResolveUDPAddr("udp", "129.241.187.255:20020")
	PrintError(err)

    con_udp, err := net.DialUDP("udp", nil, serverAddr_udp)
    PrintError(err)

	for {
		time.Sleep(100 * time.Millisecond)
		_, err2 := con_udp.Write([]byte("IAmAlive"))
		PrintError(err2)
	}
}



func udp_reciever() (bool) {
    newBuf := make([]byte, 1024)
    addr, err := net.ResolveUDPAddr("udp", "129.241.187.145:20017")
    PrintError(err)
	
    sock, err := net.ListenUDP("udp", addr)
	PrintError(err)
	
	for {
        _, _, err = sock.ReadFromUDP(newBuf)
	fmt.Println("her")
        PrintError(err)
		sock.Close()
		
        if (string(newBuf) != "IAmAlive"){
			return false
		}
	return true
    }
}

func PrintError(err error) {
	if err != nil{
        fmt.Println(err)
	}
}

