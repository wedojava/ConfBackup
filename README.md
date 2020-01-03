# ConfBackup
This project developed for backup config of juniper devices via Telnet.

# Build

```bash
go build -ldflags "-s -w" -o ./dist/ConfBackup.exe ./cmd/run.go
```
- `-s`'s effect is to remove the symbol information
- `-w`'s effect is to remove debugging information