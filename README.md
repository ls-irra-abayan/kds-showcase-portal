# Golang Service Template

This repository should contain the files required to get started with
a go service on lightspeed-hospitality.

The goal is to make it easy for developers to start experimenting with
go projects.


<br/>

## Getting Started With a New Project

1. Edit src/go.mod to reflext the module name and requirements
1. Edit the .circleci/config.yml - uncomment the deploy section and replace "go-service-template" and "Serivice Name" with the actual values
1. Edit the .github/CODEOWNERS file
1. Move the helm chart in charts/go-service-template to a folder reflecting the service name
1. Edit Chart.yaml - replace "go-service-template" and "Serivice Name" with the actual values
1. Edit this README.md

## Local Development

### Installation

1. Clone this repository.
2. Change directory to `src`.
3. Install dependencies: `go get`.

### Testing

Run tests:
```
cd src
go test ./...
```

### Test coverage

```
cd src
go test -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

<br/>

## Repository Structure

- `src` contains package files to be shared with executables.

<br/>

## Deployment with fleet

If a fleet environment linked to an inventory ephemeral environment is desired,
replace step 3 of
[Inventory Ephemeral Environment with Fleet SBX](https://confluence.atlightspeed.net/pages/viewpage.action?spaceKey=LSK&title=Inventory+Ephemeral+Environment+with+Fleet+SBX)
with the steps below.

1. Create a PR with commands to create an equivalent branch and PR in fleet to
trigger fleet deployment.

```md
!fleet:branch=<branch-name>
!fleet:ttl=<number-of-hours> hours
```

<img src="./doc/images/deployment/pr-inventory-create.png" alt="pr-inventory-create" width="600">
<img src="./doc/images/deployment/pr-inventory-open.png" alt="pr-inventory-open" width="600">

> Only the branch command is required.
> See [Fleet PR commands](https://docs.lsk.lightspeed.app/Development/Deployments/Fleet/How-To/07-pr-commands/).

2. Wait for the [CircleCI workflow for the branch](https://app.circleci.com/pipelines/github/lightspeed-hospitality/go-service-template)
to complete. Re-run pipeline if it finished before the PR is created.

> Ensure that there is only one PR opened per branch. The utility to automate
> fleet deployments can only parse a single PR.

3. Verify that the PR is created in [fleet repository](https://github.com/lightspeed-hospitality/fleet).

<img src="./doc/images/deployment/pr-fleet.png" alt="pr-fleet" width="600">

> You may get the link from the `fleet/triger-update` CircleCI workflow, under
> `Link to fleet PR` step.

4. Wait for the [Fleet CircleCI workflow](https://app.circleci.com/pipelines/github/lightspeed-hospitality/fleet) to complete.

<img src="./doc/images/deployment/circleci.png" alt="circleci" width="600">

<br/>

## Logging

Logs are piped into [logz.io](https://logz.io/) under the following workspaces:
- hospitality-k-dev
- hospitality-k-prd

Logs specific to the repository may be filtered under the field
`kubernetes.container.name` with the container name.

You may also open the template with name `Template Name`.
