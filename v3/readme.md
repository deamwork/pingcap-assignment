# pingcap assignment

## requirements
```text
题目
有一个 100GB 的文件，里面内容是文本，要求：

1. 找出第一个不重复的词
2. 只允许扫一遍原文件
3. 尽量少的 IO
4. 内存限制 16G

时间一周，完成后反馈 GitHub 的链接即可，做题过程中如果有问题也可以集中下发来，我去问下出题人建议
```

## usage
run:
```bash
go run . -file=path_to_file
```

test:
```bash
go test -v .
```

## solution

the key is not load (read) the whole file into the memory, which will cause out-of-memory.

1. initialize a io.Reader (like bufio.NewReader) load file as byte flow, scan every byte in range of `[A-Za-z]`.
2. using fnv hash it as int, get remainder by slices, create a hashmap to count word and sequence, and save to temp file.
3. read a slice, merge those hashmaps which referred from step 2.
4. read all slice and preform step 3, then pop the minimum sequence of word from each slice.
5. compare all slices, get the minimum sequence of word between them.

## I/O consume

1. read source file (1 time: 1 read; 0 write)
2. write to slice(s) (1 time: 0 read; 1 write)
3. read from slice(s) (1 time: 1 read; 0 write)

## optimize

merge occurrence and count for each occurrence of a word will tremendously reduce the size of slice file. 
and it will reclaim that part of memory heap to reuse on next loop cycle, until reach the end of file.

meanwhile, im using [gob](https://golang.org/pkg/encoding/gob/) to serialize/unserialize (encode/decode) structure of a word in 
order to reduce the file io.

## further

load a pprof to profile heap usage and cpu cycle analysis

1. using goroutine to simultaneously read file in order to optimize read performance. (be aware of single process file read/write pointer shift)
2. manually control memory profile to avoid golang runtime garbage collection
