int* twoSum(int* nums, int numsSize, int target) {
    int *a = (int*)malloc(2 * sizeof(int));
    for(int i = 0;i < numsSize;i++){
        for(int j = i + 1;j < numsSize;j++){
            if(nums[j] == target - nums[i]){
                a[0] = i;
                a[1] = j;
            }
        }
    }

    return a;
}

