#!/bin/sh
t=x86_64-unknown-linux-musl;cargo test&&cargo clean&&cargo build --release --target $t&&strip target/$t/release/secpass
