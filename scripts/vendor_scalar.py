#!/usr/bin/env python3
"""Vendor a pinned Scalar standalone bundle; normal Go builds stay offline."""
import argparse
import base64
import gzip
import hashlib
import io
import json
from pathlib import Path
import tarfile
import urllib.request

ROOT = Path(__file__).resolve().parents[1]
DEST = ROOT / 'internal/transport/http/router/docs/assets'
VERSION = '1.69.0'
URL = f'https://registry.npmjs.org/@scalar/api-reference/-/api-reference-{VERSION}.tgz'
INTEGRITY = 'N+prEf0YjQByZ9JmiQhrJoSrv2UgIPcD9VPE4JTf3rWkdRNESiVBM4J/rUi9hO6JGo7eDcdEMLhQ3Md26LiXdg=='

def main():
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument('--archive', type=Path, help='Previously downloaded npm archive')
    p.add_argument('--check', action='store_true')
    args = p.parse_args()
    bundle = DEST / f'scalar-{VERSION}.js.gz'
    lock = DEST / 'scalar.lock.json'
    if args.check:
        pin = json.loads(lock.read_text())
        assert pin['version'] == VERSION
        assert pin['integrity'] == 'sha512-' + INTEGRITY
        assert hashlib.sha256(bundle.read_bytes()).hexdigest() == pin['gzip_sha256']
        assert hashlib.sha256(gzip.decompress(bundle.read_bytes())).hexdigest() == pin['js_sha256']
        print('Scalar bundle checksum verified')
        return
    data = args.archive.read_bytes() if args.archive else urllib.request.urlopen(URL, timeout=60).read()
    if base64.b64encode(hashlib.sha512(data).digest()).decode() != INTEGRITY:
        raise SystemExit('Scalar npm archive integrity mismatch')
    with tarfile.open(fileobj=io.BytesIO(data), mode='r:gz') as archive:
        js = archive.extractfile('package/dist/browser/standalone.js').read()
    compressed = gzip.compress(js, mtime=0)
    DEST.mkdir(parents=True, exist_ok=True)
    bundle.write_bytes(compressed)
    lock.write_text(json.dumps({'package': '@scalar/api-reference', 'version': VERSION,
        'url': URL, 'integrity': 'sha512-' + INTEGRITY,
        'js_sha256': hashlib.sha256(js).hexdigest(),
        'gzip_sha256': hashlib.sha256(compressed).hexdigest()}, indent=2) + '\n')
    print(f'Vendored Scalar {VERSION}: {len(compressed)} compressed bytes')

if __name__ == '__main__':
    main()
