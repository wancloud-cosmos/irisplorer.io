package m

import (
	"github.com/cosmos/cosmos-sdk/modules/coin"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

const (
	DocsNmStakeTx = "stake_tx"
)

//Stake交易
type StakeTx struct {
	TxHash string    `json:"tx_hash" bson:"tx_hash"`
	Time   time.Time `json:"time" bson:"time"`
	Height int64     `json:"height" bson:"height"`
	From   string    `json:"from" bson:"from"`
	PubKey string    `json:"pub_key" bson:"pub_key"`
	Type   string    `json:"type" bson:"type"`
	Amount coin.Coin `json:"amount" bson:"amount"`
}

func (c StakeTx) Name() string {
	return DocsNmStakeTx
}

func (c StakeTx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": c.TxHash}
}

func (c StakeTx) Index() mgo.Index {
	return mgo.Index{
		Key:        []string{"from"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     false,            // 唯一索引 同mysql唯一索引
		DropDups:   false,            // 索引重复替换旧文档,Unique为true时失效
		Background: true,             // 后台创建索引
	}
}

func QueryAllStakeTx() []StakeTx {
	result := []StakeTx{}
	query := func(c *mgo.Collection) error {
		err := c.Find(nil).Sort("-time").All(&result)
		return err
	}

	if store.ExecCollection(DocsNmStakeTx, query) != nil {
		log.Printf("CoinTx is Empry")
	}
	return result
}

//
func QueryStakeTxsByAccount(account string) []StakeTx {
	result := []StakeTx{}
	query := func(c *mgo.Collection) error {
		err := c.Find(bson.M{"$or": []bson.M{{"from": account}, {"to": account}}}).Sort("-time").All(&result)
		return err
	}

	if store.ExecCollection(DocsNmStakeTx, query) != nil {
		log.Printf("StakeTx is Empry")
	}

	return result
}

func QueryPageStakeTxsByAccount(account string, page, pagesize int) []StakeTx {
	result := []StakeTx{}
	query := func(c *mgo.Collection) error {
		skip := (page - 1) * pagesize
		err := c.Find(bson.M{"$or": []bson.M{{"from": account}, {"to": account}}}).Sort("-time").Skip(skip).Limit(pagesize).All(&result)
		return err
	}

	if store.ExecCollection(DocsNmStakeTx, query) != nil {
		log.Printf("StakeTx is Empry")
	}

	return result
}

func QueryPageStakeTxs(page, pagesize int) []StakeTx {
	result := []StakeTx{}
	query := func(c *mgo.Collection) error {
		skip := (page - 1) * pagesize
		err := c.Find(nil).Sort("-time").Skip(skip).Limit(pagesize).All(&result)
		return err
	}

	if store.ExecCollection(DocsNmStakeTx, query) != nil {
		log.Printf("StakeTx is Empry")
	}

	return result
}
