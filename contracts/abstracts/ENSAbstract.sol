pragma solidity ^0.4.0;


contract ENSAbstract {
    function owner(bytes32 node) constant returns (address);
    function resolver(bytes32 node) constant returns (address);
    function ttl(bytes32 node) constant returns (uint64);
    function setOwner(bytes32 node, address owner);
    function setSubnodeOwner(bytes32 node, bytes32 label, address owner);
    function setResolver(bytes32 node, address resolver);
    function setTTL(bytes32 node, uint64 ttl);
}


contract ResolverAbstract {
    function content(bytes32 node) constant returns (bytes32 content);
}
