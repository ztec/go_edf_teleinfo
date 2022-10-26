package go_edf_teleinfo

// Teleinfo (all) data returnable by EDF teleinfo. Remember that I cannot test all cases,
// meaning some labels are missing. Please contribute to add them
type Teleinfo struct {
	OPTARIF string `json:"OPTARIF"` //abonement
	ISOUSC  int64  `json:"ISOUSC"`  //abonement_puissance
	HCHC    int64  `json:"HCHC"`    //index_heures_creuses
	HCHP    int64  `json:"HCHP"`    //index_heures_pleines
	IINST   int64  `json:"IINST"`   //intensitee_instantanee
	IMAX    int64  `json:"IMAX"`    //intensitee_max
	PAPP    int64  `json:"PAPP"`    //puissance_apparente
	HHPHC   string `json:"HHPHC"`   //groupe_horaire
	PTEC    string `json:"PTEC"`    //Période Tarifaire en cours
	RAW     []byte `json:"RAW"`     //Raw edf teleinfo payload
}
