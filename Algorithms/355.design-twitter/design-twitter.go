type Twitter struct {
    //follow
    FollowedMap map[int]map[int]int
    //twitter
    TwitterMap []map[int]int
}


/** Initialize your data structure here. */
func Constructor() Twitter {
    TwitterMap := []map[int]int{}
    FollowedMap := map[int]map[int]int{}

    return Twitter{
        TwitterMap: TwitterMap,
        FollowedMap: FollowedMap,
    }
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.TwitterMap = append(this.TwitterMap, map[int]int{userId:tweetId})
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
    twitterArray := []int{}
    for  i := len(this.TwitterMap) - 1 ; i >=0 ; i-- {
        for authId,twitterId := range  this.TwitterMap[i] {
            if userId == authId || this.FollowedMap[userId][authId] > 0{
                twitterArray = append(twitterArray, twitterId)
            }   
        }
        if len(twitterArray) == 10 {
            return twitterArray
        }
    }
    return twitterArray
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.FollowedMap[followerId] == nil {
        this.FollowedMap[followerId] = map[int]int{}
    }
    this.FollowedMap[followerId][followeeId] = 1
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if this.FollowedMap[followerId] != nil {
        this.FollowedMap[followerId][followeeId] = 0
    }
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
