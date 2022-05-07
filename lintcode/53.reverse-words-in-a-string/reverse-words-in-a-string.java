public class Solution {
    /*
     * @param s: A string
     * @return: A string
     */
    public String reverseWords(String s) {
        if(s.trim().equals(""))
        return s;
        String[] strs = s.split(" ");
        StringBuilder sb = new StringBuilder();
        for(int i=strs.length-1; i>=0; i--){
            if(!strs[i].equals(""))
            sb.append(strs[i]+" ");
        }
        return sb.toString().trim();

    }
}