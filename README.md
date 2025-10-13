# BUNKR

A lightweight, secure, LAN-only file-sharing tool written entirely in Go. It operates in two modes: **host** or **client** using a single binary.

* **Host Mode:** Upload files (encrypted locally), generate LAN URL/QR code.
* **Client Mode:** Connect via URL, provide optional password, download and decrypt files.
* Files exist **only while the server is running**.

## Features

* **Single Binary:** Run as host or client.
* **Host-only uploads:** Only the host can add files.
* **Client downloads:** Clients can download files with optional password.
* **LAN-only operation:** Works over same Wi-Fi or local network.
* **TUI interface:** Simple UI for host and client.

**Host Mode:**
* Upload/Manage files using the host TUI.
* LAN URL and QR code will be displayed for clients.

**Client Mode:**
* Enter server URL and optional password.
* List available files and download them.

## A Quick Look

https://github.com/user-attachments/assets/14adc4d5-6239-4af5-9a6e-2dc09b2001f2

## Build Instructions

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
