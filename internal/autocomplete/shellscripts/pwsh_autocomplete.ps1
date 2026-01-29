Register-ArgumentCompleter -Native -CommandName __APPNAME__ -ScriptBlock {
  param($wordToComplete, $commandAst, $cursorPosition)

  $elements = $commandAst.CommandElements
  $completionArgs = @()

  # Extract each of the arguments
  for ($i = 0; $i -lt $elements.Count; $i++) {
    $completionArgs += $elements[$i].Extent.Text
  }

  # Add empty string if there's a trailing space (wordToComplete is empty but cursor is after space)
  # Necessary for differentiating between getting completions for namespaced commands vs. subcommands
  if ($wordToComplete.Length -eq 0 -and $elements.Count -gt 0) {
    $completionArgs += ""
  }

  $output = & {
    $env:COMPLETION_STYLE = 'pwsh'
    __APPNAME__ __complete @completionArgs 2>&1
  }
  $exitCode = $LASTEXITCODE

  switch ($exitCode) {
    10 {
      # File completion behavior
      Get-ChildItem -Path "$wordToComplete*" | ForEach-Object {
        $completionText = if ($_.PSIsContainer) { $_.Name + "/" } else { $_.Name }
        [System.Management.Automation.CompletionResult]::new(
          $completionText,
          $completionText,
          'ProviderItem',
          $completionText
        )
      }
    }
    11 {
      # No reasonable suggestions
      [System.Management.Automation.CompletionResult]::new(' ', ' ', 'ParameterValue', ' ')
    }
    default {
      # Default behavior - show command completions
      $output | ForEach-Object {
        [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
      }
    }
  }
}
