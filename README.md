<!-- omit in toc -->
# Commenter CLI

<!-- omit in toc -->
## Contents

- [📘 Description](#-description)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
- [🚢 Installation](#-installation)
- [🔧 Usage](#-usage)
  - [Commands](#commands)
  - [Global Flags](#global-flags)
  - [Examples](#examples)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)

## 📘 Description

Commenter CLI is a command-line tool designed to perform comment operations on GitHub pull request (PR) issues.
It simplifies creating, updating, and retrieving comments on GitHub PRs directly from your terminal.

### Features

- **Create Comments**: Add new comments to a specific PR.
- **Update Comments**: Modify existing comments on a PR.
- **Get Comment ID**: Retrieve the ID of a comment based on its text content.
- **Autocompletion**: Generate autocompletion scripts for various shells.

### Prerequisites

Before using this tool, ensure you have the following:

Go: Install Go programming language.
GitHub Token: A valid GitHub personal access token with appropriate permissions.

## 🚢 Installation

To install Commenter CLI from the source, follow these steps:

1. Clone the Repository:

```bash
git clone https://github.com/yourusername/commenter-cli.git
cd commenter-cli
```

2. Download Dependencies:

```bash
go mod download
```

3. Build the Binary:

```bash
go build -o commenter
```

## 🔧 Usage

Once installed, you'll have to set the `GITHUB_TOKEN` environment then can use the commenter command to interact
with GitHub PR issues.

### Commands

- **create**: Create a new comment on a PR.
- **get**: Get the message ID based on text.
- **help**: Display help information about any command.
- **update**: Update an existing comment on a PR.

### Global Flags

-h, --help: Show help for the commenter command.
-o, --org string: Specify the GitHub organization name.
-p, --pull int: Specify the target pull request number.
-r, --repo string: Specify the target repository name.

### Examples

To create a new comment on a specific PR:

```bash
./commenter create -o your-org -r your-repo -p 123 -m "This is a comment"
```

To update an existing comment:

```bash
./commenter update -o your-org -r your-repo -i 456 -m "Updated comment text"
```

To retrieve the ID of a comment based on its text:

```bash
./commenter get -o your-org -r your-repo -p 123 -t "Comment text to find"
```

## 🤝 Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (git checkout -b feature-branch).
3. Make your changes.
4. Commit your changes (git commit -m 'Add new feature').
5. Push to the branch (git push origin feature-branch).
6. Open a Pull Request.

## 📄 License

This project is licensed under the MIT License. See the LICENSE file for details.
