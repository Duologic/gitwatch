# gitwatch

Simple tool that can pulls in the latest of a git branch every minute.

The intend is to run this as a sidecar to provide configuration or data on the filesystem in a GitOps-like fashion.

## Usage:

```
Usage of gitwatch:
  -branch string
    	remote branch name to watch (default "master")
  -config string
    	path to the config file
  -dir string
    	remote branch name to watch (default "/repo")
  -repository string
    	git repository (http/https/ssh)
```

The `-config` is a json file for all arguments, making the other arguments optional. Using other arguments will override
the `-config`.

## Usage example with SSH:

```
docker run --rm -ti \
    -v ${HOME}/.ssh:/ssh \
    -e GIT_SSH_COMMAND='ssh -i /ssh/id_rsa' \
    duologic/gitwatch \
    --repository git@github.com:example/repo.git
```
