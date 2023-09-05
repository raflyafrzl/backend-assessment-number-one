package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(encodeBaju("AT", "YH", 70000))
	fmt.Println(encodeBaju("ESH", "DTT", 150000))
	fmt.Println(encodeBaju("DET", "DHT", 250000))
}

func encodeBaju(limitPrice string, idealPrice string, bargainedPrice int) string {

	var limit int
	limit, _ = strconv.Atoi(sumTotal(limitPrice))
	var ideal int
	ideal, _ = strconv.Atoi(sumTotal(idealPrice))

	if bargainedPrice == (ideal * 1000) {
		return "GOOD, customer terbaik gak pake nawar"
	}
	if bargainedPrice >= (limit * 1000) {
		return "ACCEPT, terimakasih sudah berbelanja"
	}

	return "REJECT, belum balik modal nih!"
}
func sumTotal(data string) string {
	var result string = ""
	var num int = 0
	for _, data := range data {

		num = strings.Index("TEDUHASYIK", string(data))
		if num >= 4 {
			result = result + strconv.Itoa(num+1)
		} else {
			result = result + strconv.Itoa(num)
		}

	}

	return result

}
