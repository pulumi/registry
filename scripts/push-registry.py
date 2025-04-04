#!/usr/bin/env python3

# Can be run without dependencies with
#
#	uv run --with requests,pyyaml scripts/push-registry.py

import os
import tempfile
import requests
import json
import yaml
import sys
from logging import info, basicConfig, INFO

yaml_db_path = "themes/default/data/registry/packages"
docs_path = "themes/default/content/registry/packages"
tmpdir = tempfile.mkdtemp()

basicConfig(level=INFO)

info(f"Tempdir: {tmpdir}")

run_script = []

target = None
if len(sys.argv) >= 2:
    target = sys.argv[1]

publishers = {
    "1Password": "1Password",
    "airbytehq": "airbytehq",
    "Aviatrix": "Aviatrix",
    "azaurus1": "azaurus1",
    "castai": "castai",
    "celest-dev": "celest-dev",
    "Checkly": "checkly",
    "Christian Nunciato": "christian_nunciato",
    "Chronosphere": "chronosphere",
    "Cisco": "cisco",
    "civo": "civo",
    "constellix": "constellix",
    "CrowdStrike": "crowdstrike",
    "CTFer.io": "ctfer",
    "Daniel Muehlbachler-Pietrzykowski": "daniel_muehlbachler_pietrzykowski",
    "DataRobot, Inc.": "datarobot",
    "Defang": "defang",
    "Descope": "descope",
    "dirien": "dirien",
    "dmacvicar": "dmacvicar",
    "Equinix": "equinix",
    "EventStore": "eventstore",
    "fivetran": "fivetran",
    "fortinetdev": "fortinetdev",
    "Genesiscloud": "genesiscloud",
    "goauthentik": "goauthentik",
    "Grapl Security": "grapl_security",
    "hashicorp": "hashicorp",
    "honeycombio": "honeycomb",
    "Ian Wahbe": "ian_wahbe",
    "Impart Security": "impart_security",
    "incident-io": "incident_io",
    "jfrog": "jfrog",
    "joneshf": "joneshf",
    "komminarlabs": "komminarlabs",
    "kong": "kong",
    "Koyeb": "koyeb",
    "labd": "labd",
    "lbrlabs": "lbrlabs",
    "Lee Zen": "lee_zen",
    "littlejo": "littlejo",
    "lucky3028": "lucky3028",
    "netlify": "netlify",
    "Nuage-Studio": "nuage_studio",
    "onelogin": "onelogin",
    "oun": "oun",
    "outscale": "outscale",
    "OVHcloud": "ovhcloud",
    "Paul Stack": "paul_stack",
    "pgEdge": "pgedge",
    "Piers Karsenbarg": "piers_karsenbarg",
    "pinecone-io": "pinecone",
    "planetscale": "planetscale",
    "port-labs": "port_labs",
    "prefecthq": "prefecthq",
    "Prodvana": "prodvana",
    "propelauth": "propelauth",
    "Pulumi": "pulumi",
    "pulumiverse - Marcel Arns": "pulumiverse",
    "pulumiverse": "pulumiverse",
    "Pulumiverse": "pulumiverse",
    "qdrant": "qdrant",
    "rancher": "rancher",
    "RedisLabs": "redislabs",
    "redpanda-data": "redpanda_data",
    "Rootly": "rootly",
    "Runpod": "runpod",
    "splightplatform": "splightplatform",
    "supabase": "supabase",
    "Symbiosis": "symbiosis",
    "temporalio": "temporalio",
    "terraform-lxd": "terraform_lxd",
    "Theo Gravity": "theo_gravity",
    "Three141": "three141",
    "Threefold": "threefold",
    "timescale": "timescale",
    "Twingate": "twingate",
    "UpCloudLtd": "upcloudltd",
    "Upstash": "upstash",
    "vantage-sh": "vantage-sh",
    "Vates": "vates",
    "Volcengine": "volcengine",
    "Wttech": "wttech",
    "zenduty": "zenduty",
    "Zscaler": "zscaler",
}

for filename in os.listdir(yaml_db_path):
    if not filename.endswith('.yaml'):
        continue
    if target and target != filename.removesuffix(".yaml"):
        continue

    filepath = os.path.join(yaml_db_path, filename)
    with open(filepath) as file:
        data = yaml.safe_load(file)

    version = data.get('version')

    schema_file_path = data.get('schema_file_path')
    schema_url = data.get(
        'schema_file_url',
        f"https://raw.githubusercontent.com/{data.get('repo_url').removeprefix("https://github.com/")}/refs/tags/{version}/{schema_file_path}",
    )
    
    source = 'pulumi' if 'registry.opentofu.org' not in schema_url else "opentofu"
    name = data.get('name')
    publisher_display = data.get('publisher')


    if publisher_display == "DEPRECATED":
        continue
    elif publisher_display not in publishers:
        raise Exception(f"Missing publisher entry for {publisher_display}")
    publisher = publishers[publisher_display]


    package_version_name = f"{source}/{publisher}/{name}@{version}"
    info(f"Preparing command for {package_version_name}")

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

        with open(schema_file, 'w') as sf:
            response = requests.get(schema_url)
            response.raise_for_status()

            schema = None
            if (schema_file_path is not None and schema_file_path.endswith(".yaml")) or \
               (schema_url is not None and schema_url.endswith(".yaml")):
                schema = yaml.safe_load(response.content)
            else:
                schema = json.loads(response.content)

            if "version" not in schema:
                schema["version"] = version.removeprefix("v")
            json.dump(schema, sf)

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
        info("Package already exists in the registry DB")
    else:
        info(f"Unable to check on package {package_version_name}")
        exit(1)

for cmd in run_script:
    print(cmd)
