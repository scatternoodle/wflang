#!/bin/bash
git ls-files | grep -E '(.ts|.go)$' | xargs wc -l