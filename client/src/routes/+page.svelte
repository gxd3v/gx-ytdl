<script lang="ts">
  import Sidebar from '$lib/components/Sidebar.svelte';
  import Header from '$lib/components/Header.svelte';
  import Download from '$lib/components/Download.svelte';
  import '$lib/style.css';
  import {AppShell} from '@skeletonlabs/skeleton';
  import Session from '$lib/components/Session.svelte';
  import {connectionClosed, connectionFailed, currentTile, socket} from '$lib/stores';
  import {get} from 'svelte/store';
  import List from '$lib/components/List.svelte';
  import WS from '$lib/ws';
  import {browser} from '$app/environment';
  import Loader from '$lib/components/Loader.svelte';

  let ws: WS;
  let connectingMessage: string = '';

  let selectedTile: number = get(currentTile);
  let tries: number = 1;
  let maxTries: number = tries + 1;
  let connectionEstablished: boolean = false;

  currentTile.subscribe((value: number) => {
    if (value === undefined) return;
    selectedTile = value;
  });

  if (browser) {
    ws = new WS();
    maxTries = ws.maxRetries;
    connectingMessage = `Connecting ${tries}/${maxTries}`;

    connectionFailed.subscribe((value: boolean) => {
      if (value === undefined) return;
      if (value === false) return;

      if (tries <= maxTries) {
        setTimeout(() => {
          connectionFailed.set(false);
          ws = new WS();
          connectingMessage = `Connecting ${tries}/${maxTries}`;
          tries++;
        }, 5000);
      } else {
        connectingMessage = `Connection failed :(`;
      }
    });

    let connectionIntegretyChecker = setInterval(() => {
      if (ws !== undefined) {
        connectionEstablished = ws.IsConnected();
      }
    }, 1000);

    connectionClosed.subscribe((value: boolean) => {
      if (value === undefined) return;
      if (!value) return;

      clearInterval(connectionIntegretyChecker);
      connectionEstablished = false;
      tries = 0;
    });
  }

</script>

{#if ws !== undefined && !connectionEstablished}
  <Loader message={connectingMessage} show={tries <= maxTries} />
{:else if connectionEstablished}
  <AppShell class="h-screen">
    <svelte:fragment slot="header">
      <Header />
    </svelte:fragment>

    <svelte:fragment slot="sidebarLeft">
      <Sidebar />
    </svelte:fragment>

    <slot>
      {#if selectedTile === 0}
        <Download />
      {:else}
        <List />
      {/if}
    </slot>

    <svelte:fragment slot="pageFooter">
      <Session />
    </svelte:fragment>
  </AppShell>
{/if}