# AWS Assume Role

A simple utility to assume an AWS role and print the credentials to stdout. Statically linked, no dependencies.

Uses the default AWS credential resolution mechanisms, so should support anything supported by profiles (including SSO,
MFA etc.) as well as environment variables and instance roles.

# Usage

Use a profile:

```sh
export $(AWS_PROFILE='my-profile' aws-assume-role -role-arn arn:aws:iam::123456789:role/role/the-role)
```

Use environment variables:

```sh
export AWS_ACCESS_KEY_ID=...
export AWS_SECRET_ACCESS_KEY=...
export AWS_SESSION_TOKEN=...
export $(aws-assume-role -role-arn arn:aws:iam::123456789:role/role/the-role)
```

Use instance roles:

```sh
export $(aws-assume-role -role-arn arn:aws:iam::123456789:role/role/the-role)
```

Validate and test the assumed role by running a command like `aws sts get-caller-identity`:

```sh
aws sts get-caller-identity
{
    "UserId": "AROAYU3CAIBXDSONTDJIN:aws-assume-role",
    "Account": "123456789",
    "Arn": "arn:aws:sts::123456789:assumed-role/the-role/aws-assume-role"
}
```
