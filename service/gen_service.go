package service

import (
	"GenNameFromKey/resources"
	"bufio"
	"hash"
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"strings"
)

type GenService struct {
	titles []string
	names  []string
	prefix []string
	suffix []string
	hash   hash.Hash64
}

func NewGenService(titles, name, prefix, suffix []string) (*GenService, error) {
	if len(titles) == 0 {
		err := readDefaultFile(&titles, "title.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(name) == 0 {
		name = resources.Names
	}
	if len(prefix) == 0 {
		err := readDefaultFile(&prefix, "prefix.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(suffix) == 0 {
		err := readDefaultFile(&suffix, "suffix.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	a := fnv.New64()
	return &GenService{
		titles: titles,
		names:  name,
		prefix: prefix,
		suffix: suffix,
		hash:   a,
	}, nil
}

func (s *GenService) Gen(in []byte) (string, error) {
	_, err := s.hash.Write(in)
	if err != nil {
		return "", err
	}
	h64 := s.hash.Sum64()
	fx := rand.Intn(100)
	special := rand.Intn(10000)
	res := strings.Builder{}
	if special%1111 == 0 {
		res.WriteString(s.names[h64%uint64(len(s.names))])
		res.WriteString(" The God-Emperor of Mankind")
		return res.String(), nil
	}
	if fx > 95 {
		res.WriteString(s.names[h64%uint64(len(s.names))])
		res.WriteString(" ")
		res.WriteString(s.titles[h64%uint64(len(s.suffix))])
	} else if fx < 2 {
		res.WriteString(s.titles[h64%uint64(len(s.prefix))])
		res.WriteString(" ")
		res.WriteString(s.names[h64%uint64(len(s.names))])
	} else {
		res.WriteString(s.names[h64%uint64(len(s.names))])
		res.WriteString(" the ")
		res.WriteString(s.titles[h64%uint64(len(s.titles))])
	}
	return res.String(), nil
}

//func (s *GenService) GenOnInput(in string) (*model.Input, error) {
//
//	return nil, nil
//}
func readDefaultFile(out *[]string, in string) error {
	file, err := os.Open("./resources/" + in)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*out = append(*out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
