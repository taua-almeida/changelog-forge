# Changelog Forge

`changelog-forge` is a CLI tool for generating changelog entries and managing version updates based on semantic versioning.

## Features

- Generate JSON entries for changelogs.
- Update changelog files with semantic versioning.
- Automatically update the CLI version based on the latest changelog entry.

## Installation

### Using Precompiled Binaries

Precompiled binaries for Linux, macOS, and Windows are available on the [Releases](https://github.com/taua-almeida/changelog-forge/releases) page.

1. Download the binary for your platform.
2. Move the binary to your `$PATH`. Example:
```bash
   mv changelog-forge-linux-amd64 /usr/local/bin/changelog-forge
   chmod +x /usr/local/bin/changelog-forge
```

### Install via `go install`
Ensure you have Go installed (v1.20+ recommended).
```bash
go install github.com/taua-almeida/changelog-forge@latest
```

### Platform-Specific Installation Scripts

#### Linux:
```bash
#!/bin/bash

VERSION=$(curl -s https://api.github.com/repos/taua-almeida/changelog-forge/releases/latest | grep tag_name | cut -d '"' -f 4)
URL="https://github.com/taua-almeida/changelog-forge/releases/download/$VERSION/changelog-forge-linux-amd64"

echo "Downloading changelog-forge $VERSION for Linux..."
curl -L $URL -o changelog-forge
chmod +x changelog-forge
sudo mv changelog-forge /usr/local/bin/

echo "Installation complete. Run 'changelog-forge --version' to verify."
```

#### MacOS:
```bash
#!/bin/bash

VERSION=$(curl -s https://api.github.com/repos/taua-almeida/changelog-forge/releases/latest | grep tag_name | cut -d '"' -f 4)
URL="https://github.com/taua-almeida/changelog-forge/releases/download/$VERSION/changelog-forge-darwin-amd64"

echo "Downloading changelog-forge $VERSION for macOS..."
curl -L $URL -o changelog-forge
chmod +x changelog-forge
sudo mv changelog-forge /usr/local/bin/

echo "Installation complete. Run 'changelog-forge --version' to verify."
```

#### Windows:
```PowerShell
$version = Invoke-RestMethod -Uri https://api.github.com/repos/taua-almeida/changelog-forge/releases/latest | Select-String -Pattern '"tag_name":' | ForEach-Object { $_ -replace '"tag_name": "', '' -replace '",', '' }
$url = "https://github.com/taua-almeida/changelog-forge/releases/download/$version/changelog-forge-windows-amd64.exe"

Write-Host "Downloading changelog-forge $version for Windows..."
Invoke-WebRequest -Uri $url -OutFile "changelog-forge.exe"
Move-Item -Path "changelog-forge.exe" -Destination "C:\Windows\System32\"

Write-Host "Installation complete. Run 'changelog-forge --version' to verify."
```

### Build from Source
1. Clone the repository:
```bash
git clone https://github.com/taua-almeida/changelog-forge.git
cd changelog-forge
```

2. Build the binary:
```bash
go build -ldflags "-X 'main.Version=$(git describe --tags --abbrev=0)'" -o changelog-forge
```

3. Move the binary to your `$PATH`:
```bash
mv changelog-forge /usr/local/bin/
```

## Usage

### Generate a New Changelog Entry

Run the CLI to generate a new changelog entry:
```bash
changelog-forge --generate-json
```
> Follow the prompts to select the version type and enter descriptions.

### Update the Changelog
Update the `CHANGELOG.m` with the generated JSON:

```bash
changelog-forge --update-changelog
```

### Pre-Commit Hook
This repository includes a pre-commit hook to ensure that:

- The `.changeset` folder contains valid JSON files before committing.

Install `pre-commit`:
```bash
pip install pre-commit
```

Install hooks in the repository:
```bash
pre-commit install
```