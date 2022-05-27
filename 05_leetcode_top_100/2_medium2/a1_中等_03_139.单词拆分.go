package main

import "fmt"

/*
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
wordDict 中的所有字符串互不相同
s 和 wordDict[i] 仅有小写英文字母组成

示例 1：
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

示例 2：
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
注意，你可以重复使用字典中的单词。

示例 3：
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
*/

/*
动态规划
时间复杂度：O(n²)
空间复杂度：O(n)
转移方程：dp[i]=dp[j] && check(s[j..i−1])
https://leetcode.cn/problems/word-break/solution/dan-ci-chai-fen-by-leetcode-solution/
*/
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, v := range wordDict {
		wordDictSet[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	//s = "catsandog"
	//wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}
