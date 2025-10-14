# BUNKR

**Encrypted file sharing over your local network.**

A Go tool that lets you share files securely within your LAN. Host files from one device, download them from another — all encrypted.

Why route files through the cloud when your devices are already on the same network?

## How It Works

**Host Mode:**  
- Upload and encrypt files via a terminal UI
- Get a shareable LAN URL.
- Files exist only while the server runs

**Client Mode:**  
- Connect using the host's URL
- Enter the password (if set)
- Download and decrypt files directly

## Features

- **Cross-platform** — Linux, macOS, Windows
- **Password-protected transfers** — optional but recommended
- **LAN-only** — never leaves your network
- **Terminal UI** — powered by Bubble Tea

https://github.com/user-attachments/assets/14adc4d5-6239-4af5-9a6e-2dc09b2001f2

## Quick Start

Simply clone, build, and run:

```bash
git clone https://github.com/rrajvardhan/bunkr
cd bunkr
go build -o bunkr ./cmd/bunkr

# Run the binary
./bunkr
```
or, skip build:

```bash
go run ./cmd/bunkr/main.go
```
