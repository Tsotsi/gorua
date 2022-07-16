package main

import (
	"github.com/peterh/liner"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	history_fn = filepath.Join(os.TempDir(), ".liner_example_history")
	names      = []string{}
)

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)
	line.SetTabCompletionStyle(liner.TabPrints)
	line.SetCompleter(func(line string) (c []string) {
		for _, n := range names {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	if f, err := os.Open(history_fn); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	for {
		if name, err := line.Prompt("> "); err == nil {
			if name == ":q" {
				break
			}
			//r := strings.NewReader(name)
			//l := parser.NewLexer(r)
			line.AppendHistory(name)
		} else if err == liner.ErrPromptAborted {
			log.Print("Aborted")
		} else {
			log.Print("Error reading line: ", err)
		}
	}

	if f, err := os.Create(history_fn); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
}
