class Solution {
public:
    vector<vector<int>> findContinuousSequence(int target) {
        vector<vector<int>> result;
        vector<int> temp;
        for(int n=2;n*n+n<=2*target;n++){
            float t = (2.0*target/n-n+1)/2;
            cout<<t<<endl;
            if(int(t)==t){
                int sum = 0;
                int i = t;
                while(sum!=target){
                    temp.emplace_back(i);
                    sum += i;
                    i++;
                }
                result.insert(result.begin(),temp);
                temp.clear();
            }
        }
        return result;
    }
};
