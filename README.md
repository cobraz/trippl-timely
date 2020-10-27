# Copy Timely events to Tripletex

This utility is written in Go and features functionality to copy Timely events to Tripletex, a Norwegian accounting software.

## Install

```shell
brew install cobraz/tools/trippl-timely
```

**Notes**: The library is not tested on Linux or Windows. There are [executables available](https://github.com/cobraz/jira-to-tripletex/releases/latest) at every release >1.0.3

## Help

```shell
> trippl-timely --help
NAME:
   Trippl Timely - A new cli application

USAGE:
   main [global options] command [command options] [arguments...]

DESCRIPTION:
   Send Timely events to Tripletex

COMMANDS:
   get-config        
   set-config        
   get-timely-token  
   add-timesheet     
   help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Contribute

Please, oh pretty please do contribute! If you feel this helps you out, but you want to increase the quality of this software, please submit pull requests. 🎉 PS: There are no tests 🤷
