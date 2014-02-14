// primary:


package main

import (
	"net"
	"fmt"
	"time"
	"os"
	"os/exec"
	"bufio"
	"strconv"

)

func main(){
	
	cmd := exec.Command("mate-terminal","-x", "go", "run", "backup.go")
	cmd.Run()
	
	tall := make([]int, 1)
    	tall[0] = 0
	

	primary("counting.txt",tall)


}

func primary(path string, tall []int) (){
	
	for;;{
		fmt.Println(tall[0])
		tall[0] = tall[0] + 1
		time.Sleep(1000 * time.Millisecond)
		writeLines(tall, "counting.txt")
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
    serverAddr_udp, err := net.ResolveUDPAddr("udp", "129.241.187.255:20022")
	PrintError(err)

    con_udp, err := net.DialUDP("udp", nil, serverAddr_udp)
    PrintError(err)

	for {
		time.Sleep(100 * time.Millisecond)
		_, err2 := con_udp.Write([]byte("IAmAlive"))
		PrintError(err2)
	}
}




func PrintError(err error) {
	if err != nil{
        fmt.Println(err)
	}
}
