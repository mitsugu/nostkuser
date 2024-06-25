package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/yosuke-furukawa/json5/encoding/json5"
	"log"
	"os"
	"time"
)

// type declaration
// const
const (
	readWriteFlag = 0
	readFlag      = 1
	writeFlag     = 2
)


// for contact lists structure
type CONTACT struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

//

/*
main
*/
func main() {
	/* for debug
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	path += "/w.err"
	startDebug(path)
	*/

	if len(os.Args) < 2 {
		dispHelp()
		os.Exit(0)
	}

	// load config.json
	var cc confClass
	if err := cc.existConfiguration(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := cc.loadConfiguration(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		dispHelp()
		os.Exit(0)
	case "--help":
		dispHelp()
		os.Exit(0)
	case "-h":
		dispHelp()
		os.Exit(0)
	}

	f, err := cc.openJSON5(cc.ConfData.Filename.Contacts)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	var data interface{}
	dec := json5.NewDecoder(f)
	err = dec.Decode(&data)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	b, err := json5.Marshal(data)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	c := make(map[string]CONTACT)
	if err := json5.Unmarshal([]byte(b), &c); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(c[os.Args[1]].Name)
}
//

/*
readStdIn {{{
*/
func readStdIn() (string, error) {
	cn := make(chan string, 1)
	go func() {
		sc := bufio.NewScanner(os.Stdin)
		var buff bytes.Buffer
		for sc.Scan() {
			fmt.Fprintln(&buff, sc.Text())
		}
		cn <- buff.String()
	}()
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
	select {
	case text := <-cn:
		return text, nil
	case <-timer.C:
		return "", errors.New("Time out input from standard input")
	}
}

// }}}

/*
debugPrint {{{
*/
func startDebug(s string) {
	f, err := os.OpenFile(s, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.Println("start debug")
}

// }}}
