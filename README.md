# 37037

Attempted reproduction for [renovate discussion 37037](https://github.com/renovatebot/renovate/discussions/37037)

## Current behavior

Renovate is failing to update dependencies from a golang module proxy behind Google's Artifact Registry (GAR)

## Expected behavior

Golang dependencies in a Google Artifact Registry should be updated.

## Link to the Renovate issue or Discussion

https://github.com/renovatebot/renovate/discussions/37037

## Steps to reproduce

1. Clone this repo and all submodules
2. In a separate tab/window/tmux/screen, run the local goproxy

```sh
cd goproxy
GOPROXY=direct GOMODCACHE=/tmp/goproxy-gomodcache go run ./cmd/goproxy/... server --address localhost:8080
```

3. Run renovate

```sh
LOG_LEVEL=debug GOPROXY=http://localhost:8080 renovate --platform local --enabled-managers gomod
```

4. Inspect logs
