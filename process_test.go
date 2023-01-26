package dataparse

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var available = map[string]string{
	"TextMap_en":             "2578607577_BD97DB82911325B6BF85A5814F81121E87320C1E.unity3d",
	"DialogueData_en":        "816421683_538CD02972BF773E137A948B782BA1201D69E373.unity3d",
	"DormitoryFurnitureData": "606639383_85D489170438C0E9771B16D847D3C42973D6FDD2.unity3d",
	"DormitoryEventSequence": "1435715286_9AE44CFA4F10A04BC7CEB9BE0C0C1E179ACC02C4.unity3d",
	"StigmataPositionData":   "612893459_F3E9E7FCE5C79743B12ECEBC321D9C21327C3241.unity3d",
	"OWActivityBossData":     "11165326_78DF5976C4C27D635A80D264FFAC0B7A5882DFC4.unity3d",
	"WorldMap":               "36243594_FAFCC74C44BCFF00CAA3DB2C5783EA4AE9C1C636.unity3d",
	"BossRushBuffList":       "53370678_63F590EFFB7D0B6F03E78839B9128992A372F604.unity3d",

	"50428269":  "50428269_009441833134EA893A25ABB6416CB2CF1391ED70.unity3d",
	"112154430": "112154430_010083C199E5CE5CFE3055B7602B1875CD67724A.unity3d",
	"137850209": "137850209_E79B5956D3694332DFD47C23FDA4BE34F8812840.unity3d",
	"150008597": "150008597_06713F90473CEAE5C50EFE98AC1D19BB43EEE201.unity3d",
	"169620141": "169620141_BC9355EAFF565E4912611978482A74084F0BA04A.unity3d",
	"175952655": "175952655_0420D78F0FC7E3C9C8B3885BB39E8CD3B76DE233.unity3d",
	"201380970": "201380970_A51004E07DCBD51B75A01F4885FD984809A3CA6A.unity3d",
	"208715169": "208715169_4ACD72DA8A882D4D56015997876327E2AA466BF1.unity3d",
	"213532048": "213532048_937657EA9456DBDDE21EADC47466E1201AA489CD.unity3d",
}

func TestProcessStructAll(t *testing.T) {
	for name, file := range available {
		t.Run(name, func(t *testing.T) {
			proc(t, name, file)
		})
	}
}

func TestProcess(t *testing.T) {
	name := "213532048"

	file := available[name]

	proc(t, name, file)
}

func proc(t *testing.T, name, file string) {
	a := GetAsset(file)

	result, err := ProcessStruct(file, a.Parser)
	if err != nil {
		t.Errorf("ProcessStruct() error = %v", err)
		return
	}

	out := filepath.Join("result", fmt.Sprintf("%s.json", a.Name))

	err = os.WriteFile(out, result, os.ModePerm)
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}
}
