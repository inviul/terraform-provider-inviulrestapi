terraform {
  required_providers {
    inviulrestapi = {
      version = "0.1.6"
      source  = "inviul/inviulrestapi"
    }
  }
}

provider "inviulrestapi" {}

data "inviulrestapi" "myFirstRestCall" {
  base_uri = "https://username:password@dev.test.com?"
  path = "trailpath"
  http_rest_method = "PATCH"
  json_payload = "[{\"key\":\"value\"}]"
}



# Returns Output
output "myRestCallOutput" {
  value = data.inviulrestapi.myFirstRestCall.rest_out
}

