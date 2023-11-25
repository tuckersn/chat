# Metadata Schemas
Multiple tables have a JSONB field called 'metadata' for unstructured storage in that entity's record directly using Postgres JSON operators. These are the definitions of the fields for each table.

# Universal Fields
```js
{
    "embeddingsIds": {
        [columnName]: /[0-9]+/ // see table_openai_embeddings.go
    }
}
```

# Specific Table Fields

## user

## room

## page

## message

