package clients

type DataField struct {
	Key   string
	Value string
}

type Request struct {
	Service string
	Data    []DataField
}

type Response struct {
	Service string
	Status  string
	Data    []DataField
}
