# ImJustAFish

"I'm just a fish" meme in fullscreen Google Chrome kiosk window.

## Requirements:

- `Google Chrome` installed.
- `Go Programming Language` installed.

## How to setup and run:

Install the requirements. Run `setup.sh` or enter the given commands manually: 
```bash
cd project-dir
go build -ldflags "-H=windowsgui" -o fish.exe src/main.go
fish.exe
```
To leave use `ALT + F4`.

## FAQ
- **Why I'm not providing the ready .exe file?** Transparency and... Windows Defender. The program isn't certified and will be captured by Defender unless it has been made on your computer. I'm not sure what exactly is causing this.
