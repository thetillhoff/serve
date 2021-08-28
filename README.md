# serve
A minimal webserver for local development.

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

> If you messed something up, you can delete local tags with `git tag -d v0.0.3`.
> And you can delete remote tags with `git push --delete origin v0.0.3`.
