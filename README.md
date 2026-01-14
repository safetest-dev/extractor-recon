# 🔎 Extractor Recon

> A minimal, deterministic HTTP reconnaissance tool for security engineers and bug hunters.

**Extractor Recon** is a lightweight CLI tool written in Go for **HTTP status inspection and link response analysis**.  
It fetches a target URL, extracts links from HTML content, and performs **meta-requests (curl-like checks)** on each discovered link to observe server behavior.

This tool is intentionally **non-invasive**, **non-recursive**, and **deterministic** — designed for **early-stage reconnaissance and analysis**, not exploitation.

---

## ✨ Key Features

- ✅ HTTP status inspection (200 / 301 / 403 / 404 / 500 / etc)
- 🔗 HTML link extraction using a real parser (no regex)
- 🧭 Relative → absolute URL normalization
- 🔍 Meta-request for each extracted link (curl-style GET)
- 📄 JSON or human-readable output
- 📂 Single URL or file-based multi-target input
- ⚙️ Single static binary (easy to deploy)

---

## 🧠 Design Philosophy

- Observation over exploitation  
- Behavior over payloads  
- Explicit over magical  
- Readable over clever  

Built with long-term maintainability and auditability in mind.

---

## 📦 Download

```bash
git clone https://github.com/safetest-dev/extractor-recon.git
cd extractor-recon
```

---

## 🛠️ Requirements

- Go 1.20+ (1.21 / 1.22 recommended)
- Linux, macOS, or Windows

Verify Go installation:

```bash
go version
```

---

## 🧱 Compile

Initialize dependencies (first time only):

```bash
go mod init github.com/safetest-dev/extractor-recon
go get golang.org/x/net/html
```

Build the binary:

```bash
go build -o extractor ./cmd/extractor
```

(Optional)

```bash
chmod +x extractor
```

---

## ▶️ Usage

### Scan a single URL

```bash
./extractor https://example.com
```

### Status check only

```bash
./extractor --status-only https://example.com
```

### Follow redirects

```bash
./extractor --follow-redirect https://example.com
```

### JSON output

```bash
./extractor --json https://example.com
```

### Multiple targets

```bash
./extractor urls.txt
```

---

## 🌍 Cross-Platform Build

Linux:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o extractor-linux ./cmd/extractor
```

macOS:

```bash
GOOS=darwin GOARCH=arm64 go build -o extractor-macos ./cmd/extractor
```

Windows:

```bash
GOOS=windows GOARCH=amd64 go build -o extractor.exe ./cmd/extractor
```

---

## 🔐 Intended Use

- Bug bounty reconnaissance
- HTTP behavior analysis
- Status-code mapping
- Security research & education

---

## 🚫 Non-Goals

- No recursive crawling
- No fuzzing or exploitation
- No auth bypass attempts

---

## 📜 Disclaimer

This tool is intended for **authorized security testing and educational purposes only**.
