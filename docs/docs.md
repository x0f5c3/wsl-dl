# wsl-dl

## Usage
> This cli tool will download WSL distros as zip files and unpack them to allow for custom installations

wsl-dl [output directory]

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|
|`-w, --workers int`|number of workers to use for downloading (default 2)|

## Commands
|Command|Usage|
|-------|-----|
|`wsl-dl completion`|Generate the autocompletion script for the specified shell|
|`wsl-dl help`|Help about any command|
# ... completion
`wsl-dl completion`

## Usage
> Generate the autocompletion script for the specified shell

wsl-dl completion

## Description

```
Generate the autocompletion script for wsl-dl for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`wsl-dl completion bash`|Generate the autocompletion script for bash|
|`wsl-dl completion fish`|Generate the autocompletion script for fish|
|`wsl-dl completion powershell`|Generate the autocompletion script for powershell|
|`wsl-dl completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`wsl-dl completion bash`

## Usage
> Generate the autocompletion script for bash

wsl-dl completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(wsl-dl completion bash)

To load completions for every new session, execute once:

#### Linux:

	wsl-dl completion bash > /etc/bash_completion.d/wsl-dl

#### macOS:

	wsl-dl completion bash > $(brew --prefix)/etc/bash_completion.d/wsl-dl

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`wsl-dl completion fish`

## Usage
> Generate the autocompletion script for fish

wsl-dl completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	wsl-dl completion fish | source

To load completions for every new session, execute once:

	wsl-dl completion fish > ~/.config/fish/completions/wsl-dl.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`wsl-dl completion powershell`

## Usage
> Generate the autocompletion script for powershell

wsl-dl completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	wsl-dl completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`wsl-dl completion zsh`

## Usage
> Generate the autocompletion script for zsh

wsl-dl completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(wsl-dl completion zsh); compdef _wsl-dl wsl-dl

To load completions for every new session, execute once:

#### Linux:

	wsl-dl completion zsh > "${fpath[1]}/_wsl-dl"

#### macOS:

	wsl-dl completion zsh > $(brew --prefix)/share/zsh/site-functions/_wsl-dl

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... help
`wsl-dl help`

## Usage
> Help about any command

wsl-dl help [command]

## Description

```
Help provides help for any command in the application.
Simply type wsl-dl help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 05 April 2023**
