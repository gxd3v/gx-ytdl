<script lang="ts">
  import Sidebar from '$lib/components/Sidebar.svelte';
  import Header from '$lib/components/Header.svelte';
  import Download from '$lib/components/Download.svelte';
  import '$lib/style.css';
  import {AppShell} from '@skeletonlabs/skeleton';
  import Session from '$lib/components/Session.svelte';
  import {currentTile} from '$lib/stores';
  import {get} from 'svelte/store';
  import List from '$lib/components/List.svelte';

  let selectedTile: number = get(currentTile);
  currentTile.subscribe((value: number) => {
    if (value === undefined) return;
    selectedTile = value;
  });
</script>

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
