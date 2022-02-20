# Required resource variables
variable "resource_group_name" {
  description = "Name of resource group"
  type        = string
  default     = "WebServices"
}

variable "traffic_manager_name" {
  description = "Traffic manager name"
  type        = string
}

variable "traffic_manager_location" {
  description = "Traffic manager location"
  type        = string
}

variable "vm_environment" {
  description = "Environment of vm"
  type        = string
  default     = "Production"
}

variable "vm_name" {
  description = "Name of vm"
  type        = string
}