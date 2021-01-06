# Go Viper Tutorial

A  Go Viper package tutorial with examples. It includes integration with Cobra and Pflag packages.

## Note this tutorial is still Under Development 

## Viper - What is it
Viper is a Go package designed to support and simplify the management of program configuration values.

Viper is a complete configuration solution for Go applications including 12-Factor applications, where it meets the requirements for program configuration management. It is designed to work within an application, and can handle all types of configuration needs and formats. Viper supports:

- Setting defaults.
- Reading from JSON, TOML, YAML, HCL, envfile, INI and Java properties files.
- Live watching and re-reading of config files (not covered in tutorial).
- Reading from environment variables.
- Reading from remote config systems (etcd or Consul), and watching changes (not covered in tutorial).
- Reading from command line flags.
- Reading from buffer  (not covered in tutorial).
- Setting explicit values programatically.
- Writing current configuration values to a new or existing configuration file.

Note that Viper also allows the program to save its current configuration values for reloading or reference at a later time, or by a different program in a suite of programs. 

For general use, the tutorial provides enough options for flexible configuration management.
 
Viper can be thought of as a registry for all the application's configuration needs.

Viper can be used without Cobra. Viper has few linkages with Cobra.  Cobra provides a configuration file initialization component for Viper when CLI commands are created. Cobra command line flags, can be linked to Viper configuration values, and this capability will be seen in this tutorial.

To understand the roles of Viper and Cobra, it is best to think of them separately, and then when they are both clear as to their individual functionality, the joint linkages are easier to understand. Read this document first, and then the document 'How to Use Cobra' at https://github.com/dsbitor/GOCobraTutorial .  

This tutorial contains five examples, each of which builds-on or explains the various features of Viper. The last example, explains the integration of Viper with Cobra.

The notes are based on using Go version go1.14.3 darwin/amd64, and the Viper and Cobra packages available at httpe://github.com/spf13/  on an Apple Mac running OS X Mojave, (version 10.14.6).

Some of the above is taken from the Viper documentation at https://github.com/spf13/viper . 

## Installing Viper

To get viper issue the following command:

    go get github.com/spf13/viper

## Examples

All the code for the examples are stored at https://github.com/dsbitor/GOViperTutorial . 


DSB check above

## Example 1 - Reading a Configuration File

### Using a Configuration File

Example1 uses a configuration file as a source for program configuration values. The first example sets up configuration values in a small TOML configuration file for important program infrastructure namely, debugging, logging and program identification. For more information on TOML (Tom's Obvious Minimal Language) syntax see https://github.com/toml-lang/toml.  The code is presented in aspexv2.go and the configuration file is in aspexv2.toml. Both, for simplicity, are in the same source directory.

If you have the choice of configuration file format, TOML and YAML provide the simplest, clearest formats for configuration purposes. JSON is a better format for data transfer. Java Property Files are only useful where interchange with Java environments is required for syntax see https://en.wikipedia.org/wiki/.properties .  For clarity stay away from INI format files, unless specific requirements dictate their use.  

#### INI and Key=Value Files

Note that Viper supports `key=value` files which are the basis of `ini` files. If you create your own key value files, you can test them and ensure they work in your Viper environment. If you are using `ini` files from another computing, OS or application/program environment you must be careful. Viper **does not support**  all `ini` files, due to the variability in their formatting standards and the lack of accurate rules/methods of validating their content. See Appendix A, for further discussion about INI files.

### Reading a Configuration File

To allow flexability over, file type, naming and location, several Viper parameters can be set  to direct access to the configuration file.

If you are starting from scratch, and you are trying to understand the principles, try using TOML or YAML files first, as these are the simplist to create, and will cover most of your needs. If the syntax of your configuration file is incorrect Viper will lilely report the error.
https://www.toml-lint.com/ offers an online syntax validator for TOML files; and for YAML files try https://codebeautify.org/yaml-validator, which also provides a formatter and guidance. If you are using Visual Studio Code you can install add-ins to format and validate TOML, YAML JSON etc.

#### Setting-up To Read a Configuration File

Note that the type can be deduced from reading the file name extension, or can be specifically set; and several different possible paths to the file can be configured to:

- Set file configuration name use `viper.SetConfigName("config")` where the file root name is `config` but the type is not specified.
- Set the configuration file suffix type use `viper.SetConfigType("toml")` which is required if the configuration file does not have a suffix.
- Set the path or paths to  searched for the configuration file use `viper.AddConfigPath("path/to/conf/file")` as ahown in these examples:
    - `viper.AddConfigPath("$HOME/.appname")` // relative to `$HOME`
    - `viper.AddConfigPath(".")` // Use current (working) directory
    - `viper.AddConfigPath("/etc/appname/")` // common conf file location
    - `viper.SetConfigFile("./aspexv2.toml")` // Specific ref to path, name, type

 - You can call `AddConfigPath` multiple times to add multiple search paths.
 
To read the file simply execute:

```go
err := viper.ReadInConfig() // Find and read the config file
if err != nil { // Handle errors reading the config file
	panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
```

### TOML Configuration File

Here is a demonstration of configuration file use.  We are going to demonstrate using a TOML file. See TOML description at https://github.com/toml-lang/toml. Viper handles TOML smoothly, and is sometimes preferred to YAML and JSON configuration files due to its simplicity. The TOML file is loaded from the users home directory.

The TOML configuration file contains the following initialization data.

```go
# This is a TOML document.
# It provides configuration data for program aspexv1.go

title = "aspexv1 configuration file"

[logging]
filename = "aspexv1.log"
dir = "./logs/"

[devparms]
debug = false

[appident]
appshortname = "aspexv1"
```

The code to read and use or output data from the configuration file is shown in Example 1 below:

```go
//Example 1
// Code to read viper configuraion from TOML file in aspexv1.go
package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// Set config file path including file name and extension
	viper.SetConfigFile("./aspexv1.toml")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	// Confirm config file used
	fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())

	debug := viper.Get("devparms.debug") // returns string
	fmt.Printf("devparms.debug Value: %v, Type: %T\n", debug, debug)

	logfilename := viper.Get("logging.filename") // returns string
	fmt.Printf("logging.filename Value: %v, Type: %T\n", logfilename, logfilename)

	// Check if a particular key is set print if avail
	if !viper.IsSet("title") {
		log.Fatal("missing title")
	} else {
		fmt.Printf("Configuration File title: %v \n", viper.Get("title"))
	}
}
```

As you can see a considerable amount of overhead associated with defining struct's to store the data from the TOML/JSON/YAML etc. types of file is removed, as well as the overhead of opening, reading and parsing the configuration file.

If you just need to load and use values from a configuration file, you have everything required in the above example.

## Example 2 - A Multi-source Sample 

The simple flexability using Viper with configuration files was demonstrated above, but lets now create a more realistic example that uses multiple sources for confiuration values and the associated Viper features. 

### Requirements 

- Get working directory from environment variable PWD.
- Get the hostname from environment variable HOSTNAME.
- Default values to ensure that without the configuration file the program will run in production mode.
    - Data, results and logs files are written to the users home directory.
    - Minimal copyright, support, and software licence data will be set to default values in the code.
    - Working directory will be set to `.` if environment variable HOME is not set.
- From the configuration file:
    - The program will provide configuration values for a flexible, parameterized logging capability.
    - Ability to switch on various debugging components. 
    - Minimal application identification data will be set beyond those, and overwriting those in  default.

### Home DIrectory - Useful Helper Function

Note that we use a package called `go-homedir` available at github.com/mitchellh/go-homedir. It is renamed to `homedir` in the import. Viper uses this facility to safely get the home directory no matter which OS you are using. `homedir` will be used along with a small amount of adjustment to get the configuration file.

### Configuration File Design
The configuration file should be designed prior to building the code, because:
- This will ensure the requirements are accommodated.
- Configuration item names will appear in the code.
- Plan the value types needed. 
- Plan the hierarchy, as this adds understanding and naming clarity.
- Configuration file type needs to be considered.

We will build on top of the TOML configuration file in Example 1 above, although the configuration could be equally well represented in a YAML or JSON files. The file is `aspexv2.toml` and is in the examples folder.

```toml
# This is a TOML document.
# It provides configuration data for - aspexv2.go

title = "aspexv2 configuration file"

hostname = "ITOR"
wrkdir = "./data"

[logging]
filename = "aspexv2.log"
# logs are stored in the logs subdirectory of the users home directory
dir = "/logs/"
loglevel = "TRACE"
log = true

[devparms]
debug = false
tracing = false

[app]
    [app.id]
    name = "aspexv2 - A simple Example of a Viper Configuration File"
    shortname = "aspexv2"
    softlicencetype   = "The Apache 2 License"
    [app.developer]
    shortorg = "ITOR"
    org = "I.T. Operational Risk Ltd., Toronto Canada"
    emailaddr = "itoperaionalrisk@gmail.com"
    website = "www.itor.io"

```

## Precedence for Overiding Values

Since Viper offers six different types of sources for configuration values, a value can be be overidden by another value. Viper uses the following precedence order. Each setting capability takes precedence over the capability below it:

- explicit call to Set
- flag value
- environment varaible
- configuration file value
- key/value store
- default

## Default Values
Better programming practice requires that variables that control program operation should have default values, so that that operation of the program is predictable, and missing values in a configuration file do not result in unpredictable results.

The definition of default values using Viper is simplified and integrated into the use of the other subsequent setting capabilities such as environment variables, configuration files and flags.

### Adding Default Values 

Example 2, `aspexv2.go` and its TOML configuration file `aspexv2.toml` show how to incorporate default values into the configuration environment. 

Lines XXX and XXX in `aspexv2.go` show two default configuration values declared and value set. Note that the varaible `hostname` does not have an equivalent value in the TOML configuration file `aspexv2.toml`.

## Environment Variables

### Why Use Environment Variables

An environment variable is a variable whose value is set outside the program, typically through functionality built into the operating system, microservice or a program. An environment variable is made up of a name/value pair, and any number may be created and available for reference at a point in time. They are often set at:

- Operating system initiation
- User sign-on
- Initiation of a shell service such as Unix (Bash, Bourne, C shells),  Windows PowerShell.
- Installation of application, microservice or program(s)
- Users changing expected operating characteristics

Environment variables are easy to create, access, and change.

#### Examples of Environment Variable Use

Use cases for environment variables include but are not limited to data such as:

- Execution mode (e.g., production, development, staging, etc.)
- Domain names
- API URL/URI’s
- Public and private authentication keys (only secure in server applications)
- Group mail addresses, such as those for marketing, support, sales, etc.
- Service account names
- Sensitive information that cannot be saved on external servers and code repositories
- Variables common to a group of programs or applications
- Database set-up and access

What these have in common are their data values change infrequently and the application logic treats them like constants, rather than mutable variables.

### Viper - Access to Environmental Variables

Viper makes it easy to set configuration values from environment variables. 

The function `viper.AutomaticEnv()` will read all available environment variables and make them available to the program. They will override variables with the same name if a configuration file (YAML, TOML, JSON etc.) and was initialized by Viper. (YAML, TOML, JSON etc.)

Environment variable names can have unusual capitalization and conventions which may not be easy to use in Go programs where variable names follow expected conventions, but this can be easily remedied. `viper.BindEnv` takes one or two parameters.  I recommend you use the two parameter version to map the Go viper configuration name to the external environment variable name, thus providing consistency of naming across the configuration.  

Example 2, `aspexv3.go` and its TOML configuration file `aspexv3.toml` show how to incorporate Environment variable values into the configuration environment. 

Lines XXX and XXX in `aspexv2.go` shows default configuration values declared and value set.

Lines XXX and XXX in `aspexv2.go` shows environment variables assigned to  configuration values.

## Example 3 - A Real-life Example

Most of the flexability of Viper was demonstrated above, but lets now create a real-life example that uses all of the major sources for confiuration values and the associated Viper features. 

This example deals with flags, and getting values from a set of configuration values managed by Viper. 

The final example, Example 4, below deals with taking this example and integrating it with Cobra based caommand definitions with associated flags.

### Requirements

- Get working directory from environment variable PWD.
- From the configuration file:
    - The program requires a flexible, parameterized logging capability.
    - Ability to switch on various debugging components.
    - Application identification data.
- Default values to ensure that without the configuration file the program will run in production mode.
    - Data, results and logs files are written to the users home directory.
    - Minimal copyright, support, and software licence data is embedded in the code.
    - Working directory will be set to  if environment variable HOSTNAME is not set.
- Flags will over-ride default and configuration data values.
- Programmatically set the current IP address which in the config file is set to 0.0.0.0 (‘no particular address’ placeholder).

## Using Flags

Viper has the ability to bind to flags. Flags offer a method of changing configuration variables at run time. Specifically, Viper supports Pflags as also used by the Cobra library. pflag is a drop-in replacement for Go's flag package, but also implements POSIX/GNU-style `-f` single character flags as well as longname flags such as `--debug`.  This is important where there is an expectation that GNU flags are available for use, such as in Unix and othe operating system utilities. 

Flags are handled in a similar manner to Environment Variables. The configuration variable is bound to to a flag.  Like `BindEnv`, the value is not set when the binding method is called, but when it is accessed. This means you can bind as early as you want, even in an `init()` function. For individual flags, the `BindPFlag()` method provides this functionality.

### The flag and pflag Packages

#### The flag Package

Flag (and pflag) set-up encompasses two steps:

1. Declare each flag's structure.
2. Issue a  `flag.Parse()` to instantiate the flag structure(s) with the current values supplied when the program was invoked.

The Parse() function will continue to parse flags that it encounters until it detects a non-flag argument. The flag package makes these available through the Args() and Arg() functions.

The declaration can be made in several forms, two of which are:

    var ip *int = flag.Int("flagname", 1234, "help message for flagname")

Or you can bind the flag to a variable using the Var() functions as shown below using `flag.IntVar`. the alternative functions in the pflag package is `pflag.IntVar` and `pflag.IntP`.

    var flagvar int
    func init() {
        flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
    }

To get the values from the flags when all flags have been declared use flag.Parse().

The flag package provides methods also to parse non-flag parameters. The remaining parameters entered when the program was invoked are provided using the `flag.Args()` function. It returns a slice of strings with the parameters not parsed as flags.

#### The pflag Package Capability and Example

Refer to the example code in the source file `aspexv3f.go`. This example uses the `pflag` package. The file contains only the code necessary to show flag definition, Viper default and flag overide code, along with output showing flag value states at various points along the the source code. To run this example use the two terminal commands below:

    go build aspexv3f.go 
    ./aspexv3f --debug=true --linelist=42
    
flags in `pflag` are set-up in the same manner as the flags description above. The major difference is that pflag allows the implemention of POSIX/GNU-style `--flags`.

`--` flags are implemented by using an alternate version of the flag declaration
such as `flag.Int()` as `pflag.IntP()` which adds a `P` suffix to the function name. Note that `pflag` also implements `pflag.Int()`, which acts exactly the same as the `flag` version. Two examples of the extended notation declarations are:

    pflag.IntP("linelist", "l", 5, "Number of lines to list")
    debug := pflag.BoolP("debug", "d", false, "Switch debugging off/on")


`pflag` has the same flag parsing mechanism as the `flag` package, and is called after all flags have been decalred. 

    // Call pflag.Parse() after all flags have been defined
    pflag.Parse()

Executing the program `aspexv3f` allows flags to be entered in several ways:

- --debug=true
- --debug
- -d=true
- -d true
- -d
- --linelist=42
- --linelist 42
- -l=42
- -l 42

Note that the -d and --debug flag switches the value of the boolean to the alternate of the current value. 

**See https://godoc.org/github.com/spf13/pflag for full pflag detailed developer documentation.**

### Viper's Use of Flags

Viper uses the pflag package, which allows the command line flags to be bound to configuration variables.

	viper.BindPFlags(pflag.CommandLine)

## Getting Values From Viper

Most of our effort above was associated with getting values into Viper using the various Viper's configuration entry functions. We must also be able to access those values easily and effectively. You have already seen some of these access mechanisms, but in this section we will describe all the major capabilities.

### Getting Values by Key Name

Keys (names) are case insensitive. 

Viper documentation states 'Viper merges configuration values from various sources, many of which are either case insensitive or uses different casing than the rest of the sources (eg. env vars). In order to provide the best experience when using multiple sources, the design decision was made to make all keys case insensitive.'

We recommend that when you use a key you use all lower case for consistency and ease of reading.

You only need to provide the key  for the value you want, and the type you expect that value to be. Viper provides a set of GetType methods where  'Type' is the data type such as Int, String, Float, Bool etc.. Here is the list:

- GetBool(key string) : bool
- GetFloat64(key string) : float64
- GetInt(key string) : int
- GetIntSlice(key string) : []int
- GetString(key string) : string
- GetStringMap(key string) : map[string]interface{}
- GetStringMapString(key string) : map[string]string
- GetStringSlice(key string) : []string
- GetTime(key string) : time.Time
- GetDuration(key string) : time.Duration

Each Get function will return a zero value if it’s not found, not an error. If you are not sure the key exists, then use `IsSet` described below.

In addition there are four other important access functions:

- Get(key string) : interface{}  - This function returns the string value for the provided key.
- IsSet(key string) : bool  - Function returns true if the specified key exists.
- AllKeys() : []string - returns all available keys as string slice.
- AllSettings() : map[string]interface{} - Function returns a map of keys and strings representing the values of all defined.

## Example 4 - A Real-life Example With Cobra Defined Flags

This example shows how to integrate Example 3 with Cobra based caommand definitions with associated flags. Finally this example will ensure configuration parameters will be in place prior to any other program activity, by using the function `init`.

### Requirements

- Requirements from Example 3.
- Create Cobra commands and associated flags.
- Saves the configuration to a new file after adding/changing a value from a flag.
- Demonstrates accessing configuration values.

### Marsalling
Marshaling of configuration variables allows them to be written out:
- As a new or updated configuration file, where you want to store all modifications made at run time
- As a structure for use either within the program
- As output in some form other than as a configuration file

### Unmarshalling

Viper offers an option of un-marshalling all, or a specific value to a struct, map, etc. Normally using the named configuration variable form offered by Viper is more simple and clear for programming, but there could be some unique cases where it is useful. 

Steve Francia the designer and author of Viper (and Cobra) stated in a chat are "Typically with Viper you just use things like `Viper.Get("key")` to get the values you need out and don't marshall it into a struct. There's nothing wrong with the struct approach and sometimes is the right way to do it, but it's definitely more upfront work and less flexible." 

## Cobra Integration

Cobra and Viper work together as long as the relationship is understood. And the steps to integrate the two components are taken in an orderly manner.

Complete the Cobra work to create the command and sub commands required for the project first. Make sure you are using Go Modules, and a source code management system, such as Git or SVCS.

Part of the Cobra design decisions is to identify the Cobra defined flags for the defined commands and sub-commands. Of the identified command flags, those that will modify the program configuration. Those configuration flags will have to be connected to the configuration file via Viper. 

Cobra provides a linkage to Viper services in the `root.go` file created when the Cobra command `cobra init` is created. And also provides similar capabiity in each of the command and sub-command definitions. 

### Cobra Defined Flags

Any flag defined within a Cobra command environment including `root.cmd` can be linked to a Viper configuration variable. A flag defined in Cobra is linked to an environment variable using the function  `viper.BindPFlag` which we have seen above, when dealing with pflags and configuration variables. This is possible, because Cobra flags are really just 'enhanced' pflags.

A small example implements two Cobra flags which will be linked to Python configuration variables. 

Lets assume we want to control trace messages through all Cobra commands in an multiple command environment. The second flag will be active with only one command,
and will change the working directory for both current and future use, by recording, when used, a new value for the working directory.

In the Cobra component, a flag is defined which is boolean and switches tracing messages on and off. It is a a global flag as it is defined as the flag in  the file `cmd/root.go`. The flag is `--trace` or short name `-t`, so it can be invoked with either a long or short flag name. The flag trace will have an equivalent configuration value which can be read fron a TOML configuration fle, and has a default configuration setting of false.

The second configuration flag used in this example is, a local flag to the command initialize. The initialize command creates the run time and session computing environment for the other commands. A flag is provided to change the working directory, and that working directory will overide working directory configuration values already set up by default or via the configuration file values.

The source files for this example are Cobra environments, with their own particular layout of files. If you are not familiar with Cobra, and how it works, and have never created a Cobra environment, please read the Tutorial 'How to Use Cobra' and run the examples in that tutorial first. They are avaialable at https://github.com/dsbitor/GOCobraTutorial . 

#### Example mpm1

mpm1 implements the first three components of a Cobra based comprehensive solution. It implements configuration variables from: 

- defined default values
- a configuration file
- environment variables

The added or changed values of the variables are displayed after each code comonent has beeen executed.

#### Example mpm2

mpm2 implements full a Cobra based comprehensive solution, adding to the mpm1. It implements configuration variables from: 

- defined default values
- a configuration file
- environment variables
- global and local flags associated with commands
- writing a changed configuration file
- setting configuration values programatically

The added or changed values of the variables are displayed after each code comonent has beeen executed.

### Executing the Cobra Examples mpm1 and mpm2

To **compile** the Go code in folders `mpm1` or `mpm2` use **`go install mpm1`** or **`go install mpm2`**. The command `go install` which does a compile/link/save-binary-to-binlib can be executed from the root of `mpm1` or `mpm2`. The resultant binary program will normally, by default, be stored in the `$GOPATH/bin/`. If you want to be sure list that directory after getting a clean-compile from `go install`.

To execute the code run at the terminal prompt, issue  **mpm1** or **mpm2* with appropriate comannd and flags. For example **for mpm1 use**, `mpm2 initialize` as flags are not implemeted in this partial version, and **for mpm2 use** `mpm2 initialize --trace=true` . Note that issuing just the bare **mpm1 or mpm2 will produce help information as output**.

The descriptive output go code can be safely commented out without affecting functionalty.

## Saving Configuration Files

Viper offers flexability in saving a configuration file, which can be very useful where program or application settings need to be carried over from one execution of a program to the next execution. 

Formats for saving configuration files are JSON, TOML, YAML, HCL, and INI. Saving is as simple as configuring 

DSB XXXXXX above

Four functions, each with its own capabilities, available for writing the configuration file are:

- WriteConfig - writes the current viper configuration to the predefined path, if exists. Errors if no predefined path. Will overwrite the current config file, if it exists.
- SafeWriteConfig - writes the current viper configuration to the predefined path. Errors if no predefined path. Will not overwrite the current config file, if it exists.
- WriteConfigAs - writes the current viper configuration to the given filepath. Will overwrite the given file, if it exists.
- SafeWriteConfigAs - writes the current viper configuration to the given filepath. Will not overwrite the given file, if it exists.

Examples from the documentaion show:

    viper.WriteConfig() // writes current config to predefined path set by     'viper.AddConfigPath()' and 'viper.SetConfigName'
    viper.SafeWriteConfig()
    viper.WriteConfigAs("/path/to/my/.config")
    viper.SafeWriteConfigAs("/path/to/my/.config") // will error since it has already been written in statement above
    viper.SafeWriteConfigAs("/path/to/my/.other_config")

`viper.WriteConfigAs` is shown in the final example when used with Cobra in the `initialize` command to save an altered configuration.

### Accessing Configuration Values using Viper Functions

Using Example 4 above all of the statements below are valid:

```go
if viper.IsSet("wrkdir") {fmt.Printf("Working Directory Exists \n") }

workingdir := viper.Get("wrkdir")
fmt.Printf("Working Directory = [%s]\n", workingdir)

# Getting a value from within a structure
devemail := viper.Get("app.developer.emailaddr")
fmt.Printf("Developer eMail = [%s]\n", devemail)

AllConfig:= viper.AllSettings() 
for k, v := range AllConfig { 
    	fmt.Printf("Config key = [%s] Config value = [%s]\n", k, v)
}
```

## Appendix A - Further Notes `ini` Files

Since some of the more complicated `ini` files are basically key value pairs with groups, delimited by square `[]` or curly `{}` brackets, and they do not work with Viper, you could build your own parser if essential,  using the `io.Reader` feature to read the files. 

Viper uses the package `gopkg.in/ini.v1` to read and write ini files, and which as of October 2020, was actively supported and updated. It is described as 'The package `ini` provides INI file read and write functionality in Go'. 

Note that Viper will read configuration data in INI format, but it is not clear from the documentation what support is provided for writing INI files from configuration data.

If Viper does not support a particular feature direct use of `gopkg.in/ini.v1` in some custom code may assist in resolving the problem. If you want to see what `gopkg.in/ini.v1` is capable of reading and writing, see https://github.com/go-ini/ini/blob/master/testdata/full.ini .

## DSB Fix-up

The code to add the initialization file is provided as a prototype in `root.go`. The prototype code needs some modification. The directory where the config file is stored, is not the users home directory. A more common practice is to have a hidden configuration direcory on Unix like systems. In this implementation it is the `.config` subdirectory of the users home directory. The configuration file name also must change. instead of `.add3` we have chosen `add3.config.toml`. 

The changes to root.go are in the `initConfig` function that follows the comment:

    // initConfig reads in config file and ENV variables if set.

Following the comment:

    // Search config in home directory with name ".add3" (without extension).
    
This comment and two subsequent lines of code will be changed 

```go
		// Search config in home directory with name ".add3" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".add3")
```