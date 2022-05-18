Kubernetes Go types that can be used with [TinyGo](tinygo.org/) to build WebAssembly
modules meant to be run outside of the browser.

> These files are generated automatically via
> [`k8s-objects-generator`](https://github.com/kubewarden/k8s-objects-generator).

## Usage

Create a new Go project using Go modules:

```console
mkdir demo
cd demo
go mod init demo
```

Run the following command to ensure a TinyGo compatible version of
the `github.com/go-openapi/strfmt` module:

```console
go mod edit -replace github.com/go-openapi/strfmt=github.com/kubewarden/strfmt@v0.1.0
```

Now create a `main.go` with the following contents:

```go
package main

import (
	"fmt"

	corev1 "github.com/kubewarden/k8s-objects/api/core/v1"
	metav1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

func main() {
	containerName := "nginx"
	var httpPort int32 = 80

	pod := corev1.Pod{
		Metadata: metav1.ObjectMeta{
			Name: "nginx",
		},
		Spec: &corev1.PodSpec{
			Containers: []*corev1.Container{
				{
					Name:  &containerName,
					Image: "nginx:latest",
					Ports: []*corev1.ContainerPort{
						{
							ContainerPort: &httpPort,
						},
					},
				},
			},
		},
	}

	b, err := pod.MarshalJSON()
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("%s\n", string(b))
	}
}
```

Next, run `go mod tidy` and, if you want, vendor all the dependencies:

```console
go mod tidy
go mod vendor
```

Finally, build the code using TinyGo to a WebAssembly module targeting the
[WASI interface](https://wasi.dev/):

```console
docker run --rm -v `pwd`:/src -w /src tinygo/tinygo:0.23.0 tinygo build -o demo.wasm -target=wasi -no-debug .
```

Now, you can run the code using [Wasmtime](https://wasmtime.dev/):

```console
wasmtime run demo.wasm
```

If you pipe the output through something like the `jq` utility, you will
obtain some like that:

```json
{
  "metadata": {
    "creationTimestamp": "0001-01-01T00:00:00.000Z",
    "deletionTimestamp": "0001-01-01T00:00:00.000Z",
    "finalizers": null,
    "managedFields": null,
    "name": "nginx",
    "ownerReferences": null
  },
  "spec": {
    "containers": [
      {
        "args": null,
        "command": null,
        "env": null,
        "envFrom": null,
        "image": "nginx:latest",
        "name": "nginx",
        "ports": [
          {
            "containerPort": 80
          }
        ],
        "volumeDevices": null,
        "volumeMounts": null
      }
    ],
    "ephemeralContainers": null,
    "hostAliases": null,
    "imagePullSecrets": null,
    "initContainers": null,
    "readinessGates": null,
    "tolerations": null,
    "topologySpreadConstraints": null,
    "volumes": null
  }
}
```
