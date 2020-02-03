func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
    len_nums1,len_nums2 := len(nums1),len(nums2)  //获取num1和num2数组的长度
    
    len_num := len_nums1 + len_nums2  //两个数组加起来的总长度
    var middle int = len_num / 2  //中位数
    
    var m float64 = 0.0  //声明m为浮点型变量
    if len_num % 2 == 1{  //如果总长度是奇数
        i,j,t := 0,0,0 //
        for {
            if t > middle || i > len_nums1 - 1 || j > len_nums2 - 1{  
                break
            }
            if nums1[i] <= nums2[j]{
                m = float64(nums1[i])
                i++
                t++
            }else{
                m = float64(nums2[j])
                j++
                t++  
            }  
        } 
        if t <= middle {
            if i > len_nums1 - 1{
                j += middle - t
                m = float64(nums2[j])
            }else if j > len_nums2 - 1{
                i += middle - t
                m = float64(nums1[i])
            }
        }
    }else{  //如果总长度为偶数
        i,j,t,l,r := 0,0,0,0,0
        for {
            if t > middle  || i > len_nums1 - 1 || j > len_nums2 - 1{
                break
            }
            if nums1[i] <= nums2[j]{
                if t == middle - 1{
                    l = nums1[i]
                } else if t == middle{
                    r = nums1[i]
                } 
                i++
                t++
            }else{
                if t == middle - 1{
                    l = nums2[j]
                } else if t == middle{
                    r = nums2[j]
                } 
                j++
                t++
            }
        }
        if t <= middle {
            if i > len_nums1 - 1{
                j += middle - t
                if l == 0{
                    l = nums2[j - 1]
                    r = nums2[j]
                    
                }else{
                    r = nums2[j]
                }
               
            }else if j > len_nums2 - 1{
               i += middle - t
                if l == 0{
                    l = nums1[i - 1]
                    r = nums1[i]
                    
                }else{
                    r = nums1[i]
                }
            }
        }
        m = float64(l + r) / 2
    }
    
    return m
    
}

