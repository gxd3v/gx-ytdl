import {writable} from 'svelte/store';
import * as util from '$lib/util';

export const socket = writable<WebSocket>();
export const darkMode = writable<boolean>(false);
export const currentTile = writable<number>(0);
