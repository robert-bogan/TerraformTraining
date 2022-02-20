# Provision public IP for load balancer
resource "azurerm_public_ip" "public_ip" {
  name                = "${var.vm_name}-ip"
  location            = var.resource_location
  resource_group_name = var.resource_group_name
  allocation_method   = "Static"
  sku                 = "Standard"
  availability_zone   = var.location_zone_support == true ? "Zone-Redundant" : "No-Zone"
}

# Create load balancer
resource "azurerm_lb" "load_balancer" {
  name                = "${var.vm_name}-lb"
  location            = var.resource_location
  resource_group_name = var.resource_group_name
  sku                 = "Standard"

  frontend_ip_configuration {
    name                 = azurerm_public_ip.public_ip.name
    public_ip_address_id = azurerm_public_ip.public_ip.id
  }
}

# Provision backend pool for load balancer
resource "azurerm_lb_backend_address_pool" "backend_address_pool" {
  loadbalancer_id = azurerm_lb.load_balancer.id
  name            = var.vm_name
}

# Provision addresses within backend pool
resource "azurerm_lb_backend_address_pool_address" "backend_address_pool_address" {
  count                   = var.vm_instance_count
  name                    = "${var.vm_name}${count.index + 1}"
  backend_address_pool_id = azurerm_lb_backend_address_pool.backend_address_pool.id
  virtual_network_id      = var.vm_network_id
  ip_address              = cidrhost(var.network_address, 4 + count.index)
}