# go-embeddings

[![Build Status](https://github.com/milosgajdos/go-embeddings/workflows/CI/badge.svg)](https://github.com/milosgajdos/go-embeddings/actions?query=workflow%3ACI)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/milosgajdos/go-embeddings)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This project provides an implementation for fetching embeddings from various LLMs.

Currently supported APIs:
* [x] [OpenAI](https://platform.openai.com/docs/api-reference/embeddings)
* [x] [Cohere AI](https://docs.cohere.com/reference/embed)
* [x] [Google Vertex AI](https://cloud.google.com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings)

There are also simple command line tools provided by this project that let you query the APIs for text embeddings passed in via cli flags.

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
