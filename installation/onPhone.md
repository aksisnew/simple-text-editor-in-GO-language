
> CAUTION : i do this on my secondary phone dedicated for development and not on main phone with personal stuff

# 1. Update Termux packages
pkg update && pkg upgrade -y

# 2. Install Git and Go
pkg install git golang -y

# 3. Check that Go is installed
go version

# 4. Create the project directory
mkdir simple-editor
cd simple-editor

# 5. Create the Go module
go mod init simple-editor

# 6. Put your source files in this directory:
#    main.go
#    editor.go
#    open.go
#    save.go
#    exit.go
#    change-extension.go
#
# If you already have the files somewhere, copy them here.

# 7. Download/verify Go dependencies
#    (There are currently no external dependencies, but this is safe)
go mod tidy

# 8. Compile the editor
go build -o simple-editor

# 9. Run it
./simple-editor