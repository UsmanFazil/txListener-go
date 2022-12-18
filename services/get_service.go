package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/block-listener/models"
	"github.com/block-listener/models/mysql"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	useraddress := r.URL.Query().Get("useraddress")
	fmt.Println("useraddress =>", useraddress)

	if len(useraddress) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "No response")
		return
	}

	resp, err := mysql.SharedStore().GetTxBurnbyUserAddr(useraddress)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "No response")
		return
	}
	var getAllInfo []*models.GetInfoVo

	for _, product := range resp {
		getAllInfo = append(getAllInfo, models.UserInfoVo(product))
	}
	body, err := json.Marshal(getAllInfo)

	if err != nil {
		fmt.Println(err)
	}

	w.Write(body)
	return
}

func GetTxwithSignature(w http.ResponseWriter, r *http.Request) {
	txHash := r.URL.Query().Get("txHash")
	fmt.Println("txHash =>", txHash)

	if len(txHash) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "No response")
		return
	}

	resp, err := mysql.SharedStore().GetTxByHash(txHash)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "No response")
		return
	}
	respBurn, err := mysql.SharedStore().GetTxBurnInfo(txHash)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "No response")
		return
	}

	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(struct {
		Contractaddress string `json:"Contractaddress"`
		Originchainid   int64  `json:"Originchainid"`
		Tochainid       int64  `json:"Tochainid"`
		Amount          string `json:"Amount"`
		Transaction     string `json:"Transaction"`
		Signature       string `json:"Signature"`
	}{
		Contractaddress: resp.Contractadd,
		Originchainid:   respBurn.Originchainid,
		Tochainid:       respBurn.Tochainid,
		Amount:          respBurn.Amount,
		Transaction:     respBurn.Txhash,
		Signature:       respBurn.Signature,
	})

	if err != nil {
		fmt.Println(err)
	}

	w.Write(body)

	return
}
