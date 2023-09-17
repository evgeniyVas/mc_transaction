package psql

const queryUpdateBalance = `
UPDATE balance
SET amount = $1
WHERE user_id = $2
`
