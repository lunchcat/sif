<div align="center">
<img src="https://github.com/lunchcat/sif/blob/main/assets/banner.png?raw=true">
</div>

<div align="center">

![Go version](https://img.shields.io/github/go-mod/go-version/dropalldatabases/sif)
[![Go Report Card](https://goreportcard.com/badge/github.com/dropalldatabases/sif)](https://goreportcard.com/report/github.com/dropalldatabases/sif)
[![Version](https://img.shields.io/github/v/tag/dropalldatabases/sif)](https://github.com/dropalldatabases/sif/tags)

[![All Contributors](https://img.shields.io/github/all-contributors/lunchcat/sif?color=ee8449&style=flat-square)](#contributors)

</div>

## 📖 Table of Contents

- [Features](#-features)
- [Installation](#-installation)
- [Quick Start](#-quick-start)
- [Usage](#-usage)
- [Modules](#-modules)
- [Performance](#-performance)
- [Contributing](#-contributing)
- [Contributors](#-contributors)
- [Acknowledgements](#-acknowledgements)

## 🚀 Features

sif is a comprehensive pentesting suite designed for efficiency and ease of use. Our key features include:

- 📂 **Directory/file fuzzing/scanning**: Quickly discover hidden directories and files
- 📡 **DNS subdomain enumeration**: Uncover subdomains associated with target domains
- 🐾 **Common Web scanning**: Identify common web vulnerabilities and misconfigurations
- 🖥️ **Port/service scanning**: Detect open ports and running services on target systems
- 🦠 **Vulnerability scanning**:
  - Support for pre-existing nuclei templates
  - Metasploit emulation for execution
- 🔎 **Automated Google dorking**: Leverage search engines for information gathering
- 💘 **Shodan integration**: Harness the power of Shodan for additional reconnaissance
- 📦 **CMS detection**: Identify content management systems used by target websites
- 🔍 **HTTP Header Analysis**: Analyze HTTP headers for potential security issues
- ☁️ **C3 Misconfiguration Scanner**: Detect common cloud configuration vulnerabilities
- 🔍 **Subdomain Takeover Checks**: Identify potential subdomain takeover vulnerabilities

## 📦 Installation

### Using pre-built binaries

Visit our [Releases](https://github.com/dropalldatabases/sif/releases) page to download the latest pre-built binary for your operating system.

### Building from source

1. Ensure you have Go 1.23+ installed on your system.
2. Clone the repository:
   ```
   git clone https://github.com/lunchcat/sif.git
   cd sif
   ```
3. Build using the Makefile:
   ```
   make
   ```
4. The binary will be available in the root directory.

## 🚀 Quick Start

1. Run a basic scan:
   ```
   ./sif -u example.com
   ```
2. For more options and advanced usage, refer to the help command:
   ```
   ./sif -h
   ```

## 🛠 Usage

sif offers a wide range of commands and options to customize your pentesting workflow. Here are some common usage examples:

- Directory fuzzing

```
./sif -u http://example.com -dirlist medium
```

- Subdomain enumeration

```
./sif -u http://example.com -dnslist medium
```

- Supabase/Firebase and C3 Vulnerability scanning

```
./sif -u https://example.com -js -c3
```

- Port scanning
  
```
./sif -u https://example.com -ports common
```

For a complete list of commands and options, run `./sif -h`.

## 🧩 Modules

sif is built with a modular architecture, allowing for easy extension and customization. Some of our key modules include:

- 📂 Directory/file fuzzing/scanning
- 📡 DNS subdomain enumeration
- 🐾 Common Web scanning
- 🖥️ Port/service scanning
- 🦠 Vulnerability scanning
  - Support for pre-existing nuclei templates
  - Metasploit emulation for execution
- 🔎 Automated Google dorking
- 💘 Shodan integration
- 📦 CMS detection
- 🔍 HTTP Header Analysis
- ☁️ C3 Misconfiguration Scanner
- 🔍 Subdomain Takeover Checks

## ⚡ Performance

sif is designed for high performance and efficiency:

- Written in Go for excellent concurrency and speed
- Optimized algorithms for minimal resource usage
- Supports multi-threading for faster scans
- Efficient caching mechanisms to reduce redundant operations

## 🤝 Contributing

We welcome contributions from the community! Please read our [Contributing Guidelines](CONTRIBUTING.md) before submitting a pull request.

Areas we're particularly interested in:
- New scanning modules
- Performance improvements
- Documentation enhancements
- Bug fixes and error handling improvements

## 🌟 Contributors

Thanks to these wonderful people who have contributed to sif:

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://vmfunc.re"><img src="https://avatars.githubusercontent.com/u/59031302?v=4?s=100" width="100px;" alt="mel"/><br /><sub><b>mel</b></sub></a><br /><a href="#maintenance-vmfunc" title="Maintenance">🚧</a> <a href="#mentoring-vmfunc" title="Mentoring">🧑‍🏫</a> <a href="#projectManagement-vmfunc" title="Project Management">📆</a> <a href="#security-vmfunc" title="Security">🛡️</a> <a href="#test-vmfunc" title="Tests">⚠️</a> <a href="#business-vmfunc" title="Business development">💼</a> <a href="#code-vmfunc" title="Code">💻</a> <a href="#design-vmfunc" title="Design">🎨</a> <a href="#financial-vmfunc" title="Financial">💵</a> <a href="#ideas-vmfunc" title="Ideas, Planning, & Feedback">🤔</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://projectdiscovery.io"><img src="https://avatars.githubusercontent.com/u/50994705?v=4?s=100" width="100px;" alt="ProjectDiscovery"/><br /><sub><b>ProjectDiscovery</b></sub></a><br /><a href="#platform-projectdiscovery" title="Packaging/porting to new platform">📦</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/macdoos"><img src="https://avatars.githubusercontent.com/u/127897805?v=4?s=100" width="100px;" alt="macdoos"/><br /><sub><b>macdoos</b></sub></a><br /><a href="#code-macdoos" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://epitech.eu"><img src="https://avatars.githubusercontent.com/u/75166283?v=4?s=100" width="100px;" alt="Matthieu Witrowiez"/><br /><sub><b>Matthieu Witrowiez</b></sub></a><br /><a href="#ideas-D3adPlays" title="Ideas, Planning, & Feedback">🤔</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/tessa-u-k"><img src="https://avatars.githubusercontent.com/u/109355732?v=4?s=100" width="100px;" alt="tessa "/><br /><sub><b>tessa </b></sub></a><br /><a href="#infra-tessa-u-k" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="#question-tessa-u-k" title="Answering Questions">💬</a> <a href="#userTesting-tessa-u-k" title="User Testing">📓</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/xyzeva"><img src="https://avatars.githubusercontent.com/u/133499694?v=4?s=100" width="100px;" alt="Eva"/><br /><sub><b>Eva</b></sub></a><br /><a href="#blog-xyzeva" title="Blogposts">📝</a> <a href="#content-xyzeva" title="Content">🖋</a> <a href="#research-xyzeva" title="Research">🔬</a> <a href="#security-xyzeva" title="Security">🛡️</a> <a href="#test-xyzeva" title="Tests">⚠️</a> <a href="#code-xyzeva" title="Code">💻</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

## 🙏 Acknowledgements

We'd like to thank the following projects and communities for their inspiration and support:

- [ProjectDiscovery](https://projectdiscovery.io/) for their amazing open-source security tools
- [Shodan](https://www.shodan.io/)
- [Malcore](https://www.malcore.io/), for providing us direct API support at Lunchcat.

---

<div align="center">
  <strong>Happy Hunting! 🐾</strong>
  <p>
    <sub>Built with ❤️ by the lunchcat team and contributors worldwide</sub>
  </p>
</div>
