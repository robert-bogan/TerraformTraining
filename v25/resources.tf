/* module "vm_deployment" {
  for_each           = toset(var.resource_location)
  source             = "./Modules/Deployments/Linux_virtual_machine"
  vm_name            = var.vm_name
  vm_instance_count  = var.vm_instance_count
  vm_size            = var.vm_size
  vm_data_disk_count = var.vm_data_disk_count
  vm_data_disk_size  = var.vm_data_disk_size
  resource_location  = each.value
  network_address    = lookup(var.network_address, each.value, null)
}

module "vm_deployment_traffic_manager" {
  source                             = "./Modules/Deployments/Traffic_manager"
  vm_name                            = var.vm_name
  traffic_manager_name               = "${var.vm_name}-TM"
  traffic_manager_location           = var.traffic_manager_location
  traffic_manager_endpoints          = module.vm_deployment
  traffic_manager_endpoint_locations = var.resource_location
}

module "vm_deployment_recovery_services_vault" {
  for_each                           = toset(var.recovery_services_location)
  source                             = "./Modules/Deployments/Recovery_services_vault"
  vm_name                            = var.vm_name
  recovery_services_name             = "${var.vm_name}-RSV"
  recovery_services_location         = each.value
  recovery_services_instance_count   = var.recovery_services_instance_count
  recovery_services_virtual_machines = module.vm_deployment[each.value]
} */