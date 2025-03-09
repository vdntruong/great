# product service

Many Rust programmers call themselves “Rustaceans,”. And Ferris is the unofficial mascot of the Rust Community.

## Project desc

`/src/main.rs` is the entry point of the application.
## Rust command

Formats all bin and lib files of the current crate:
```bash
cargo fmt
```

Run the app:
```bash
cargo run
```

Update Rust and the toolchains:
```bash
rustup update
```

From `services` directory, to create product project:
```bash
cargo new product-ms
```

Adding a crate
```bash
cargo add ferris-says serde serde_json serder_derive
```

(optional) we can build the project with new added crate:
```bash
cargo build
```

(optional) another way to build project:
```bash
rustc main.rs
```
