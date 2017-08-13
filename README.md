# tt - A simple command-line time tracker

## Usage

### Starting a Task
    $ tt start "Adding something"
    Started task "Adding something"

### Starting a Task while one is in progress
    $ tt start "Doing something else"
    Stopped "Adding something", total time: 02:24:45
    Started task "Doing something else"

### Stopping the currently running task
    $ tt stop  
    Stopped "Doing something else", total time: 01:22:33

### Listing tasks 
    $ tt
    #  Name                            Started          Ended            Time     
    1  Doing something else            08/13/2017 14:00                  05:00:23 
    2  Adding Something                08/13/2017 13:22 08/13/2017 14:22 04:20:22     
    ...

## Storage/Data Model

The data model is a list of objects with 3 attributes:

    name    :  The name of the task
    started :  Task start time (in unix epoch)
    ended   :  Task end time (in unix epoch)

The data is stored as json in a file named `tt.json` in the home directory.  If tt does not find the file, it will create one.