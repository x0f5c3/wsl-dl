# Quick Start - Install wsl-dl

> [!TIP]
> wsl-dl is installable via [instl.sh](https://instl.sh).\
> You just have to run the following command and you're ready to go!

<!-- tabs:start -->

#### ** Windows **

### Windows Command

```powershell
iwr instl.sh/x0f5c3/wsl-dl/windows | iex
```

#### ** Compile from source **

### Compile from source with Golang

?> **NOTICE**
To compile wsl-dl from source, you have to have [Go](https://golang.org/) installed.

Compiling wsl-dl from source has the benefit that the build command is the same on every platform.\
It is not recommended to install Go only for the installation of wsl-dl.

```command
go install github.com/x0f5c3/wsl-dl@latest
```

<!-- tabs:end -->
