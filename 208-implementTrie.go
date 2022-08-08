//A trie (pronounced as "try") or prefix tree is a tree data structure used to
//efficiently store and retrieve keys in a dataset of strings. There are various
//applications of this data structure, such as autocomplete and spellchecker.
//
// Implement the Trie class:
//
//
// Trie() Initializes the trie object.
// void insert(String word) Inserts the string word into the trie.
// boolean search(String word) Returns true if the string word is in the trie (
//i.e., was inserted before), and false otherwise.
// boolean startsWith(String prefix) Returns true if there is a previously
//inserted string word that has the prefix prefix, and false otherwise.
//
//
//
// Example 1:
//
//
//Input
//["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
//[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
//Output
//[null, null, true, false, true, null, true]
//
//Explanation
//Trie trie = new Trie();
//trie.insert("apple");
//trie.search("apple");   // return True
//trie.search("app");     // return False
//trie.startsWith("app"); // return True
//trie.insert("app");
//trie.search("app");     // return True
//
//
//
// Constraints:
//
//
// 1 <= word.length, prefix.length <= 2000
// word and prefix consist only of lowercase English letters.
// At most 3 * 10⁴ calls in total will be made to insert, search, and
//startsWith.
//
//
// Related Topics Hash Table String Design Trie 👍 7812 👎 96

package main

//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	IsEnd    bool
	Children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		i := ch - 'a'
		if node.Children[i] == nil {
			node.Children[i] = &Trie{}
		}
		node = node.Children[i]
	}
	node.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		i := ch - 'a'
		if node.Children[i] == nil {
			return false
		}
		node = node.Children[i]
	}
	if node.IsEnd == true {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		i := ch - 'a'
		if node.Children[i] == nil {
			return false
		}
		node = node.Children[i]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
