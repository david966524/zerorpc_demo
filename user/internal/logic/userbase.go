package logic

type localUser struct {
	id   int64
	name string
	desc string
}

var users = map[int]*localUser{
	1: &localUser{
		id:   1,
		name: "david",
		desc: "david user",
	},
	2: &localUser{
		id:   2,
		name: "baha",
		desc: "baha user",
	},
	3: &localUser{
		id:   1,
		name: "jack",
		desc: "jack user",
	},
}
