extern crate kern;

static CARGO_TOML: &'static str = include_str!("../Cargo.toml");

fn main() {
    println!(
        "secpass {} (c) 2018 Lennart Heinrich",
        kern::version(CARGO_TOML).1
    );
}
