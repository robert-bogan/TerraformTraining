# TerraformTraining
Iterative Terraform training module

| Version | Description                                                                                                                                                     |
| :-----: | :-------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|   v1    | Hard coded configuration                                                                                                                                        |
|   v2    | Hard coded configuration but taking into account resource (implicit) dependencies                                                                               |
|   v3    | Use of variables instead of hard coding                                                                                                                         |
|   v4    | Reorganise files to make them easier to digest                                                                                                                  |
|   v5    | Add a count argument to create multiple virtual machines                                                                                                        |
|   v6    | Add an availability set resource                                                                                                                                |
|   v7    | Add VMs to the availability set                                                                                                                                 |
|   v8    | Convert to a module to use a for_each to deploy to multiple region                                                                                              |
|   v9    | Use of locals to calculate values, switch to region as unique identifier (VM & resource group)                                                                  |
|   v10   | Add a conditional expression to decide when to create an availability set                                                                                       |
|   v11   | Add a conditional expression to decide when to add to an availability zone                                                                                      |
|   v12   | Modularise - Compute and Network now in separate modules                                                                                                        |
|   v13   | Add more network resources - unintentionally introduce deadlock                                                                                                 |
|   v14   | Add explicit dependency for network (to remove deadlock of interfaces being associated at the same time as route tables/NSGs) & remove duplicate resource group |