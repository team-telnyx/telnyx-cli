#!/usr/bin/env fish

function ____APPNAME___fish_autocomplete
    set -l tokens (commandline -xpc)
    set -l current (commandline -ct)

    set -l cmd $tokens[1]
    set -l args $tokens[2..-1]

    set -l completions (env COMPLETION_STYLE=fish $cmd __complete -- $args $current 2>>/tmp/fish-debug.log)
    set -l exit_code $status

    switch $exit_code
        case 10
            # File completion
            __fish_complete_path "$current"
        case 11
            # No completion
            return 0
        case 0
            # Use returned completions
            for completion in $completions
                echo $completion
            end
    end
end

complete -c __APPNAME__ -f -a '(____APPNAME___fish_autocomplete)'

