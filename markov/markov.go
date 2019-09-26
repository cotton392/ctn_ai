package markov

import (
	"fmt"
	"github.com/bluele/mecab-golang"
	"math/rand"
	"os"
	"time"
)

func ParseToNode(m *mecab.MeCab, input string) []string {
	words := []string{}
	tg, err := m.NewTagger()
	if err != nil {
		fmt.Printf("New tagger error. err: %v", err)
		os.Exit(-1)
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(input)
	if err != nil {
		fmt.Printf("New Lattice error. error: %v", err)
		os.Exit(-1)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		if node.Surface() != "" {
			words = append(words, node.Surface())
		}
		if node.Next() != nil {
			break
		}
	}
	return words
}

func GetMarkovBlocks(words []string) [][]string {
	res := [][]string{}
	resHead := []string{}
	resEnd := []string{}

	if len(words) < 3 {
		return res
	} //長さが3以下の場合作らないようにする

	resHead = []string{"#This is begin#", words[0], words[1]}
	res = append(res, resHead)

	for i := 1; i < len(words)-2; i++ {
		markovBlock := []string{words[i], words[i+1], words[i+2]}
		res = append(res, markovBlock)
	}

	resEnd = []string{words[len(words)-2], words[len(words)-1], "#This is end#"}
	res = append(res, resEnd)

	return res
}

func FindBlocks(array [][]string, target string) [][]string {
	blocks := [][]string{}
	for _, s := range array {
		if s[0] == target {
			blocks = append(blocks, s)
		}
	}

	return blocks
}

func ConnectBlocks(array [][]string, dist []string) []string {
	rand.Seed((time.Now().Unix()))
	i := 0

	for _, word := range array[rand.Intn(len(array))] {
		if i != 0 {
			dist = append(dist, word)
		}
		i += 1
	}

	return dist
}

func MarkovChainExec(array [][]string) []string {
	ret := []string{}
	block := [][]string{}
	count := 0

	block = FindBlocks(array, "#This is begin#")
	ret = ConnectBlocks(block, ret)

	for ret[len(ret)-1] != "#This is end#" {
		block = FindBlocks(array, ret[len(ret)-1])
		if len(block) == 0 {
			break
		} //　blockが空だった場合break
		ret = ConnectBlocks(block, ret)

		count++
		if count == 200 {
			break
		} // 無限ループ対策
	}

	return ret
}

func TextGenerate(array []string) string {
	ret := ""
	for _, s := range array {
		if s == "#This is end#" {
			continue
		}

		if len([]rune(ret)) >= 70 {
			break
		}

		ret += s
	}

	return ret
}
