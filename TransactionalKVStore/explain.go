package main

// Build an interactive shell that allows access to a "transactional in-memory key/value store".

//COMMANDS:
//  SET key value - Set the value at key to value
//  GET key - Get the value at key
//  DEL key - Delete the value at key
//  COUNT value - Count the number of keys with the given value
//  BEGIN - Start a new transaction
//  ROLLBACK - Rollback the most recent transaction
//  COMMIT - Commit all open transactions
//  END - Exit the program

// DATA IS NOT PERSISTED BETWEEN RUNS
// Key Values can only be strings

// Understanding a Transaction
// A transaction is a sequence of operations that are treated as a single unit.
// A transaction is a logical unit of work that is performed as a single, indivisible operation on the data.
// A transaction is a set of operations that must either all be completed or all be aborted.

// > BEGIN // Creates a new transaction
// > SET X 200
// > SET Y 14
// > GET Y

// Until the active transaction is committed using the COMMIT command, all changes are held in memory.
// > ROLLBACK // Rolls back the active transaction

// > BEGIN //Creates a new active transaction
// > SET X 5
// > SET Y 19
// > BEGIN //Spawns a new transaction in the context of the previous transaction and now this is currently active
// > GET Y
// Y = 19 //The new transaction inherits the context of its parent transaction**
// > SET Y 23
// > COMMIT //Y's new value has been persisted to the key-value store**
// > GET Y
// Y = 23 // Changes made by SET Y 19 have been discarded**
