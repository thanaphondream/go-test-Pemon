
package repository

type View struct {
    Time   string `json:"Time"`
    ACSTR  int    `json:"ACSTR"`
    DCSTR  int    `json:"DCSTR"`
    IOSTR  int    `json:"IOSTR"`

    ACs  []int `json:"ACs"`
    DCF  []int `json:"DCF"`
    ACF  []int `json:"ACF"`
    IOF  []int `json:"IOF"`
    HiF  []int `json:"HiF_"`
    LoD  []int `json:"LoD_"`
    LoF  []int `json:"LoF_"`

    VDC  []int `json:"VDC_"`
    IDC  []int `json:"IDC_"`
    ACtype []int `json:"ACtype"`

    VAC  []any `json:"VAC"`
    FVAC []any `json:"FVAC"`
    IAC  []any `json:"IAC"`
    SYSF []int `json:"SYSF"`

    TT  []any `json:"TT"`
    H   []any `json:"H"`
    SHT []any `json:"SHT"`
    S0  []any `json:"S0"`

    Vi   []any `json:"Vi"`
    Vo   []any `json:"Vo"`
    Pi   []any `json:"Pi"`
    Po   []any `json:"Po"`
    Vd   []any `json:"Vd"`
    Id   []any `json:"Id"`
    Trem []any `json:"Trem"`

    EqTyp int `json:"EqTyp"`
    EqS   int `json:"EqS"`
    EqDP  int `json:"EqDP"`
    Et    int `json:"Et"`
    EqR   int `json:"EqR"`
}


type RerositoryView interface {
	View() ([]View, error)
}