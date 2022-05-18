Kubernetes Go types that can be used with [TinyGo](tinygo.org/) to build WebAssembly
modules meant to be run outside of the browser.

> These files are generated automatically via
> [`k8s-objects-generator`](https://github.com/kubewarden/k8s-objects-generator).

## Marshal a Kubernetes object to JSON

Create a new Go project using Go modules:

```console
mkdir demo
cd demo
go mod init demo
```

**Important:** run the following command to ensure a TinyGo compatible version of
the `github.com/go-openapi/strfmt` module is being used:

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

## Unmarshal from JSON to a Kubernetes object

Create a new Go project using Go modules:

```console
mkdir demo
cd demo
go mod init demo
```

**Important:** run the following command to ensure a TinyGo compatible version of
the `github.com/go-openapi/strfmt` module is being used:

```console
go mod edit -replace github.com/go-openapi/strfmt=github.com/kubewarden/strfmt@v0.1.0
```

Now create a `main.go` with the following contents:

```go
package main

import (
	"fmt"
	"os"

	networkingv1 "github.com/kubewarden/k8s-objects/api/networking/v1"
	"github.com/mailru/easyjson"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: demo <path to file with json ingress>")
		os.Exit(1)
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	ingress := &networkingv1.Ingress{}
	if err := easyjson.Unmarshal(data, ingress); err != nil {
		fmt.Fprintf(os.Stderr, "cannot decode contents of file %s into an Ingress object: %v\n", filename, err)
		os.Exit(1)
	}

	ingress.Metadata.Name = "DEMO"

	b, err := ingress.MarshalJSON()
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("%s\n", string(b))
	}
}
```

This program loads an Ingress definition from a JSON file, changes its name to
be `DEMO` and then prints the object back to the standard output using
the JSON format.

Let's create a file named `ingress.json`, which contains the definition of a
simple Ingress object:

```json
{
  "apiVersion": "networking.k8s.io/v1",
  "kind": "Ingress",
  "metadata": {
    "name": "minimal-ingress",
    "annotations": {
      "nginx.ingress.kubernetes.io/rewrite-target": "/"
    }
  },
  "spec": {
    "ingressClassName": "nginx-example",
    "rules": [
      {
        "http": {
          "paths": [
            {
              "path": "/testpath",
              "pathType": "Prefix",
              "backend": {
                "service": {
                  "name": "test",
                  "port": {
                    "number": 80
                  }
                }
              }
            }
          ]
        }
      }
    ]
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

Now, you can run the code using [Wasmtime](https://wasmtime.dev/). This time
however we have to provide some extra runtime flags to allow the Wasm module
to read the `ingress.json` file we previously created. This is needed because
all the Wasm modules are run inside of a dedicated sandbox that, by default,
does not grant access to the host file system to the running code.

```console
wasmtime run \
  --mapdir /demo::`pwd` \
  demo.wasm /demo/ingress.json
```

The code maps the current directory on the host (`pwd`) into the guest
under the `/demo` path. Because of that, the guest code will find the
`ingress.json` file at the `/demo/ingress.json` location.

> **Tip:** the `mapdir` statement works exactly like sharing a Volume between
> the host system and a container.

If you pipe the output through something like the `jq` utility, you will
obtain some like that:

```json
{
  "apiVersion": "networking.k8s.io/v1",
  "kind": "Ingress",
  "metadata": {
    "annotations": {
      "nginx.ingress.kubernetes.io/rewrite-target": "/"
    },
    "creationTimestamp": "0001-01-01T00:00:00.000Z",
    "deletionTimestamp": "0001-01-01T00:00:00.000Z",
    "finalizers": null,
    "managedFields": null,
    "name": "DEMO",
    "ownerReferences": null
  },
  "spec": {
    "ingressClassName": "nginx-example",
    "rules": [
      {
        "http": {
          "paths": [
            {
              "backend": {
                "resource": {
                  "kind": null,
                  "name": null
                },
                "service": {
                  "name": "test",
                  "port": {
                    "number": 80
                  }
                }
              },
              "path": "/testpath",
              "pathType": "Prefix"
            }
          ]
        }
      }
    ],
    "tls": null
  }
}
```
