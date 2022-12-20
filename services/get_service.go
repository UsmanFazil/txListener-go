package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/block-listener/models"
	"github.com/block-listener/models/mysql"
	"github.com/pkg/errors"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	useraddress := r.URL.Query().Get("userAddress")
	fmt.Println("useraddress =>", useraddress)

	if len(useraddress) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		writeJSON(w, nil, errors.New("userAddress Parameter Required"))
		return
	}

	resp, err := mysql.SharedStore().GetTxBurnbyUserAddr(useraddress)
	fmt.Println("resp, err:", resp, err, len(resp))
	if err != nil || len(resp) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		writeJSON(w, nil, errors.New("record not found"))
		return
	}
	var getAllInfo []*models.GetInfoVo

	for _, product := range resp {
		getAllInfo = append(getAllInfo, models.UserInfoVo(product))
	}

	writeJSON(w, getAllInfo, nil)
	return
}

func GetTxwithSignature(w http.ResponseWriter, r *http.Request) {
	txHash := r.URL.Query().Get("txHash")
	fmt.Println("txHash =>", txHash)

	if len(txHash) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		writeJSON(w, nil, errors.New("txHash Parameter Required"))
		return
	}

	resp, err := mysql.SharedStore().GetTxByHash(txHash)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		writeJSON(w, nil, err)
		return
	}
	respBurn, err := mysql.SharedStore().GetTxBurnInfo(txHash)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		writeJSON(w, nil, err)
		return
	}

	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")

	writeJSON(w, struct {
		Contractaddress string `json:"Contractaddress"`
		Originchainid   int64  `json:"Originchainid"`
		Tochainid       int64  `json:"Tochainid"`
		Amount          string `json:"Amount"`
		Transaction     string `json:"Transaction"`
		Signature       string `json:"Signature"`
		UserAddress     string `json:"UserAddress"`
	}{
		Contractaddress: resp.Contractadd,
		Originchainid:   respBurn.Originchainid,
		Tochainid:       respBurn.Tochainid,
		Amount:          respBurn.Amount,
		Transaction:     respBurn.Txhash,
		Signature:       respBurn.Signature,
		UserAddress:     respBurn.Address,
	}, nil)

	if err != nil {
		fmt.Println(err)
	}
	return
}

type errResp struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

func writeJSON(w http.ResponseWriter, v interface{}, err error) {
	var respVal interface{}
	if err != nil {
		msg := err.Error()

		w.WriteHeader(http.StatusBadRequest)
		var e errResp
		e.Error.Message = msg
		respVal = e
	} else {
		respVal = v
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(respVal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("io.Copy: %v", err)
		return
	}
}
