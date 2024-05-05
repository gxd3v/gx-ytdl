import {connectionClosed, connectionFailed, message, socket} from '$lib/stores';
import type {Message} from '$lib/models';
import {constants} from '$lib/constants';

export default class WS {
  private ws: WebSocket | null = null;
  public maxRetries: number = 5;

  constructor() {
    const ws = this.connect();

    ws.addEventListener('open', (ev: Event) => {
      console.log('connection opened', ev);
      socket.set(ws);
    });

    ws.addEventListener('message', ({data}) => {
      try {
        message.set(data);
      } catch (e) {}
    });

    ws.addEventListener('close', (ev: Event) => {
      console.log('connection closed', ev);
      this.close();
      connectionClosed.set(true);
    });

    ws.addEventListener('error', (ev: Event) => {
      console.log('websocket error', ev);
      connectionFailed.set(true);
    });
  }

  private connect(): WebSocket {
    let oldSession = localStorage.getItem(constants.sessionStorageKey);

    if (oldSession === null) {
      this.ws = new WebSocket('ws://127.0.0.1:7000/connect');
    } else {
      this.ws = new WebSocket(`ws://127.0.0.1:7000/connect/${oldSession}`);
    }

    return this.ws;
  }

  public SendMessage(message: Message): boolean {
    if (this.ws !== null) {
      this.ws.send(message.toString());
      return true;
    } else {
      return false;
    }
  }

  public IsConnected(): boolean {
    if (this.ws !== null) {
      if (this.ws.readyState === WebSocket.OPEN) {
        return true;
      }
    }

    return false;
  }

  private close() {
    if (this.ws !== null) {
      this.ws.close();
    }
  }
}
