import collections
class Solution:
    def maxWeight(self, edges: List[List[int]], weight: List[int]) -> int:
        n = len(weight)
        self.weight = weight
        point_set = collections.defaultdict(set) # 记录和 i相连且编号大于i的所有点
        for x,y in edges:
            if x>y:
                x,y = y,x
            point_set[x].add(y)

        max_triangle_point_dict = collections.defaultdict(list)  # 点i能构成的最大三角形
        triangle_point_dict = collections.defaultdict(list) # 点i能构成的所有三角形
        triangle_edge_dict = collections.defaultdict(list)  # 边(i,j)能构成的 Top3 三角形

        # 查找三角形
        for i in range(0,n):
            for j in point_set[i]:
                all_points_list = list(point_set[i]&point_set[j]) # 能与 i,j 构成三角形的点
                for k in all_points_list:
                    # 由于 point_set 结构， 满足 i<j<k
                    sum_weight = weight[i]+weight[j]+weight[k]
                    triangle_info = [i,j,k,sum_weight]
                    # i,j,k 三个点 记录和更新三角形信息
                    for lm in [i,j,k]:
                        if not max_triangle_point_dict[lm] or max_triangle_point_dict[lm][-1]<sum_weight:
                            max_triangle_point_dict[lm] = triangle_info
                        triangle_point_dict[lm].append([i,j,k])
                    # 三个条边 记录和更新三角形信息
                    for edge in [(i,j),(i,k),(j,k)]:
                        triangle_edge_dict[edge] = self.get_top3(triangle_edge_dict[edge],triangle_info)
 
        res = 0
        for i in range(0,n):
            # 点无三角形的情况
            if not max_triangle_point_dict[i]:
                continue
            max_triange = max_triangle_point_dict[i]
            # 两个三角形完全重合的情况，即一个三角形
            res = max(res,max_triange[-1])

            # 最大三角形 max_triange 和所有包含 i 的三角形一一组合
            for info in triangle_point_dict[i]:
                res = max(res,self.count_val(max_triange,info))
        
            # 两个包含max_triangle边（i,x),(i,y) 的 Top3 三角形一一组合
            max_points = max_triange[:3]
            max_points.remove(i)
            edge1,edge2 = [(i,x) if i<x else (x,i) for x in max_points]
            for info1 in triangle_edge_dict[edge1]:
                for info2 in triangle_edge_dict[edge2]:
                    res = max(res,self.count_val(info1,info2))
        
        return res
    
    def count_val(self, info1, info2):
        # 计算 两个三角的val总和， 过滤重复点
        all_points = set(info1[:3])|set(info2[:3])
        return sum([self.weight[x] for x in all_points])
        
    def get_top3(self, triangle_list, triangle_info):
        # 更新 同一条边 的 top3 的三角形
        if not triangle_list:
            return [triangle_info]
        if triangle_list[-1][-1]>=triangle_info[-1]:
            triangle_list.append(triangle_info)
            return triangle_list
        for index in range(0,len(triangle_list)):
            if triangle_list[index][-1]<triangle_info[-1]:
                triangle_list.insert(index,triangle_info)
                break
        triangle_list = triangle_list[:3]
        return triangle_list