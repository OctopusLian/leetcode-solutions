func generateParenthesis(n int) []string {
	res:=make([]string,0)
	helpGenerateParenthesis(0,0,n,"",&res)
	return res
}

func helpGenerateParenthesis(l,r,n int ,s string,res *[]string){
	if l==r && l==n{
		*res= append(*res, s)
	}
	if l<n{
		helpGenerateParenthesis(l+1,r,n,s+"(",res)
	}
	if r<l{
		helpGenerateParenthesis(l,r+1,n,s+")",res)
	}
}