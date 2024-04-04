UPDATE drafts
SET status = {{.Status}}
WHERE id = {{.DraftID}};