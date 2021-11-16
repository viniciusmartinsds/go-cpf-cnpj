package cpf

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Generate() (cpfString string) {
	rand.Seed(time.Now().UTC().UnixNano())
	cpf := rand.Perm(9)
	cpf = append(cpf, verify(cpf, len(cpf)))
	cpf = append(cpf, verify(cpf, len(cpf)))

	for _, c := range cpf {
		cpfString += strconv.Itoa(c)
	}

	return
}

func GeneratePretty() (cpf string) {
	digits := Generate()

	re := regexp.MustCompile(`(?P<1>\d{3})(?P<2>\d{3})(?P<3>\d{3})(?P<4>\d{2})`)

	names := re.SubexpNames()
	formatted := fmt.Sprintf("${%s}.${%s}.${%s}-${%s}",
		names[1], names[2], names[3], names[4])

	cpf = re.ReplaceAllString(digits, formatted)
	return
}

func sanitize(data string) string {
	data = strings.Replace(data, ".", "", -1)
	data = strings.Replace(data, "-", "", -1)
	return data
}

func valid(data string) (bool, error) {
	data = sanitize(data)

	if len(data) != 11 {
		return false, errors.New("Invalid length")
	}

	if strings.Contains(blacklist, data) || !check(data) {
		return false, errors.New("Invalid value")
	}

	return true, nil
}

const blacklist = `00000000000
11111111111
22222222222
33333333333
44444444444
55555555555
66666666666
77777777777
88888888888
99999999999
12345678909`

func stringToIntSlice(data string) (res []int) {
	for _, d := range data {
		x, err := strconv.Atoi(string(d))
		if err != nil {
			continue
		}
		res = append(res, x)
	}
	return
}

func verify(data []int, n int) int {
	var total int

	for i := 0; i < n; i++ {
		total += data[i] * (n + 1 - i)
	}

	total = total % 11
	if total < 2 {
		return 0
	}
	return 11 - total
}

func check(data string) bool {
	return checkEach(data, 9) && checkEach(data, 10)
}

func checkEach(data string, n int) bool {
	final := verify(stringToIntSlice(data), n)

	x, err := strconv.Atoi(string(data[n]))
	if err != nil {
		return false
	}
	return final == x
}
