
loosely following [this introductory tutorial](https://go.dev/doc/tutorial/create-module)

---

info overview:
- modules vs packages
- importing from modules, locally or not
- creating local modules
- minimal use of math/rand
- minimal use of error package
- minimal use of fmt package
- minimal use of slices & arrays

---

- package: directory/collection of .go files
- module: collection of package(s)
	- generally, you have one module per repository

[helpful stackoverflow response](https://stackoverflow.com/a/57314494)

---

Testing

	Go has built-in support for testing
	implemented through "testing" package, naming conventions, and cli command

	naming conventions
		filenames ending with "_test.go" tells cli tool that there are test funcs inside
		testing functions are prefixed by "Test"
			[doesnt say if this is purely convention or necessary]

	testing functions' first parameter must be: (*testing.T) type

	shows testing functions as implemented within the same package they are testing

	run tests in module dir using "go test" command

---

build/install

"go build" in module directory builds binary there.
"go install" in dir builds binary in go installations directory
	find the directory w/: go list -f '{{.Target}}'
	may need to add it to PATH
