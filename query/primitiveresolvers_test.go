package query

import (
	"testing"
	"time"
)

func TestIntClauseResolver(t *testing.T) {
	type args struct {
		clause Clause
		val    int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{Clause{0, Eq, 10}, 10}, true},
		{args{Clause{0, Eq, 10}, 13}, false},
		{args{Clause{0, NEq, 10}, 10}, false},
		{args{Clause{0, Gt, 10}, 115}, true},
		{args{Clause{0, In, []int{10, 1, 2}}, 10}, true},
		{args{Clause{0, In, []int{1, 2}}, 10}, false},
		{args{Clause{0, In, []int{}}, 10}, false},
		{args{Clause{0, In, nil}, 10}, false},
	}
	for _, tt := range tests {
		if got := IntClauseResolver(tt.args.clause, tt.args.val); got != tt.want {
			t.FailNow()
		}
	}
}

func TestUint64ClauseResolver(t *testing.T) {
	type args struct {
		clause Clause
		val    uint64
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{Clause{0, Eq, uint64(10)}, uint64(10)}, true},
		{args{Clause{0, Eq, uint64(10)}, uint64(13)}, false},
		{args{Clause{0, NEq, uint64(10)}, uint64(10)}, false},
		{args{Clause{0, Gt, uint64(10)}, uint64(115)}, true},
		{args{Clause{0, In, []uint64{10, 1, 2}}, uint64(10)}, true},
		{args{Clause{0, In, []uint64{1, 2}}, uint64(10)}, false},
		{args{Clause{0, In, []uint64{}}, uint64(10)}, false},
		{args{Clause{0, In, nil}, uint64(10)}, false},
	}
	for _, tt := range tests {
		if got := Uint64ClauseResolver(tt.args.clause, tt.args.val); got != tt.want {
			t.FailNow()
		}
	}
}

func TestUint8lauseResolver(t *testing.T) {
	type args struct {
		clause Clause
		val    uint8
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{Clause{0, Eq, uint8(10)}, uint8(10)}, true},
		{args{Clause{0, Eq, uint8(10)}, uint8(13)}, false},
		{args{Clause{0, NEq, uint8(10)}, uint8(10)}, false},
		{args{Clause{0, Gt, uint8(10)}, uint8(115)}, true},
		{args{Clause{0, In, []uint8{10, 1, 2}}, uint8(10)}, true},
		{args{Clause{0, In, []uint8{1, 2}}, uint8(10)}, false},
		{args{Clause{0, In, []uint8{}}, uint8(10)}, false},
		{args{Clause{0, In, nil}, uint8(10)}, false},
	}
	for _, tt := range tests {
		if got := Uint8ClauseResolver(tt.args.clause, tt.args.val); got != tt.want {
			t.FailNow()
		}
	}
}

func TestUintClauseResolver(t *testing.T) {
	type args struct {
		clause Clause
		val    uint
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{Clause{0, Eq, uint(10)}, uint(10)}, true},
		{args{Clause{0, Eq, uint(10)}, uint(13)}, false},
		{args{Clause{0, NEq, uint(10)}, uint(10)}, false},
		{args{Clause{0, Gt, uint(10)}, uint(115)}, true},
		{args{Clause{0, In, []uint{10, 1, 2}}, uint(10)}, true},
		{args{Clause{0, In, []uint{1, 2}}, uint(10)}, false},
		{args{Clause{0, In, []uint{}}, uint(10)}, false},
		{args{Clause{0, In, nil}, uint(10)}, false},
	}
	for _, tt := range tests {
		if got := UintClauseResolver(tt.args.clause, tt.args.val); got != tt.want {
			t.FailNow()
		}
	}
}

func TestFloatClauseResolver(t *testing.T) {
	type args struct {
		clause Clause
		val    float32
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{Clause{0, Eq, float32(10)}, float32(10)}, true},
		{args{Clause{0, Eq, float32(10)}, float32(13)}, false},
		{args{Clause{0, NEq, float32(10)}, float32(10)}, false},
		{args{Clause{0, Gt, float32(10)}, float32(115)}, true},
		{args{Clause{0, In, []float32{10, 1, 2}}, float32(10)}, true},
		{args{Clause{0, In, []float32{1, 2}}, float32(10)}, false},
		{args{Clause{0, In, []float32{}}, float32(10)}, false},
		{args{Clause{0, In, nil}, float32(10)}, false},
	}
	for _, tt := range tests {
		if got := FloatClauseResolver(tt.args.clause, tt.args.val); got != tt.want {
			t.FailNow()
		}
	}
}

func TestTimeClauseResolver(t *testing.T) {
	type args struct {
		clause Clause
		val    time.Time
	}

	now := time.Now()
	nowPlMin := time.Now().Add(time.Minute)

	tests := []struct {
		args args
		want bool
	}{
		{args{Clause{0, Eq, now}, now}, true},
		{args{Clause{0, Eq, now}, nowPlMin}, false},
		{args{Clause{0, NEq, now}, now}, false},
		{args{Clause{0, Gt, nowPlMin}, now}, false},
		{args{Clause{0, Lt, nowPlMin}, now}, true},
	}
	for _, tt := range tests {
		if got := TimeClauseResolver(tt.args.clause, tt.args.val); got != tt.want {
			t.FailNow()
		}
	}
}

func TestBoolClauseResolver(t *testing.T) {
	type args struct {
		clause Clause
		val    bool
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{Clause{0, Eq, true}, true}, true},
		{args{Clause{0, Eq, false}, false}, true},
		{args{Clause{0, Eq, false}, true}, false},
		{args{Clause{0, NEq, false}, true}, true},
		{args{Clause{0, In, []bool{}}, true}, false},
		{args{Clause{0, In, []bool{false}}, true}, false},
		{args{Clause{0, In, []bool{false}}, false}, true},
	}
	for _, tt := range tests {
		if got := BoolClauseResolver(tt.args.clause, tt.args.val); got != tt.want {
			t.FailNow()
		}
	}
}

func TestStringClauseResolver(t *testing.T) {
	type args struct {
		clause Clause
		val    string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{Clause{0, Eq, ""}, ""}, true},
		{args{Clause{0, Eq, "str"}, "str"}, true},
		{args{Clause{0, Eq, "Str"}, "str"}, false},
		{args{Clause{0, NEq, "str"}, ""}, true},
		{args{Clause{0, In, []string{}}, ""}, false},
		{args{Clause{0, In, []string{"str"}}, "str"}, true},
		{args{Clause{0, In, []string{"STR"}}, "str"}, false},
		{args{Clause{0, Like, ""}, ""}, true},
		{args{Clause{0, Like, "A"}, "abc"}, true},
		{args{Clause{0, Like, "AC"}, "abc"}, false},
		{args{Clause{0, Like, "bC"}, "abc"}, true},
	}
	for _, tt := range tests {
		if got := StringClauseResolver(tt.args.clause, tt.args.val); got != tt.want {
			t.FailNow()
		}
	}
}
