<script lang="ts">
  import {AppBar, LightSwitch} from '@skeletonlabs/skeleton';
  import iconGithub from '$lib/icons/github.svg';
  import {darkMode} from '$lib/stores';
  import {get} from 'svelte/store';

  let dark: boolean = true;
  darkMode.subscribe((value: boolean) => {
    if (value === undefined) return;
    dark = value;
  });
</script>

<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
  <svelte:fragment slot="lead">
    <a href="https://github.com/gxd3v/gx-ytdl" target="_blank">
      <img src={iconGithub} alt="github" class={`w-full center h-10 pl-1 ${dark ? 'blackSVG' : 'whiteSVG '}`} />
    </a>
  </svelte:fragment>

  <div class="2xl:text-3xl xl:text-2xl lg:text-xl md:text-lg sm:text-md text-sm text-center">Youtube music & video downloader</div>

  <svelte:fragment slot="trail">
    <LightSwitch
      on:click={() => {
        darkMode.set(!get(darkMode));
      }} />
  </svelte:fragment>
</AppBar>
