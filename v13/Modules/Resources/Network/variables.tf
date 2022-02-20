variable "resource_group_name" {
  description = "Name of resource group"
  type        = string
}

variable "resource_location" {
  description = "Location of resources"
  type        = string
}

variable "network_name" {
  description = "Name of network"
  type        = string
  default     = "WebServices-vnet"
}

variable "subnet_name" {
  description = "Name of subnet"
  type        = string
  default     = "WebServicesSubnet"
}

variable "network_address" {
  description = "Network address space"
  type        = string
}