output "nginx_endpoint" {
  value = azurerm_container_app.nginx.ingress[0].fqdn
}