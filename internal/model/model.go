package model

import (
	"time"
	//"fmt"
)

//type JSONTime time.Time

/* func (t JSONTime)MarshalJSON() ([]byte, error) {
    //do your serializing here
    stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
    return []byte(stamp), nil
} */

type Contract struct {
	Id        uint64    `json:"id"` //
	NN        string    `json:"number"`
	Owner     uint64    `json:"owner"`
	OwnerName string    `json:"name"`
	Begin     time.Time `json:"begin"`
	End       time.Time `json:"end"`
	Status    uint64    `json:"status"`
	StatusH   uint64    `json:"statusHistory"`
	DateH     time.Time `json:"dateHistory"`
}

func CreateContract(id *uint64, n *string, owner *uint64, name *string, beg, end *time.Time, status, statH *uint64, dH *time.Time) *Contract {
	t := new(Contract)
	if id != nil {
		t.Id = *id
	}
	if n != nil {
		t.NN = *n
	}
	if owner != nil {
		t.Owner = *owner
	}
	if name != nil {
		t.OwnerName = *name
	}
	if beg != nil {
		t.Begin = *beg
	}
	if end != nil {
		t.End = *end
	}
	if status != nil {
		t.Status = *status
	}
	if statH != nil {
		t.StatusH = *statH
	}
	if dH != nil {
		t.DateH = *dH
	}

	return t
}

type Contracts struct {
	c []Contract
}

func New() Contracts {
	c := Contracts{c: make([]Contract, 0, 10)}
	return c
}

func (con *Contracts) Add(t ...Contract) {
	con.c = append(con.c, t...)
}

func (con *Contracts) Get() []Contract {
	return con.c
}

func (con *Contracts) Len() int {
	return len(con.c)
}
