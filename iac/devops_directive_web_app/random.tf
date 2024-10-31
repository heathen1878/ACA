resource "random_id" "this" {
  keepers = {
    proj = "ddd"
    key  = data.azurerm_client_config.current.subscription_id
  }

  byte_length = 4
}