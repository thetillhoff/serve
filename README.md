# serve

A minimal webserver for local development.

## How to use

`serve` will just run a webserver on 0.0.0.0:3000, which serves the local directory.
`serve --verbose` does the same, but will also print the path every requests URI.
`serve --port <portnumber>` changes the port.
`serve --ipaddress <bind-ip>` changes the ipaddress where `serve` will bind to.
`serve --directory <path>` changes the directory which is served.
`serve --help` will display the shortcuts for these flags as well.

## How to release

1. Add information about new version to `CHANGELOG.md` & commit.
2. Push latest changes with `git push`.
3. List existing tags with `git tag`.
   ```
   v0.0.1
   v0.0.2
   ```
4. Select the next available tag and apply it with `git tag v0.0.3`.
5. Push tag with `git push origin v0.0.3`.

> If the build fails, the action will delete the tag itself. Not so much for when the precondition fails.
> If you messed something up, you can delete local tags with `git tag -d v0.0.3`.
> And you can delete remote tags with `git push --delete origin v0.0.3`.
