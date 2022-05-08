/*
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-05-08 16:32:06
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-08 16:32:08
 */
class Solution {
    public int minArray(int[] numbers) {
        //设置left，right指针分别指向 numbers 数组左右两端
        //left指向当前区间的最左边位置，所以初始化为0
        int left = 0;
        //right指向当前区间的最右边位置，所以初始化为nums.length - 1
        int right = numbers.length - 1;

        //循环进行二分查找，直到左端点位置超过右端点
        while (left < right) {
            //mid为中点
            int mid = (left + right) / 2;
            //当mid点所在元素大于数组末端的元素时，由于原来的数组是递增有序的，此时出现了异常，大的数在前面
            //所以旋转点在[mid + 1,end]区间里面
            if (numbers[mid] > numbers[right]) {
                //所以旋转点在[mid + 1,end]区间里面，更新left的位置为mid + 1
                left = mid + 1;
            } else if (numbers[mid] < numbers[right]) {
                //当mid点所在元素小于数组末端的元素时，由于原来的数组是递增有序的
                //所以旋转点在[left,mid]区间里面，更新right的位置为mid
                right = mid;
            } else {
                //采取遍历的方式
                return findMin(numbers,left,right);
            }
        }
        return numbers[left];
    }
    //从头到尾遍历numbers，获取最小值
    public int findMin(int[] numbers,int left,int right) {
        //默认为数组的第一个元素为最小值
        int result = numbers[left];
        //从头到尾遍历numbers
        for (int i = left;i <= right;i++) {
            //当发现此时遍历的元素值小于result
            if (numbers[i] < result) {
                //更新result
                result = numbers[i];
            }
        }
        return result;
    }
}