#!/usr/bin/env python3
import unittest

from verify_cli_release import (
    AvailabilityError,
    expected_assets,
    validate_checksum_manifest,
    validate_release,
)


class CLIReleaseAvailabilityTests(unittest.TestCase):
    def payload(self, tag="v1.2.3"):
        assets = [
            {"name": name, "size": 42, "browser_download_url": "https://example.test/" + name}
            for name in expected_assets(tag)
        ]
        return {"tag_name": tag, "draft": False, "prerelease": False, "assets": assets}

    def test_exact_published_release_with_expected_assets_is_valid(self):
        self.assertEqual(validate_release(self.payload(), "v1.2.3"), expected_assets("v1.2.3"))

    def test_wrong_draft_prerelease_or_tag_fails_closed(self):
        for mutation in (
            {"tag_name": "v9.9.9"},
            {"draft": True},
            {"prerelease": True},
        ):
            payload = self.payload()
            payload.update(mutation)
            with self.subTest(mutation=mutation), self.assertRaises(AvailabilityError):
                validate_release(payload, "v1.2.3")

    def test_missing_empty_or_non_downloadable_asset_fails_closed(self):
        payload = self.payload()
        payload["assets"].pop()
        with self.assertRaisesRegex(AvailabilityError, "missing release assets"):
            validate_release(payload, "v1.2.3")

        for mutation in ({"size": 0}, {"browser_download_url": "http://unsafe.test/x"}):
            payload = self.payload()
            payload["assets"][0].update(mutation)
            with self.subTest(mutation=mutation), self.assertRaises(AvailabilityError):
                validate_release(payload, "v1.2.3")

    def test_checksum_manifest_must_cover_primary_archives(self):
        required = expected_assets("v1.2.3")
        manifest = "\n".join("abc  %s" % name for name in required[1:])
        validate_checksum_manifest(manifest, required)
        with self.assertRaisesRegex(AvailabilityError, "checksum manifest omits"):
            validate_checksum_manifest(manifest.replace(required[-1], "missing.zip"), required)

    def test_tag_shape_is_strict(self):
        with self.assertRaises(AvailabilityError):
            expected_assets("1.2.3")


if __name__ == "__main__":
    unittest.main()
