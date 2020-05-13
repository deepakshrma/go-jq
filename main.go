/*
Copyright © 2020 Deepak Vishwakarma

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func parseJSON(args ...string) {
	spaceCount := 2
	str := "{}"
	if len(args) == 0 {
		return
	}
	if len(args) == 1 {
		str = args[0]
	}
	if len(args) == 2 {
		str = args[1]
		n, _ := strconv.ParseInt(args[0], 10, 64)
		spaceCount = int(n)
	}
	spaces := ""
	for i := 0; i < spaceCount; i++ {
		spaces += " "
	}
	var jsonStr map[string]interface{}
	err := json.Unmarshal([]byte(strings.TrimRight(str, "\n")), &jsonStr)
	if err == nil {
		str, _ := json.MarshalIndent(jsonStr, "", spaces)
		fmt.Println(string(str))
	}
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("%v", os.Args)
	argCount := len(os.Args)
	if contains(os.Args, "--url") {
		ack := make(chan bool, 1) // Acknowledgement channel
		go func() {
			resp, err := http.Get(os.Args[2])
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			parseJSON(string(body))
			ack <- true
		}()
		<-ack
		return
	}
	if argCount == 2 {
		parseJSON(os.Args[1])
	}
	if argCount == 3 {
		parseJSON(os.Args[1], os.Args[2])
	}
	if argCount == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)
		var file string
		for scanner.Scan() {
			file = file + scanner.Text() + "\n"
		}
		parseJSON("2", file)
	}
}
