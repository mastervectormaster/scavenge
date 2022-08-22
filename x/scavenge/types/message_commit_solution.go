package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCommitSolution = "commit_solution"

var _ sdk.Msg = &MsgCommitSolution{}

func NewMsgCommitSolution(creator string, solutionHash string, solutionScavengeHash string) *MsgCommitSolution {
	return &MsgCommitSolution{
		Creator:              creator,
		SolutionHash:         solutionHash,
		SolutionScavengeHash: solutionScavengeHash,
	}
}

func (msg *MsgCommitSolution) Route() string {
	return RouterKey
}

func (msg *MsgCommitSolution) Type() string {
	return TypeMsgCommitSolution
}

func (msg *MsgCommitSolution) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCommitSolution) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCommitSolution) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
