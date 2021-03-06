package m

import (
	"github.com/irisnet/irisplorer.io/server/modules/store"
)

func init() {
	store.RegisterDocs(new(Account))
	store.RegisterDocs(new(Block))
	store.RegisterDocs(new(CoinTx))
	store.RegisterDocs(new(Delegator))
	store.RegisterDocs(new(StakeTx))
	store.RegisterDocs(new(SyncTask))
	store.RegisterDocs(new(Candidate))
}
