package keeper

import (
	"context"
	"strings"

	"github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RejectGame(goCtx context.Context, msg *types.MsgRejectGame) (*types.MsgRejectGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Handling the message
	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.IdValue)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game not found %s", msg.IdValue)
	}
	if storedGame.Winner != rules.NO_PLAYER.Color {
		return nil, types.ErrGameFinished
	}
	if strings.Compare(storedGame.Red, msg.Creator) == 0 {
		if 1 < storedGame.MoveCount {
			return nil, types.ErrRedAlreadyPlayed
		}
	} else if strings.Compare(storedGame.Black, msg.Creator) == 0 {
		if 1 < storedGame.MoveCount {
			return nil, types.ErrBlackAlreadyPlayed
		}
	} else {
		return nil, types.ErrCreatorNotPlayer
	}
	k.Keeper.MustRefundWager(ctx, &storedGame)
	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("Nextgame not found")
	}
	k.Keeper.RemoveFromFifo(ctx, &storedGame, &nextGame)
	k.Keeper.RemoveStoredGame(ctx, msg.IdValue)
	ctx.GasMeter().ConsumeGas(types.RejectGameGas, "Reject game")
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "checkers"),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.RejectGameEventKey),
			sdk.NewAttribute(types.RejectGameEventCreator, msg.Creator),
			sdk.NewAttribute(types.RejectGameEventIdValue, msg.IdValue),
		),
	)

	return &types.MsgRejectGameResponse{}, nil
}
