# app

A tiny Go module the sandbox's CI tests (`go vet` and `go test`), so the
merge queue has real code to break.

- `tasks`: the task model and how a task is shown.
- `notify`: messages about tasks.
- `webhooks`: task events sent to registered URLs.
- `config`: settings.
