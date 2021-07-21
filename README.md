<p align="center"><img src="logo.svg" align="center" width="100"/></p>
<h1 align="center">Filestack Go</h1>

[![Go Report Card](https://goreportcard.com/badge/github.com/filestack/filestack-go?style=flat-square)](https://goreportcard.com/report/github.com/filestack/filestack-go)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.16-61CFDD.svg?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/filestack/filestack-go)](https://pkg.go.dev/mod/github.com/filestack/filestack-go)

This is the official Go SDK for Filestack - API and content management system that makes it easy
to add powerful file uploading and transformation capabilities to any web or mobile application.

## Resources

To learn more about this SDK, please visit our API Reference

## Install

Install the package with:

```bash
go get github.com/filestack/filestack-go
```

Import it with:

```go
import "github.com/filestack/filestack-go"
```

## Quickstart

The Filestack SDK allows you to upload files, perform transformations and handle filelinks.

## The Client
For most features provided by this SDK you will need a `Client` instance. It can be initialized with the following statement:
```go
client.NewClient(yourApiKey)
```
The most basic use case:
```go
package main

import (
	"fmt"
	"log"

	"github.com/filestack/filestack-go/client"
)

func main() {
	cli, err := client.NewClient(yourApiKey)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	fileLink := cli.MustNewFileLink(yourHandle)
	fmt.Print(fileLink.AsString())
}
```

Often you will need to adjust settings to your needs. `NewClient` constructor accepts optional arguments. For example if your account has enabled security, you will need to add an optional argument like here: 

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/filestack/filestack-go/client"
	clientOptions "github.com/filestack/filestack-go/options/client"
	"github.com/filestack/filestack-go/security"
)

func main() {
	sec := security.NewSecurity(YourSecret, &security.Policy{
		Expiry: time.Now().Add(time.Duration(24 * time.Hour)).Unix(), // the only required parameter
	})
	cli, err := client.NewClient(YourApiKey, clientOptions.SecurityPolicy(sec))
	if err != nil {
		log.Fatalf("failed to initialize the client: %v", err)
	}

	fileLink := cli.MustNewFileLink(YourHandle)
	fmt.Print(fileLink.AsString())
}
```

#### Overwrite the default http client
API requests are handled by `RequestHandler` service. It uses an instance of `http.Client` by default.\
To fulfill your requirements, a custom client can be injected using options:
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/filestack/filestack-go/client"
	clientOptions "github.com/filestack/filestack-go/options/client"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Minute,
	}

	cli, err := client.NewClient(yourApiKey, clientOptions.HTTPClient(httpClient))
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	fileLink := cli.MustNewFileLink(yourHandle)
	fmt.Print(fileLink.AsString())
}
```

### Uploads

One of the basic features of this SDK are uploads.\
The following example shows the easiest way to upload a file.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/filestack/filestack-go/client"
)

func main() {
	file, err := os.Open(yourFilePath)
	if err != nil {
		log.Fatal("cannot read the file")
	}
	defer file.Close()

	cli, err := client.NewClient(yourApiKey)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	fileLink, err := cli.Upload(context.Background(), file)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(fileLink.AsString())
}
```

#### Upload URL
When you are dealing with external files, `UploadURL` method can be used:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/filestack/filestack-go/client"
)

func main() {
	cli, err := client.NewClient(yourApiKey)
	if err != nil {
		log.Fatalf("failed to initialize the client: %v", err)
	}

	fileLink, err := cli.UploadURL(context.Background(), yourURL)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(fileLink.AsString())
}
```

### Transformations

You can perform a number of transformations on your files. The source files for transformations are called <i>resources</i> and they can be one of:
* external url
* handle
* storage alias

Most of the transformations work in a context of a single resource, but there are exceptions where more than one file reference is required.

### Resources

For a handle string, you can create a new transformation with:
```go
cli.MustNewTransformation(resource.NewHandle(yourHandle))
```
For an external URL use:
```go
cli.MustNewTransformation(resource.NewExternalURL(yourExternalURL))
```
And for storage aliases use:
```go
cli.MustNewTransformation(resource.NewStorageAlias(yourAlias, yourPath))
```

### Performing transformations

```go
package main

import (
	"fmt"

	"github.com/filestack/filestack-go/client"
	"github.com/filestack/filestack-go/resource"
)

func main() {
	cli := client.NewClient(yourApiKey)
	transformation := cli.MustNewTransformation(resource.NewExternalURL("https://www.google.com/image.jpg"))
	transformation.Flip() // or any other image transformation

	fmt.Println(transformation.BuildURL())
}
```

Transformations can be chained, since the API provides possibility to apply multiple tasks at once. For example:
```go
transformation.Flip().Negative() // more than one transformation task at once
```

Transformation tasks can have optional parameters. For example `Resize` needs <i>width</i>, <i>height</i> or both dimensions together. Therefore majority of transformation methods expect `Args` as a parameter. This dedicated type wraps all parameters and provides setters for ease of use. Each method has its corresponding `Args` type and can be initialized with a constructor, as in the following example:

```go
transformation.Upscale(args.NewUpscaleArgs().SetNoise("low"))

transformation.BlackWhite(args.NewBlackWhiteArgs().SetThreshold(5))
```

### Workflows
Workflows allow you to perform a sequence of tasks and include a conditional logic.\
They can be defined on your account. You refer to a workflow by its unique identifier.\
More information about the concept of workflows can be found here: https://www.filestack.com/docs/workflows/overview/

The following example shows how to schedule a job:
```go
externalURL := resource.NewExternalURL(yourImageURL)
workflow := client.NewWorkflow()
response, err := workflow.Run(context.Background(), yourWorkflowsId, externalURL)
if err != nil {
    // handle error
}
fmt.Printf("job id = %s", response.JobID)
```

The `workflow` type provides `Run` and `Status` methods. Workflows processing is asynchronous, therefore you don't get an instant result upong running it. You get a `JobID` instead and based on that value, you can check the processing status:
```go
response, err := workflow.Status(context.Background(), workflowJobID)
if err != nil {
    // handle error
}
fmt.Println(response.Results)
```