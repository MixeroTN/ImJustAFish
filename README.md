# ImJustAFish

"I'm just a fish" meme in fullscreen Google Chrome kiosk window.

## Requirements

- `Google Chrome` installed.
- `Go Programming Language` installed (to build by yourself).

## How to setup and run

Install the requirements. Run `setup.sh` or enter the given commands manually:

```bash
cd project-dir
go build -ldflags "-H=windowsgui" -o fish.exe src/main.go
./fish.exe
```

To leave use <kbd>ALT</kbd> + <kbd>F4</kbd>.

## FAQ

- **Windows Defender warns about .exe downloaded from the Releases page!** The program isn't certified and can be even captured by Defender unless it has been made on your computer. We also know this code can be understood as malicious.
