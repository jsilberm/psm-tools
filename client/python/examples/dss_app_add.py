import click
import os
import pensando_dss
import pensando_dss.psm
import yaml
from pensando_dss.psm.api import security_v1_api
from pensando_dss.psm.models.security import *
from pprint import pprint
from dss_common import *

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = pensando_dss.psm.Configuration(
    psm_config_path = os.environ["HOME"] + "/.psm/config.json",
    interactive_mode = True
)
configuration.verify_ssl = False

@click.command()
@click.option('--input_file', help='yaml_file for input fields to add vpc/vrf object', prompt=True)

def create_app(input_file):
    if not input_file:
        print("\n**input yaml file is missing**\n")
        os.system("python3 dss_app_create.py --help")
        exit()
    with open(input_file, 'r') as fp:
        input = yaml.safe_load(fp)
        validate_input_file(input, 'App')
    # Enter a context with an instance of the API client
    with pensando_dss.psm.ApiClient(configuration) as api_client:
        # Create an instance of the API class
        api_instance = security_v1_api.SecurityV1Api(api_client)
        # example passing only required values which don't have defaults set
        try:
            for item in input['App']:
                body = get_app_body(item)
                api_response = api_instance.add_app1(body)
                pprint(api_response.to_dict())

        except pensando_dss.psm.ApiException as e:
            print("Exception when calling SecurityV1Api->add_app1: %s\n" % e)

if __name__ == '__main__':
    create_app()
