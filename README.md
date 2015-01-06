
github\_auto\_pull
==================

gitub\_auto\_pull is a small service listening for github push hook notifications, and updating the corresponding folder.

## Install

You need (go)[http://golang.org/] to be installed.

```bash
go get -u github.com/yazgazan/github_auto_pull
```

## Usage

Example :

```bash
github_auto_pull
```

```toml
# config.toml

Listen = ":8454"

# no default listen port

```

```toml
# repos.toml

# repo1_name should be the repository name on github
[[repo1_name]]
Folder = "./repo1_prod"
Remote = "http_origin" # the remote url should be an http url.
Branch = "master"

# when triggered, the following shell command will be executed : `sh -c "cd './repo1_prod'; git pull 'http_origin' 'master'"`

[[repo1_name]]
Folder = "./repo1_test"
Remote = "origin"
Branch = "test"

[[repo2_name]]
Folder = "./repo2"
Remote = "origin"
Branch = "master"

```

To configure your github repository, you need to add this service address to the webhooks in the github settings (settings > webhooks & Serices > Add webhook).
Be sure to leave the "Content type" field to application/json, the "Secret" field empty and the "Just the push event" checked.

