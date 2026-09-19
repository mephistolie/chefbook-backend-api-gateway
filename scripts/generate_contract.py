#!/usr/bin/env python3
"""Generate Gin routes and types from the pinned OpenAPI snapshot."""

import argparse
from pathlib import Path
import subprocess

ROOT = Path(__file__).resolve().parents[1]
OUTPUT = ROOT / "internal/transport/http/contract/api.gen.go"


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--check", action="store_true")
    args = parser.parse_args()
    subprocess.run(["python3", str(ROOT / "scripts/sync_contract.py"), "--check"], check=True)
    generated = subprocess.check_output([
        "go", "run", "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1",
        "-generate", "types,gin", "-package", "contract", "contracts/chefbook.yaml",
    ], cwd=ROOT)
    if args.check:
        if OUTPUT.read_bytes() != generated:
            raise SystemExit("Generated Go contract is stale; run scripts/generate_contract.py")
        print("Go contract generation verified")
    else:
        OUTPUT.write_bytes(generated)


if __name__ == "__main__":
    main()
