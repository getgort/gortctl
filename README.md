# Gortctl

`gortctl` is a CLI tool for administering a
[Gort](https://github.com/getgort/gort) chatops server installation.

# Configuring

`gortctl` uses an INI-formatted configuration file, conventionally
named `.gortctl` in your home directory. This is where you can store
connection credentials to allow `gortctl` to interact with Gort's REST
API.

An example file might look like this:
```
[defaults]
profile = gort

[gort]
password = "seekrit#password"
url = https://gort.mycompany.com:4000
user = me

[preprod]
password = "anotherseekrit#password"
url = https://gort.preprod.mycompany.com:4000
user = me
```

Comments begin with a `#` character; if your password contains a `#`,
surround the entire password in quotes, as illustrated above.

You can store multiple "profiles" in this file, with a different name
for each (here, we have `gort` and `preprod`). Whichever one is noted
as the default (in the `defaults` section) will be used by
`gortctl`. However, you can pass the `--profile=$PROFILE` option to
`gortctl` to use a different set of credentials.

While you can add profiles to this file manually, you can also use the
`gortctl profile create` command to help.

# Getting Help

The `gortctl` executable contains a number of commands and
sub-commands. Help is available for all of them by passing the `--help`
option. Start with `gortctl --help`, and go from there.
