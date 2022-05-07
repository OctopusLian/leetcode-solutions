/**
 * Definition of Interval:
 * class Interval {
 *     int start, end;
 *     Interval(int start, int end) {
 *         this->start = start;
 *         this->end = end;
 *     }
 * }
 */

class Solution {
public:
    /**
     * @param intervals: an array of meeting time intervals
     * @return: if a person could attend all meetings
     */
    bool static cmp1(Interval &a,Interval &b){
    if(a.start!=b.start)        return a.start<b.start;
    else                        return a.end<b.end;
}

    bool canAttendMeetings(vector<Interval> &intervals) {
        sort(intervals.begin(),intervals.end(),cmp1);
        for(int i=1;i<intervals.size();i++){
            if(intervals[i-1].end>intervals[i].start)   return false;
        }
        return true;
    }
};