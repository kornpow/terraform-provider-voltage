terraform {
  required_providers {
    voltage = {
      source = "registry.terraform.io/qustavo/voltage"
    }
  }
}

provider "voltage" {}

resource "voltage_node" "testnet" {
  network = "testnet"
  purchased_type = "ondemand"
  type = "lite"
  name = "qustavo1"
  settings = {
    autopilot = false
    grpc = true
    rest = true
    keysend = true
    whitelist = [""]
    alias = "qustavo1"
    color = "#000000"
  }
}

resource "voltage_dashboard" "thunderhub" {
  node_id = "39869a8f-b812-4581-af31-f5a25f5e843fN"
  # node_id = voltage_node.testnet.node_id
  type    = "thunderhub"
}

output node {
  value = voltage_node.testnet.node_id
}

output dashboard {
  value = voltage_dashboard.thunderhub.dashboard_id
}
