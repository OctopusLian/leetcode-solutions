#include <bits/stdc++.h>

using namespace std;

struct Point {
    int x, y, z;
    bool isOk;

    Point() {}
    Point(int x, int y, int z, bool isOk) : x(x), y(y), z(z),
        isOk(isOk) {}
};

#define MAX_SIZE (1 << 12)

int main()
{
    // 打开两个输入的文件
    FILE *fp1 = fopen("in", "r");
    FILE *fp2 = fopen("in2", "r");

    // 输入询问序号，并对其进行排序去重，使得处理工作可以方便进行
    vector<int> tag;
    for (int x; EOF != fscanf(fp2, "%d", &x); tag.push_back(x)) {}
    sort(tag.begin(), tag.end());
    tag.erase(unique(tag.begin(), tag.end()), tag.end());

    // 关闭第二个文件
    fclose(fp2);
    fp2 = NULL;

    // 使用map可以处理多个无人机
    int top = 0;
    map<string, Point> mp;
    char s[MAX_SIZE];
    for (int index = 0; EOF != fscanf(fp1, "%s", s); index++) {
        bool f = false;
        int x, y, z, offsetX, offsetY, offsetZ;
        fscanf(fp1, "%d%d%d", &x, &y, &z);
        if (fgetc(fp1) != '\n')
            fscanf(fp1, "%d%d%d", &offsetX, &offsetY, &offsetZ),
            f = true;

        if (mp.find(s) == mp.end()) {
            if (!f)
                mp[s] = Point(x, y, z, true);
            else
                mp[s] = Point(0, 0, 0, false);
        } else {
            if (mp[s].isOk && mp[s].x == x && mp[s].y == y && mp[s].z == z) {
                mp[s] = Point(x + offsetX, y + offsetY, z + offsetZ, true);
            } else
                mp[s].isOk = false;
        }

        if (tag[top] == index) {
            if (!mp[s].isOk)
                printf("Error: %d\n", index);
            else
                printf("%s %d %d %d %d\n", s, index, mp[s].x, mp[s].y, mp[s].z);
            ++top;
        }
    }

    // 剩下的询问一定是超过序号的，因此，依次输出
    for (; top < (int)tag.size(); ++top)
        printf("cannot find %d\n", tag[top]);

    fclose(fp1);
    fp1 = NULL;
    return 0;
}