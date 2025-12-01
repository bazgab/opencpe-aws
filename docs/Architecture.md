## Architecture


### Examples

Sample Usage

```sh
opencpe notify --config="configuration.json" --policy="instance-age" --region="us-east-1" 
```

Example Configuration file:

```json
{
    "authentication": {
        "aws_profile": "db-prod" 
    },
    "notification": {
        "smtp_host": "smtp.example.com",
        "smtp_port": "587",
        "email_from": "admin@example.com",
        "email_password":"your-secure-password"
    },
    "ignored_tags": {
        "owner": "admin",
        "project": "current-project-name"
    }
}
```
Important: The email will be sent to the resource owner's email, which should be tagged in the resource as "OwnerEmail" as the key and their respective email as the value


## Reference
| Key               |  Type          | Description                                                                                                                     |
|-------------------|:--------------:|---------------------------------------------------------------------------------------------------------------------------------|
| `"aws_profile"`   | String         | the aws profile information required to authenticate with IAM Identity Center. This is the profile section on `~/.aws/config`, it is recommended to set this file up with the AWS CLI tool (`aws sso config`), otherwise the authentication credentials will look for the `default` profile in `~/.aws/config`, the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` Environment Variables or a shared credentials file (`~/.aws/credentials`) in that order     |
| `"ignored_tags"`  | Object         | This is an object that takes in key-value pairs to ignore tags that have the corresponding value when filtering through resources. In the example above, OpenCPE will ignore all resources that have the `"owner"` tag value of `"admin"`, as well as ignore all resources that have the `"project"` tag value of `"current-project-name"` |
