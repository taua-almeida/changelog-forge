# Changelog Forge

`changelog-forge` is a CLI tool for generating changelog entries and managing version updates based on semantic versioning.

## Features

- Generate JSON entries for changelogs.
- Update changelog files with semantic versioning.
- Automatically update the CLI version based on the latest changelog entry.

## Installation

### Build from Source
1. Clone the repository:
```bash
   git clone https://github.com/taua-almeida/changelog-forge.git
   cd changelog-forge
```

2. Build the binary:
```bash
go build -o changelog-forge ./cmd
```

3. Move the binary to your $PATH:
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


### Update the Changelog:

Update the CHANGELOG.md with the generated JSON:
```bash
changelog-forge --update-changelog
```

### Pre-Commit Hook
This repository includes a pre-commit hook to ensure that:

- The changelog is updated before each commit.
- The CLI version is synchronized with the latest changelog entry.

```yaml
### Pre-Commit Hook

**1. Install `pre-commit`:**
```bash
pip install pre-commit
```