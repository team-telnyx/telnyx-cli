#!/usr/bin/env python3
"""Verify an exact CLI GitHub Release and its checksum manifest are available."""
from __future__ import annotations

import argparse
import json
import os
import subprocess
import tempfile
import time
from typing import Mapping, Sequence


class AvailabilityError(RuntimeError):
    pass


def expected_assets(tag: str) -> tuple[str, ...]:
    if not tag.startswith("v") or len(tag) < 2:
        raise AvailabilityError("release tag must start with v")
    version = tag[1:]
    return (
        "telnyx_%s_checksums.txt" % version,
        "telnyx_%s_linux_amd64.tar.gz" % version,
        "telnyx_%s_linux_arm64.tar.gz" % version,
        "telnyx_%s_macos_amd64.zip" % version,
        "telnyx_%s_macos_arm64.zip" % version,
        "telnyx_%s_windows_amd64.zip" % version,
    )


def validate_release(payload: Mapping[str, object], tag: str) -> tuple[str, ...]:
    if payload.get("tag_name") != tag:
        raise AvailabilityError("release tag mismatch")
    if payload.get("draft") is not False or payload.get("prerelease") is not False:
        raise AvailabilityError("release is not a final published release")
    raw_assets = payload.get("assets")
    if not isinstance(raw_assets, list):
        raise AvailabilityError("release assets payload is invalid")
    assets = {}
    for raw in raw_assets:
        if not isinstance(raw, Mapping):
            raise AvailabilityError("release asset is invalid")
        name, size, url = raw.get("name"), raw.get("size"), raw.get("browser_download_url")
        if isinstance(name, str):
            assets[name] = (size, url)
    required = expected_assets(tag)
    missing = [name for name in required if name not in assets]
    if missing:
        raise AvailabilityError("missing release assets: %s" % ", ".join(missing))
    invalid = [
        name
        for name in required
        if not isinstance(assets[name][0], int)
        or assets[name][0] <= 0
        or not isinstance(assets[name][1], str)
        or not assets[name][1].startswith("https://")
    ]
    if invalid:
        raise AvailabilityError("release assets are not downloadable: %s" % ", ".join(invalid))
    return required


def validate_checksum_manifest(text: str, required: Sequence[str]) -> None:
    missing = [name for name in required[1:] if name not in text]
    if missing:
        raise AvailabilityError("checksum manifest omits: %s" % ", ".join(missing))


def api_release(repository: str, tag: str) -> Mapping[str, object]:
    completed = subprocess.run(
        ["gh", "api", "repos/%s/releases/tags/%s" % (repository, tag)],
        text=True,
        capture_output=True,
        env=os.environ,
    )
    if completed.returncode != 0:
        raise AvailabilityError("release API lookup failed")
    payload = json.loads(completed.stdout)
    if not isinstance(payload, Mapping):
        raise AvailabilityError("release API payload is invalid")
    return payload


def verify(repository: str, tag: str, attempts: int = 20, delay: int = 15) -> None:
    last_error: Exception | None = None
    for attempt in range(attempts):
        try:
            required = validate_release(api_release(repository, tag), tag)
            with tempfile.TemporaryDirectory() as directory:
                completed = subprocess.run(
                    [
                        "gh", "release", "download", tag, "--repo", repository,
                        "--pattern", required[0], "--dir", directory,
                    ],
                    text=True,
                    capture_output=True,
                    env=os.environ,
                )
                if completed.returncode != 0:
                    raise AvailabilityError("checksum asset download failed")
                with open(os.path.join(directory, required[0]), encoding="utf-8") as handle:
                    validate_checksum_manifest(handle.read(), required)
            return
        except (AvailabilityError, json.JSONDecodeError, OSError) as exc:
            last_error = exc
            if attempt + 1 < attempts:
                time.sleep(delay)
    raise AvailabilityError("release did not become available: %s" % last_error)


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--repository", required=True)
    parser.add_argument("--tag", required=True)
    args = parser.parse_args()
    verify(args.repository, args.tag)
    print("verified published CLI release %s" % args.tag)
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except AvailabilityError as exc:
        print("release availability failed: %s" % exc)
        raise SystemExit(1)
