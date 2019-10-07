How to use KO

1. Install `GO111MODULE=on go get github.com/google/ko/cmd/ko`
2. Set env variable:

export KO_DOCKER_REPO=docker.io/rinormaloku

3. ko apply -f deploy.yaml

(Ko will notice that this is within the path github.com/rinormaloku/hello for that reason it will build and push the container and replace the line: `image: github.com/rinormaloku/hello` and apply to the cluster for rapid development)
