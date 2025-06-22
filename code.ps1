function bm {
    param(
        [Parameter(Position=0)]
        [string]$Command,

        [Parameter(Position=1, ValueFromRemainingArguments=$true)]
        [string[]]$Arguments
    )

    # Handle 'go' command separately
    if ($Command -eq 'go') {
        if (-not $Arguments) {
            Write-Host "Error: Bookmark name is required" -ForegroundColor Red
            return
        }

        # Capture output and errors
        $result = bm.exe go $Arguments 2>&1
        $exitCode = $LASTEXITCODE

        if ($exitCode -eq 0) {
            # Success - change directory to the captured path
            Set-Location -Path $result
        }
        else {
            # Display error message in red
            Write-Host $result -ForegroundColor Red
        }
    }
    else {
        # Pass other commands directly to the executable
        bm.exe $Command $Arguments
    }
}