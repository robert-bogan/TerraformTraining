# Create network dependencies
module "vm_network" {
  source              = "../Resources/Network"
  network_name        = var.network_name
  resource_location   = module.vm_group.location
  resource_group_name = module.vm_group.name
  network_address     = var.network_address
}
