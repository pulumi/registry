#!/usr/bin/env python3

# Can be run without dependencies with
#
#	uv run --with requests,pyyaml scripts/ci/push-registry.py

import os
import tempfile
import requests
import subprocess
import json
import yaml
import argparse
from logging import info, warning, basicConfig, WARNING

basicConfig(level=WARNING)

yaml_db_path = "themes/default/data/registry/packages"
docs_path = "themes/default/content/registry/packages"
backend_url = os.getenv("PULUMI_BACKEND_URL") or "https://api.pulumi.com/api"
tmpdir = tempfile.mkdtemp()

token = os.getenv("PULUMI_ACCESS_TOKEN")

parser = argparse.ArgumentParser(
    prog='push-registry',
    description="""Push the contents of the pulumi/registry GitHub repo to the live registry service

    PULUMI_ACCESS_TOKEN must be set to publish.
    PULUMI_BACKEND_URL may be set. Defaults to "https://api.pulumi.com/api".
    """,
    epilog='This is a Pulumi internal tool - it is not intended for external use')
parser.add_argument("--target", default=None, help='Only focus on one provider. This should just be the name, like "aws"')
parser.add_argument("--dry-run", action='store_true', help="Don't actually publish provider, but do everything else")

args = parser.parse_args()

info(f"Tempdir: {tmpdir}")

publishers = {
    "1Password": "1Password",
    "Aviatrix": "Aviatrix",
    "CTFer.io": "ctfer",
    "Checkly": "checkly",
    "Christian Nunciato": "christian_nunciato",
    "Chronosphere": "chronosphere",
    "Cisco": "cisco",
    "CrowdStrike": "crowdstrike",
    "Daniel Muehlbachler-Pietrzykowski": "daniel_muehlbachler_pietrzykowski",
    "DataRobot, Inc.": "datarobot",
    "Defang": "defang",
    "Descope": "descope",
    "Equinix": "equinix",
    "EventStore": "eventstore",
    "Genesiscloud": "genesiscloud",
    "Grapl Security": "grapl_security",
    "Ian Wahbe": "ian_wahbe",
    "Impart Security": "impart_security",
    "Koyeb": "koyeb",
    "Lee Zen": "lee_zen",
    "Nuage-Studio": "nuage_studio",
    "OVHcloud": "ovhcloud",
    "Paul Stack": "paul_stack",
    "Piers Karsenbarg": "piers_karsenbarg",
    "Prodvana": "prodvana",
    "Pulumi": "pulumi",
    "Pulumiverse": "pulumiverse",
    "RedisLabs": "redislabs",
    "Rootly": "rootly",
    "Runpod": "runpod",
    "Symbiosis": "symbiosis",
    "Theo Gravity": "theo_gravity",
    "Three141": "three141",
    "Threefold": "threefold",
    "Twingate": "twingate",
    "UpCloudLtd": "upcloudltd",
    "Upstash": "upstash",
    "Vates": "vates",
    "Volcengine": "volcengine",
    "Wttech": "wttech",
    "Zscaler": "zscaler",
    "airbytehq": "airbytehq",
    "akeyless-community": "akeyless-community",
    "aptible": "aptible",
    "athenz": "athenz",
    "azaurus1": "azaurus1",
    "castai": "castai",
    "celest-dev": "celest-dev",
    "chainguard-dev": "chainguard-dev",
    "checkpointsw": "checkpointsw",
    "ciscodevnet": "ciscodevnet",
    "ciscodevnet": "ciscodevnet",
    "civo": "civo",
    "cloudfoundry-community": "cloudfoundry-community",
    "coder": "coder",
    "constellix": "constellix",
    "coralogix": "coralogix",
    "cox-automotive": "cox-automotive",
    "cyralinc": "cyralinc",
    "datadrivers": "datadrivers",
    "dell": "dell",
    "dell": "dell",
    "dell": "dell",
    "denouche": "denouche",
    "dirien": "dirien",
    "dmacvicar": "dmacvicar",
    "dome9": "dome9",
    "drfaust92": "drfaust92",
    "e-breuninger": "e-breuninger",
    "edge-center": "edge-center",
    "elastic": "elastic",
    "elastic-infra": "elastic-infra",
    "ferlab-ste-justine": "ferlab-ste-justine",
    "ferlab-ste-justine": "ferlab-ste-justine",
    "fivetran": "fivetran",
    "flexibleenginecloud": "flexibleenginecloud",
    "fortinetdev": "fortinetdev",
    "g-core": "g-core",
    "glesys": "glesys",
    "goauthentik": "goauthentik",
    "hashicorp": "hashicorp",
    "honeycombio": "honeycomb",
    "hpe": "hpe",
    "ibm-cloud": "ibm-cloud",
    "imperva": "imperva",
    "incident-io": "incident_io",
    "infobloxopen": "infobloxopen",
    "ionos-cloud": "ionos-cloud",
    "iterative": "iterative",
    "jdamata": "jdamata",
    "jfrog": "jfrog",
    "joneshf": "joneshf",
    "juju": "juju",
    "k-yomo": "k-yomo",
    "kislerdm": "kislerdm",
    "komminarlabs": "komminarlabs",
    "kong": "kong",
    "labd": "labd",
    "lacework": "lacework",
    "lbrlabs": "lbrlabs",
    "littlejo": "littlejo",
    "logdna": "logdna",
    "logzio": "logzio",
    "lucky3028": "lucky3028",
    "mastercard": "mastercard",
    "maxlaverse": "maxlaverse",
    "megaport": "megaport",
    "mrolla": "mrolla",
    "nats-io": "nats-io",
    "netapp": "netapp",
    "netlify": "netlify",
    "octopusdeploylabs": "octopusdeploylabs",
    "onelogin": "onelogin",
    "opennebula": "opennebula",
    "opensearch-project": "opensearch-project",
    "opentelekomcloud": "opentelekomcloud",
    "oun": "oun",
    "outscale": "outscale",
    "paloaltonetworks": "paloaltonetworks",
    "paloaltonetworks": "paloaltonetworks",
    "pan-net": "pan-net",
    "paultyng": "paultyng",
    "pgEdge": "pgedge",
    "philips-software": "philips-software",
    "pinecone-io": "pinecone",
    "planetscale": "planetscale",
    "port-labs": "port_labs",
    "prefecthq": "prefecthq",
    "propelauth": "propelauth",
    "pulumiverse - Marcel Arns": "pulumiverse",
    "pulumiverse": "pulumiverse",
    "qdrant": "qdrant",
    "rancher": "rancher",
    "redpanda-data": "redpanda_data",
    "rollbar": "rollbar",
    "selectel": "selectel",
    "spectrocloud": "spectrocloud",
    "splightplatform": "splightplatform",
    "supabase": "supabase",
    "sysdiglabs": "sysdiglabs",
    "temporalio": "temporalio",
    "tencentcloudstack": "tencentcloudstack",
    "terraform-lxd": "terraform_lxd",
    "terraform-routeros": "terraform-routeros",
    "timescale": "timescale",
    "ucloud": "ucloud",
    "vantage-sh": "vantage-sh",
    "vk-cs": "vk-cs",
    "vmware": "vmware",
    "vmware": "vmware",
    "vmware": "vmware",
    "vmware": "vmware",
    "vmware": "vmware",
    "zenduty": "zenduty",
}

has_failed = False

for filename in os.listdir(yaml_db_path):
    if not filename.endswith('.yaml'):
        continue
    if args.target and args.target != filename.removesuffix(".yaml"):
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
        raise Exception(f'Missing publisher entry for "{publisher_display}"')
    publisher = publishers[publisher_display]


    package_version_name = f"{source}/{publisher}/{name}@{version.removeprefix("v")}"
    info(f"Preparing command for {package_version_name}")

    # We skip azure-native-v1, since it will (1) never be updated again and (2) isn't
    # actually a package, it's an alias to the azure native package.
    #
    # If azure-native-v2 ships (with azure-native now at v3), we will want to exclude
    # azure-native-v2 also.
    if name.startswith("azure-native") and name != "azure-native":
        continue


    api_url = f"{backend_url}/preview/registry/packages/{source}/{publisher}/{name}/versions/{version.removeprefix("v")}"
    existence_check = requests.get(api_url, headers={
        "Authorization": f"token {token}",
    }).status_code

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


            # We *correct* the version if it isn't present *or* if it disagrees with the
            # version that has been published in github.com/pulumi/registry.
            if "version" not in schema or schema["version"] != version.removeprefix("v"):
                schema["version"] = version.removeprefix("v")
            json.dump(schema, sf)


        cmd = [
            "pulumi", "package", "publish",
            schema_file,
            "--readme", f"{docs_path}/{name}/_index.md",
            "--source", source,
            "--publisher", publisher,
        ]

        if os.path.isfile(os.path.join(docs_path, name, "_installation-configuration.md")):
            cmd.extend(["--installation-configuration", f"{docs_path}/{name}/installation-configuration.md"])

        if args.dry_run:
            print(" ".join(cmd))
        else:
            try:
                subprocess.run(cmd, shell=False, check=True)
            except subprocess.CalledProcessError:
                warning(f"Failed to run {" ".join(cmd)}")
                has_failed=True

        os.remove(schema_file)
        os.removedirs(os.path.dirname(schema_file))
    elif existence_check == 200:
        info("Package already exists in the registry DB")
    else:
        warning(f"Unable to check on package {package_version_name}")
        has_failed=True

if has_failed:
    exit(1)
