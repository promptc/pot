# Pot - Prompt Manager

`pot`, short for Prompt OJBK Tool, is a prompt manager tool.
It works with `promptc`.


## Command

- `pot update`: Update source
- `pot upgrade`: Upgrade prompts
- `pot install <prompt>`: Install specific prompt
- `pot remove <prompt>`: Remove installed prompt
- `pot list`: List all prompt
- `pot search <name>`: Search prompt

## Structure

```
+example-repo.dev
  +info.json
   {
      "db": "https://cdn.example-repo.dev/store.db",
      "id", "example-repo.dev",              |
      "prompt": "https://cdn.er.dev/prompts" |
   }                             |           |
                                 |           |
+cdn.cdn.er.dev                  |           |
  +store.db <--------------------------------+
  +prompts <---------------------+
    +prompt1
    | +info.json
    | +main.prompt
    +prompt2
      +info.json
      +main.prompt
```

```
+config path
  +promptc
    +pot
      +db
      | +example-repo.dev.db
      | +updating.dev.db-new
      | +info.json
      +prompts
      | +prompt1
      | | +example-repo.dev
      | |   +info.json
      | |   +main.prompt
      | +prompt2
      |     +example-repo.dev
      |     |   +info.json
      |     |   +main.prompt
      |     +ex2.dev
      |         +info.json
      |         +main.prompt
      +config.json

```