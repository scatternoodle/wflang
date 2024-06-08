#!/bin/bash
alias codeline="git ls-files | grep -E '(.ts|.go)$' | xargs wc -l"