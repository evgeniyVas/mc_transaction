package psql

const queryUpdateBalance = `
UPDATE balance
SET amount = $1, updated_at = now()
WHERE user_id = $2
`
