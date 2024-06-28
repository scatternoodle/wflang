import { ExtensionContext, commands, window, workspace } from "vscode";

import {
  LanguageClient,
  LanguageClientOptions,
  TransportKind,
  ServerOptions,
} from "vscode-languageclient/node";

const serverPath = "editors/vscode/extension/server/bin/wflang";
const logPath = "editors/vscode/extension/server/logs/server.log";
const selector = { pattern: "**/*.wflang", scheme: "file", language: "wflang" }; // TODO - change to wflang once implemented

let client: LanguageClient;

const outputChannel = window.createOutputChannel("wflang");
const clientOptions: LanguageClientOptions = {
  documentSelector: [selector],
  synchronize: {
    fileEvents: workspace.createFileSystemWatcher("**/.clientrc"),
  },
  traceOutputChannel: outputChannel,
};

const serverOptions: ServerOptions = {
  run: {
    command: serverPath,
    args: [logPath],
    transport: TransportKind.stdio,
  },
  debug: {
    command: serverPath,
    args: [logPath],
    transport: TransportKind.stdio,
  },
};

export function activate(context: ExtensionContext) {
  client = new LanguageClient("wflsrv", serverOptions, clientOptions);
  client.start();
}

export function deactivate() {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
