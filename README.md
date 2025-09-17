# flake-store-cli

![flake-store-cli](https://img.shields.io/badge/status-beta-yellow)

`flake-store-cli` is an interactive CLI application to browse, search, and install [Nix flakes](https://nixos.wiki/wiki/Flakes) from the `flake-store-flakes` repository. With a user-friendly TUI powered by Bubble Tea, you can select and install flakes seamlessly into your project directories.

---

## Demo

![TUI Demo](https://github.com/user-attachments/assets/f3baa88d-4b54-42d3-a7b2-2dd8a4f79d33)

---

## Installation

Make sure you have Go installed:

```bash
git clone https://github.com/SaladinAyyub/flake-store-cli.git
cd flake-store-cli
go build -o flake-store-cli
```

Or install using `go install`:

```bash
go install github.com/SaladinAyyub/flake-store-cli@latest
```

---

## Usage

### Launch Interactive TUI (default)

Simply run:

```bash
flake-store-cli list
```

This will launch the interactive TUI to browse and install flakes.  

- Use **arrow keys** to navigate.  
- Press **Enter** to install a flake.  
- Press **q** or **Esc** to quit.

---

### Commands

| Command | Description |
|---------|-------------|
| `flake-store-cli list` | Launch the interactive TUI to browse flakes |
| `flake-store-cli install <flake-name> [--nodirenv]` | Install `flake.nix` and optional `shell.nix` from the specified flake into the current directory. Creates `.envrc` unless `--nodirenv` is provided |
| `flake-store-cli search <query>` | Search flakes by name and display matching flakes |

---

### Install a Flake via CLI

```bash
flake-store-cli install hikari-crescent-python
```

- Installs `flake.nix` and optional `shell.nix` from the specified flake into the current directory.
- Creates a `.envrc` file for `direnv` automatically unless `--nodirenv` is specified.

---

### Search Flakes

```bash
flake-store-cli search <query>
```

Search flakes by name and print matching results.

Example:

```bash
flake-store-cli search raylib
```

---

## Contributing

Want to add a new flake? Contributions are welcome!  

Please submit flakes as **Pull Requests** to the [`flake-store-flakes`](https://github.com/SaladinAyyub/flake-store-flakes) repository.  

Steps:

1. Fork the [`flake-store-flakes`](https://github.com/SaladinAyyub/flake-store-flakes) repository.  
2. Add your flake in a new folder with `flake.nix` (and optional `shell.nix`).  
3. Update the description in `flake.nix`.  
4. Submit a Pull Request.  

Once merged, your flake will automatically appear in `flake-store-cli`.

---

## Supporting

Support me via ko-fi !

<a href='https://ko-fi.com/G2G8TLR6I' target='_blank'><img height='50' style='border:0px;height:50px;' src='https://storage.ko-fi.com/cdn/brandasset/kofi_s_tag_white.png?' border='0' alt='Donate and Support' /></a>

## License

GPL V3
