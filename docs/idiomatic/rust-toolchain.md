# The Rustup

### Installing
```bash
curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh
```

### Upgrading
```bash
rustup update
```

### Uninstalling
```bash
rustup self uninstall
```

# Rustfmt - The formatter

# Rustc - The compiler

### Run from a file
```bash
rustc main.rs
```

### Version checking
```bash
rustc --version
```

# Cargo - The Rust's build system and package manager

### Version checking 
```bash
cargo --version
```

### Creating project/crate
```bash
cargo new project_a --vcs=none
```

### Building project
```bash
cargo build
```

Or, for release:
```bash
cargo build --release
```

### Running project
```bash
cargo run
```

### Checking project
```bash
cargo check
```
Often, `cargo check` is much faster than `cargo build` because it skips the step of producing an executable.

> [!INFO]
> With simple projects, Cargo doesn’t provide a lot of value over just using rustc, 
> but it will prove its worth as your programs become more intricate.\
> Once programs grow to multiple files or need a dependency, 
> it’s much easier to let Cargo coordinate the build.


