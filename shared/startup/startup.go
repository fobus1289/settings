package startup

import (
	_ "embed"
	b64 "encoding/base64"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//go:embed example.docker-compose.yml
var dockerExample string

//go:embed env.example
var envExample string

const (
	DOCKER_NAME = "example.docker-compose.yml"
	ENV_NAME    = ".env.example"
)

func init() {
	serviceName, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	main(serviceName)
}

func Gen(servicePath string) {
	main(servicePath)
}

func main(servicePath string) {
	// %secret%
	// %service_name%

	createDockeFile(servicePath)

	createEnvFile(servicePath)
}

func createEnvFile(servicePath string) {
	serviceName := filepath.Base(servicePath)

	d1 := strings.Replace(envExample, "%service_name%", serviceName, 1)

	rndStr := b64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(rand.Int63(), 10)))

	d2 := strings.Replace(d1, "%secret%", rndStr, 1)

	if err := os.WriteFile(fmt.Sprintf("%s/%s", servicePath, ".env"), []byte(d2), fs.ModePerm); err != nil {
		panic(err)
	}
}

func createDockeFile(servicePath string) {
	serviceName := filepath.Base(servicePath)

	d1 := strings.Replace(dockerExample, "%service_name%", serviceName, 1)

	if err := os.WriteFile(fmt.Sprintf("%s/%s", servicePath, "dev.docker-compose.yml"), []byte(d1), fs.ModePerm); err != nil {
		panic(err)
	}
}
