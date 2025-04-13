# Let Me Google That (command line tool)

Sometimes junior engineers ask you frustrating questions ("How to pass arguments to a function"), questions that can be easily found using Google. So, instead of opening your browser to visit https://letmegooglethat.com/ you can stay in your terminal and get the same effect.

## How to Install

1. Download the latest executable from **Releases**
2. Rename the executable to ```lmgt```
3. Move the executable to your ```/usr/local/bin``` directory
4. Restart your shell: ```exec zsh``` or equivalent

Note: You may need to update your ```$PATH```

## How to Use

```zsh
Usage: lmgt "prompt here" [options]
Example: lmgt "when is presidents day"
Example (copy to clipboard): lmgt "when is presidents day" copy
```

## Install clipboard tool (Linux)

```zsh
sudo apt install xclip -y
```
