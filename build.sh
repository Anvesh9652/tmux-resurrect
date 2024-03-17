#!/usr/bin/env bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
bin_name="clean_tmux_pane_contents"

cd $CURRENT_DIR/handle_pane_content && go build -o $bin_name main.go && mv $bin_name ~/.custom-codes/bin/