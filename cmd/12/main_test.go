package main

import (
	"strings"
	"testing"
)

const input = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`

const input2 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`

func TestTask1a(t *testing.T) {
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	if res := task1(lines); res != 19 {
		t.Errorf("expected: 19 finalized paths, got: %d", res)
	}
}

func TestTask1b(t *testing.T) {
	lines := strings.Split(strings.Trim(string(input2), "\n"), "\n")
	if res := task1(lines); res != 226 {
		t.Errorf("expected: 226 finalized paths, got: %d", res)
	}
}
