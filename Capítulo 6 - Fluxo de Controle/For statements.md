## For statements
A "for" statement specifies repeated execution of a block. There are three forms: The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.

ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
Condition = Expression .
For statements with single condition¶
In its simplest form, a "for" statement specifies the repeated execution of a block as long as a boolean condition evaluates to true. The condition is evaluated before each iteration. If the condition is absent, it is equivalent to the boolean value true.

for a < b {
	a *= 2
}

## For statements with for clause
A "for" statement with a ForClause is also controlled by its condition, but additionally it may specify an init and a post statement, such as an assignment, an increment or decrement statement. The init statement may be a short variable declaration, but the post statement must not.

ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
InitStmt = SimpleStmt .
PostStmt = SimpleStmt .
for i := 0; i < 10; i++ {
	f(i)
}
If non-empty, the init statement is executed once before evaluating the condition for the first iteration; the post statement is executed after each execution of the block (and only if the block was executed). Any element of the ForClause may be empty but the semicolons are required unless there is only a condition. If the condition is absent, it is equivalent to the boolean value true.

for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }
Each iteration has its own separate declared variable (or variables) [Go 1.22]. The variable used by the first iteration is declared by the init statement. The variable used by each subsequent iteration is declared implicitly before executing the post statement and initialized to the value of the previous iteration's variable at that moment.

var prints []func()
for i := 0; i < 5; i++ {
	prints = append(prints, func() { println(i) })
	i++
}
for _, p := range prints {
	p()
}
prints

1
3
5
Prior to [Go 1.22], iterations share one set of variables instead of having their own separate variables. In that case, the example above prints

6
6
6