resource "aws_iam_role" "github_action_role" {
  name = "github_action_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      # ruleid: aws-oidc-role-policy-duplicate-condition
      {
        "Effect" : "Allow",
        "Principal": {
          "Federated": "arn:aws:iam::123456123456:oidc-provider/token.actions.githubusercontent.com"
        },
        "Action" : "sts:AssumeRoleWithWebIdentity",
        "Condition" : {
          "StringEquals" : {
            "token.actions.githubusercontent.com:sub" : [
              "repo:octo-org/octo-repo:ref:refs/heads/main"
            ]
          },
          "StringEquals" : {
            "token.actions.githubusercontent.com:aud": "sts.amazonaws.com"
          },
        }
      }
    ]
  })
}

resource "aws_iam_role" "github_action_role" {
  name = "github_action_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      # ruleid: aws-oidc-role-policy-duplicate-condition
      {
        "Effect" : "Allow",
        "Principal": {
          "Federated": "arn:aws:iam::123456123456:oidc-provider/token.actions.githubusercontent.com"
        },
        "Action" : "sts:AssumeRoleWithWebIdentity",
        "Condition" : {
          "StringEquals" : {
            "token.actions.githubusercontent.com:sub" : [
              "repo:octo-org/octo-repo:ref:refs/heads/main"
            ],
            "token.actions.githubusercontent.com:sub": "sts.amazonaws.com"
          },
        }
      }
    ]
  })
}

resource "aws_iam_role" "github_action_role" {
  name = "github_action_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      # ruleid: aws-oidc-role-policy-duplicate-condition
      {
        "Effect" : "Allow",
        "Principal": {
          "Federated": "arn:aws:iam::123456123456:oidc-provider/token.actions.githubusercontent.com"
        },
        "Action" : "sts:AssumeRoleWithWebIdentity",
        "Condition" : {
          "StringLike" : {
            "token.actions.githubusercontent.com:sub" : [
              "repo:octo-org/octo-repo:ref:refs/heads/main"
            ],
            "token.actions.githubusercontent.com:sub": "sts.amazonaws.com"
          },
        }
      }
    ]
  })
}

resource "aws_iam_role" "github_action_role" {
  name = "github_action_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      # ok: aws-oidc-role-policy-duplicate-condition
      {
        "Effect" : "Allow",
        "Principal": {
          "Federated": "arn:aws:iam::123456123456:oidc-provider/token.actions.githubusercontent.com"
        },
        "Action" : "sts:AssumeRoleWithWebIdentity",
        "Condition" : {
          "StringEquals" : {
            "token.actions.githubusercontent.com:sub" : [
              "repo:octo-org/octo-repo:ref:refs/heads/main"
            ],
            "token.actions.githubusercontent.com:aud": "sts.amazonaws.com"
          },
        }
      }
    ]
  })
}
