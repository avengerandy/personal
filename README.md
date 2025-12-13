# Personal Website Generator

A simple static site generator for my personal website.

The goal is to quickly and easily generate a deployable static website while keeping things simple (KISS principle).

Go was chosen because it is fast, type-safe, and its standard library provides template rendering. No additional frameworks or dependencies are required.

## Personal Website

**PC**

<p align="center">
  <img src = "https://github.com/avengerandy/personal/blob/master/screenshot_pc.png" width="90%"/>
</p>

**Mobile**

<p align="center">
  <img src = "https://github.com/avengerandy/personal/blob/master/screenshot_ph.png" width="90%"/>
</p>

[**Link**](https://avengerandy.github.io/personal/index.html)

## Architecture

The generator follows a clear flow:

```
JSON → Handler → Template → Docs
```

### Component Overview

* **Main**: Iterates through all handlers and generates HTML files.
* **Handler**: Reads JSON and passes data to templates.
* **RenderTemplateFromJSON[T any]**: Generic function that renders templates from any JSON-defined struct type.
* **Template**: HTML templates with placeholders for JSON data.

### Sequence
```mermaid
sequenceDiagram
    participant Main
    participant Handler
    participant RenderTemplate
    participant Template

    Main->>Handler: Initialize
    Handler->>Handler: Read JSON data
    Handler->>RenderTemplate: Call RenderTemplateFromJSON[T any]
    RenderTemplate->>Template: Render HTML
    Template->>RenderTemplate:
    RenderTemplate->>Main:
    Main->>Main: Write output to ./docs
```

## Build & Deployment

### Requirements

- Go 1.22 (or compatible version)

### Generate Site

```bash
go run main.go
```

All HTML files are generated in the ./docs folder, ready for deployment.
