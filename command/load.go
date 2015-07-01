package command

import (
	"flag"
	"log"
	"fmt"
	"github.com/couchbaselabs/gocb"
	"github.com/mitchellh/cli"
	"os"
	"strings"
)

const (
	// HTTPAddrEnvName defines an environment variable name which sets
	// the HTTP address if there is no -http-addr specified.
	HTTPAddrEnvName = "COUCHBASE_ADMIN_IP"
)

// LoadCommand is a Command implementation that loads JSON into Couchbase.
type LoadCommand struct {
	Ui cli.Ui
}

// TestStruct is document structure we'll use to insert test data.  Ideally
// we only do this if the user hasn't provided a json file to import.
type TestStruct struct {
  Name string
  Age  int
}

func (c *LoadCommand) Run(args []string) int {
	var bucket string 
	var pass string 

	cmdFlags := flag.NewFlagSet("load", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&bucket, "bucket", "default", "bucket")
	cmdFlags.StringVar(&pass, "pass", "", "pass")
	httpAddr := HTTPAddrFlag(cmdFlags)
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	
	url := fmt.Sprintf("http://%s", *httpAddr)
	log.Println(url)

	//cluster, _ := gocb.Connect("couchbase://127.0.0.1")
	//cluster, _ := gocb.Connect(*url)
	//bucket, _ := cluster.OpenBucket(*bucket, *pass)

	couchBucket, err := CouchbaseBucket(url, bucket, pass)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error connecting to Couchbase server: %s", err))
		return 1
	} else if couchBucket == nil {
		c.Ui.Error(fmt.Sprintf("Error connecting to Couchbase server: %s", "No valid bucket found"))
		return 1
	} else {
		log.Println("Got a valid Couchbase bucket")
	}

	dataSource := "dummy"
	docCount, err := LoadDummyData(couchBucket)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error loading %s data into Couchbase server: %s", dataSource, err))
		return 1
	}

	c.Ui.Output(fmt.Sprintf(
		"Successfully loaded %d documents of %s data into Couchbase server", docCount, dataSource))
	return 0
}

func (c *LoadCommand) Synopsis() string {
	return "Load json into Couchbase"
}

func (c *LoadCommand) Help() string {
	helpText := `
Usage: couchloader load [options] ...

  Loads some data into a running Couchbase server.  You should specify your 
  own json data file to load.  But it is not required.

Options:

  -http-addr=127.0.0.1:8091   Admin address of the Couchbase.
  -pass                       Bucket password.
  -bucket                     Bucket to load data into. 
`
	return strings.TrimSpace(helpText)
}

// HTTPAddrFlag returns a pointer to a string that will be populated
// when the given flagset is parsed with the HTTP address of the Consul.
func HTTPAddrFlag(f *flag.FlagSet) *string {
	defaultHTTPAddr := os.Getenv(HTTPAddrEnvName)
	if defaultHTTPAddr == "" {
		defaultHTTPAddr = "127.0.0.1:8091"
	}
	return f.String("http-addr", defaultHTTPAddr,
		"HTTP address of the Couchbase server")
}

func CouchbaseBucket(url, bucket, password string) (cb *gocb.Bucket, err error) {
	if c, err := gocb.Connect(url); err == nil {
		if b, err := c.OpenBucket(bucket, password); err == nil {
			return b, nil
		}
	}
	return nil, err
}

func LoadDummyData(b *gocb.Bucket) (int, error) {
	count := 0
	
	testOutOne := TestStruct{
		Name: "Andy",
		Age:  39,
	}
	
	testOutTwo := TestStruct{
		Name: "Brian",
		Age:  34,
	}
	
	testOutThree := TestStruct{
		Name: "Craig",
		Age:  31,
	}

	testOutFour := TestStruct{
		Name: "Hardeep",
		Age:  28,
	}

	testOutFive := TestStruct{
		Name: "Jarvis",
		Age:  24,
	}

	testOutSix := TestStruct{
		Name: "Nick",
		Age:  22,
	}
	
	testOutSeven := TestStruct{
		Name: "Scott",
		Age:  19,
	}

	if _, err := b.Upsert(testOutOne.Name, &testOutOne, 0); err == nil {
		count++
	} else {
		return count, err
	}

	if _, err := b.Upsert(testOutTwo.Name, &testOutOne, 0); err == nil {
		count++
	} else {
		return count, err
	}

	if _, err := b.Upsert(testOutThree.Name, &testOutOne, 0); err == nil {
		count++
	} else {
		return count, err
	}

	if _, err := b.Upsert(testOutFour.Name, &testOutOne, 0); err == nil {
		count++
	} else {
		return count, err
	}

	if _, err := b.Upsert(testOutFive.Name, &testOutOne, 0); err == nil {
		count++
	} else {
		return count, err
	}

	if _, err := b.Upsert(testOutSix.Name, &testOutOne, 0); err == nil {
		count++
	} else {
		return count, err
	}

	if _, err := b.Upsert(testOutSeven.Name, &testOutOne, 0); err == nil {
		count++
	} else {
		return count, err
	}

	var testIn TestStruct
	b.Get("Andy", &testIn)
 
	fmt.Printf("%s is %d\n", testIn.Name, testIn.Age)	
	return count, nil 
}