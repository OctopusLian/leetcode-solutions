class Solution {
public:
    /**
     * @param numbers: An array of Integer
     * @param target: target = numbers[index1] + numbers[index2]
     * @return: [index1 + 1, index2 + 1] (index1 < index2)
     */
    vector<int> twoSum(vector<int> &numbers, int target) {
        // write your code here
        int i,j;
        int len=numbers.size();
        vector<int> result;
        for(i=0;i<len-1;i++)
        { 
            for(j=i+1;j<len;j++)
            { 
                if(numbers[i]+numbers[j]==target)
                {
                    result.push_back(i);
                    result.push_back(j);
                    //错误代码在这里写作v[0]=i;v[1]=j;              
                }
            }
        }
        return result;

    }
};