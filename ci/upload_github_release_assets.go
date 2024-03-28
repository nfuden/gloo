package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/version"
	"github.com/solo-io/go-utils/githubutils"
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/go-utils/versionutils"
)

func main() {
	dryRun := false
	if len(os.Args) > 1 {
		var err error
		dryRun, err = strconv.ParseBool(os.Args[1])
		if err != nil {
			log.Fatalf("Unable to parse `dry_run` boolean argument, was provided %v", os.Args[1])
		}
	}
	const buildDir = "_output"
	const repoOwner = "solo-io"
	const repoName = "gloo"

	validateReleaseVersionOfCli(dryRun)

	assets := []githubutils.ReleaseAssetSpec{
		{
			Name:       "glooctl-linux-amd64",
			ParentPath: buildDir,
			UploadSHA:  true,
		},
		{
			Name:       "glooctl-linux-arm64",
			ParentPath: buildDir,
			UploadSHA:  true,
		},
		{
			Name:       "glooctl-darwin-amd64",
			ParentPath: buildDir,
			UploadSHA:  true,
		},
		{
			Name:       "glooctl-darwin-arm64",
			ParentPath: buildDir,
			UploadSHA:  true,
		},
		{
			Name:       "glooctl-windows-amd64.exe",
			ParentPath: buildDir,
			UploadSHA:  true,
		},
	}

	spec := githubutils.UploadReleaseAssetSpec{
		Owner:             repoOwner,
		Repo:              repoName,
		Assets:            assets,
		SkipAlreadyExists: true,
	}
	if !dryRun {
		githubutils.UploadReleaseAssetCli(&spec)
	}
}

func validateReleaseVersionOfCli(dryRun bool) {
	releaseVersion := getReleaseVersionOrExitGracefully(dryRun).String()[1:]
	name := fmt.Sprintf("_output/glooctl-%s-amd64", runtime.GOOS)

	cmd := exec.Command(name, "version")
	bytes, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error while trying to validate artifact version. Error was: %s", err.Error())
	}

	var ver version.Version
	if err := json.Unmarshal(bytes, &ver); err != nil {
		log.Fatalf("Unexpected version output for glooctl: %s", string(bytes))
	}

	if dryRun {
		ver.Client.Version, err = validVersionFromPrVersion(ver.Client.Version)
		if err != nil {
			log.Fatalf("unable to parse valid semver from PR version: %v", err)
		}
	}

	if ver.Client.Version != releaseVersion {
		log.Fatalf("Expected to release artifacts for version %s, glooctl binary reported version %s", releaseVersion, ver.Client.Version)
	}
}

// stolen from "github.com/solo-io/go-utils/versionutils", but changed the hardcoding of "TAGGED_VERSION" to "VERSION"
func getReleaseVersionOrExitGracefully(dryRun bool) *versionutils.Version {
	versionStr, present := os.LookupEnv("VERSION")
	if !present || versionStr == "" {
		fmt.Printf("VERSION not found in environment.\n")
		os.Exit(1)
	}

	if dryRun {
		var err error
		versionStr, err = validVersionFromPrVersion(versionStr)
		if err != nil {
			log.Fatalf("unable to parse valid semver from PR version: %v", err)
		}
	}
	tag := "v" + versionStr

	version, err := versionutils.ParseVersion(tag)
	if err != nil {
		fmt.Printf("VERSION %s is not a valid semver version.\n", tag)
		os.Exit(1)
	}
	return version
}

// we have to force the version into semver in order to test this. We do this by chopping the PR designation off
// when running as dry_run
func validVersionFromPrVersion(ver string) (string, error) {
	expectedLen := 2
	if strings.Contains(ver, "beta") || strings.Contains(ver, "rc") {
		expectedLen = 3
	}
	splitVer := strings.Split(ver, "-")
	if len(splitVer) != expectedLen {
		return "", errors.New("invalid format for PR version; expected v<MAJOR>.<MINOR>.<PATCH>-(beta|rc)<N>-<PR> for beta or rc and v<MAJOR>.<MINOR>.<PATCH>-<PR> for LTS")
	}
	return strings.Join(splitVer[:expectedLen-1], ""), nil

}
