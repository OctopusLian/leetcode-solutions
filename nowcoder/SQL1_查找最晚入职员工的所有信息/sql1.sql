SELECT *
FROM employees
WHERE hire_date = (SELECT MAX(hire_date) FROM employees)