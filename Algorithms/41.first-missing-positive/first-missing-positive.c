int firstMissingPositive(int* nums, int numsSize) {
    if(numsSize == 0)
        return 1;
    
    //第i位存放i+1的数值
    for(int i = 0;i < numsSize;i++){
        if(nums[i] > 0){
            //如果交换的数据大于0且小于i+1且数据不相等，则进行交换（放在合适的位置上）
            while(nums[i] > 0 && nums[i] < i+1 && nums[i] != nums[nums[i] - 1]){
                int temp = nums[nums[i] - 1];
                nums[nums[i] - 1] = nums[i];
                nums[i] = temp;
            }
        }
    }

    //循环寻找不符合要求的数据
    for(int i = 0;i < numsSize;i++){
        if(nums[i] != i+1){
            return i+1;
        }
    }

    //如果都符合要求，则返回长度+1的值
    return numsSize + 1;
}
