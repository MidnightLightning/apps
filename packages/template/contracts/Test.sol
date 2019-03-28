pragma solidity ^0.4.24;

contract Test {  

    string public text;

    function setText(string _text) public {
        text = _text;
    }

    function getText() public view returns (string) {
        return text;
    }

}
