<script lang="ts">
  import {AppRail, AppRailAnchor, AppRailTile} from '@skeletonlabs/skeleton';
  import iconDownload from '$lib/icons/download.svg';
  import iconFolder from '$lib/icons/folder.svg';
  import iconDelete from '$lib/icons/delete.svg';
  import {currentTile, darkMode} from '$lib/stores';
  import * as util from '$lib/util';

  let selectedTile: number = 0;
  let dark: boolean = true;

  darkMode.subscribe((value: boolean) => {
    if (value === undefined) return;
    dark = value;

  });
</script>

<AppRail>
  <AppRailTile
    bind:group={selectedTile}
    name="Download"
    value={0}
    title="Download"
    on:click={() => {
      currentTile.set(0);
      util.SaveTile(0);
    }}>
    <span>Download</span>
    <svelte:fragment slot="lead">
      <img src={iconDownload} alt="download icon" class={`w-full center h-7 ${dark ? 'blackSVG' : 'whiteSVG '}`} />
    </svelte:fragment>
  </AppRailTile>

  <AppRailTile
    bind:group={selectedTile}
    name="My downloads"
    value={1}
    title="My downloads"
    on:click={() => {
      currentTile.set(1);
      util.SaveTile(1);
    }}>
    <svelte:fragment slot="lead">
      <img src={iconFolder} alt="my downloads icon" class={`w-full center h-7 ${dark ? 'blackSVG' : 'whiteSVG '}`} />
    </svelte:fragment>
    <span>My downloads</span>
  </AppRailTile>

  <svelte:fragment slot="trail">
    <AppRailTile
      bind:group={selectedTile}
      name="Delete session"
      value={2}
      title="Delete session"
      on:click={() => {
        currentTile.set(2);
        util.SaveTile(2);
      }}>
      <svelte:fragment slot="lead">
        <img src={iconDelete} alt="delete session icon" class={`w-full center h-7 ${dark ? 'blackSVG' : 'whiteSVG '}`} />
      </svelte:fragment>
      <span>Delete &nbsp;session</span>
    </AppRailTile>
  </svelte:fragment>
</AppRail>
