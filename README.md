# flutter-tool (community-distro)

A zero-dependency CLI tool that downloads, installs, and manages Flutter SDK on ARM64 Linux — powered by community-maintained GitHub releases.

> Default source: [`zhzhzhy/Flutter-SDK-ARM64`](https://github.com/zhzhzhy/Flutter-SDK-ARM64)

---

## System Requirements

| Requirement | Details |
|-------------|---------|
| OS | Linux |
| Architecture | ARM64 |
| Shell | Bash |
| Dependencies | `wget`, `unzip`, `tar` |
| Optional | `jq` (faster GitHub API parsing) |

---

## Installation

```bash
curl -L https://github.com/<your-org>/flutter-tool/releases/latest/download/flutter-tool \
  -o flutter-tool

chmod +x flutter-tool
sudo mv flutter-tool /usr/local/bin/
```

---

## Usage

```bash
flutter-tool <command> [flags]
```

### Commands

| Command | Description |
|---------|-------------|
| `install` | Install Flutter ARM64 |
| `upgrade` | Upgrade existing Flutter installation |
| `update` | Check for a newer version |
| `doctor` | Run `flutter doctor` |
| `purge` | Completely uninstall Flutter |
| `version`, `-v` | Show installer version |
| `help`, `-h` | Show help menu |

### Flags

| Flag | Description |
|------|-------------|
| `--yes`, `-y` | Auto-confirm all prompts |
| `--no-api` | Skip GitHub API; use fallback download URL |
| `--purge-cache` | Also remove the cache directory during purge |
| `--upgrade` | Upgrade Flutter if already installed |
| `--fix-dart` | Replace bundled Dart SDK with the official ARM64 binary |

---

## Examples

```bash
# Install Flutter
flutter-tool install

# Install and fix Dart SDK for ARM64 compatibility
flutter-tool install --fix-dart

# Auto-upgrade if a newer version is available
flutter-tool update --yes

# Install without using the GitHub API
flutter-tool install --no-api

# Completely remove Flutter including cache
flutter-tool purge --purge-cache
```

---

## Post-Installation

After installation, two helper functions are added to your `~/.bashrc`:

```bash
# Reload your terminal
source ~/.bashrc

# Add Flutter to PATH
flutter-on

# Remove Flutter from PATH
flutter-off
```

---

## Environment Variables

Use a custom source repo or pin a specific version:

```bash
export FLUTTER_REPO_URL="username/my-flutter-arm64-repo"
export FLUTTER_VERSION="3.19.0"

flutter-tool install
```

| Variable | Default | Description |
|----------|---------|-------------|
| `FLUTTER_REPO_URL` | `zhzhzhy/Flutter-SDK-ARM64` | GitHub `owner/repo` |
| `FLUTTER_VERSION` | `latest` | Specific release tag or `latest` |

---

## Installation Paths

| Purpose | Path |
|---------|------|
| Flutter SDK | `~/.flutter-tool/flutter/` |
| Flutter binary | `~/.flutter-tool/flutter/bin/flutter` |
| Cache | `~/.cache/flutter-tool/` |
| Bin dir | `~/.local/bin/` |

---

## When to Use `--fix-dart`

Some Flutter releases bundle a Dart SDK that is not compatible with ARM64. If you encounter Dart-related errors after installation, run:

```bash
flutter-tool install --fix-dart
```

This downloads the official ARM64 Dart SDK from Google's storage and replaces the bundled one.

---

## How It Works

```
flutter-tool install
  │
  ├── 1. Detect latest release tag via GitHub API
  ├── 2. Find and download the ARM64 asset (.zip or .tar.xz)
  ├── 3. Extract to ~/.flutter-tool/flutter/
  ├── 4. Run Flutter preflight (cache warmup)
  ├── 5. Optionally replace Dart SDK (--fix-dart)
  ├── 6. Run flutter doctor
  └── 7. Inject flutter-on / flutter-off into ~/.bashrc
```

---

## 🙏 Special Thanks

This tool wouldn't exist without the incredible work of the community members who build and maintain ARM64 Flutter SDK releases. Huge, heartfelt thanks to:

### [`zhzhzhy`](https://github.com/zhzhzhy) — [`zhzhzhy/Flutter-SDK-ARM64`](https://github.com/zhzhzhy/Flutter-SDK-ARM64)

The backbone of this project. **zhzhzhy** consistently builds and publishes Flutter SDK releases compiled for ARM64 Linux — something the official Flutter team doesn't provide. Without this repository, running Flutter on ARM64 hardware would be a painful, manual process for thousands of developers. This tool is essentially a polished wrapper around their incredible effort.

If you find their work useful, please consider starring their repository and showing your appreciation.

---

## License

See the `LICENSE` file for license details.
