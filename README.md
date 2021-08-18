# gud-git

opinionated git commit message style checker

# usage

## as a standalone tool

```console
$ go get github.com/alxbse/gud-git
$ cd /my/git/repository
$ gud-git
```

## as a container

```console
$ podman run --volume /my/git/repository/:/repository:ro --workdir /repository -it ghcr.io/alxbse/gud-git
```

## as a tekton task

```yaml
kind: TaskRun
metadata:
  name: test-gud-git
spec:
  taskRef:
    name: gud-git
    bundle: ghcr.io/alxbse/gud-git-tekton-bundle
```
