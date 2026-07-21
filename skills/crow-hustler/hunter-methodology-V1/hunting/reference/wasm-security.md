> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/hunting/00-index|← Back to Hunting]]

# WebAssembly (WASM) Security

Emerging attack surface as WASM adoption grows.

## Vulnerability Classes

| Class | Description |
|-------|-------------|
| **Memory safety issues** | In WASM modules |
| **Sandbox escape** | Via WASI (WebAssembly System Interface) APIs |
| **Side-channel attacks** | Timing analysis |
| **Import/export manipulation** | Tamper with module interfaces |

Impact: Can escalate to RCE in server-side WASM environments.

Tools: `wasmtime-fuzzer`, `wasm-decompiler`, `wasm-security-analyzer`
