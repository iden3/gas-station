package endpoint

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/iden3/tx-forwarder/eth"
	log "github.com/sirupsen/logrus"
)

func fail(c *gin.Context, err error) {
	log.Error("error: " + err.Error())
	c.JSON(400, gin.H{
		"error": err.Error(),
	})
}

func handleGetInfo(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func handleGetTx(c *gin.Context) {
	txHashString := c.Param("txhash")
	txHash := common.HexToHash(txHashString)
	tx, receipt, isPending, err := ethsrv.GetTx(txHash)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"tx":        tx,
		"receipt":   receipt,
		"isPending": isPending,
	})
}

func handlePostTxSampleContract(c *gin.Context) {
	var d eth.SampleCallData
	err := c.BindJSON(&d)
	if err != nil {
		fail(c, err)
		return
	}

	ethTx, err := ethsrv.ForwardTxToSampleContract(d)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"ethTx": ethTx.Hash().Hex(),
	})
}

func handlePostTxFullVerifier(c *gin.Context) {
	var d eth.FullVerifierCallData
	err := c.BindJSON(&d)
	if err != nil {
		fail(c, err)
		return
	}

	ethTx, err := ethsrv.ForwardTxToFullVerifierContract(d)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"ethTx": ethTx.Hash().Hex(),
	})
}

func handlePostTxDisableId(c *gin.Context) {
	var d eth.DisableIdCallData
	err := c.BindJSON(&d)
	if err != nil {
		fail(c, err)
		return
	}

	ethTx, err := ethsrv.ForwardTxToDisableIdContract(d)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"ethTx": ethTx.Hash().Hex(),
	})
}

func handleGetIdInWhitelist(c *gin.Context) {
	ethAddrString := c.Param("ethaddr")
	ethAddr := common.HexToAddress(ethAddrString)

	res, err := ethsrv.GetIdInWhitelist(ethAddr)
	if err != nil {
		fail(c, err)
		return
	}

	c.JSON(200, gin.H{
		"inWhitelist": res,
	})
}
