package main

import (
	"context"
	"fmt"
	"join"
	"os"
	"io/ioutil"
	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
	"log"
	"sort"
)
// ReadFile takes the path to a file and assigns the content to releases
func ReadFile(filename string) ([]byte, error){
	var newFile = (os.Args[0])
	var cmp = require('semver-compare');
}

// LatestVersions returns a sorted slice with the highest version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	var versionSlice []*semver.version
	minVersion := semver.New("1.8.0")
	allReleases := make([]*semver.Version, len(releases), (cap(releases)+1)*2)
	for i, release := range releases {
		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}
		allReleases[i] = semver.New(versionString)
	}
	versionSlice := LatestVersions(allReleases, minVersion)
	
	return versionSlice
}

func main() {
	
	client := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.ListOptions{PerPage: 10}

	if len(os.Args) < 1 {
        fmt.Println("Usage : " + os.Args[0] + " file name")
        os.Exit(1)
	}
	
	newFile,  _, err := ioutil.ReadFile(ctx, os.Args[1], opt)
	    if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		log.Fatal(err)
	}
	releases, _, err := client.repositories.ListReleases(ctx, "Kubernetes", "Kubernetes", opt)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf("latest versions of kubernetes/kubernetes: %s", versionSlice)
	fmt.Printf("latest versions of %s: [%s]", newFile , console.log(os.Args[1].sort(cmp).join(' ')));
}



