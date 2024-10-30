resource "random_id" "this" {
  keepers = {
    proj = "fbc"
    key = data.azurerm_client_config.current.subscription_id
  }

  byte_length = 4
}