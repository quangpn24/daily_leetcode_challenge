(select distinct salary as SecondHighestSalary from Employee
 order by salary desc
 limit 1 offset 1)
 union
(SELECT null)
limit 1
