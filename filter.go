package sensitive

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

// Filter 敏感词过滤器
type Filter struct {
	trie  *Trie
	noise *regexp.Regexp
}

var (
	noise *regexp.Regexp
)

func init() {
	noise = regexp.MustCompile(`[/\\|\s&%$@*]+`)
}

// New 返回一个敏感词过滤器
func New() *Filter {
	return &Filter{
		trie:  NewTrie(),
		noise: noise,
	}
}

// UpdateNoisePattern 更新去噪模式
func (filter *Filter) UpdateNoisePattern(pattern string) {
	filter.noise = regexp.MustCompile(pattern)
}

// Load common method to add words
func (filter *Filter) Load(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		filter.trie.add(string(line))
	}

	return nil
}

// AddWord 添加敏感词
func (filter *Filter) AddWord(word string) {
	filter.trie.add(strings.ToLower(word))
}

// DelWord 删除敏感词
func (filter *Filter) DelWord(word string) {
	filter.trie.del(strings.ToLower(word))
}

// Filter 过滤敏感词
func (filter *Filter) Filter(text string) string {
	return filter.trie.Filter(text)
}

// Replace 和谐敏感词
func (filter *Filter) Replace(text string, repl rune) string {
	return filter.trie.Replace(text, repl)
}

// FindAll 找到所有匹配词
func (filter *Filter) FindAll(text string) []string {
	return filter.trie.FindAll(text)
}

// Validate 检测字符串是否合法
func (filter *Filter) Validate(text string) (bool, string) {
	text = filter.RemoveNoise(text)
	return filter.trie.Validate(text)
}

// RemoveNoise 去除空格等噪音
func (filter *Filter) RemoveNoise(text string) string {
	return filter.noise.ReplaceAllString(text, "")
}
