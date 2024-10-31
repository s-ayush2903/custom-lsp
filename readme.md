# custom-lsp
A custom LSP server written in 100% Go. It's purpose was learning about how LSP
servers actually work at a deeper level, gaining understanding and implementing
a minimal version of it.

Built with go version 1.22.1 on darwin/arm64, and tested on NEOVIM `v0.9.5`
(Release) - [link](https://github.com/neovim/neovim/releases/tag/v0.9.5). Should
work with other editors as well given proper config of lsp at the client end.
A sample config used for testing this LSP server is linked [here](https://github.com/ayush-oyorooms/.dotfiles/blob/5a7ab2ba1c965cc1bc91de9a646a8a82d68976d1/nvim/.config/nvim/after/plugin/custom_lsp_server_test.lua)

A huuuge thanks to @tjdevries for sparking the interest in LSPs : )

## Functionalities Supported
### 1. Hover
This functionality in general displays the function / method / class synopsis as
a popup in usual editors. It is invoked when cursor is kept on a word for a
while. Here I have implemented a simple version of it which handles state and
dynamically tells the char count.

* LSP Spec: [Hover](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_hover)
* Invocation: `:lua vim.lsp.buf.hover()`

### 2. Go To Definition
This takes to the definition of the function / method(it is different from
implementation). This is invoked usually in editors when you 
`Ctrl / Cmd + click` on a keyword. In this implementaion it takes you to a line
above where the cursor is at the moment.

* LSP Spec: [Go To Definition](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_definition)
* Invocation: `:lua vim.lsp.buf.definition()`

### 3. Code Actions
When there are errors in your project and editor shows options to fix them,
these are the code actions. The server returns with the possible potential
solutions to the problems in the current context. This implementation takes it 
quite seriously xD.

* LSP Spec: [Code Actions](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_codeAction)
* Invocation: `:lua vim.lsp.buf.code_action()` _or_ `:Telescope diagnostics`

### 4. Completions
These are simply the suggestions which appear as you type. This implementation
provides a couple of suggestions based on typing.

* LSP Spec: [Completion](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion)
* Invocation: `:lua vim.lsp.buf.completion()`

### 4. Diagnostics (Notification)
These are simply the errors, warnings, hints along with some more info related
to them that editor keeps generating async-ly and shows them in the top right
section (usually). This implementation takes a few funny takes on editors.

* LSP Spec: [Diagnostics](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_publishDiagnostics)
* Invocation: `:lua vim.lsp.diagnostic.get_line_diagnostics()`  (these btw get generated in insert mode only)
