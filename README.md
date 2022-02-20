# TerraformTraining
Iterative Terraform training module with automated CI/CD Pipeline deployment
## [Infrastructure as Code: CI/CD Pipeline][pipeline-summary]
| Version |                                                                         main                                                                          |
| :-----: | :---------------------------------------------------------------------------------------------------------------------------------------------------: |
|   v27   | [![Build Status](https://dev.azure.com/wesleytrust/TerraformTraining/_apis/build/status/ENV-P%3B%20TerraformTraining?branchName=main)][pipeline-link] |

## [Version history][github-link]

| Version | Description                                                                                                                                           |
| :-----: | :---------------------------------------------------------------------------------------------------------------------------------------------------- |
|   v1    | Hard coded minimal virtual machine and network configuration                                                                                          |
|   v2    | Hard coded configuration but taking into account resource (implicit) dependencies                                                                     |
|   v3    | Use of variables instead of hard coding                                                                                                               |
|   v4    | Reorganise files to make them easier to digest                                                                                                        |
|   v5    | Creation of random VM password for security                                                                                                           |
|   v6    | Add a count argument to create multiple virtual machines, ensuring unique naming                                                                      |
|   v7    | Add an availability set resource                                                                                                                      |
|   v8    | Add VMs to the availability set, which must be done at creation                                                                                       |
|   v9    | Convert to a module so it's possible to use a for_each to deploy to multiple regions for redundancy, in a single configuration                        |
|   v10   | Use of locals to calculate values, and use of region as a unique identifier (for VMs & resource groups)                                               |
|   v11   | Add a conditional expression to decide when to create an availability set                                                                             |
|   v13   | Add a conditional expression to decide when to add VMs to an availability zone, and ensuring each VM is in a different zone (up to maximum available) |
|   v13   | Modularise - Compute and Network now in separate modules                                                                                              |
|   v14   | Add more networking resources, including route tables and network security groups (NSGs) - introduce possible resource deadlock                       |
|   v15   | Add explicit dependency for network (to remove deadlock of interfaces being associated at the same time as route tables/NSGs)                         |
|   v16   | Create module to add multiple data disks to multiple virtual machines (with support for regions with and without availability zones)                  |
|   v17   | Create module for Linux virtual machine                                                                                                               |
|   v18   | Use function to calculate and set static IP address for virtual machines                                                                              |
|   v19   | Create module for provisioning public IP address and public load balancer (with zonal support)                                                        |
|   v20   | Add additional load balancer resources for backend configuration                                                                                      |
|   v21   | Add health probe and load balancer rule                                                                                                               |
|   v22   | Creation of Traffic Manager profile for global load balancing                                                                                         |
|   v23   | Addition of Traffic Manager endpoints for the public IP of each regional load balancer                                                                |
|   v24   | Creation of Key Vault module, and storage of VM passwords as encrypted secrets                                                                        |
|   v25   | Creation of Recovery Services Vault module, to allow for backups to be created                                                                        |
|   v26   | Addition of backup policies for VMs in the Recovery Services Vault                                                                                    |
|   v27   | Automate deployment in a CI/CD Pipeline with Azure DevOps using remote state                                                                          |

[pipeline-link]: https://dev.azure.com/wesleytrust/TerraformTraining/_build/latest?definitionId=49&branchName=main
[pipeline-summary]: https://dev.azure.com/wesleytrust/TerraformTraining
[github-link]: https://github.com/wesley-trust/TerraformTraining