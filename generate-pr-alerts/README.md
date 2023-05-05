# generate-pr-alerts

GitHub's Pull Request feature is one of the best things to ever happen to
collaborative software development.  Unfortunately, PRs can also be a parking
lot for code.  There is a feature that can remind you of open PRs if configured
properly, but I had a need to generate a minor report.  Basically, I'm surveying
a very large number of repositories in an organization and needed birdseye view
at the org level rather than a Teams level.

This is where `generate-pr-alerts` comes in.  It can check arbitrary repos or it
can fetch repo urls via those attached to an Organization.  This makes it simple
and easy to monitor many repos at the org level.

By default, `generate-pr-alerts` will fetch information from GitHub and print to
STDOUT.  The text format is in Markdown and provides a short summary of open PRs
as well as convenient links directly to them.

## Configuration

First thing first, you'll need to generate a Personal Access Token.  You can
create a PAT in the [Token Settings page][ts].  That will be a special token
that will give the app access to everything that you have access to (unless you
choose to restrict those options when generating the token).

[ts]: https://github.com/settings/tokens

You can also choose to have `generate-pr-alerts` look for specific repos, pull
repo information from organizations, or both.  Repos are designated using the
format `username/reponame`.  For example, specifying this repo would be
`Grayson/dashboard`.  Organizations are specified by their "login" name.

In addition to printing information to stdout, you can also generate a JSON file
that contains the top level information about these PRs and Issues.  By default,
this file is not created.  However, you can specify a file path using the `json`
key and one will be created at that location.

You can specify this information either as a CLI flag or in a `config.yml` file
placed in the same folder as the `generate-pr-alerts` executable.  The `token`
property is a string that can only be set once, but the repos and orgs can be
specified as arrays (in the YAML) or defined multiple times as CLI flags.

These two example configurations should be effectively identical:

```
token: pat_github_token
repos:
- Grayson/dashboard
- Grayson/code-clone-tool
orgs:
- objectiveceo
json: /Users/grayson/Desktop/output.json
```

`$ ./generate-pr-alerts -token pat_github_token -repo Grayson/dashboard -repo
Grayson/code-clone-tool -orgs objectiveceo -json
/Users/grayson/Desktop/output.json`

Here's a brief table summarizing what is an array in the yaml or is repeatable
on the command line:

| Config file key | Is Array | CLI flag | Can repeat |
|-----------------|----------|----------|------------|
|`token`          |          |`token`   |            |
|`repos`          |    ✓     |`repo`    |      ✓     |
|`orgs`           |    ✓     |`org`     |      ✓     |
|`json`           |          |`json`    |            |

There may be times when you want to set the token as an environment variable.
If no token is specified on the command line or in the configuration file,
`generate-pr-alerts` will check to see if the `GITHUB_TOKEN` environment
variable is set.
