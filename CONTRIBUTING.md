# Contributing to Gophers Go SDK

## Creating a new Release

* Add the changelod of the new tag/release in `CONTRIBUTING.md` file
* Create a new tag: `git tag v0.0.5` and push it in the repository `git push --tags`
* A GitHub action will be executed that will read the `CONTRIBUTING.md` file and generate a new GitHub Release
* The pkg.go.dev documentation website will be forced to index the new tag
