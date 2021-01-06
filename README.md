This document is a Go Viper package tutorial with examples. It includes integration with the Cobra and Pflag packages.

## Viper - What is it
Viper is a Go package designed to support and simplify program configuration variables and their associated values.

The Viper documentation on the Viper GitHub site is documentary, most other tutorials and examples miss out steps, exclude important points regarding set-up, implementation, usage or testing.

This tutorial is long but attempts to cover all of the most important aspects of configuration value management, in a series of steps, each adding additional complexity and features available in Viper. The tutorial also covers connectivity to the Cobra package. Both Viper and Cobra rely on the pflag package for flag management, and therefore this tutorial discusses some aspects of pflag.

Viper is a complete configuration solution for Go applications including 12-Factor applications, where it meets one of those factor requirements, namely program configuration management. It is designed to work within an application and can handle all types of configuration needs and formats. Viper supports:

- Setting defaults.
- Reading from JSON, TOML, YAML, HCL, environment file, INI and Java properties files.
- Live watching and re-reading of config files (not covered in tutorial).
- Reading from environment variables.
- Reading from remote config systems (etcd or Consul), and watching changes (not covered in tutorial).
- Reading from command line flags.
- Reading from a buffer  (not covered in tutorial).
- Setting explicit values programmatically.
- Writing current configuration variable and associated values to a new or existing configuration file.

Viper acts as a registry for all the application's configuration needs.

Note that Viper also allows the program to save its current configuration values, allowing them to be reloaded or referenced later or by a different program in a suite of programs.
For general use, the tutorial provides enough options for flexible configuration management.


You can use Viper without Cobra as Viper has few linkages with Cobra.  CLI commands in Cobra can provide a configuration value initialization component for Viper, and this capability will be seen in this tutorial.

###Relationship to the Cobra Package

It is best to think of  Viper and Cobra separately and understand their roles. When those roles are clear as to their functionality, the joint linkages are easier to understand. Read this document first, and then the tutorial document 'How to Use Cobra' at https://github.com/dsbitor/GoCobraTutorial , at the appropriate time identified below in the section titled Cobra Integration.

This tutorial contains five examples, each of which builds on the earlier example and shows most of Viper's various features. The last example explains the integration of Viper with Cobra.

The notes use Go version go1.14.3 darwin/amd64, the Viper and Cobra packages available at httpe://github.com/spf13/, and ran on an Apple Mac running OS X Mojave, (version 10.14.6).

## Installing Viper

To get Viper issue the following command:

    go get github.com/spf13/viper

## Examples and Their Execution

All the code for the examples, and this file as a markdown document, are stored at https://github.com/dsbitor/GoViperTutorial . 

**When executing the examples here, make sure you call the program from the directory that contains the source code file, which for Examples 1 to 4, is directories aspex1 to aspex4. Example 5 includes two examples associated with Viper and Cobra integration. They are in two directories called mpm1 and mpm2. For Example 5 start from directory mpm1 or mpm2 as appropriate. Further details on running Example 5 are preceded in the section 'Executing the Cobra Examples' in mpm1 and mpm2** 

To keep the examples simple, most access the configuration file *.toml, stored in the same directory (folder) as the source code.

Example 1 to 4 can be executed by issuing a `go run sourcefile.go` and appending CLI parameters as required for the given example.

## Example 1 - Reading a Configuration File

### Using a Configuration File

Example1 uses a configuration file as a source for program configuration values. The first example sets up configuration values in a small TOML format configuration file containing necessary program infrastructure such as debugging, logging and program identification. For more information on TOML (Tom's Obvious Minimal Language) syntax see https://github.com/toml-lang/toml. The code is in aspex1/aspexv1.go and the configuration file is in aspex1/aspexv1.toml. Both files, for simplicity, are in the same source directory.

You can find a set of good examples of TOML at https://learnxinyminutes.com/docs/toml/ on the page entitled 'Learn TOML in Y minutes'. One comment seen in several web pages is that TOML is evolving, but readers should note that TOML now has an official definition. The Go package go-toml at https://github.com/pelletier/go-toml , is fully compliant with the official definition and is the package used by Viper.

Configuration files can use various formats. In this tutorial, we will use TOML, but configuration file formats such as YAML, JSON, INI are also available.

You can find YAML syntax and use descriptions at https://yaml.org/spec/1.2/spec.html. The web page https://rollout.io/blog/yaml-tutorial-everything-you-need-get-started/,
titled 'YAML Tutorial: Everything You Need to Get Started in Minutes' is useful for understanding YAML components.

If you choose the configuration file format, TOML and YAML provide the most straightforward, clearest arrangements for configuration purposes.

JSON is a better format for data transfer. Java Property Files are only useful where interchange with Java environments are required. For syntax see https://en.wikipedia.org/wiki/.properties.

Please stay away from INI format configuration files for clarity and simplicity unless specific requirements dictate their use. Formatting rules are variable for INI files.

#### INI and Key=Value Files

Note that Viper supports key=value files, which are the basis of INI files. If you create your key-value files, you can test them and ensure they work in your Viper environment. If you are using INI files from another computing, OS or application/program environment, you must be careful. Viper does not support all files with a file type of INI, due to the variability in their formatting standards and the lack of accurate syntax rules and therefore validating their content. See Appendix A, for further discussion about INI files.

### Reading a Configuration File
Viper provides flexibility over configuration file-type, naming and location; Viper provides functions to set access to the configuration file. If you are starting from scratch and trying to understand the principles, try using TOML files first, as these are the simplest to create, and will cover most of your needs. If the syntax of your configuration file is incorrect Viper will likely report the error.

https://www.toml-lint.com/ offers an online syntax validator for TOML files; and for YAML files try https://codebeautify.org/yaml-validator, which also provides a formatter and usage guidance. If you are using Visual Studio Code, and some other source code editors, you can install add-ins to format and validate TOML, YAML JSON files.

#### Setting-up To Read a Configuration File

Viper can deduce the type by reading the file name extension. Viper also allows the extension to be specified, and the specification of multiple possible paths to the file. 

Please stay away from INI format configuration files for clarity and simplicity unless specific requirements dictate their use. Formatting rules are variable for INI files.

#### Setting-up To Read a Configuration File

Note that the type can be deduced from reading the file name extension, specified, and different possible paths to the file, using: 

- Set file configuration name use `viper.SetConfigName("config")` where the file root name is `config` but the type is not specified.
- Set the configuration file suffix type use `viper.SetConfigType("toml")` and is required if the configuration file does not have a suffix.
- To set the path or paths to  searched for the configuration file use `viper.AddConfigPath("path/to/conf/file")` as ahown in these examples:
    - `viper.AddConfigPath("$HOME/.appname")` // relative to `$HOME`
    - `viper.AddConfigPath(".")` // Use current (working) directory
    - `viper.AddConfigPath("/etc/appname/")` // common conf file location
    - `viper.SetConfigFile("./aspexv2.toml")` // Specific ref to path, name, type

 - You can call `AddConfigPath` multiple times to provide multiple search paths.
 
To read the file simply execute:

```go
err := viper.ReadInConfig() // Find and read the config file
if err != nil { // Handle errors reading the config file
	panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
```

### TOML Configuration File

Here is an example of the TOML configuration file, similar to those used in Examples 1 to 4. The TOML file is loaded from the users home directory.

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
softlicencetype   = "The Apache 2 License"
```

The code to read and use, the configuration file, is available on GitHub at https://github.com/dsbitor/GoViperTutorial and since this is the first example, it is also shown below.

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

As you can see a considerable amount of overhead associated with defining struct's to store the data from the TOML/JSON/YAML configuration types of file has been removed and the overhead of opening, reading and parsing the configuration file.

**If you just need to load and use values from a configuration file, you have everything required in the above example.**

## Example 2 - Multi-source Configuration Variables

The example above demonstrated the simple flexibility of using Viper with configuration files. Now let's create a more realistic example that uses multiple sources for configuration values and the associated Viper features. 

### Requirements 

- Default values to ensure that without the configuration file the program will run in production mode.
    - Data, results and logs files are written to the users home directory.
    - Minimal copyright, support, and software licence data will be set to default values in the code.
    - Working directory will be set to `.` if environment variable HOME is not set.
- From the configuration file:
    - The program will provide configuration values for a flexible, parameterized logging capability.
    - Ability to switch on various debugging components. 
    - Minimal application identification data will be set beyond those, and for overwriting those set by default.

### Home DIrectory - Useful Helper Function

Note that we use a package called `go-homedir` available at github.com/mitchellh/go-homedir. It is renamed to `homedir` in the import. Viper uses this facility to safely get the home directory no matter which OS you are using. `homedir` will be used to get the configuration file.

### Configuration File Design
The configuration file should be designed prior to building the code, because:
- This will ensure the requirements are accommodated.
- Identify configuration item names will appear in the code.
- Plan the value types needed. 
- Plan the hierarchy, as this adds understanding and naming clarity.
- Configuration file type also needs to be considered, using the criteria described above.

We will build on top of the TOML configuration file in Example 1 above, although the configuration could be equally well represented in a YAML or JSON file. The file is `aspexv2.toml` and is in the examples directory with the `aspex2.go` source file.

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

Viper offers six different types of sources for configuration values, and another value can override an existing value. Viper uses the following precedence order for defining the final value. Each setting source takes precedence over the setting below it:

- explicit call to Set (programmatic setting)
- flag value
- environment variable
- configuration file value
- key/value store
- default

## Default Values
Better programming practice requires that variables that control program operation should have default values to predict the program's execution. Missing values in a configuration file should not result in unpredictable results.

The definition of default values using Viper is simplified and integrated into the other subsequent setting capabilities such as environment variables, configuration files and flags.

### Adding Default Values 

Example 2, `aspexv2.go` and its TOML configuration file `aspexv2.toml` show how to incorporate default values into the configuration environment. 

Code in  `aspexv2.go` shows eight default configuration values declared and values set.

## Example 3 - Environment Variables

Access to operating system environment variables is a useful feature for values that have been set for use across the whole operating environment, important infrastructure, or an application.

The requirements here are:

- Get working directory from environment variable PWD.
- Get the hostname from environment variable HOSTNAME.

 Note that the variable `hostname` does not have an equivalent value in the TOML configuration file `aspexv2.toml`. The environment variable HOSTNAME does not exist on my Mac OS as configured. Example 5, directory `mpm2/cmd/root.go` shows how to get and set `hostname`. 

### Why Use Environment Variables

An environment variable is a variable whose value is set outside the program, typically through functionality built into the operating system, microservice or a program. An environment variable consists of a name/value pair.  Any number of variables may be created and are available for reference at a point in time. They are often used and set at:

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

What these have in common are their data values change infrequently, and the application logic treats them like constants, rather than mutable variables.

### Viper - Access to Environmental Variables

Viper makes it easy to set configuration values from environment variables. 

The function `viper.AutomaticEnv()` will read all available environment variables and make them available to the program. They would override variables with the same name in a configuration file  (YAML, TOML, JSON etc.) and be reassigned by Viper.

Environment variable names can have unusual capitalization and naming conventions. These names do not make them easy to use in Go programs, where variable names are expected to follow Go naming conventions. It is easy to remedy by using `viper.BindEnv` which takes one or two parameters. Use the two-parameter version to map the Go viper configuration name to the external environment variable name, thus providing consistency of naming across the configuration which is compatible with the Go naming standards.  

Example 3, `aspexv3.go` and its TOML configuration file `aspexv3.toml` show how to incorporate Environment variable values into the configuration environment. 

## Example 4 - A Real-life Example

The flexibility of Viper is demonstrated above.  Now let's create a real-life example that uses all of the significant sources for configuration values and their associated Viper features. 

This example deals with flags and getting flag values into a set of configuration values managed by Viper. 

## Using Flags

Viper can bind to Command Line Interface (CLI) flags. Flags offer a method of changing configuration variables at run time. Specifically, Viper supports Pflags, which are also used by the Cobra library. The pflag package is a drop-in replacement for Go's `flag` package and implements POSIX/GNU-style such as  `-d`, that give one-letter shorthand flags as well as long-name flags such as `--debug`.  If programming to POSIX/GNU standards, this is important, where there is an expectation that GNU flags are available for use, such as in Unix and other operating system utilities. 

**If you are not fully aware or don't understand CLI components and the role of flags in the command line, read 'CLI Briefly Explained' in Appendix B.**

Flags are handled similarly to Environment Variables. The configuration variable is bound to a flag.  Like `BindEnv`, the value is not set when the binding method is called, but when it is accessed.  By enabling this feature, it means you can bind as early as you want, even in an `init()` function. For individual flags, the `BindPFlag()` method provides this functionality.

Example 4 shows how to use flags to change or set Viper configuration values.

### The flag and pflag Packages

#### The flag Package

Flag (and pflag) set-up encompasses two steps:

1. Declare each flag's structure.
2. Issue a  `flag.Parse()` to instantiate the flag structure(s) with the current values supplied when the program was invoked.

The Parse() function will continue to parse flags that it encounters until it detects a non-flag argument. The flag and pflag packages makes these non-flag arguments, available through the `pflag Args()` and `Arg()` functions.

The flag declaration can be made in several forms, two of which are:

    var ip *int = flag.Int("flagname", 1234, "help message for flagname")

    var flagvar int
    func init() {
        flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
    }

The three parameters are:

- "flagname" - the flag name consisting of an alphanumeric string
- 1234 - the default value if no value is provided with the flag
- "help message for the flag name"

To get the values from the flags when all flags have been declared use flag.Parse().

The flag package provides methods also to parse non-flag parameters. The remaining parameters, entered when the program was invoked,  are provided using the `flag.Args()` function. It returns a slice of strings with the parameters not parsed as flags.

#### The pflag Package Capability and Example

The equivalent `pflag` functions to the flag functions shown above are `pflag.Int` and `pflag.IntVar`.

Or you can bind the flag to a variable using the Var() functions as shown above using `flag.IntVar`. The alternative functions in the pflag package is `pflag.IntVar` and `pflag.IntP`.

Refer to the example code in the source file `aspexv4.go`. This example uses the `pflag` package. The file contains only the code necessary to show flag definition, Viper default and flag overide code, and output showing flag value states at various points along the source code. To run this example use the two terminal commands below:

    go build aspexv4.go 
    ./aspexv3f --debug=true --linelist=42
    
Flags in `pflag` are set-up in the same manner as the flags description above. The significant difference is that pflag allows the implementation of POSIX/GNU-style short-form flags such as, `-d`, `-l66`.


Short-flags are implemented by using an alternate version of the flag declaration, where `flag.Int()` uses `pflag.IntP()` which adds a `P` suffix to the function name. Note that `pflag` also implements `pflag.Int()`, which acts exactly the same as the `flag` version. Two examples of the extended notation declarations are:

    pflag.IntP("linelist", "l", 5, "Number of lines to list")
    debug := pflag.BoolP("debug", "d", false, "Switch debugging off/on")

Note that in each of the above declarations, the second parameter `"l"` and `"d"` respectively are the short-form names. All other parameters are the names as for flag parameters as described above.

`pflag` has the same flag parsing mechanism as the `flag` package, and is called after all flags have been declared. 

    // Call pflag.Parse() after all flags have been defined
    pflag.Parse()

Executing the program `aspex4.go` allows flags to be entered in several ways:

- --debug=true
- --debug
- -d=true
- -d true
- -d
- --linelist=42
- --linelist 42
- -l=42
- -l 42

Note that the -d and --debug flag switches the boolean's value to the alternate of the current value. 

**See https://godoc.org/github.com/spf13/pflag for full pflag detailed developer documentation.**

### Viper's Use of Flags

**Viper uses the pflag package** and allows command-line flags to be bound to configuration variables. 

	viper.BindPFlags(pflag.CommandLine)
	
Expanding on the examples above, and using the long-name as the configuration variable key, we can connect the flag to the configuration variable using the two viper functions:

- `viper.BindPFlags(pflag.CommandLine)` binds all the command line flags as viper configuration variables. 
- `Viper.BindPFlag("debug", CommandLine.Flags().Lookup("debug"))` binds a specific key to a named pflag.
	
See Example 4 uses the program `aspex4.go` as an example of `viper.BindPFlags` usage.

## Getting Values From Viper

Most of our effort above was associated with Viper's getting configuration variables using various Viper functions. We must also be able to access those values efficiently and effectively. You have already seen some of these access mechanisms, but we will describe all the significant capabilities in this section.

### Getting Values by Key Name

**Key (names) are case insensitive.** 

Viper documentation states "Viper merges configuration values from various sources, many of which are either case insensitive or uses different casing than the rest of the sources (e.g. env vars). In order to provide the best experience when using multiple sources, the design decision was made to make all keys case insensitive."

We recommend that you use all lower case names for consistency and ease of reading when you use a key.

You only need to provide the key  for the value you want, and the type you expect that value to be. Viper provides seventeen GetType methods where  'Type' is the data type such as Int, String, Float, Bool etc.. Here is the list of the most important functions:

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

Each Get function will return a zero value if it’s not found, not an error. If you are not sure the key exists, then use `IsSet` described below along with the three other important access functions:

- Get(key string) : interface{}  - This function returns the string value for the provided key.
- IsSet(key string) : bool  - Function returns true if the specified key exists.
- AllKeys() : []string - returns all available keys as string slice.
- AllSettings() : map[string]interface{} - Function returns a map of keys and strings representing the values of all defined flags.

##Notable Points Not Included in Examples

The following capabilities are noted but are not shown in examples.

Steve Francia the designer and author of Viper (and Cobra) stated in a chat are "Typically with Viper you just use things like `Viper.Get("key")` to get the values you need out and don't marshall it into a struct. There's nothing wrong with the struct approach and sometimes is the right way to do it, but it's definitely more upfront work and less flexible." 

### Marsalling
Marshalling of configuration variables allows them to be written out:
- As a new or updated configuration file, where you want to store all modifications made at run time
- As a structure for use either within the program
- As output in some form other than as a configuration file

### Unmarshalling

Viper offers an option of un-marshalling all, or a specific value to a struct, map, etc. Normally using the named configuration variable form offered by Viper is more simple and clear for programming, but there could be some unique cases where it is useful. 

## Example 5 - A Real-life Example With Cobra Defined Flags

This example shows how to integrate Example 3 with Cobra based command definitions with associated flags. Finally, this example will ensure configuration parameters will be in place before any other program activity, by using the function `init`.

### Requirements

- Create Cobra commands and associated flags.
- Use some flags to change/set Viper configuration values.
- Saves the configuration to a new file after adding/changing a value from a flag.
- Demonstrates accessing configuration values.
- Use the init function in Cobra for Viper configuration handling.
- Get a working directory from environment variable PWD.
- From the configuration file:
    - The program requires a flexible, parameterized logging capability.
    - Ability to switch on various debugging components.
    - Application identification data.
- Default values to ensure that without the configuration file, the program will run in production mode.
    - Data, results, and logs files written to the users home directory.
    - Minimal copyright, support, and software licence, default configuration values embedded in the code.
    - The working directory will be set, but will be overwritten by a flag variable if present.
- Programmatically set the computer name.

## Saving Configuration Files

Viper offers flexibility in saving a configuration file, which can be very useful where program or application settings need to preserve from one execution of a program to the next. 

Formats for saving configuration files are JSON, TOML, YAML, HCL, and INI. Saving is as simple as configuring the ability to overide existing file types, location and file name, or creating a new file, with a specific location and name.

Four functions, each with its capabilities, available for writing the configuration file are:

- WriteConfig - writes the current viper configuration to the predefined path, if it exists. Errors if no predefined path. Will overwrite the existing config file, if it exists.
- SafeWriteConfig - writes the current viper configuration to the predefined path. Errors if no predefined path. Will not overwrite the existing config file, if it exists.
- WriteConfigAs - writes the current viper configuration to the given file path. Will overwrite the given file, if it exists.
- SafeWriteConfigAs - writes the current viper configuration to the given file path. Will not overwrite the given file, if it exists.

Examples from the documentaion show:

    viper.WriteConfig() // writes current config to predefined path set by     'viper.AddConfigPath()' and 'viper.SetConfigName'
    viper.SafeWriteConfig()
    viper.WriteConfigAs("/path/to/my/.config")
    viper.SafeWriteConfigAs("/path/to/my/.config") // will error since it has already been written in statement above
    viper.SafeWriteConfigAs("/path/to/my/.other_config")

`viper.WriteConfigAs` is shown in Example5, when used with Cobra in the `initialize.go` command to save an altered configuration. See line:

    viper.WriteConfigAs("./writtenconf.toml")  // Write updated configuration values to file

### Accessing Configuration Values using Viper Functions

Using Example 5 above all of the statements below are valid:

```go
if viper.IsSet("wrkdir") {fmt.Printf("Working Directory Config Variable Exists \n") }

workingdir := viper.Get("wrkdir")
fmt.Printf("Working Directory = [%s]\n", workingdir)

# Getting a value from within a structure
devemail := viper.Get("app.developer.emailaddr")
fmt.Printf("Developer Email = [%s]\n", devemail)

# Getting a defined type of value from within a structure
hopperBug := viper.GetBool("devparms.debug")
fmt.Printf("Config key = devparms.debug  Config type = %T Config value = %v \n",  hopperBug, hopperBug )

AllConfig:= viper.AllSettings() 
for k, v := range AllConfig { 
    	fmt.Printf("Config key = [%s] Config value = [%s]\n", k, v)
}
```

## Cobra Integration

Cobra and Viper work well together, as long as the relationship is understood, and the steps to integrate the two components are taken in an orderly manner. If you are not familiar with Cobra, read the Copra Tutorial, identified at the beginning of this document in the section 'Relationship to Cobra'.

Complete the Cobra work to create the command and sub-commands required for the project first. Make sure you are using Go Modules, and a source code management system, such as Git or SVCS, are very useful to manage the source code as you make changes.

Part of the Cobra design decisions is to identify the flags for each of the defined commands and sub-commands. Of those identified command flags, separately identify those flags that will modify the program configuration variables.  Those configuration flags will be connected to the configuration variables using Viper functions. 

Cobra provides a linkage to Viper services in the `root.go` file created when the Cobra command `cobra init` was executed and provides similar capability in each command and sub-command definition, created when `cobra add` commands are executed. 

### Cobra Defined Flags

Any flag defined within a Cobra command environment, generated as a Go source file, including `cmd/root.go` can be linked to a Viper configuration variable. A flag defined in Cobra is linked to an environment variable using the function  `viper.BindPFlag`. You  have seen a similar binding when dealing with pflags and configuration variables. This binding is enabled because Cobra flags are just 'enhanced' pflags.

This small example implements three Cobra flags linked to Viper configuration variables. Let us assume we want to control trace messages through all Cobra commands (a global flag) in a multiple command environment. The second and third flags will be active within only one command, and will change the working directory, or set the number of lines of output on one page, by changing the  appropriate configuration variable to a new value.

1. In the Cobra component, a boolean flag switches tracing messages on and off.  It is a global flag, implemented in the file `cmd/root.go`. The flag is `--trace` or short name `-t`, so it can be used with either a long or short name. The flag trace will have an equivalent configuration value which will be read from a TOML configuration file, and has a default configuration setting of false.
2. The second configuration flag used in this example is, a local flag to the command initialize. The initialize command creates the run time and session computing environment for the run command. A flag provides access to change the working directory, and when present,  will override working directory configuration values already set up by default or via the configuration file values. This flag will only be available in the `initialize` command.
3. A flag will be implemented only for the run command, which will control a run time variable, and will over-ride the configuration file for the number of print lines per page.

### Cobra Implementation

**The source files for this example are a Cobra environment, with Cobra's particular layout for files. If you are not familiar with Cobra, and how it works, or have never created a Cobra environment, please read the Tutorial 'How to Use Cobra' and run the examples in that tutorial first, and they are available at https://github.com/dsbitor/GOCobraTutorial**. 


#### Example mpm1

mpm1 implements the first three components of a Cobra based comprehensive solution. It implements configuration variables from: 

- defined default values
- a configuration file
- environment variables

The added or changed values of the variables are displayed after each code component is executed.

1. **Default Values** are implemented using the `setConfigDefaults` function. `setConfigDefaults` has two components. First a call to `homedir.Dir` to get the home directory no matter which operating system is being used.  Second, calls to `viper.SetDefault` for each required default value. `setConfigDefaults` is called by `initConfig`, which is called by `init`. Thus all code within `init` will  be executed before any code in `main`, ensuring that all configuration values are set before running application functions.
2.  **Configuration File** is dealt with by infrastructure provided by Cobra when the command `cobra init project-directory` creates  `project-directory/cmd/root.go`. Within `cmd/root.go`, the code, starting with the comment `// Configuration File defined by flag?` is in  the function `initConfig`. To this code was added, the Viper configuration file set-up function calls:       `viper.AddConfigPath(home)`   
`viper.AddConfigPath(".")`   
`viper.SetConfigName(".mpm2")`      
This was followed by the code to read the TOML configuration file `.mpm*.toml`, consisting of:  
`if err := viper.ReadInConfig(); err == nil {`    
	`fmt.Println("Using config file:", viper.ConfigFileUsed())`    
`} else {`    
	`fmt.Println("Failed to read configuration file:", viper.ConfigFileUsed())`   
}`   
3. **Environment Variables** are set up next, in the code following the configuration file code in the function `initConfig`. The environment variables are selected using the Viper function 
`viper.BindEnv(_configValueName_,_environmentValueName_)`. When all have been identified a call is made to `viper.AutomaticEnv()` to initialize environment variables that match config variables.

#### Example mpm2

mpm2 adds to the mpm1 example and implements a full Cobra based comprehensive solution. It implements configuration variables from: 

The mpm2 example adds:
- global and local flags associated with commands
- setting configuration values programmatically
- writing a changed configuration file

The mpm1 description above covers the code in Example mpm1. The code is copied to mpm2, and will not be explained further.

The added or changed values of the variables are displayed after each code component has been executed.

**Command Global and Local Flags** 

Global and local flags are associated with Cobra commands. The flag definitions can use either `flag` or `pflag` notation, but the examples in `mpm2` use `pflag` notation. A flag can be defined in `root.go` or any other go source file generated by Cobra in the `cmd` subdirectory.

In `mpm2` flags are defined in `root.go`, `initialize.go`, and `run.go`. The flags defined in `root.go` are global flags, and can be used with any defined command.  In the mpm2 example, they are available to the commands `initialize` and `run`. The two pflags defined in `cmd/root.go` are, and allow:

 - `--trace` or `-t` which support a boolean value, and are used to control trace messages in the programs
 - `--config` (it has no single character short name) and supports a string that is a path to a valid configuration file, offering a way to read a specified configuration file at run-time.     

In source file `initialize.go` a local flag is defined. The flag named `--workdir`, or `-w`, supports a string value which changes the working directory at initialization. initialize.go also writes a new updated configuration file called after using the value entered using the `--workdir` flag altering the working directory path, and all the preceding configuration file definitions or changes.  This update ensures the configuration parameters in place at run-time are saved.

A local flag defined in run.go. It is called `--linecount` and has the single character `-l` short name. It overwrites the value provided in the configuration file. 

**Note that after defining the flags, they must be connected to configuration values.**   Since `--trace`, `--workdir`, and `--linect` are all planned to change configuration values, and they must be connected to the appropriate Viper configuration items. The configuration item keys are respectively devparms.trace, wrkdir, and linect.  To connect the trace flag to the configuration variable use the function `viper.BindPflag`,  which takes two parameters, the specific full key for the Viper configuration item such as `devparms.trace`, and the full pflag reference `rootCmd.PersistentFlags().Lookup("trace")`.  `Lookup` is a pflag function that finds the appropriate configuration variable and appears in all calls of this type. 

The three pairs of program lines for defining flags trace, workdir, and binding them them to the viper configuration variable are:

root.go

    rootCmd.PersistentFlags().BoolP("trace", "t", true, "Switches on/off Trace messages in mpm")
    viper.BindPFlag("devparms.trace", rootCmd.PersistentFlags().Lookup("trace"))
    
initialize.go

    	initializeCmd.PersistentFlags().StringP("workdir", "w", "", "Input Run-time Path to Working Directory ")
	    viper.BindPFlag("wrkdir", initializeCmd.PersistentFlags().Lookup("workdir"))
	    
run.go

    	runCmd.PersistentFlags().IntP("linecount", "l", 65, "Number of Lines on a Printed Page ")
	    viper.BindPFlag("linect", runCmd.PersistentFlags().Lookup("linecount"))
	    
	
As shown in the examples above, to keep clear which flags are bound to Viper configuration key/values, it is wise to declare the flag, and on the next line the associated `viper.BindPflag` statement.

You will notice if you run the code. When you enter a flag with a parameter that modifies the Viper configuration, the displayed values will show the configuration value entered with by the flag at all points where the configuration values are displayed. 

**Programmatically Setting Configuration Values** 

Configuration values can be set programmatically at any time, and have they have the highest precedence. Two examples provide, in the code in example contained in the directory `aspex5/mpm2`. The hostname of the computer is extracted and stored in the configuration variable `hostname`, and is set programmatically in the function `initConfig` in `root.go` where the code below is executed:

    // Setting a configuration variable programatically
    // Get and set hostname
    hname, _ := os.Hostname()
    viper.Set("hostname", hname)

A second example of a programmatic setting is shown in `run.go`. The code shown below,  setting the logging timestamp variable to `highres`

    // set new configuration variable logging.tstamp to highres (nanosec)
    viper.Set("logging.tstamp", "highres")

**Writing a  Configuration File**

The file is simply written in `initialize.go` using:

    viper.WriteConfigAs("./writtenconf.toml")

### Executing the Cobra Examples in mpm1 and mpm2
 
To **compile** the Go code in directories `mpm1` or `mpm2` use **`go install mpm1`** or **`go install mpm2`**. The command `go install` which does a compile/link/save-binary-to-binlib can be executed from the root of `mpm1` or `mpm2` depending on the code you are working. The resultant binary program will typically, by default, be stored in the `$GOPATH/bin/`. If you want to be sure list that directory after getting a clean-compile from `go install`.

To execute the code run at the terminal prompt, issue  **mpm1** or **mpm2* with appropriate command and flags. For example **for mpm1 use**, `mpm2 initialize` as flags are not implemented in this partial version, and **for mpm2 use** `mpm2 initialize --trace=true` . Note that issuing just the bare **mpm1 or mpm2 will produce help information as output**.

The descriptive output code can be safely commented out without affecting functionality.

## Appendix A - Further Notes `INI` Files

Since some of the more complicated `ini` files are key-value pairs with groups, delimited by square `[]` or curly `{}` brackets, and they do not work with Viper, you could build your own parser if essential,  using the `io.Reader` feature to read the files. 

Viper uses the package `gopkg.in/ini.v1` to read and write ini files, and which as of October 2020, is supported and updated. It is described as 'The package `ini` provides INI file read and write functionality in Go'. 

Note that Viper will read configuration data in INI format, but it is not clear from the documentation what support is provided for writing INI files from configuration data.

If Viper does not support a particular feature direct use of `gopkg.in/ini.v1` with some custom code, may help resolve the problem. If you want to see what `gopkg.in/ini.v1` can read and write, see https://github.com/go-ini/ini/blob/master/testdata/full.ini.

##  Appendix B - CLI Briefly Explained

A program executed via the Command Line Interface  (CLI) can consist of a program name, commands, flags and arguments. Taking an example from the GO compile and execute environment, the command `go install 'package'` where package is the folder containing the package required. 

In the Go documentation, the command is described as `go install [-i] [build flags] [packages]`. To complete this example, here are the components: 
- `go` is the  named program to be executed
- `install` is the command
- `-i` is a flag, and packages is an argument
- `build flags` are numerous additional flags and are listed if you execute the command `go help build` 

The command `go help build`, also serves as a second example of a CLI action, where go is the program, help the command, and build the sub-command, but there are no flags or arguments.

Groups of flags or a single flag are sometimes described as options.   

Flags provide further control, change predefined parameters, or supplying additional parameters to the program or command. Examples of flags are `-e`, `-V 42`, `--print`, `--font courier` etc. Flags starting with `--` are sometimes referred to as LongOpt format. Flag values are `42` and `courier`. Flag values can be strings, integers, real values, booleans, etc. 

If you are not sure about all the basic syntax, rules and conventions for CLI flags (options), here is a summary based on the POSIX standard, GNU conventions and flags often seen in Mac OS X, Unix and MS Windows environments, and was taken from GNU, where flags are refereed to as options:

- An option is a hyphen followed by a single alphanumeric character, like this: -o.
- An option may also be a double hyphen (GNU style flag) followed by an alphanumeric string of two or more characters, like --list,  --help or --filetype.
- An option may require an argument (which must appear immediately after the option); for example, -o argument or -oargument or --list 60.
- Options that do not require arguments, and are a single alphanumeric character, can be grouped after a hyphen, so, for example, -lst is equivalent to -t -l -s.
- Options can appear in any order; thus -lst is equivalent to -tls.
- Options can appear multiple times.
- Options precede other non-option arguments: -lst non-option.
- The -- argument terminates options.
- The - option can be used to represent one of the standard input streams.

## Appendix C - Help Information Output by Programs

Typically, your help output should include:

1. Description of what the app does
2. Usage syntax, which:
    - Uses [options] to indicate where the options go
    - arg_name for a required, singular arg
    - [arg_name] for an optional, singular arg
    - arg_name... for a required arg of which there can be many (this is rare)
    - [arg_name...] for an arg for which any number  of values can be supplied  
    note that arg_name should be a descriptive, short name, in lower, snake case
3. A nicely-formatted list of options, each providing:
    - having a short description
    - showing the default value, if there is one
    - showing the possible values, if that applies
4. Note that if an option can accept a short form (e.g. -l) or a long form (e.g. --list), include them together on the same line, as their descriptions will be the same
5. Brief indicator of the location of config files or environment variables that might be the source of command-line arguments, e.g. GREP_OPTS
6. If there is a man page, indicate as such, otherwise, a brief description of where to find more detailed help.

**It is good practice to accept both -h and --help to trigger the help message and that the program should also show the help message if the user messes up the command-line syntax, e.g. omits a required argument.**