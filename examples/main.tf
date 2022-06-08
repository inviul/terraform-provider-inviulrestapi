terraform {
  required_providers {
    inviulrestapi = {
      version = "0.1"
      source  = "inviul/terraform/inviulrestapi"
    }
  }
}

provider "inviulrestapi" {}

data "inviulrestapi" "myFirstRestCall" {
  base_uri = "https://azure:32oejceu6jqqcfra5tfmgdelrlrfwaslmolnx27ig65emoeu4wkq@dev.azure.com/msci-otw/analytics-apps/_apis/pipelines/pipelinepermissions?"
  path = "api-version=7.1-preview.1"
  http_rest_method = "PATCH"
  json_payload = "[{\"resource\":{\"type\":\"endpoint\",\"id\":\"3536788c-a807-4a60-ab0a-cc1cd475fc19\",\"name\":\"Default\"},\"pipelines\":[{\"id\":13894,\"authorized\":true}]}]"
}

locals {
  json_data = jsondecode(data.inviulrestapi.myFirstRestCall.rest_out)
}

# Returns all Todos
output "myRestCallOutput" {
  value = local.json_data
}

//# Returns the title of todo
output "myRestCallTitle" {
  value = local.json_data.title
}
