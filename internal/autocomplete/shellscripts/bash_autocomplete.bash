#!/bin/bash

____APPNAME___bash_autocomplete() {
  if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    local cur completions exit_code
    local IFS=$'\n'
    cur="${COMP_WORDS[COMP_CWORD]}"

    completions=$(COMPLETION_STYLE=bash "${COMP_WORDS[0]}" __complete -- "${COMP_WORDS[@]:1:$COMP_CWORD-1}" "$cur" 2>/dev/null)
    exit_code=$?

    case $exit_code in
        10) mapfile -t COMPREPLY < <(compgen -f -- "$cur") ;;   # file
        11) COMPREPLY=() ;;                                     # no completion
        0)  mapfile -t COMPREPLY <<< "$completions" ;;          # use returned completions
    esac
    return 0
  fi
}

complete -F ____APPNAME___bash_autocomplete __APPNAME__
