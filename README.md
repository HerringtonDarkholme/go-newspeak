# go-newspeak

A type-safe, injection-free extension to GORM.

> _War is Peace. Freedom is Slavery. Ignorance is Strength._
>                                  
>                                      —— George Orwell, 1984              

## Introduction
Newspeak provides a set of type-safe and SQL injection-free API to GORM. 
Users can construct a SQL-like DSL queries by utlitizing Newspeak's API.

```go
dbb.
    Table(user)
    Select(user.Age, fn.Avg(user.WorkYears)).
    Where(user.Age.Gt(35)).
    Group(user.Age).
    Having(user.Job.Eq("programmer"))
    Order.Desc(user.Age).
    Scan(&result)
```

The above statement generates equivalent SQL as below.

```sql
SELECT user.age, avg(user.work_years)
FROM user
WHERE user.age > 35
GROUP BY user.age
HAVING user.job = "programmer"
ORDER BY user.age DESC
```


## Trivia
The name "Newspeak" is heavily inspired by George Orwell's famous novel, in the sense that, by limiting the API surface, we constrains developers' ability to construct a vulnerable query. 
In the novel,  the language "Newspeak" abolishes the concept of freedom and guarantees unconscious orthodoxy at the level of language and mind. 

And the famous quote "War is Peace. Freedom is Slavery. Ignorance is Strength." reflects exactly Newspeak's design principle:

* War (to fight with migration and code gen) is Peace (of runtime security).
* Freedom (to write arbitrary SQL) is Slavery (of system vulnerability).
* Ignorance (of injection caveats) is Strength (of type-safe framework).
