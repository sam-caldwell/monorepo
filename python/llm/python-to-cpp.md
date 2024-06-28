Notes on Porting GPT-2 to C++
=============================

## Dependencies to port?
* Hugging Face Transformers
* Pytorch
* NumPy


## Effort Overview
The following efforts would be required to port GPT-2 to C++ from Python:

1. Understanding GPT-2 Architecture:
  1. Study the original implementation in Python.
  2. Understand the model architecture, layers, and training procedures.

2. Core Library Development:
  1. Develop or port necessary libraries for tensor operations (e.g., matrix multiplications, convolutions).
  2. Implement or integrate existing libraries for autograd (automatic differentiation).

3. Model Implementation:
  1. Recreate the GPT-2 model architecture in Go.
  2. Implement all components such as multi-head attention, feed-forward neural networks, layer normalization, and positional encodings.

4. Data Handling:
  1. Implement data preprocessing utilities.
  2. Develop or port tokenization and encoding/decoding methods.

5. Training and Inference:
  1. Implement training loop and backpropagation.
  2. Optimize inference performance.

6. Optimization and Performance:
  1. Optimize code for performance, considering Go's concurrency model.
  2. Ensure efficient memory usage and garbage collection.

7. Testing and Validation:
  1. Validate the model against the original implementation to ensure correctness.
  2. Implement comprehensive testing (unit tests, integration tests).

8. Documentation and Examples:
  1. Document the code and usage instructions.
  2. Provide examples and tutorials for training and using the model.

9. Community and Ecosystem Integration:
  1. Engage with the Go community for support and feedback.
  2. Consider contributing to or leveraging existing Go ML libraries.

10. Maintenance and Updates:
  1. Regularly update the codebase to keep up with advancements in NLP and Go.
  2. Fix bugs and improve features based on user feedback.


## Model Architecture
The Transformer decoder architecture, as used in models like GPT-2, is a key component for generating text in an autoregressive manner. 

Here's an explanation of its components and functionality:

### Input Embeddings:
1. Token Embeddings: Converts input tokens (words or subwords) into dense vectors of fixed size.
2. Positional Embeddings: Adds positional information to the token embeddings to incorporate the order of the sequence.

To illustrate, let's look at the Hugging Face `transformers` library.  The code examples are in python and use PyTorch, which would need to be translated to golang for a full port.

The following illustrates how token and positional embeddings are handled:
```python
import torch
import torch.nn as nn

class GPT2Embeddings(nn.Module):
    def __init__(self, vocab_size, hidden_size, max_position_embeddings):
        super().__init__()
        self.word_embeddings = nn.Embedding(vocab_size, hidden_size)
        self.position_embeddings = nn.Embedding(max_position_embeddings, hidden_size)

    def forward(self, input_ids, position_ids=None):
        if position_ids is None:
            position_ids = torch.arange(input_ids.size(-1), dtype=torch.long, device=input_ids.device)
            position_ids = position_ids.unsqueeze(0).expand_as(input_ids)
        
        input_embeddings = self.word_embeddings(input_ids)
        position_embeddings = self.position_embeddings(position_ids)
        return input_embeddings + position_embeddings
```



### Stack of Decoder Layers:
* The Transformer decoder consists of multiple identical layers (e.g., 12, 24, or more layers).

### Components of a Single Decoder Layer:
1. Masked Multi-Head Self-Attention:
  * Self-Attention Mechanism:
    *  Allows each position in the input sequence to attend to all positions in the sequence before it, enabling the model to capture dependencies across the entire sequence.
  * Masked Attention:
    * Ensures that the prediction for a position only depends on the known outputs up to that position, preventing information leakage from future tokens.
  * Multi-Head:
    * Uses multiple attention heads to allow the model to focus on different parts of the sequence simultaneously, capturing various aspects of the input.
2. Layer Normalization:
  * Applied before and after the self-attention mechanism and feed-forward network. It stabilizes and accelerates training by normalizing the inputs to these sub-layers.

3. Residual Connections:
  * Skip connections that add the input of a sub-layer to its output. These help mitigate the vanishing gradient problem and improve the flow of gradients during backpropagation.

4. Feed-Forward Neural Network:
  * Consists of two linear transformations with a ReLU activation in between. Applied to each position independently, it allows for complex transformations of the input representations.

The following illustrates the `MultiHeadSelfAttention`:
```python
class MultiHeadSelfAttention(nn.Module):
    def __init__(self, hidden_size, num_attention_heads):
        super().__init__()
        self.num_attention_heads = num_attention_heads
        self.attention_head_size = hidden_size // num_attention_heads
        self.all_head_size = self.num_attention_heads * self.attention_head_size

        self.query = nn.Linear(hidden_size, self.all_head_size)
        self.key = nn.Linear(hidden_size, self.all_head_size)
        self.value = nn.Linear(hidden_size, self.all_head_size)
        self.out = nn.Linear(hidden_size, hidden_size)

    def forward(self, hidden_states, attention_mask=None):
        mixed_query_layer = self.query(hidden_states)
        mixed_key_layer = self.key(hidden_states)
        mixed_value_layer = self.value(hidden_states)

        query_layer = self.transpose_for_scores(mixed_query_layer)
        key_layer = self.transpose_for_scores(mixed_key_layer)
        value_layer = self.transpose_for_scores(mixed_value_layer)

        attention_scores = torch.matmul(query_layer, key_layer.transpose(-1, -2))
        attention_scores = attention_scores / math.sqrt(self.attention_head_size)

        if attention_mask is not None:
            attention_scores = attention_scores + attention_mask

        attention_probs = nn.Softmax(dim=-1)(attention_scores)
        context_layer = torch.matmul(attention_probs, value_layer)

        context_layer = context_layer.permute(0, 2, 1, 3).contiguous()
        new_context_layer_shape = context_layer.size()[:-2] + (self.all_head_size,)
        context_layer = context_layer.view(*new_context_layer_shape)

        output = self.out(context_layer)
        return output

    def transpose_for_scores(self, x):
        new_x_shape = x.size()[:-1] + (self.num_attention_heads, self.attention_head_size)
        x = x.view(*new_x_shape)
        return x.permute(0, 2, 1, 3)
```
And this illustrates the `Feed-Forward Neural Network`:
```python
class FeedForward(nn.Module):
    def __init__(self, hidden_size, intermediate_size, activation_fn):
        super().__init__()
        self.dense_1 = nn.Linear(hidden_size, intermediate_size)
        self.dense_2 = nn.Linear(intermediate_size, hidden_size)
        self.activation_fn = activation_fn

    def forward(self, hidden_states):
        hidden_states = self.dense_1(hidden_states)
        hidden_states = self.activation_fn(hidden_states)
        hidden_states = self.dense_2(hidden_states)
        return
```



#### Example of a Single Decoder Layer
Here's a simplified view of one layer in the Transformer decoder:

```text

Input Embedding
      |
Positional Encoding
      |
    +--------------------+
    |    Layer Norm      |
    |     (Input)        |
    +---------|----------+
              |
    Masked Multi-Head Self-Attention
              |
    +---------|----------+
    |    Layer Norm      |
    |  (Post Attention)  |
    +---------|----------+
              |
    Feed-Forward Neural Network
              |
    +---------|----------+
    |    Layer Norm      |
    |  (Post Feed-Forward)|
    +--------------------+
              |
            Output
```

## Processing Steps in the Decoder
### Input and Positional Embeddings:
 * The input tokens are first converted to embeddings, which are then combined with positional encodings to incorporate the order of the sequence.

### Masked Multi-Head Self-Attention:
Each token attends to all previous tokens in the sequence through the self-attention mechanism, with masking applied to prevent attending to future tokens.

### Feed-Forward Neural Network:
The output from the self-attention mechanism is passed through a feed-forward neural network to add non-linearity and further transform the embeddings.

### Layer Normalization and Residual Connections:
Applied after both the self-attention and feed-forward layers, these techniques help in stabilizing and improving the training process.

### Stacking Layers:
The output of one decoder layer is fed as input to the next layer. This stacking allows the model to learn hierarchical representations.

### Output Projection:
The final output is projected onto the vocabulary size using a linear layer, followed by a softmax function to generate probabilities for the next token in the sequence.

### Autoregressive Generation
* Autoregressive Nature: 
  * During generation, the decoder produces one token at a time, conditioning on the previously generated tokens. This process is repeated until the desired sequence length or an end-of-sequence token is generated.

By understanding these components and processing steps, one can appreciate how the Transformer decoder architecture enables powerful and efficient sequence generation, as seen in models like GPT-2.
