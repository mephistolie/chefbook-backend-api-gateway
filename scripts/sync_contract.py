#!/usr/bin/env python3
"""Explicitly vendor a committed ChefBook contract; normal builds never fetch it."""

import argparse
import hashlib
import json
from pathlib import Path
import subprocess
import tempfile

ROOT = Path(__file__).resolve().parents[1]
DESTINATION = ROOT / "contracts"
SPEC = DESTINATION / "chefbook.yaml"
LOCK = DESTINATION / "chefbook.lock.json"


def git(source, *args):
    return subprocess.check_output(["git", "-C", str(source), *args])


def atomic_write(path, data):
    with tempfile.NamedTemporaryFile(dir=path.parent, delete=False) as stream:
        temporary = Path(stream.name)
        stream.write(data)
    try:
        temporary.replace(path)
    finally:
        temporary.unlink(missing_ok=True)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--source", type=Path, help="Local chefbook-contracts Git checkout")
    parser.add_argument("--ref", help="Explicit Git commit or tag (never an implicit latest)")
    parser.add_argument("--working-tree", action="store_true", help="Explicitly vendor a local draft with SHA-256 and base commit provenance")
    parser.add_argument("--check", action="store_true", help="Check the vendored snapshot without Git/network")
    args = parser.parse_args()
    if args.check:
        if args.source or args.ref or args.working_tree:
            parser.error("--check cannot be combined with --source or --ref")
        lock = json.loads(LOCK.read_text())
        actual = hashlib.sha256(SPEC.read_bytes()).hexdigest()
        if actual != lock["sha256"]:
            raise SystemExit("Contract checksum mismatch; update with this script, do not edit the snapshot")
        print(f"Contract {lock['commit'][:12]}: checksum verified")
        return
    if not args.source or (not args.ref and not args.working_tree):
        parser.error("updating requires --source and either --ref or --working-tree")
    if args.ref and args.working_tree:
        parser.error("--ref and --working-tree are mutually exclusive")
    commit = git(args.source, "rev-parse", "--verify", "--end-of-options", (args.ref or "HEAD") + "^{commit}").decode().strip()
    data = (args.source / "openapi/chefbook.yaml").read_bytes() if args.working_tree else git(args.source, "show", commit + ":openapi/chefbook.yaml")
    lock = {
        "repository": "https://github.com/mephistolie/chefbook-contracts.git",
        "commit": commit,
        "path": "openapi/chefbook.yaml",
        "sha256": hashlib.sha256(data).hexdigest(),
    }
    if args.working_tree:
        lock["workingTree"] = True
    DESTINATION.mkdir(parents=True, exist_ok=True)
    atomic_write(SPEC, data)
    atomic_write(LOCK, (json.dumps(lock, indent=2) + "\n").encode())
    print(f"Pinned contract {commit}; source: {'working tree draft' if args.working_tree else 'committed content'}")


if __name__ == "__main__":
    main()
