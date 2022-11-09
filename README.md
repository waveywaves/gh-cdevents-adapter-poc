# CDEvents Adapter PoC for Github

This adapter is a Proof of Concept to show that Github Events from the PR can be converted to CDEvents using a Github Action. The process involves receiving an event from Github through a Github Actions Workflow and then mapping it to the respective CDEvent types.

Run the following commands to test this locally

```
export GITHUB_CONTEXT="$(cat event.json)"
go run gh-cdevents.go
```