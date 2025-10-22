package backend

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type Myconfig struct {
	Filename string     `toml:"-" json:"-"`
	Title    string     `toml:"title" json:"title"`
	License  string     `toml:"license" json:"license"`
	Nats     NatsConfig `toml:"nats" json:"nats"`
}
type NatsConfig struct {
	Servers  []string `toml:"servers" json:"servers"`
	User     string   `toml:"user" json:"user"`
	Password string   `toml:"password" json:"password"`
	Timeout  int      `toml:"timeout" json:"timeout"`
}

func LoadConfig(filename string) (*Myconfig, error) {
	// check filename is exists
	if _, err := os.Stat(filename); err != nil {
		fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return nil, err
		}
		defer fp.Close()
		fp.WriteString(CONFIG_FILE_TEMPLATE)
	}

	myconfig := &Myconfig{Filename: filename}
	if _, err := toml.DecodeFile(filename, myconfig); err != nil {
		return nil, err
	}

	// if password begin with "BASE64$...", then decode weith base64
	if len(myconfig.Nats.Password) > len(PASSWORD_PREFIX) && strings.Index(myconfig.Nats.Password, PASSWORD_PREFIX) == 0 {
		b, err := base64.StdEncoding.DecodeString(myconfig.Nats.Password[7:])
		if err != nil {
			return nil, err
		}
		myconfig.Nats.Password = string(b)
	}

	return myconfig, nil
}

// save myconfig to filename
func SaveConfig(myconfig *Myconfig, filename string) error {
	var tmpfile = filename + ".tmp"
	os.Remove(tmpfile)
	fp, err := os.OpenFile(tmpfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	fp.WriteString("# save by app, on " + time.Now().Format(time.RFC3339))
	fp.WriteString("\n\n\n")

	// encode password with base64
	myconfig.Nats.Password = PASSWORD_PREFIX + base64.StdEncoding.EncodeToString([]byte(myconfig.Nats.Password))

	buf := new(bytes.Buffer)
	if err = toml.NewEncoder(buf).Encode(myconfig); err != nil {
		return err
	}
	n, err := fp.Write(buf.Bytes())

	if err != nil {
		return err
	}
	if n != buf.Len() {
		return fmt.Errorf("write not enough bytes, %d < %d", n, buf.Len())
	}
	if err = fp.Close(); err != nil {
		return err
	}

	if err = os.Rename(tmpfile, filename); err != nil {
		return err
	}
	return nil
}

func (p *Myconfig) Dump() []byte {
	b, _ := json.MarshalIndent(p, "", "  ")
	return b
}

const (
	DEFAULT_CONFIG_FILE  = "natsui.toml"
	PASSWORD_PREFIX      = "BASE64$"
	CONFIG_FILE_TEMPLATE = `
# Natsui config file template


title = "Natsui"
license = "Copyright @ 2025"
	
	
[nats]
	name = "localhost"
	servers = [ "nats://127.0.0.1:4222" ]
	user = ""
	password = ""
	# timeout for connect and read timeout in seconds, default 10 seconds
	timeout = 10
`
)
