# TerraformTraining
Iterative Terraform training module

| Version | Description                                                                                                                                                     |
| :-----: | :-------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|   v1    | Hard coded minimal virtual machine and network configuration                                                                                                    |
|   v2    | Hard coded configuration but taking into account resource (implicit) dependencies                                                                               |
|   v3    | Use of variables instead of hard coding                                                                                                                         |
|   v4    | Reorganise files to make them easier to digest                                                                                                                  |
|   v5    | Add a count argument to create multiple virtual machines, ensuring unique naming                                                                                |
|   v6    | Add an availability set resource                                                                                                                                |
|   v7    | Add VMs to the availability set, which must be done at creation                                                                                                 |
|   v8    | Convert to a module so it's possible to use a for_each to deploy to multiple regions for redundancy, in a single configuration                                  |
|   v9    | Use of locals to calculate values, and use of region as a unique identifier (for VMs & resource groups)                                                         |
|   v10   | Add a conditional expression to decide when to create an availability set                                                                                       |
|   v11   | Add a conditional expression to decide when to add VMs to an availability zone, and ensuring each VM is in a different zone (up to maximum available)           |
|   v12   | Modularise - Compute and Network now in separate modules                                                                                                        |
|   v13   | Add more networking resources, including route tables and network security groups (NSGs) - unintentionally introduce deadlock                                   |
|   v14   | Add explicit dependency for network (to remove deadlock of interfaces being associated at the same time as route tables/NSGs) & remove duplicate resource group |
|   v15   | Create module to add multiple data disks to multiple virtual machines (with support for regions with and without availability zones)                            |
|   v16   | Create module for Linux virtual machine                                                                                                                         |
|   v17   | Use function to calculate and set static IP address for virtual machines                                                                                        |
|   v18   | Create module for provisioning public IP address and public load balancer (with zonal support)                                                                  |
|   v19   | Add additional load balancer resources for backend configuration                                                                                                |
|   v20   | Add health probe and load balancer rule                                                                                                                         |
|   v21   | Addition of traffic manager for global load balancing (with traffic manager profile and endpoints)                                                              |
|   v22   | Addition of recovery services vault module, to allow for backups to be created                                                                                  |
|   v23   | Creation of backup policies for VMs in the recovery services vault                                                                                              |