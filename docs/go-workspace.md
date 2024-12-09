# Go workspace

What is Workspace? [here](https://go.dev/ref/mod#workspaces)

## Add the module to the workspace

```bash
go work use ./libs/go-commons ./services/auth-ms ./services/user-ms
```

Ref: https://go.dev/doc/tutorial/workspaces

```bash
go work edit -dropuse=<module>
```

```bash
go work edit -dropuse=<module> <go_work_file>
```

```bash
go work sync
```
