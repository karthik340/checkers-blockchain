package keeper

import (
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) RemoveFromFifo(ctx sdk.Context, game *types.StoredGame, info *types.NextGame) {
	if game.BeforeId != types.NoFifoIdKey {
		beforeElement, found := k.GetStoredGame(ctx, game.BeforeId)
		if !found {
			panic("element before in fifo was not found")
		}
		beforeElement.AfterId = game.AfterId
		k.SetStoredGame(ctx, beforeElement)
		if game.AfterId == types.NoFifoIdKey {
			info.FifoTail = beforeElement.Index
		}
	}

	if game.AfterId != types.NoFifoIdKey {
		afterElement, found := k.GetStoredGame(ctx, game.AfterId)
		if !found {
			panic("element after in fifo was not found")
		}
		afterElement.BeforeId = game.BeforeId
		k.SetStoredGame(ctx, afterElement)
		if game.BeforeId == types.NoFifoIdKey {
			info.FifoHead = afterElement.Index
		}
	}
	game.BeforeId = types.NoFifoIdKey
	game.AfterId = types.NoFifoIdKey
}

func (k Keeper) sendToFifoTail(ctx sdk.Context, game *types.StoredGame, info *types.NextGame) {
	if info.FifoHead == types.NoFifoIdKey && info.FifoTail == types.NoFifoIdKey {
		game.BeforeId = types.NoFifoIdKey
		game.AfterId = types.NoFifoIdKey
		info.FifoHead = game.Index
		info.FifoTail = game.Index
	} else if info.FifoHead == types.NoFifoIdKey || info.FifoTail == types.NoFifoIdKey {
		panic("fifo should have both head and tail or none")
	} else if info.FifoTail == game.Index {

	} else {
		k.RemoveFromFifo(ctx, game, info)
		curentTail, found := k.GetStoredGame(ctx, info.FifoTail)
		if !found {
			panic("current fifo tail was not found")
		}
		curentTail.AfterId = game.Index
		game.BeforeId = curentTail.Index
		info.FifoTail = game.Index
	}

}
