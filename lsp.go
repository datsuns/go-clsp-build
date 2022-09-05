// qcc -Vgcc_ntox86_64 -dM -E - < /dev/null
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

func writeJson(opt *Option, entries *[]ClangdOptionEntry) {
	f, err := os.Create(opt.DestPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//json.NewEncoder(f).Encode(entries)
	raw, _ := json.Marshal(entries)
	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(raw), "", "  ")
	if err != nil {
		panic(err)
	}
	f.WriteString(buf.String())
}

func main() {
	opt := parseOption()
	ret, err := loadConfFromFile(opt.ConfPath)
	if err != nil {
		panic(err)
	}

	size := len(ret.Entries)
	entries := make([]*[]ClangdOptionEntry, size)
	var wg sync.WaitGroup
	i := 0
	for k, v := range ret.Entries {
		wg.Add(1)
		go func(n int) {
			log.Println(">> start ", k)
			var err error
			defer wg.Done()
			entries[n], err = buildStructure(&v, ret.CommonFlags)
			if err != nil {
				panic(err)
			}
			log.Println("<< done ", k)
		}(i)
		i += 1
		wg.Wait()
	}
	total := []ClangdOptionEntry{}
	for _, e := range entries {
		total = append(total, *e...)
	}
	writeJson(opt, &total)

	fmt.Println(">> done")
}
