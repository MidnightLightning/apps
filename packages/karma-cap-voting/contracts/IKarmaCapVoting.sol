pragma solidity ^0.4.24;

interface IKarmaCapVoting {
    function CREATE_VOTES_ROLE() constant public returns(bytes32);
    function initialize(
        address _token,
        address _karma,
        uint64 _supportRequiredPct,
        uint64 _minAcceptQuorumPct,
        uint64 _voteTime
    );
}
