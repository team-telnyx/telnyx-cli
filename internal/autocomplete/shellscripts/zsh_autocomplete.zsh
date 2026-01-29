#!/bin/zsh
compdef ____APPNAME___zsh_autocomplete __APPNAME__

____APPNAME___zsh_autocomplete() {

  local -a opts
  local temp
  local exit_code

  temp=$(COMPLETION_STYLE=zsh "${words[1]}" __complete "${words[@]:1}")
  exit_code=$?

  case $exit_code in
    10)
      # File completion behavior
      _files
      ;;
    11)
      # No completion behavior - return nothing
      return 1
      ;;
    0)
      # Default behavior - show command completions
      opts=("${(@f)temp}")
      _describe 'values' opts
      ;;
  esac
}
