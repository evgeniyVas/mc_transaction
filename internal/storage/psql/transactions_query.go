package psql

const queryInsertTransaction = `
INSERT INTO transactions
	(user_id, amount, status, pay_id, locked, token)
VALUES 
	($1, $2, $3, $4, $5, $6)
RETURNING id;
`

const querySelectTransactionWithLock = `
SELECT id, pay_id, amount, user_id, created_at
FROM transactions
WHERE status=created AND locked=false
FOR UPDATE SKIP LOCKED LIMIT 1;
`

const queryUpdateTransactionLocked = `
UPDATE transactions
SET locked = true
WHERE id = $1
`
