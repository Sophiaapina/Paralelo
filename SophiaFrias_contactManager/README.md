# Contact Manager (Go)

## Run
From this folder:

```bash
go run .
```

This will:
- Persist contacts in `contacts.json`
- Log activity to `activity.log`

## Features
- Add contact (saved immediately to JSON)
- Edit contact (updates JSON)
- Delete contact (removes from JSON)
- Show contacts (reads JSON and prints)

## Error handling
- If `contacts.json` doesn't exist, it starts with an empty list.
- Read/write errors are printed to stderr and logged to `activity.log`.
