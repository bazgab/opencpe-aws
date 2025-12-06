## Policy Reference

OpenCPE aims to be the easiest and simplest way to monitor cloud resources and notify users, so unlike other solutions, policies come pre-defined at the software-level, requiring users to only pick which policy to apply to a certain instance type.

There is an understanding that this approach might feel limiting, however the policy files are ever-evolving and with the help of the community (in always suggesting new policy ideas, or contributing with their own policy code) we believe we can have a comprehensive set of policies that will satisfy most users' needs.

Another way to handle policies is as well to use the `ignore_tags` feature of the configuration to achieve results specific to certain resources.  


| Policy                    |  Resource  | Description                                                                                                                                               |
|---------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `"instance-age-2-days"`   | ec2        | Notifies resource owners by email once their instance has had the state of `running` for over 2 days.         |
| `"instance-age-7-days"`   | ec2        | Notifies resource owners by email once their instance has had the state of `running` for over 7 days.         |
