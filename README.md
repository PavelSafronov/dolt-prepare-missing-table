# Intro

This repository reproduces a bug in dolt: when preparing a statement that uses a non-existent table, dolt returns an error code 1105 (ER_UNKNOWN_ERROR), but it should be returning 1146 (ER_NO_SUCH_TABLE)

## Running

1. Start MySQL
2. `make test`
3. Note that the test passes
4. Stop MySQL
5. Start dolt
6. `make test`
7. Note that the test fails with error "still getting 1105"
