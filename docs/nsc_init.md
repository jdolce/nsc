## nsc init

Init a configuration directory

### Synopsis

Init a configuration directory

```
nsc init [flags]
```

### Examples

```
init --name operatorname
init --interactive

```

### Options

```
      --account-key string     account keypath (default generated)
      --account-name string    name for the account (default '<name>_account')
      --cluster-key string     cluster keypath (default generated)
      --cluster-name string    name for the cluster (default '<name>_cluster')
      --create-cluster         create a cluster
      --create-server          create a server
  -h, --help                   help for init
  -n, --name string            name for the configuration environment (default "test")
      --operator-key string    operator keypath (default generated)
      --operator-name string   operator name (default '<name>_operator')
      --server-key string      server keypath (default generated)
      --server-name string     name for the server (default "localhost")
      --user-key string        user keypath (default generated)
      --user-name string       name for the user (default '<name>_user')
```

### Options inherited from parent commands

```
  -i, --interactive          ask questions for various settings
  -W, --long-ids             display long ids
  -K, --private-key string   private key
```

### SEE ALSO

* [nsc](nsc.md)	 - NSC enables you to create and manage NATS accounts and user configurations

###### Auto generated by spf13/cobra on 29-Apr-2019