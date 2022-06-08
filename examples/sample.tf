terraform {
  required_providers {
    inviulrestapi = {
      version = "0.1"
      source  = "inviul/inviulrestapi"
    }
  }
}

provider "inviulrestapi" {}

data "inviulrestapi" "myFirstRestCall" {
  base_uri = "https://username:password@dev.test.com?"
  path = "trailpath"
  http_rest_method = "PATCH"
  json_payload = "[{\"key\":\"value\"}]
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
