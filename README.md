# sonm-marketplace

### Installing from deb package
*Note: all the steps below were tested successfully on
Ubuntu Server 17.10*
+ [Download the latest deb package](https://github.com/sonm-io/marketplace/releases)
```
$wget https://github.com/sonm-io/marketplace/releases/download/1.4/sonm-marketplace_1.4_amd64.3.deb
```
+ Install the package
```
$sudo dpkg -i sonm-marketplace_1.4_amd64.3.deb
```
+ Make sure the package is installed correctly:
```
$sonmmarketplace -version
SONM Marketplace build version "release-2017_12_22_p1"
```
+ Adjust the settings to suit your needs
```
sudo vim /etc/sonm/marketplace-default.yaml
```

You should set at least data_dir and key_store values.
These options must contain absolute paths.

Below is an example config:

```
address: "0.0.0.0:9095"

data_dir: "/etc/sonm"

# blockchain-specific settings.
ethereum:
  # path to keystore
  key_store: ".etc/sonm/keys"
  # passphrase for keystore
  pass_phrase: "any"
```

+ Start the service
```
$sudo service sonm-marketplace start
```

+ Check the logs
```
$ journalctl -fu sonm-marketplace
```

The output should be something like:
```
dec 23 15:59:58 ubuntu-ci systemd[1]: Started SONM Marketplace.
dec 23 15:59:58 ubuntu-ci sonmmarketplace[4799]: 2017/12/23 15:59:58 Starting SONM Marketplace service
dec 23 15:59:58 ubuntu-ci sonmmarketplace[4799]: Using /etc/sonm/keys as KeyStore directory
dec 23 15:59:58 ubuntu-ci sonmmarketplace[4799]: KeyStore directory does not exists, try to create it...
dec 23 15:59:58 ubuntu-ci sonmmarketplace[4799]: 2017-12-23T15:59:58.311+0300        INFO        Config loaded form        {"path": "/etc/sonm/marketplace-default.yaml"}
dec 23 15:59:58 ubuntu-ci sonmmarketplace[4799]: 2017-12-23T15:59:58.311+0300        INFO        Public key        {"address": "0x493c1E27ea5D23C3c2F547048e4eAAd66c2af97F"}
dec 23 15:59:58 ubuntu-ci sonmmarketplace[4799]: 2017-12-23T15:59:58.311+0300        INFO        Data dir        {"path": "/etc/sonm"}
dec 23 15:59:58 ubuntu-ci sonmmarketplace[4799]: 2017-12-23T15:59:58.311+0300        INFO        Importing database schema        {"schema": "/etc/sonm/schema.sql"}
dec 23 15:59:58 ubuntu-ci sonmmarketplace[4799]: 2017-12-23T15:59:58.312+0300        INFO        Database schema successfully imported
dec 23 15:59:58 ubuntu-ci sonmmarketplace[4799]: 2017-12-23T15:59:58.640+0300        INFO        Ready to serve        {"addr": "0.0.0.0:9095"}
```


### Building from source

*Note: all the steps below were tested successfully on
Ubuntu Server 17.10 with go lang 1.9.2.*

If you want to build **sonm-marketplace** on your own you need to do the following:
+ [Install Go lang](https://golang.org/doc/install)
Recommended version is 1.9.2
+ Make sure the GOPATH is set correctly
```
$ echo $GOPATH
/home/screwyprof/go
```
+ Clone the repository
```
$ git clone git@github.com:sonm-io/marketplace.git $GOPATH/src/github.com/sonm-io/marketplace
```

+ Run build
```
cd $GOPATH/src/github.com/sonm-io/marketplace
make build
```

The last lines of the output will be something like:
```
cd cli/cmd && /usr/local/go/bin/go build -i -ldflags "-X 'main.AppVersion=master' -X 'main.GitRev=76b5253'
-X 'main.GoVersion=1.9.2' -X 'main.BuildDate=2017-12-23 14:56:25'
-X 'main.GitLog=76b5253 (HEAD -> master, origin/master, origin/HEAD) Merge remote-tracking branch  origin/release-2017_12_22_p1 '"
-o /home/screwyprof/go/bin/marketplace
```

+ Check installation
Given that your $GOPATH/bin is in $PATH run the following:
```
$marketplace -build-info
```

The output may look as follows:
```
SONM Marketplace  build version "master"
Built at 2017-12-23 14:56:25 with compiler "1.9.2"
From git rev 76b5253

From git commit 76b5253 (HEAD -> master, origin/master, origin/HEAD) Merge remote-tracking branch  origin/release-2017_12_22_p1
```
