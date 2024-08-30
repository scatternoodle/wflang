import { ExtensionContext, window, workspace } from "vscode";

import {
  LanguageClient,
  LanguageClientOptions,
  TransportKind,
  ServerOptions,
  Trace,
  SetTraceNotification,
} from "vscode-languageclient/node";

import fs from "fs";
import path from "path";

const serverPath = "editors/vscode/extension/server/bin/wflang";
const logPath = "editors/vscode/extension/server/logs/server.log";

const selector = { pattern: "**/*.wflang", scheme: "file", language: "wflang" };

let client: LanguageClient;

const outputChannel = window.createOutputChannel("wflang");
const clientOptions: LanguageClientOptions = {
  documentSelector: [selector],
  synchronize: {
    fileEvents: workspace.createFileSystemWatcher("**/.clientrc"),
  },
  traceOutputChannel: outputChannel,
  outputChannel: outputChannel,
  initializationOptions: {
    LogLevel: "verbose",
  },
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

function validateLogPath(): boolean {
  const logDir = path.dirname(logPath);

  if (!fs.existsSync(logDir)) {
    console.log(`Log directory does not exist, creating it at ${logDir}`);
    try {
      fs.mkdirSync(logDir, { recursive: true });
    } catch (e) {
      console.error(`Error creating logDir: ${e}`);
      return false;
    }
  }
  return true;
}

export function activate(context: ExtensionContext) {
  console.log(`Extension starting, CWD: ${process.cwd()}`);
  if (!validateLogPath()) {
    console.log("Unable to validate log path, exiting");
    return;
  }

  client = new LanguageClient("wflsrv", serverOptions, clientOptions);
  client.start();

  // I don't know why, but this appears to be the only way to get a server trace value of anything other than "off".
  // If we set initial setting in initialization, it's ignored. Using client.SetTrace() also appears to have no effect.
  client.sendNotification(SetTraceNotification.type.method, { value: Trace.Verbose });
}

export function deactivate() {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
