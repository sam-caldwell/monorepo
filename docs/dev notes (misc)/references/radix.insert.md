Radix Tree Insert
=================

## Objective 
Define steps for radix tree insert operation.

## Process

### Starting at the Root
Insertion begins at the root of the radix tree. If the tree is empty, 
a new root node is created. 

### Traversing the Tree
For each character in the key being inserted, the algorithm traverses 
the tree, starting from the root. It follows the path corresponding 
to the characters in the key. 

### Creating Nodes
Along the traversal path, if a node corresponding to the current character 
does not exist, a new node is created. If a node already exists for the character, the algorithm simply moves to the next character in the key and continues the traversal. 

### Termination
Once the algorithm reaches the end of the key, it marks the last node as
a terminal node (indicating the end of a valid key) and optionally stores any associated value with it. 

### Handling Overlapping Prefixes
Radix trees optimize storage by sharing common prefixes among keys. If 
during the insertion process it encounters a node where the key being 
inserted shares a common prefix with an existing key, it may reuse the 
existing nodes to represent the shared prefix. This optimization helps 
in reducing memory consumption and improving performance. 

### Complexity
The time complexity of insertion in a radix tree is O(k), where k is 
the length of the key being inserted. This is because the algorithm 
needs to traverse the tree from the root to the leaf, considering each 
character in the key. 

### Balancing
Unlike some other tree structures like binary search trees, radix 
trees don't typically require balancing operations because they maintain 
a fixed branching factor (usually the size of the character set, such as
26 for lowercase English alphabets). This makes insertion and other 
operations more predictable in terms of time complexity.
