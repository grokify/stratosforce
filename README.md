# Stratosforce

**Multi-Agent System with Interoperable Agent Definitions**

Stratosforce is a multi-agent system framework that enables the definition of interoperable agents with shared input/output schemas defined by JSON Schema. Agents can be implemented across multiple platforms including Claude Code, AutoGen, CrewAI, LangGraph, and [trpc-agent-go](https://github.com/trpc-group/trpc-agent-go).

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.24.5-blue.svg)](go.mod)

---

## Overview

Stratosforce provides a framework for building multi-agent systems where:

- **Agents are platform-agnostic** - Define once, implement across multiple agent frameworks
- **Shared schemas** - JSON Schema definitions ensure consistent data interchange between agents
- **Interoperable** - Agents can communicate regardless of underlying implementation
- **Type-safe** - Go structs automatically generated from JSON schemas for compile-time safety
- **Extensible** - Easy to add new agent types and platforms

---

## Key Features

### 🔄 Cross-Platform Agent Definitions

Define agents that work across:
- **Claude Code** - Markdown-based agent definitions with built-in tooling
- **AutoGen** - Microsoft's multi-agent conversation framework
- **CrewAI** - Role-based agent collaboration framework
- **LangGraph** - Graph-based agent workflow system
- **trpc-agent-go** - Go-based agent framework

### 📋 JSON Schema Foundation

- Strict input/output contracts via JSON Schema
- Automatic validation and type checking
- Go code generation for type-safe operations
- Versioned schema evolution

### 🤖 Pre-Built Agent Systems

Currently includes:
- **Statistics Research System** - Multi-agent research, validation, and verification
- **News Agent** - News aggregation and analysis (in development)

---

## Architecture

```
stratosforce/
├── agents/              # Agent definitions by domain
│   ├── stats/          # Statistics research agents
│   │   ├── stats_research_agent.md
│   │   ├── stats_verification_agent.md
│   │   └── stats_orchestrator_agent.md
│   └── newsagent/      # News analysis agents
│
├── schemas/            # JSON Schema definitions
│   └── stats/
│       ├── stats_schema.json        # JSON Schema definition
│       ├── stats_schema.go          # Schema loader
│       └── gostats/
│           └── stats_research.go    # Generated Go types
│
├── agentutil/          # Shared agent utilities
│   └── definition.go   # Agent definition structures
│
└── examples/           # Example implementations
    └── spotify/        # Spotify AI statistics example
```

---

## Quick Start

### Prerequisites

- Go 1.24.5 or later
- Claude Code (for Claude Code agents)
- Access to desired agent frameworks (AutoGen, CrewAI, etc.)

### Installation

```bash
# Clone the repository
git clone https://github.com/grokify/stratosforce.git
cd stratosforce

# Install Go dependencies
go mod download

# Verify installation
go test ./...
```

### Using Claude Code Agents

1. **Install agents** to Claude Code:

```bash
# Copy to user agents directory
mkdir -p ~/claude/agents
cp agents/stats/*.md ~/claude/agents/

# Or use project-specific agents
mkdir -p .claude/agents
cp agents/stats/*.md .claude/agents/
```

2. **Run an agent**:

```bash
# Research statistics on a topic
/agent stats-research-agent "AI adoption in healthcare 2024"

# Verify statistics from research output
/agent stats-verification-agent --input research_output.json

# Full orchestrated workflow
/agent stats-orchestrator "renewable energy statistics 2024"
```

### Using Go Package

```go
package main

import (
    "fmt"
    "github.com/grokify/stratosforce/schemas/stats/gostats"
)

func main() {
    // Load statistics report
    report, err := gostats.LoadFromFile("research_output.json")
    if err != nil {
        panic(err)
    }

    // Access statistics
    fmt.Printf("Topic: %s\n", report.Topic)
    fmt.Printf("Total statistics: %d\n", len(report.Statistics))

    // Get verified statistics only
    verified := report.GetVerifiedStatistics()
    fmt.Printf("Verified statistics: %d\n", len(verified))

    // Check if ready for publication
    if report.IsReadyForPublication() {
        fmt.Println("✅ All statistics verified - ready to use!")
    }
}
```

---

## Agent Systems

### Statistics Research & Validation System

A comprehensive multi-agent system for finding, validating, and verifying statistics with strict anti-hallucination safeguards.

**Agents:**
- `stats-research-agent` - Finds and validates statistics from web sources
- `stats-verification-agent` - Verifies statistics against original sources
- `stats-orchestrator` - Coordinates research and verification workflow

**Features:**
- Multi-source validation
- Source accessibility checking
- Confidence level assessment
- Structured JSON output
- Full audit trail

**Documentation:** [agents/stats/README.md](agents/stats/README.md)

**Example Output:** [schemas/stats/examples/spotify/](schemas/stats/examples/spotify/)

---

## Schema System

### JSON Schema Definitions

All agent inputs and outputs are defined using JSON Schema (Draft 07):

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Statistics Research Report",
  "type": "object",
  "required": ["topic", "query_date", "summary", "statistics", "sources"],
  "properties": {
    "topic": {
      "type": "string",
      "description": "The topic researched"
    },
    "statistics": {
      "type": "array",
      "items": { ... }
    }
  }
}
```

### Go Type Generation

Go structs are generated from schemas for type-safe operations:

```go
type Report struct {
    Topic      string      `json:"topic"`
    QueryDate  string      `json:"query_date"`
    Summary    Summary     `json:"summary"`
    Statistics []Statistic `json:"statistics"`
    Sources    []Source    `json:"sources"`
}
```

### Schema Validation

```go
import "github.com/grokify/stratosforce/schemas/stats"

// Validate JSON against schema
schema := stats.StatsSchema()
err := schema.Validate(jsonData)
```

---

## Platform Integration

### Claude Code

Agent definitions are written in Markdown with frontmatter:

```markdown
---
name: stats-research-agent
description: Statistics Research & Validation Agent
color: orange
---

## Purpose
This agent researches statistics...

## Instructions
You are a statistics research agent...
```

### AutoGen (Coming Soon)

```python
from stratosforce.agents import StatsResearchAgent

agent = StatsResearchAgent(
    schema_path="schemas/stats/stats_schema.json"
)
result = agent.research("AI adoption statistics")
```

### CrewAI (Coming Soon)

```python
from crewai import Agent
from stratosforce.schemas import StatsSchema

researcher = Agent(
    role="Statistics Researcher",
    goal="Find accurate statistics",
    output_schema=StatsSchema
)
```

---

## Examples

### Spotify AI Statistics

A complete example of the statistics research system analyzing Spotify's AI features:

- **Input:** "Spotify AI-driven music discovery statistics"
- **Research Output:** [spotify_ai_statistics.json](schemas/stats/examples/spotify/spotify_ai_statistics.json)
- **Verified Output:** [spotify_ai_statistics_verified.json](schemas/stats/examples/spotify/spotify_ai_statistics_verified.json)
- **Verification Report:** [VERIFICATION_REPORT.md](schemas/stats/examples/spotify/VERIFICATION_REPORT.md)

**Results:**
- 14 statistics researched
- 5 verified accurate (36%)
- 3 verified inaccurate
- 3 context issues identified
- Full source traceability

---

## Development

### Adding a New Agent

1. **Define the agent** in `agents/<domain>/`:
```markdown
---
name: my-agent
description: My custom agent
---

## Purpose
...

## Instructions
...
```

2. **Create JSON schema** in `schemas/<domain>/`:
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "My Agent Output",
  ...
}
```

3. **Generate Go types** in `schemas/<domain>/go<domain>/`:
```go
type MyAgentOutput struct {
    // Define structs matching schema
}
```

4. **Test the agent**:
```bash
/agent my-agent "test query"
```

### Running Tests

```bash
# Run all tests
go test ./...

# Test specific package
go test ./schemas/stats/...

# With verbose output
go test -v ./...
```

---

## Use Cases

- **Academic Research** - Find and verify statistics for papers
- **Business Intelligence** - Extract market data with source verification
- **Journalism** - Fact-check statistics with full audit trail
- **Policy Analysis** - Research policy impact statistics
- **Grant Writing** - Find supporting statistics with citations
- **Market Research** - Multi-source validation of market data

---

## Roadmap

- [ ] **Additional Agent Systems**
  - News aggregation and analysis
  - Financial data research
  - Scientific literature review

- [ ] **Platform Integrations**
  - AutoGen adapter implementation
  - CrewAI adapter implementation
  - LangGraph integration
  - trpc-agent-go examples

- [ ] **Schema Enhancements**
  - Additional validation rules
  - Schema versioning system
  - Migration tooling

- [ ] **Developer Tools**
  - CLI for agent management
  - Schema-to-code generators
  - Testing framework

---

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

Please ensure:
- All tests pass
- Code follows Go conventions
- JSON schemas are valid
- Documentation is updated

---

## Related Projects

- [trpc-agent-go](https://github.com/trpc-group/trpc-agent-go) - Go agent framework
- [AutoGen](https://github.com/microsoft/autogen) - Multi-agent conversation framework
- [CrewAI](https://github.com/joaomdmoura/crewAI) - Role-based agent framework
- [LangGraph](https://github.com/langchain-ai/langgraph) - Graph-based agent workflows

---

## Support

- **Issues:** [GitHub Issues](https://github.com/grokify/stratosforce/issues)
- **Discussions:** [GitHub Discussions](https://github.com/grokify/stratosforce/discussions)
- **Documentation:** [Agent Documentation](agents/stats/README.md)

---

**Built for interoperable, schema-driven multi-agent systems.**
