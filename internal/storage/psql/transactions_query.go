package psql

const queryInsertTransaction = `
INSERT INTO transactions
	(user_id, amount, status, pay_id, token)
VALUES 
	($1, $2, $3, $4, $5)
RETURNING id;
`

const querySelectTransactionWithLock = `
SELECT id, pay_id, amount, user_id, created_at
FROM transactions
WHERE status='created' AND (locked_at IS NULL OR locked_at < now() - $1::interval) AND updated_at < now() - $2::interval
FOR UPDATE SKIP LOCKED LIMIT 1;
`

const queryUpdateTransactionLock = `
UPDATE transactions
SET locked_at = now(), updated_at = now()
WHERE id = $1
`

const queryUpdateTransactionUnlock = `
UPDATE transactions
SET locked_at = null, updated_at = now(), status = $1
WHERE id = $2
`
