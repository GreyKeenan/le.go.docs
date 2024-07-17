
# doc_HowToWriteGoCode

	about organizing go code projects
	following [this resource](https://go.dev/doc/code)

## Code Organization

packages

	specified with "package nameGoesHere" at top of .go file

	a single package can be split into multiple source files /aka/ can specify the same package in multiple source files
		in the same directory though
	functions, types, vars, and constants are shared between source files of the same package

modules

	collection of related packages grouped/released together
	generally one per repository

	each module has a "go.mod" file in the root
		documents/manages dependencies of the module

		declares the "module path" - import path prefix for all packages within the module

	subdirectories are part of the module recursively, until hitting another directory with a "go.mod" file (another module)

	the module path tells go command where to look to download

	"import path" - a string used to import a package
		consists of module-path + '/' + subdirectory within module

		stdlib packages have no module path prefix


## Your First Program

(walks through steps of creating modules)
(see repo/tut_createGoModule as well)

install location

	"go install" command installs to $HOME/go/bin/ by default

	can change by setting GOBIN or GOBATH environmen variables
		"go env -w GOBIN=/pathgoeshere/"
	unset:
		go env -u GOBIN


## Import packages from your module

any funcs/variables named with a capitalized first letter are exported automatically / are public
lowercase variables are private to that package


## Importing packages from remote modules

if remote module path, go will fetch it with "go mod tidy" / adds requirements to go.mod

module dependencies are downloaded to GOPATH/pkg/mod
	clean with "go clean -modcache"


## Testing

("go test" cmd & "testing" package in stdlib)
(demonstrated in repo/tut_createGoModule)
