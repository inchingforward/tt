# tt - A simple command-line time tracker

## Usage

### Starting a Task
    $ tt start "Stub project functions"
    Started task "Stub project functions"

### Starting a Task while one is in progress stops a task
    $ tt start "Add display values"
    Task "Stub project functions" ended
    Started task "Add display values"

### Stopping the currently running task
    $ tt stop  
    Task "Add display values" ended

### Listing tasks 
    $ tt
    Task                                     Started          Ended            Time    
    Stub project functions                   08/13/2017 22:38 08/13/2017 22:48 10m29s
    Add display values                       08/13/2017 22:55 08/13/2017 23:08 13m22s                
    ...

## Storage/Data Model

The data model is a list of task objects with 3 attributes:

    name    :  The name of the task
    started :  Task start time (in unix epoch)
    ended   :  Task end time (in unix epoch)

The data is stored as json in a file named `tt.json` in the home directory.  If tt does not find the file, it will create one.