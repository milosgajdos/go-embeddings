# go-embeddings

[![Build Status](https://github.com/milosgajdos/go-embeddings/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/milosgajdos/go-embeddings/actions?query=workflow%3ACI)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/milosgajdos/go-embeddings)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This project provides an implementation of API clients for fetching embeddings from various LLM providers.

Currently supported APIs:
* [x] [OpenAI](https://platform.openai.com/docs/api-reference/embeddings)
* [x] [Cohere](https://docs.cohere.com/reference/embed)
* [x] [Google Vertex](https://cloud.google.com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings)
* [x] [VoyageAI](https://docs.voyageai.com/reference/embeddings-api)
* [x] [Ollama](https://ollama.com/)
* [x] [AWS Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/titan-embedding-models.html)

You can find sample programs that demonstrate how to use the client packages to fetch the embeddings in `cmd` directory of this project.

Finally, the `document` package provides an implementation of simple document text splitters, heavily inspired by the popular [Langchain framework](https://github.com/langchain-ai/langchain).
It's essentially a Go rewrite of character and recursive character text splitters from the Langchain framework with minor modifications, but more or less identical results.

## Environment variables

> [!NOTE]
> Each client package lets you initialize a default API client for a specific embeddings provider by reading the API keys from environment variables

Here's a list of the env vars for each supported client

### OpenAI

* `OPENAI_API_KEY`: Open AI API token

### Cohere

* `COHERE_API_KEY`: Cohere API token

### Google Vertex AI

* `VERTEXAI_TOKEN`: Google Vertex AI API token (can be fetch by `gcloud auth print-access-token` once you've authenticated)
* `VERTEXAI_MODEL_ID`: Embeddings model (at the moment only `textembedding-gecko@00` or `multimodalembedding@001` are available)
* `GOOGLE_PROJECT_ID`: Google Project ID
* `VOYAGE_API_KEY`: VoyageAI API key

### Voyage

* `VOYAGE_API_KEY`: Voyage AI API key

### AWS Bedrock

> [!IMPORTANT]
> You must enable access to Bedrock embedding models
> See here: [https://docs.aws.amazon.com/bedrock/latest/userguide/model-access.html#add-model-access](https://docs.aws.amazon.com/bedrock/latest/userguide/model-access.html#add-model-access)

* `AWS_REGION`: AWS region

Usual AWS env vars as read by the AWS SDKs i.e. `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, etc.

## nix

The project provides a simple `nix` flake tha leverages [gomod2nix](https://github.com/nix-community/gomod2nix) for consistent Go environments and builds.

To get started just run
```shell
nix develop
```

And you'll be dropped into development shell.

In addition, each command is exposed as a `nix` app so you can run them as follows:
```shell
nix run ".#vertexai" -- -help
```

**NOTE:** `gomod2nix` vendors dependencies into `nix` store so every time you add a new dependency you must run `gomod2nix generate` that updates `gomod2nix.toml`

# Contributions

Yes please!
