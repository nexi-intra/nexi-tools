---
title: Nexi Tools
description: Service for managing users favourite tools
---

#

# Nexi Tools

Supporting the sharing of relevant links within the `Nexi Group`.

- Latest version is found at [https://tools.home.nexi-intra.com](https://tools.home.nexi-intra.com)

- Internal implementation at [https://christianiabpos.sharepoint.com/sites/intra365/SitePages/Tools.aspx](https://christianiabpos.sharepoint.com/sites/intra365/SitePages/Tools.aspx)

- EY Proposal [Invision](https://eyitalia.invisionapp.com/console/share/NWBER3YUFA2/988226792)

This project is organized with PowerShell files for administrating the solution in subfolders to the root of the project.

The PowerShell files does not necarylly carry any functionality, but are used for documentation only. Within each folder you will find a readme.md file describing the purpose of the folder.

## Root folders

```
.
├── 20-usecases
├── 60-provision
├── 90-health
└── 91-setup


```

Numereuos hidden folders are used for different purposes

## Hidden folders

```

.
├── .devcontainer           # Used for configuring Code Spaces on GitHub
├── .git                    # Used for version control
│   ├── hooks
│   ├── info
│   ├── logs
│   ├── objects
│   └── refs
├── .github                 # Used for controlling CI/CD etc.
│   └── workflows
├── .koksmat                # Used for custom logic and UI
│   ├── app
│   ├── kitchenroot
│   ├── sessions
│   ├── web
│   └── workdir
├── .vscode                 # Used for tasks and debug configuration in VS Code

```

## .koksmat folder

```
.
├── app                                 # Source code for the GO application
│   ├── cmds
│   ├── endpoints
│   ├── execution
│   ├── magicapp
│   ├── schemas
│   ├── services
│   └── utils
├── sessions                            # Temporay storage for session execution
│   ├── 20240305T1345452575950100
│   └── 20240307T1409077240640100
├── web                                 # Source code for web UI
│   ├── app
│   ├── components
│   ├── koksmat
│   ├── node_modules
│   └── public
└── workdir                             # Temporay storage for cross session storage

```

## .koksmat/web/app folder

This is where all customization is expected to be made

```

.
├── favicon.ico
├── global.ts
├── globals.css
├── layout.tsx
├── page.tsx
├── sso
│   └── page.tsx
├── test
│   ├── page.tsx
│   └── usegraph
│       └── page.tsx
└── tools
    └── page.tsx

```
