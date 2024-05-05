import {writable} from 'svelte/store';
import * as util from '$lib/util';
import type WS from '$lib/ws';
import type {Message} from '$lib/models';

export const socket = writable<WebSocket>();
export const darkMode = writable<boolean>(false);
export const currentTile = writable<number>(0); // export const connectionRetries = writable<number>(1)
export const connectionFailed = writable<boolean>(false);
export const connectionClosed = writable<boolean>(false);
export const message = writable<string>();
