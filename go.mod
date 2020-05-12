module github.com/cjyyb/helm-2to3

go 1.14

require (
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/cjyyb/helm-plugin-utils v0.0.0-20200512032140-45c7bc7909f2
	github.com/golang/protobuf v1.4.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	helm.sh/helm/v3 v3.2.0
	k8s.io/apimachinery v0.18.2
	k8s.io/helm v2.16.7+incompatible
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
)
