package time_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TimeTestSuite struct {
	suite.Suite
}

func TestTime(t *testing.T) {
	suite.Run(t, new(TimeTestSuite))
}

func (st *TimeTestSuite) TestTimeNow() {
	now := time.Now()
	fmt.Println(now)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

    then := time.Now()

    gap := then.Sub(now)
    fmt.Println(gap)

    fmt.Printf("%s", gap)
    fmt.Printf("%+v", gap)

	// fmt.Println(now.Unix())
	// fmt.Println(now.UnixMilli())
	// fmt.Println(now.UnixMicro())
	// fmt.Println(now.UnixNano())
}
