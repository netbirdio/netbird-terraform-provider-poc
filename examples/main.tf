terraform {
  required_providers {
    netbird = {
        source = "github.com/netbirdio/netbird"
        version = "0.1"
    }
  }
}



provider "netbird" {
  server_url="http://localhost:33073"
  token_auth="nbp_vAVubk8iRxdLmhqd7NDpqj3IV0LQ6t1LcYT0"
}

resource "netbird_setup_key" "tf_test_key_2" {
  name = "tf_linux_key_2"
  type = "one-off"
  auto_groups = ["cok2pknprlfhjmlrc9tg"]
  ephemeral = true
  revoked = false
  usage_limit = 1
  expires_in = 86400
}