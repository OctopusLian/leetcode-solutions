func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    //解法1，递归
	ori:=image[sr][sc]
	helpFloodFill(&image,sr,sc,newColor,ori)
	return image
}

func helpFloodFill(image *[][]int,sr,sc int,newColor int,ori int){

	//0 <= sr < len(image) and 0 <= sc < len(image[0]) and image[sr][sc] == oldColor and image[sr][sc] != newColor:
	if sr>=0&&sr< len(*image)&&sc>=0&&sc< len((*image)[0])&&(*image)[sr][sc]==ori&&(*image)[sr][sc]!=newColor{
		(*image)[sr][sc]=newColor
		helpFloodFill(image,sr-1,sc,newColor,ori)
		helpFloodFill(image,sr+1,sc,newColor,ori)
		helpFloodFill(image,sr,sc-1,newColor,ori)
		helpFloodFill(image,sr,sc+1,newColor,ori)
	}

}