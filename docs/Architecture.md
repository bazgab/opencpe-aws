## Architecture


### Examples

Sample Usage

```sh
opencpe notify --config="config.json" --region="us-east-1" --policy="instance-age" 
```

Example Configuration file:

```json
{
    "authentication": {
        "aws_profile": "prod",
        "aws_account_id": 0123456789,
        "aws_account_name": "Development-team" 
    },
    "notification": {
        "smtp_endpoint": "email-smtp.us-east-1.amazonaws.com",
        "smtp_port": 587,
        "smtp_user": "AKIAXXXXXXXXXX",
        "smtp_password":"XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
        "sender_email": "alert@yourcompany.com"
    },
    "ignored_tags": {
        "owner": [ "admin"],
        "project": ["current-project-name", "second-project"]
    }
}

```

### Important Notes: 

- The email will be sent to the resource owner's email, which should be tagged in the resource as "Owner" as the key and their respective email as the value
- As of the time of writing this, OpenCPE only checks for instances that have the state of "running"
- Sending Email is only supported through Amazon SES for AWS


## Reference
| Key               |  Type          | Description                                                                                                                     |
|-------------------|:--------------:|---------------------------------------------------------------------------------------------------------------------------------|
| `"aws_profile"`   | String         | the aws profile information required to authenticate with IAM Identity Center. This is the profile section on `~/.aws/config`, it is recommended to set this file up with the AWS CLI tool (`aws sso config`), otherwise the authentication credentials will look for the `default` profile in `~/.aws/config`, the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` Environment Variables or a shared credentials file (`~/.aws/credentials`) in that order     |
| `"ignored_tags"`  | Object         | This is an object that takes in key-value pairs to ignore tags that have the corresponding value when filtering through resources. In the example above, OpenCPE will ignore all resources that have the `"owner"` tag value of `"admin"`, as well as ignore all resources that have the `"project"` tag value of `"current-project-name"` |
