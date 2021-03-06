package common

import (
	"encoding/json"
	"fmt"
	"hash"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func EqualStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	as := make([]string, len(a))
	copy(as, a)
	bs := make([]string, len(b))
	copy(bs, b)
	sort.Strings(as)
	sort.Strings(bs)
	for i, v := range as {
		if v != bs[i] {
			return false
		}
	}
	return true
}

func ComputeConfigElementSha(h hash.Hash, msg interface{}) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("computeConfigItemSha: json.Marshal: %s\n", err)
	}
	h.Write(data)
}

func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}

func EqualUUIDSlice(a, b []*uuid.UUID) bool {
	as := make([]string, 0, len(a))
	bs := make([]string, 0, len(b))
	for _, u := range a {
		as = append(as, u.String())
	}
	for _, u := range b {
		bs = append(bs, u.String())
	}
	return EqualStringSlice(as, bs)
}

func CompareStringSliceMap(s []string, m map[string]bool) error {
	if s == nil && m == nil {
		return nil
	}
	if len(s) != len(m) {
		return fmt.Errorf("map '%v', slice '%v'", m, s)
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// same length, so compare
	sort.Strings(keys)
	sort.Strings(s)

	sj := strings.Join(s, "\n")
	mj := strings.Join(keys, "\n")
	if sj != mj {
		return fmt.Errorf("mismatched entries, slice '%s', map '%s'", sj, mj)
	}
	return nil
}

func MismatchedErrors(e1, e2 error) bool {
	return (e1 != nil && e2 == nil) || (e1 == nil && e2 != nil) || (e1 != nil && e2 != nil && !strings.HasPrefix(e1.Error(), e2.Error()))
}

func GetOnboardCertName(cn string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\\.\\-]`)
	return re.ReplaceAllString(cn, "_")
}
