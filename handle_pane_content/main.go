package main

import (
    "bytes"
    "io"
    "log"
    "os"
    "regexp"
    "strings"
)

var (
    // temp directories for testing
    dir  = "/Users/agali/Desktop/Work/go-lang/tryouts/test/store.txt"
    dir2 = "/Users/agali/Desktop/Work/go-lang/tryouts/test/store-str.txt"

    // pRegx = regexp.MustCompile(`➜.*•.*✗(?:.*✗).*\s$`)
)

// This is helps Resurrect pluggin to remove empty line and empty prompts which has no result before saving
/*
Replace the below function end printf of tmux resurect plugin
scripts/save.sh - capture_pane_contents

	# the printf hack below removes *trailing* empty lines
	# printf '%s\n' "$(tmux capture-pane -epJ -S "$start_line" -t "$pane_id")" > "$(pane_contents_file "save" "$pane_id")"
	res=$(printf '%s\n' "$(tmux capture-pane -epJ -S "$start_line" -t "$pane_id")" | tmux_res_empty)
	# in new tab for just startign propmt after restoring it is adding new line
	if [[ ${#res} -eq 0 ]]; then
		printf '%s' "${res}" >  "$(pane_contents_file "save" "$pane_id")"
	else
		printf '%s\n' "${res}" >  "$(pane_contents_file "save" "$pane_id")"
	fi
	# printf '%s' "$(tmux capture-pane -epJ -S "$start_line" -t "$pane_id")" > /Users/agali/Desktop/Work/go-lang/tryouts/test/store.txt

	Tip: to speedup prompt loading after restore || follow these as it is
	tmux_spinner.sh - 26 => sleep 0.01
	restore.sh - 384 => display_message "Tmux restore complete!" "50"
*/
func main() {
    removeDummyCommands()
}

func removeDummyCommands() {
    f := os.Stdin
    btData, err := io.ReadAll(f)
    if err != nil {
        log.Fatal(err)
    }

    btSplit := bytes.Split(btData, []byte("\n"))
    res := [][]byte{}
    for i, bt := range btSplit {
        if canLineBeRemoved(bt, i) {
            continue
        }
        res = append(res, bt)
    }
    finalRes := bytes.Join(res, []byte("\n"))
    os.Stdout.Write(finalRes)
}

func canLineBeRemoved(bt []byte, i int) bool {
    line := strings.TrimSpace(string(bt))
    lineSplit := strings.Split(line, "✗")
    if line == "" {
        return true
    }
    n := len(lineSplit)
    lastPromt := lineSplit[n-1]
    // with git && no change in files						// without git
    return strings.HasSuffix(lastPromt, `[0m[39m[49m`) || lastPromt == `[39m`
}

// seems working as expected
func basicHandleSpacing() {
    prmotRegex := regexp.MustCompile(`\s+$`)

    f := os.Stdin
    btData, err := io.ReadAll(f)
    if err != nil {
        log.Fatal(err)
    }

    btSplit := bytes.Split(btData, []byte("\n"))
    n := len(btSplit)
    last := string(btSplit[n-2])
    if prmotRegex.MatchString(last) {
        btSplit = btSplit[1 : n-2]
    }
    res := bytes.Join(btSplit, []byte("\n"))
    _, err = os.Stdout.Write(res)
    if err != nil {
        log.Fatal(err)
    }
}
