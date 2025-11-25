## Architecture


### Examples

Sample Usage

```sh
opencpe --config="configuration.json" --policy="instance-age" --region="us-east-1" --action="notify"
```
Example Configuration file:

```json
{
    "authentication": {
        "aws_credentials_profile": "db-prod"
    },
    "ignored_tags": {
        "owner": "admin",
        "project": "current-project-name"
    }
}
```
