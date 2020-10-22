class Solution:
    def average(self, salary: List[int]) -> float:
        maxValue = max(salary)
        minValue = min(salary)
        total = sum(salary) - maxValue - minValue
        return total / (len(salary) - 2)