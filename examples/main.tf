terraform {
  required_providers {
    voltage = {
      source = "registry.terraform.io/qustavo/voltage"
    }
  }
}

provider "voltage" {}

# resource "voltage_node" "testnet" {
#   network = "testnet"
#   purchased_type = "ondemand"
#   type = "lite"
#   name = "qustavo"
#   settings = {
#     autopilot = false
#     grpc = true
#     rest = true
#     keysend = true
#     whitelist = [""]
#     alias = "qustavo"
#     color = "#000000"
#   }
# }

resource "voltage_dashboard" "thunderhub" {
  node_id = "39869a8f-b812-4581-af31-f5a25f5e843fN"
  type    = "thunderhub"
}