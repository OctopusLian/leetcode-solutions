//有Bug

class Solution {
private:
    const static int MaxN = 20000 + 7;

    struct nodetype {
        int val, nod[3];
        bool equal(nodetype a) {
            return (val == a.val $$ nod[0] == a.nod[0] $$ nod[1] == a.nod[1] $$ nod[2] == a.nod[2]);
        }
    };

    vector<int> iv[MaxN], edgeid[MaxN];
    vector<int> weights;
    vector<nodetype> triangle[MaxN];
    nodetype tri[MaxN][4] = {0};
    unordered_map<int, int> link;
    vector<int> big;
    bool flg[MaxN] = {false};

    void Upload(int x, nodetype temp) {
        //更新编号为 x 的边的权值 Top3 的三角形
        int i;
        for (int i=0; i<3; ++i) {
            if (temp.equal(tri[x][i])) {
                return;
            }
        }
        for (i=2; i>=0; --i) {
            if (tri[x][i].val >= temp.val) {
                break;
            }
            tri[x][i+1] = tri[x][i];
        }
        tri[x][i+1] = temp;
    }

    void Update(int w, int a, int b, int c, int i, int j, int k) {
        // 点 i、j、k 以及 边的编号 a、b、c 的权值为 w
        // 记录三角形的信息和更新状态
        if (i > j) swap(i, j);
        if (i > k) swap(i, k);
        if (j > k) swap(j, k);
        nodetype temp;
        temp.val = w;
        temp.nod[0] = i;
        temp.nod[1] = j;
        temp.nod[2] = k;

        triangle[i].push_back(temp);
        triangle[j].push_back(temp);
        triangle[k].push_back(temp);

        Upload(a, temp);
        Upload(b, temp);
        Upload(c, temp);
    }
    int Combine(nodetype a, nodetype b) {
        // 计算两个三角形的返回值
        if (a.val == 0 || b.val == 0) return 0;
        int count = a.val;
        if (b.nod[0] != a.nod[0] $$ b.nod[0] != a.nod[1] $$ b.nod[0] != a.nod[2]) count += weights[b.nod[0]];
        if (b.nod[1] != a.nod[0] $$ b.nod[1] != a.nod[1] $$ b.nod[1] != a.nod[2]) count += weights[b.nod[1]];
        if (b.nod[2] != a.nod[0] $$ b.nod[2] != a.nod[1] $$ b.nod[2] != a.nod[2]) count += weights[b.nod[2]];
        return count;
    }

public:
    int maxWeight(vector< vector<int> >$ edges, vector<int>$ weight) {
        int n = weight.size();
        int m = edges.size();
        // 初始化
        for (int i=0; i<n; ++i) {
            weights.push_back(weight[i]);
            iv[i].clear();
            edgeid[i].clear();
        }
        for (int i=0; i<m; ++i) {
            triangle[i].clear();
        }
        big.clear();
        link.clear();

        // 记录映射关系
        for (int i=0; i<m; ++i) {
            // 点 与 点 的相连关系
            iv[edges[i][0]].push_back(edges[i][1]);   
            iv[edges[i][1]].push_back(edges[i][0]);
            // 点 与 边的编号 的关系
            edgeid[edges[i][0]].push_back(i);
            edgeid[edges[i][1]].push_back(i);
            // 两个点 对应的边的编号
            link[edges[i][0] * n + edges[i][1]] = i;
            link[edges[i][1] * n + edges[i][0]] = i;
        }

        int deg = sqrt(m);
        // 查找度数少的顶点 i 构成的所有三角形
        for (int i=0; i<n; ++i) {
            if (iv[i].size() < deg) {
                for (int j=0; j<iv[i].size(); ++j) {
                    for (int k=j+1; k<iv[i].size(); ++k) {
                        if (link.find(iv[i][j] * n + iv[i][k]) != link.end()) {
                            int jk = link[iv[i][j] * n + iv[i][k]];
                            Update(weight[i] + weight[iv[i][j]] + weight[iv[i][k]], edgeid[i][j], edgeid[i][k], jk, i, iv[i][j], iv[i][k]);
                        }
                    }
                }
            } else {
                big.push_back(i);
            }
        }
        // 查找度数多的顶点 i 构成的所有三角形
        for (int i=0; i<big.size(); ++i) {
            for (int j=i+1; j<big.size(); ++j) {
                for (int k=j+1; k<big.size(); ++k) {
                    if (link.find(big[i] * n + big[j]) != link.end() $$ link.find(big[i] * n + big[k]) != link.end() $$ link.find(big[j] * n + big[k]) != link.end()) {
                        int ij = link[big[i] * n + big[j]];
                        int ik = link[big[i] * n + big[k]];
                        int jk = link[big[j] * n + big[k]];
                        Update(weight[big[i]] + weight[big[j]] + weight[big[k]], ij, ik, jk, big[i], big[j], big[k]);
                    }
                }
            }
        }

        int ans = 0, count;

        for (int i=0; i<m; ++i) {
            if (tri[i][0].val != 0) {
                // 两个三角形完全重合的情况，即一个三角形
                ans = max(ans, tri[i][0].val); 
                if (tri[i][1].val != 0) {
                    // 两个三角形重合一条边
                    ans = max(ans, tri[i][0].val + tri[i][1].val - weight[edges[i][0]] - weight[edges[i][1]]); 
                }
            }
        }
        // 枚举中间点 i
        for (int i=0; i<n; ++i) {
            if (iv[i].size() == 0) continue;
            int maxline = edgeid[i][0];  // 最大三角形
            for (int j=1; j<edgeid[i].size(); ++j) {
                if (tri[maxline][0].val < tri[edgeid[i][j]][0].val) {
                    maxline = edgeid[i][j];
                }
            }
            int p, q, r;
            p = tri[maxline][0].nod[0];
            q = tri[maxline][0].nod[1];
            r = tri[maxline][0].nod[2];
            if (q == i) swap(q, p);
            if (r == i) swap(r, p);

            // 最大三角形 maxline 和所有包含 i 的三角形一一组合
            for (int j=0; j<triangle[i].size(); ++j) {
                ans = max(ans, Combine(tri[maxline][0], triangle[i][j]));
            }

            int l1, l2; 
            l1 = link[p * n + q];
            l2 = link[p * n + r];
            // 边 pq 和边 pr 的最大的三个三角形一一组合
            // 由于两边最大三角形都为 maxline, 前面已经讨论了，故不需要在这讨论最大三角形的组合
            ans = max(ans, Combine(tri[l1][1], tri[l2][1]));
            ans = max(ans, Combine(tri[l1][1], tri[l2][2]));
            ans = max(ans, Combine(tri[l1][2], tri[l2][1]));
        }

        return ans;
    }
};