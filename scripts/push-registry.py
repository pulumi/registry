#!/usr/bin/env python3

# Can be run without dependencies with
#
#	uv run --with requests,pyyaml scripts/push-registry.py

import os
import tempfile
import subprocess
import requests
import json
import yaml

yaml_db_path = "themes/default/data/registry/packages"
docs_path = "themes/default/content/registry/packages"
tmpdir = tempfile.mkdtemp()

print(f"Tempdir: {tmpdir}")

run_script = []

for filename in os.listdir(yaml_db_path):
    if not filename.endswith('.yaml'):
        break

    filepath = os.path.join(yaml_db_path, filename)
    with open(filepath) as file:
        data = yaml.safe_load(file)

    version = data.get('version')

    schema_url = data.get(
        'schema_file_url',
        f"https://raw.githubusercontent.com/{data.get('repo_url').removeprefix("https://github.com/")}/refs/tags/{version}/{data.get('schema_file_path')}",
    )
    source = 'pulumi'
    if 'registry.opentofu.org' in schema_url:
        source = 'opentofu'

    name = data.get('name')
    publisher = data.get('publisher')

    package_version_name = f"{source}/{publisher}/{name}@{version}"
    print(f"Preparing command for {package_version_name}")

    # We skip azure-native-v1, since it will (1) never be updated again and (2) isn't
    # actually a package, it's an alias to the azure native package.
    #
    # If azure-native-v2 ships (with azure-native now at v3), we will want to exclude
    # azure-native-v2 also.
    if name.startswith("azure-native") and name != "azure-native":
        continue

    api_url = f"https://api.pulumi.com/api/registry/preview/packages/{source}/{publisher}/{name}/versions/{version}"
    existence_check = requests.head(api_url).status_code

    if existence_check == 404:
        schema_file = os.path.join(tmpdir, name, "schema.json")
        os.makedirs(os.path.dirname(schema_file), exist_ok=True)

        with open(schema_file, 'wb') as sf:
            response = requests.get(schema_url)
            response.raise_for_status()
            schema = json.loads(response.content)
            if not schema["version"]:
                schema["version"] = version.removeprefix("v")
            sf.write(json.dump())

        cmd = [
            "pulumi", "package", "publish", schema_file,
            f"--readme={docs_path}/{name}/_index.md",
            f"--source={source}",
            f"--publisher={publisher}",
        ]

        if os.path.isfile(os.path.join(docs_path, name, "_installation-configuration.md")):
            cmd.extend(["--installation-configuration", f"{docs_path}/{name}/installation-configuration.md"])

        run_script.append(' '.join(cmd))
    elif existence_check == 200:
        print("Package already exists in the registry DB")
    else:
        print(f"Unable to check on package {package_version_name}")
        exit(1)

for cmd in run_script:
    print(cmd)
