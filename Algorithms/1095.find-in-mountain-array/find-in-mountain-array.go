/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
    //找到峰值的位置
    peak:=findpeak(mountainArr)
    //如果峰值和target相相等，则直接返回
    if target==mountainArr.get(peak){
        return peak
    }
    //去左边找
    lindex:=binarysearch(mountainArr,0,peak-1,target, true)
    if lindex!=-1{
        return lindex
    }
    //去右边找
    rindex:=binarysearch(mountainArr,peak+1,mountainArr.length()-1,target,false)
    return rindex
}

//先找到峰值
func findpeak(mountainArr *MountainArray)int{
    left:=0
    right:=mountainArr.length()-1
    for left+1<right{
        mid:=left+(right-left)/2
        //（因为本题要求调用get方法的次数有限制，且峰值一定存在，所以不需要在中间找到峰值时进行返回,在最后返回就行）
        // if mountainArr.get(mid) > mountainArr.get(mid+1) && mountainArr.get(mid) > mountainArr.get(mid-1){
        //     return mid
        // }
        if mountainArr.get(mid) < mountainArr.get(mid+1){
            left=mid
        }else {
            right=mid
        }
    }
    if mountainArr.get(left)>mountainArr.get(right){
        return left
    }else{
        return right
    }
}
//普通的二分查找
//多加了一个参数：flag:为true表示按照升序进行搜索，为false表示按降序进行搜索
func binarysearch(mountainArr *MountainArray,left,right int,target int,flag bool)int{
    for left+1<right{
        mid:=left+(right-left)/2
        if mountainArr.get(mid)==target{
            return mid
        }else if mountainArr.get(mid) > target{
            if flag{ //升序
                right=mid
            }else{ //降序
                left=mid
            }
        }else{
            if flag{ //升序
                left=mid
            }else{ //降序
                right=mid
            }
        }
    }
    if mountainArr.get(left)==target{
        return left
    }
    if mountainArr.get(right)==target{
        return right
    }
    return -1
}
