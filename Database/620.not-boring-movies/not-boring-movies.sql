select *
from cinema #查询所有表
where mod(id, 2) = 1 and description != 'boring' #设定条件，id为奇数且影片描述为非 boring
order by rating DESC #根据rating降序排列
