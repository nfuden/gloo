package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/jsonpb"
	structpb "github.com/golang/protobuf/ptypes/struct"
	errors "github.com/rotisserie/eris"
	plugins "github.com/solo-io/ext-auth-service/pkg/config/plugin"
	extauth "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
)

var pluginManifestFlagUsage = `A .yaml file containing information required to load the ext auth plugins. Must have the following format: 

	name: MyPlugin
	pluginFileName: Plugin.so
	exportedSymbolName: MyPlugin
`

func main() {
	ctx := contextutils.WithLogger(context.Background(), "verify-plugins")
	logger := contextutils.LoggerFrom(ctx)

	// Get command line flags
	pluginDir := flag.String("pluginDir", "", "directory containing the compiled plugin files")
	pluginManifestFile := flag.String("manifest", "", pluginManifestFlagUsage)
	debug := flag.Bool("debug", false, "if present sets log level to debug")
	flag.Parse()

	if *pluginManifestFile == "" {
		fmt.Println("Missing required -f option")
		flag.Usage()
		os.Exit(1)
	}

	if *pluginDir == "" {
		fmt.Println("Missing required -pluginDir option")
		flag.Usage()
		os.Exit(1)
	}

	if *debug {
		contextutils.SetLogLevel(zap.DebugLevel)
	}

	if err := verifyPlugin(ctx, *pluginDir, *pluginManifestFile); err != nil {
		logger.Errorw("Plugin(s) cannot be loaded by Gloo", zap.Any("error", err))
		os.Exit(1)
	}

	logger.Info("Successfully verified that plugins can be loaded by Gloo!")
	os.Exit(0)
}

func verifyPlugin(ctx context.Context, pluginDir, pluginManifestFile string) error {
	pluginConfig, err := parseManifestFile(pluginManifestFile)
	if err != nil {
		return errors.Wrapf(err, "failed to parse plugin manifest file")
	}

	return loadPlugin(ctx, pluginDir, pluginConfig)
}

func parseManifestFile(filePath string) (*extauth.AuthPlugin, error) {
	readFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	byt, err := yaml.YAMLToJSON(readFile)
	if err != nil {
		return nil, err
	}

	into := &extauth.AuthPlugin{}
	if err = jsonpb.Unmarshal(bytes.NewBuffer(byt), into); err != nil {
		return nil, err
	}
	return into, nil
}

func loadPlugin(ctx context.Context, pluginDir string, pluginConfig *extauth.AuthPlugin) error {
	sanitizeConfig(pluginConfig)
	_, err := plugins.NewPluginLoader(pluginDir).LoadAuthPlugin(ctx, pluginConfig)
	return err
}

// Loader will fail if proto is nil
func sanitizeConfig(pluginConfig *extauth.AuthPlugin) {
	if pluginConfig.Config == nil {
		pluginConfig.Config = &structpb.Struct{
			Fields: map[string]*structpb.Value{},
		}
	}
}