update salary #表名
set sex = (case sex
               when 'm' then 'f' #将m换为f
               else 'm'
           end);
