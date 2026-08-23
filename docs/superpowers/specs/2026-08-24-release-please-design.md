# Release Please Design

## Context

AronaUnflatd currently publishes GitHub releases manually and does not attach
compiled binaries. The repository is transitioning from `develop` to a
trunk-based workflow on `master`. This setup will be reviewed and merged through
`develop`, then become active after `develop` is merged into `master`.

## Goals

- Let Release Please manage release pull requests, changelogs, semantic
  versions, `vX.Y.Z` tags, and GitHub releases from Conventional Commits.
- Publish native CGO-enabled binaries for Linux x86_64 and Windows x86_64 only.
- Include the released version in the executable.
- Use the repository `GITHUB_TOKEN` without adding a personal access token.
- Keep the configuration valid after `develop` is removed.

## Non-goals

- Building ARM, macOS, or 32-bit artifacts.
- Publishing containers or language-package registries.
- Signing artifacts or generating an SBOM.
- Running release automation from `develop`.

## Branch transition

The workflow listens only for pushes to `master`, and Release Please explicitly
targets `master`. Consequently, merging this setup into `develop` does not open
a release pull request. Once `develop` is merged into `master`, the next
`master` workflow run begins release management. Removing `develop` afterward
requires no release configuration changes.

## Release lifecycle

One workflow owns both release creation and artifact publication:

1. A push to `master` runs `googleapis/release-please-action` in manifest mode.
2. Conventional Commits accumulate in a Release Please pull request targeting
   `master`.
3. Merging that pull request creates a `vX.Y.Z` tag and GitHub release.
4. The action's `release_created`, `tag_name`, `version`, and `sha` outputs gate
   two native build jobs in the same workflow run.
5. Each build job checks out the released SHA, packages its platform binary,
   creates a SHA-256 checksum, and uploads both files to the existing release.

Keeping publication in the same workflow is required because events created by
the default `GITHUB_TOKEN`, including Release Please's tag, do not start a
second workflow. This design therefore avoids an extra token and its associated
permissions.

## Release Please configuration

The repository uses manifest mode with:

- Root package path `.`.
- Release type `go`.
- Component name `arona-unflatd`.
- Component-free tags so tags remain `vX.Y.Z`.
- Initial manifest version `0.0.2`, matching the latest existing release.
- Generated changelog path `CHANGELOG.md`.

The workflow grants Release Please only the repository permissions it needs to
write contents and manage release pull requests and their labels. Repository
settings must allow GitHub Actions to create pull requests.

## Release artifacts

Builds run natively so `go-tree-sitter` and its CGO dependency use the compiler
provided by each hosted runner:

| Platform | Runner | Environment | Archive |
| --- | --- | --- | --- |
| Linux x86_64 | Ubuntu x86_64 | `CGO_ENABLED=1`, `GOOS=linux`, `GOARCH=amd64` | `arona-unflatd_VERSION_linux_amd64.tar.gz` |
| Windows x86_64 | Windows x86_64 | `CGO_ENABLED=1`, `GOOS=windows`, `GOARCH=amd64` | `arona-unflatd_VERSION_windows_amd64.zip` |

Each archive contains the executable, `README.md`, and `LICENSE`. Each archive
has a neighboring `.sha256` file. Uploads use `gh release upload` with the
workflow's `GITHUB_TOKEN`.

## Version metadata

The hard-coded development version becomes an overridable package variable with
the default value `dev`. Release builds pass the Release Please version through
Go's linker using `-X main.version=VERSION`. A focused executable-level test
builds the command with a synthetic linker version and verifies that the CLI
reports it, protecting the release contract rather than the variable's source
representation.

## Failure behavior

- No platform build runs unless Release Please reports that it created a
  release.
- A failed build, package, checksum, or upload step fails only its platform job
  and is visible in the release workflow.
- Uploads use replacement semantics so a rerun cannot fail solely because an
  asset with the same name already exists.
- No credentials are embedded in files or command arguments; GitHub supplies
  the short-lived workflow token through the environment.

## Verification

- Demonstrate the version test failing before the package variable is made
  linker-overridable, then passing afterward.
- Run the full Go test suite with CGO enabled.
- Build and execute the Windows x86_64 binary locally with an injected version.
- Validate the Release Please JSON files against their expected structure.
- Parse the workflow YAML and run an Actions-aware workflow linter when
  available.
- Review the final diff against the approved platform, branch, tag, token, and
  artifact requirements before publishing the feature branch.
