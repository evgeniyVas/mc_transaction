package psql

const queryInsertTransaction = `
INSERT INTO transactions
	(user_id, amount, status, pay_id, locked, token)
VALUES 
	($1, $2, $3, $4, $5, $6)
RETURNING id;
`
