# cogctl2

`cogctl` is a CLI tool for administering a
[Cog2](https://github.com/clockworksoul/cog2) chatops server installation.

# Configuring

`cogctl` uses an INI-formatted configuration file, conventionally
named `.cogctl` in your home directory. This is where you can store
connection credentials to allow `cogctl` to interact with Cog's REST
API.

An example file might look like this:
```
[defaults]
profile = cog

[cog]
password = "seekrit#password"
url = https://cog.mycompany.com:4000
user = me

[preprod]
password = "anotherseekrit#password"
url = https://cog.preprod.mycompany.com:4000
user = me
```

Comments begin with a `#` character; if your password contains a `#`,
surround the entire password in quotes, as illustrated above.

You can store multiple "profiles" in this file, with a different name
for each (here, we have `cog` and `preprod`). Whichever one is noted
as the default (in the `defaults` section) will be used by
`cogctl`. However, you can pass the `--profile=$PROFILE` option to
`cogctl` to use a different set of credentials.

While you can add profiles to this file manually, you can also use the
`cogctl profile create` command to help.

# Getting Help

The `cogctl` executable contains a number of commands and
subcommands. Help is available for all of them by passing the `--help`
option. Start with `cogctl --help`, and go from there.
