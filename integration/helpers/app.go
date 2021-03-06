package helpers

import (
	"archive/zip"
	"fmt"
	"github.com/onsi/gomega/gbytes"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"gopkg.in/yaml.v2"
)

// WithHelloWorldApp creates a simple application to use with your CLI command
// (typically CF Push). When pushing, be aware of specifying '-b
// staticfile_buildpack" so that your app will correctly start up with the
// proper buildpack.
func WithHelloWorldApp(f func(dir string)) {
	dir, err := ioutil.TempDir("", "simple-app")
	Expect(err).ToNot(HaveOccurred())
	defer os.RemoveAll(dir)

	tempfile := filepath.Join(dir, "index.html")
	err = ioutil.WriteFile(tempfile, []byte(fmt.Sprintf("hello world %d", rand.Int())), 0666)
	Expect(err).ToNot(HaveOccurred())

	err = ioutil.WriteFile(filepath.Join(dir, "Staticfile"), nil, 0666)
	Expect(err).ToNot(HaveOccurred())

	f(dir)
}

// WithMultiEndpointApp creates a simple application to use with your CLI command
// (typically CF Push). It has multiple endpoints which are helpful when testing
// http healthchecks
func WithMultiEndpointApp(f func(dir string)) {
	dir, err := ioutil.TempDir("", "simple-app")
	Expect(err).ToNot(HaveOccurred())
	defer os.RemoveAll(dir)

	tempfile := filepath.Join(dir, "index.html")
	err = ioutil.WriteFile(tempfile, []byte(fmt.Sprintf("hello world %d", rand.Int())), 0666)
	Expect(err).ToNot(HaveOccurred())

	tempfile = filepath.Join(dir, "other_endpoint.html")
	err = ioutil.WriteFile(tempfile, []byte("other endpoint"), 0666)
	Expect(err).ToNot(HaveOccurred())

	tempfile = filepath.Join(dir, "third_endpoint.html")
	err = ioutil.WriteFile(tempfile, []byte("third endpoint"), 0666)
	Expect(err).ToNot(HaveOccurred())

	err = ioutil.WriteFile(filepath.Join(dir, "Staticfile"), nil, 0666)
	Expect(err).ToNot(HaveOccurred())

	f(dir)
}

// WithNoResourceMatchedApp creates a simple application to use with your CLI
// command (typically CF Push). When pushing, be aware of specifying '-b
// staticfile_buildpack" so that your app will correctly start up with the
// proper buildpack.
func WithNoResourceMatchedApp(f func(dir string)) {
	dir, err := ioutil.TempDir("", "simple-app")
	Expect(err).ToNot(HaveOccurred())
	defer os.RemoveAll(dir)

	tempfile := filepath.Join(dir, "index.html")

	err = ioutil.WriteFile(tempfile, []byte(fmt.Sprintf("hello world %s", strings.Repeat("a", 65*1024))), 0666)
	Expect(err).ToNot(HaveOccurred())

	f(dir)
}

// WithMultiBuildpackApp creates a multi-buildpack application to use with the CF push command.
func WithMultiBuildpackApp(f func(dir string)) {
	f("../../assets/go_calls_ruby")
}

// WithProcfileApp creates an application to use with your CLI command
// that contains Procfile defining web and worker processes.
func WithProcfileApp(f func(dir string)) {
	dir, err := ioutil.TempDir("", "simple-ruby-app")
	Expect(err).ToNot(HaveOccurred())
	defer os.RemoveAll(dir)

	err = ioutil.WriteFile(filepath.Join(dir, "Procfile"), []byte(`---
web: ruby -run -e httpd . -p $PORT
console: bundle exec irb`,
	), 0666)
	Expect(err).ToNot(HaveOccurred())

	err = ioutil.WriteFile(filepath.Join(dir, "Gemfile"), nil, 0666)
	Expect(err).ToNot(HaveOccurred())

	err = ioutil.WriteFile(filepath.Join(dir, "Gemfile.lock"), []byte(`
GEM
  specs:

PLATFORMS
  ruby

DEPENDENCIES

BUNDLED WITH
   1.15.0
	`), 0666)
	Expect(err).ToNot(HaveOccurred())

	f(dir)
}

// WithCrashingApp creates an application to use with your CLI command
// that will not successfully start its `web` process
func WithCrashingApp(f func(dir string)) {
	dir, err := ioutil.TempDir("", "crashing-ruby-app")
	Expect(err).ToNot(HaveOccurred())
	defer os.RemoveAll(dir)

	err = ioutil.WriteFile(filepath.Join(dir, "Procfile"), []byte(`---
web: bogus bogus`,
	), 0666)
	Expect(err).ToNot(HaveOccurred())

	err = ioutil.WriteFile(filepath.Join(dir, "Gemfile"), nil, 0666)
	Expect(err).ToNot(HaveOccurred())

	err = ioutil.WriteFile(filepath.Join(dir, "Gemfile.lock"), []byte(`
GEM
  specs:

PLATFORMS
  ruby

DEPENDENCIES

BUNDLED WITH
   1.15.0
	`), 0666)
	Expect(err).ToNot(HaveOccurred())

	f(dir)
}

// WithBananaPantsApp creates a simple application to use with your CLI command
// (typically CF Push). When pushing, be aware of specifying '-b
// staticfile_buildpack" so that your app will correctly start up with the
// proper buildpack.
func WithBananaPantsApp(f func(dir string)) {
	dir, err := ioutil.TempDir("", "simple-app")
	Expect(err).ToNot(HaveOccurred())
	defer os.RemoveAll(dir)

	tempfile := filepath.Join(dir, "index.html")
	err = ioutil.WriteFile(tempfile, []byte("Banana Pants"), 0666)
	Expect(err).ToNot(HaveOccurred())

	err = ioutil.WriteFile(filepath.Join(dir, "Staticfile"), nil, 0666)
	Expect(err).ToNot(HaveOccurred())

	f(dir)
}

// AppGUID returns the GUID for an app in the currently targeted space.
func AppGUID(appName string) string {
	session := CF("app", appName, "--guid")
	Eventually(session).Should(Exit(0))
	return strings.TrimSpace(string(session.Out.Contents()))
}

func AppJSON(appName string) string {
	appGUID := AppGUID(appName)
	session := CF("curl", fmt.Sprintf("/v3/apps/%s", appGUID))
	Eventually(session).Should(Exit(0))
	return strings.TrimSpace(string(session.Out.Contents()))
}

// WriteManifest will write out a YAML manifest file at the specified path.
func WriteManifest(path string, manifest map[string]interface{}) {
	body, err := yaml.Marshal(manifest)
	Expect(err).ToNot(HaveOccurred())
	err = ioutil.WriteFile(path, body, 0666)
	Expect(err).ToNot(HaveOccurred())
}

// Zipit zips the source into a .zip file in the target dir
func Zipit(source, target, prefix string) error {
	// Thanks to Svett Ralchev
	// http://blog.ralch.com/tutorial/golang-working-with-zip/

	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	if prefix != "" {
		_, err = io.WriteString(zipfile, prefix)
		if err != nil {
			return err
		}
	}

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == source {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name, err = filepath.Rel(source, path)
		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(header.Name)

		if info.IsDir() {
			header.Name += "/"
			header.SetMode(0755)
		} else {
			header.Method = zip.Deflate
			header.SetMode(0744)
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

func ConfirmStagingLogs(session *Session) {
	Eventually(session).Should(gbytes.Say(`(?i)Creating container|Successfully created container|Staging\.\.\.|Staging process started \.\.\.|Staging Complete|Exit status 0|Uploading droplet\.\.\.|Uploading complete`))
}
