GPT-2 in C++ - Top Level Project Skeleton
=========================================

This is a high-level overview of the project structure for porting GPT-2 to C++.  Keep in mind this is a simplified representation and may require additional components depending on your specific needs.

## Environment / Toolchain
  * Commandline Arg Parser
  * Color output (AnsiColors)

### Headers
  * Install basic dev stuff: `sudo apt install build-essential dkms linux-headers-$(uname -r)`

### Linear Algebra Library:
  * Linear Algebra library: `Eigen` (for efficient matrix stuff)
  * Install with `sudo apt install libeigen3-dev`

### CUDA Integration:
  * Install with `apt-get install libnvvm4 nvidia-cuda-toolkit`

### Testing Framework:
  * Google Test

### Build System:
  * Run in X86_64 Linux Platform only
  * All artifacts built to run in Docker
    * Docker base image
    * Docker build stage
    * Docker test stage
    * Docker run stage
  * Use `cmake` to automate functions.
  * use g++ and C++ 11 (ease and compatibility if we go to ARM64)

## Project Structure:

### main.cpp:
  Entry point for the program. This could demonstrate loading a pre-trained model, performing inference on a sample input, and generating text.

### gpt2_model:
  (Contains core model definition and functionalities)

#### gpt2_layer.h:
  Header file defining the GPT-2 layer architecture (transformer block).

#### gpt2_layer.cpp:
  Implementation file for the GPT-2 layer functionalities.

#### gpt2_model.h:
  Header file defining the overall GPT-2 model architecture.

#### gpt2_model.cpp:
  Implementation file for the GPT-2 model, including building the layers and functionalities like forward pass and backward pass.

### utils: (Contains utility functions)
#### tensor.h:
  Header file defining the tensor class for representing data (may leverage existing libraries like Eigen).

#### tensor.cpp:
  Implementation file for the tensor class functionalities.

#### io.h:
  Header file for input/output functionalities (loading pre-trained model weights, text processing).

#### io.cpp:
  Implementation file for input/output functionalities.

### inference.h:
  Header file defining the inference API for using the GPT-2 model for text generation.

### inference.cpp:
  Implementation file for the inference API functionalities.
